package utils

import (
	"go.uber.org/zap"
	"io"
	"liao/global"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"time"
)

func UploadFile(file *multipart.FileHeader) (err error, localPath string, key string) {
	// 读取文件后缀
	ext := path.Ext(file.Filename)
	// 读取文件名并加密
	fileName := strings.TrimSuffix(file.Filename, ext)
	fileName = MD5V([]byte(fileName))
	// 拼接新文件名
	lastName := fileName + "_" + time.Now().Format("20060102150405") + ext
	// 读取全局变量的定义路径
	savePath := "/home/html/liao/public/uploads/miniProgramPic"
	// 尝试创建此路径
	err = os.MkdirAll(savePath, os.ModePerm)
	if err != nil {
		global.LOG.Error("upload local file fail:", zap.Any("err", err))
		return err, "", ""
	}
	// 拼接路径和文件名
	dst := savePath + "/" + lastName
	// 下面为上传逻辑
	// 打开文件 defer 关闭
	src, err := file.Open()
	if err != nil {
		global.LOG.Error("upload local file fail:", zap.Any("err", err))
		return err, "", ""
	}
	defer src.Close()
	// 创建文件 defer 关闭
	out, err := os.Create(dst)
	if err != nil {
		global.LOG.Error("upload local file fail:", zap.Any("err", err))
		return err, "", ""
	}
	defer out.Close()
	// 传输（拷贝）文件
	_, err = io.Copy(out, src)
	if err != nil {
		global.LOG.Error("upload local file fail:", zap.Any("err", err))
		return err, "", ""
	}
	resFileName := "/miniProgramPic/" + lastName
	return nil, resFileName, lastName
}
