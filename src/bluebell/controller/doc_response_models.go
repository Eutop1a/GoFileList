package controller

import "bluebell/models"

// 专门用来放接口文档用到的model
// 因为接口文档返回的数据类型格式是一致的，

type _ResponsePostList struct {
	Code    ResCode                 `json:"code"`    // 业务相应状态码
	Message string                  `json:"message"` // 提示信息
	Data    []*models.ApiPostDetail `json:"data"`    // 数据
}
