package service

import (
	"singo/model"
	"singo/serializer"

	"github.com/gin-gonic/gin"
)

// CreateVideoService 创建视频投稿服务
type CreateVideoService struct {
	Title string `form:"title" json:"title" binding:"required,min=2,max=30"`
	Info  string `form:"info" json:"info" binding:"required,max=200"`
}

// Create 创建视频
func (service *CreateVideoService) Create(c *gin.Context) serializer.Response {
	video := model.Video{
		Title: service.Title,
		Info:  service.Info,
	}
	err := model.DB.Create(&video).Error
	if err != nil {
		return serializer.Response{
			Code:  50001,
			Msg:   "视频保存失败",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildVideo(video),
	}
}
