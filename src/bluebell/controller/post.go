package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"strconv"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// CreatePostHandler 创建帖子的处理函数
func CreatePostHandler(c *gin.Context) {
	// 1、获取参数及参数校验
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Debug("c.ShouldBindJSON(p) error", zap.Any("err", err))
		zap.L().Error("Create post with invalid param")
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 从 c 中取到当前发送请求的用户ID
	userID, err := GetCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	p.AuthorID = userID
	// 2、创建帖子
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3、返回响应
	ResponseSuccess(c, nil)
}

// GetPostDetailHandler 获取帖子详情的处理函数
func GetPostDetailHandler(c *gin.Context) {
	// 1、获取参数(从URL中获取帖子的id)
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("get post detail invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 2、根据id取出帖子数据 mysql
	data, err := logic.GetPostById(pid)
	if err != nil {
		zap.L().Error("logic.GetPostById(pid) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3、返回响应
	ResponseSuccess(c, data)
}

// GetPostListHandler 获取帖子列表的接口
func GetPostListHandler(c *gin.Context) {
	// 获取分页参数
	page, size := getPageInfo(c)
	// 1、获取数据
	data, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 2、返回响应
	ResponseSuccess(c, data)
}

// GetPostListHandler2 升级版帖子列表的接口
// 根据前端传来的参数动态获取帖子列表
// 按创建时间排序 或者 按分数排序
// 1、获取请求的 query string 参数
// 2、去redis查询id列表
// 3、根据id去数据库查询帖子详细信息

// GetPostListHandler2 升级版帖子列表接口
// @Summary 升级版帖子列表接口
// @Description 可以按社区按时间或分数排序查询帖子列表接口
// @Tags 帖子相关的接口(api分组展示使用的)
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Param object query models.ParamPostList false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /post2 [get]

func GetPostListHandler2(c *gin.Context) {
	// GET请求参数：(query string参数) /api/v1/posts2?page=1&size=10&order=time
	// 指定初始参数默认值
	p := &models.ParamPostList{
		Page:  1,
		Size:  10,
		Order: models.OrderTime,
	}
	// c.ShouldBind() // 动态的选择相应的方法获取数据
	// c.ShouldBindJSON() 如果请求中携带的是json格式的数据，才能用这个方法获取到数据

	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("GetPostListHandler2 with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 更新： 合二为一
	data, err := logic.GetPostListNew(p)
	// 获取数据
	if err != nil {
		zap.L().Error("logic.GetPostList2() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 2、返回响应
	ResponseSuccess(c, data)
}

//// GetCommunityPostListHandler 根据社区查询帖子
//func GetCommunityPostListHandler(c *gin.Context) {
//	p := &models.ParamCommunityPostList{
//		ParamPostList: &models.ParamPostList{
//			Page:  1,
//			Size:  10,
//			Order: models.OrderTime,
//		},
//		//CommunityID: 0,
//	}
//
//	// c.ShouldBind() // 动态的选择相应的方法获取数据
//	// c.ShouldBindJSON() 如果请求中携带的是json格式的数据，才能用这个方法获取到数据
//
//	if err := c.ShouldBindQuery(p); err != nil {
//		zap.L().Error("GetCommunityPostListHandler with invalid param", zap.Error(err))
//		ResponseError(c, CodeInvalidParam)
//		return
//	}
//	// 1、获取数据
//	data, err := logic.GetCommunityPostList2(p)
//	if err != nil {
//		zap.L().Error("logic.GetCommunityPostList2() failed", zap.Error(err))
//		ResponseError(c, CodeServerBusy)
//		return
//	}
//	// 2、返回响应
//	ResponseSuccess(c, data)
//}
