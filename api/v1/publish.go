package v1

import (
	"github.com/gin-gonic/gin"
	"liao/global"
	"liao/global/response"
	"liao/model"
	"liao/utils"
	"strconv"
)

// @Tags publish
// @Summary 发布信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param selectType path int true "发布信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发布成功"}"
// @Router /api/addPublish [post]
func AddPublish(c *gin.Context) {
	var publishInfo model.PublishInfo
	_ = c.ShouldBindJSON(&publishInfo)
	if publishInfo.Content == "" {
		response.FailWithMessage("随记内容为空,请确认!", c)
		return
	}
	var publish model.Publish
	var publishPicInfo model.PublishPic
	tx := global.DB.Begin()
	//写入主表
	publish.Position = "四川成都"
	publish.UserId = 1
	publish.Content = publishInfo.Content
	createErr := global.DB.Create(&publish).Error
	if createErr != nil {
		tx.Callback()
		response.OkWithMessage("随记发布失败，请重试！", c)
	} else {
		if len(publishInfo.Files) > 0 {
			for _, file := range publishInfo.Files{
				//写入图片表
				publishPicInfo.Pic = file
				publishPicInfo.PubId = publish.Id
				global.DB.Create(publishPicInfo)
			}
		}
		tx.Commit()
		response.OkWithMessage("随记发布成功", c)
	}
}

// @Tags publish
// @Summary 获取发布信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param selectType path int true "获取发布信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发布信息列表"}"
// @Router /api/getCateList/{selectType} [post]
func PublishList(c *gin.Context) {
	var pageInfo model.PageInfo
	var publishInfo []model.Publish
	var total int64
	_ = c.ShouldBindJSON(&pageInfo)
	limit := pageInfo.PageSize
	offset := pageInfo.PageSize * (pageInfo.Page - 1)
	db := global.DB.Model(&model.Publish{}).Where("user_id = 1")
	db.Count(&total)
	db.Preload("Pics").Limit(limit).Offset(offset).Order("created_at DESC").Find(&publishInfo)
	if len(publishInfo) == 0 {
		response.FailWithMessage("暂无随记", c)
		return
	}
	res := utils.SetOrGetRedisKey("publishList_"+strconv.Itoa(pageInfo.Page), publishInfo, false)
	response.OkWithData(model.PageResult{
		List:     res,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, c)
}
