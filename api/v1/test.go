package v1

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"liao/global"
	"liao/global/response"
	"liao/model"
	"liao/utils"
	"time"
)

// @title 获取省级名称列表
// @version v1
// @description 获取省级地区列表
// @Tags Province
// @Summary 获取中国省级地区列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/ [get]
func Province(c *gin.Context) {
	var area []model.Area
	str, _ := global.REDIS.Get("proviceList").Result()
	if str == "" {
		db := global.DB.Model(&model.Area{}).Where("type = 1")
		db.Find(&area)
		//序列化
		info, _ := json.Marshal(area)
		if string(info) != "" {
			global.REDIS.Set("provinceList", string(info), 86400*time.Second)
		}
	} else {
		_ = json.Unmarshal([]byte(str), &area)
	}
	response.OkWithData(area, c)
}

// @Tags category
// @Summary 获取分类列表信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param selectType path int true "获取不同数据类型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getCateList/{selectType} [get]
func GetCateList(c *gin.Context) {
	selectType := c.Param("selectType")
	str, _ := global.REDIS.Get("cateList").Result()
	var cat []model.Cates
	var cates []model.Categories
	switch selectType {
	case "1":
		if str == "" {
			db := global.DB.Model(&model.Categories{})
			db.Find(&cat)
			//序列化
			info, _ := json.Marshal(cat)
			global.REDIS.Set("cateList", string(info), 86400*time.Second)
		} else {
			_ = json.Unmarshal([]byte(str), &cat)
		}
		response.OkWithData(cat, c)
	case "2":
		db := global.DB.Model(&model.Categories{})
		db.Find(&cates)
		response.OkWithData(cates, c)
	default:
		response.OkWithMessage("nothing", c)
	}
}

// @Tags category
// @Summary 获取小程序主页面分类
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getMainCategory [get]
func GetMainCategory(c *gin.Context) {
	var mainCatInfo []model.Categories
	db := global.DB.Table("categories cat").Select("distinct cat.name,cat.id").Joins("inner join articles art on art.category_id = cat.id").Where("cat.sort = 2")
	db.Find(&mainCatInfo)
	if len(mainCatInfo) == 0 {
		response.FailWithMessage("暂无分类", c)
		return
	}
	res := utils.SetOrGetRedisKey("mainCategory", mainCatInfo, true)
	response.OkWithData(res, c)
}
