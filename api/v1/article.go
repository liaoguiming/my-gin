package v1

import (
	"github.com/gin-gonic/gin"
	"liao/global"
	"liao/global/response"
	"liao/model"
	"liao/utils"
	"strconv"
)

// @Tags Article
// @Summary 获取文章分类文章
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ArticleByType true "文章分类ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getArticleList [post]
func GetArticleList(c *gin.Context) {
	var pageInfo model.ArticleByType
	var total int64
	var articles []model.Article
	_ = c.ShouldBindJSON(&pageInfo)
	pid := pageInfo.CatId
	if pid == 0 {
		response.FailWithMessage("catId不能为空", c)
		return
	}
	limit := pageInfo.PageSize
	offset := pageInfo.PageSize * (pageInfo.Page - 1)
	db := global.DB.Model(&model.Article{}).Where("status = 0")
	db.Count(&total)
	db.Limit(limit).Offset(offset).Order("created_at DESC").Where("category_id = ?", pid).Find(&articles)
	if len(articles) == 0 {
		response.FailWithMessage("未找到对应文章", c)
		return
	}
	res := utils.SetOrGetRedisKey("article_"+strconv.Itoa(pid), articles, false)
	response.OkWithData(model.PageResult{
		List:     res,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, c)
}

// @Tags Article
// @Summary 获取小程序首页推荐文章
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PageInfo true "获取小程序首页推荐文章"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getIndexArticleList [post]
func GetIndexArticleList(c *gin.Context) {
	var pageInfo model.PageInfo
	var articles []model.Article
	var total int64
	_ = c.ShouldBindJSON(&pageInfo)
	limit := pageInfo.PageSize
	offset := pageInfo.PageSize * (pageInfo.Page - 1)
	db := global.DB.Model(&model.Article{}).Where("is_index = 0")
	db.Count(&total)
	db.Limit(limit).Offset(offset).Order("created_at DESC").Find(&articles)
	if len(articles) == 0 {
		response.FailWithMessage("暂无首页推荐文章", c)
		return
	}
	res := utils.SetOrGetRedisKey("articleIndex_"+strconv.Itoa(pageInfo.Page), articles, false)
	response.OkWithData(model.PageResult{
		List:     res,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, c)
}

// @Tags Article
// @Summary 获取文章内容详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ArticleId true "根据文章ID获取文章内容"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getArticleDetail [post]
func GetArticleDetail(c *gin.Context) {
	var article []model.ArticleDetail
	var articleId model.ArticleId
	_ = c.ShouldBindJSON(&articleId)
	aId := articleId.Id
	if aId == 0 {
		response.FailWithMessage("文章id不能为空", c)
		return
	}
	db := global.DB.Model(&model.Article{}).Where("id = ?", aId)
	db.Find(&article)
	if len(article) == 0 {
		response.FailWithMessage("未找到文章内容", c)
		return
	}
	res := utils.SetOrGetRedisKey("articleDetail_"+strconv.Itoa(aId), article, true)
	response.OkWithData(res, c)
}
