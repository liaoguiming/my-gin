package basic

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	//返回URL body的内容,并且将在这个页面上找到的URl放入一个slice中
	Fetch(url string) (body string, urls []string, err error)
}

type safeState struct {
	v   map[string]bool
	mux sync.Mutex
}

//使用sync.Mutex保证不同协程间保存url不冲突
func (c *safeState) setState(key string, state bool) {
	c.mux.Lock()
	c.v[key] = state
	c.mux.Unlock()
}

//验证url是否重复
func (c *safeState) value(key string) (bool, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	v, ok := c.v[key]
	return v, ok
}

var urlState = safeState{v: make(map[string]bool)}

func Crawl(url string, depth int, fetcher Fetcher, ch chan int) {
	defer func() {
		ch <- 1
	}()
	if depth < 0 {
		return
	}
	//验证urlState是否包含url 包含则不抓取该URL
	if _, ok := urlState.value(url); ok {
		return
	}
	urlState.setState(url, false)
	defer urlState.setState(url, true)
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	subChan := make(chan int)
	fmt.Printf("found:%s %q\n", url, body)
	for _, u := range urls {
		go Crawl(u, depth-1, fetcher, subChan)
	}

	for i := 0; i < len(urls); i++ {
		<-subChan
	}
	return
}

type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher 是填充后的 fakeFetcher。
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

func main() {
	ch := make(chan int, 4)
	go Crawl("https://golang.org/", 4, fetcher, ch)
	<-ch
}
