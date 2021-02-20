package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"liao/global/response"
	"liao/utils"
)

// @Tags upload
// @Summary 上传图片
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param selectType path int true "上传图片"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"上传成功"}"
// @Router /api/upload [post]

func Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取图片失败，%v", err), c)
	} else {
		uploadErr, filePath, _ := utils.UploadFile(file)
		if uploadErr != nil {
			response.FailWithMessage(fmt.Sprintf("图片上传失败，%v", uploadErr), c)
		} else {
			response.OkWithData(filePath, c)
		}
	}
}
