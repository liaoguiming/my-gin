package v1

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"liao/global"
	"liao/global/response"
	"liao/model"
	"time"
)

// @Tags Article
// @Summary 获取文章主分类
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/mainTypeLIst [get]
func MainTypeList(c *gin.Context) {
	var mainType []model.Article
	str, _ := global.REDIS.Get("mainType").Result()
	if str == "" {
		global.DB.Where("sort = 2").Model(&model.Article{}).Find(&mainType)
		info, _ := json.Marshal(mainType)
		global.REDIS.Set("mainType", string(info), 86400*time.Second)
	} else {
		_ = json.Unmarshal([]byte(str), &str)
	}
	response.OkWithData(mainType, c)
}
