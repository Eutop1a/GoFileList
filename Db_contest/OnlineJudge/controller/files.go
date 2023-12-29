package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
)

type File struct {
	UserName string                `form:"username" json:"username"`
	Input    *multipart.FileHeader `form:"input" json:"input"`
	Output   *multipart.FileHeader `form:"output" json:"output"`
}

const Dir = "D:\\MyOjProblems"

// AddFiles 上传问题文件
// @Summary 上传问题文件
// @Description 通过管理员身份，使用表单数据上传问题的输入和输出文件，并返回操作结果
// @Accept multipart/form-data
// @Produce json
// @Param username formData string true "执行操作的管理员用户名"
// @Param id path string true "要上传文件的问题ID"
// @Param input formData file true "问题输入文件（.in）"
// @Param output formData file true "问题输出文件（.out）"
// @Success 200 {object} models._ResponseMsg "成功上传问题文件"
// @Failure 403 {object} models._ResponseMsg "Token 已超时或用户非管理员"
// @Failure 400 {object} models._ResponseMsg "接收或保存文件出错"
// @Router /problem/file/add/:id [post]
func AddFiles(c *gin.Context) {
	var err error
	// 从请求中获取执行操作的管理员用户名
	var file File
	// 从请求中获取执行操作的用户名
	tmp := c.Request.Form.Get("username")

	if tmp == "" {
		// Token 已超时
		zap.L().Error("TokenTimeOut")
		c.JSON(http.StatusForbidden, gin.H{
			"Msg": "TokenTimeOut",
		})
		return
	} else {
		file.UserName = tmp
	}
	id := c.Param("id")

	// 接收输入和输出文件
	file.Input, err = c.FormFile("input")
	if err != nil {
		zap.L().Error("recv input file error", zap.Error(err))
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"Msg": fmt.Sprintf("recv input file %s\n", err),
		})
		return
	}

	file.Output, err = c.FormFile("output")
	if err != nil {
		zap.L().Error("recv output file error", zap.Error(err))
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"Msg": fmt.Sprintf("recv output file %s\n", err),
		})
		return
	}

	dir := Dir + "\\" + "Problem_" + id + "\\"

	// 设置文件保存路径
	inputDir := dir + "p_" + id + ".in"
	outputDir := dir + "p_" + id + ".out"

	// 保存文件
	err = c.SaveUploadedFile(file.Input, inputDir)
	if err != nil {
		zap.L().Error("save input file error", zap.Error(err))
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"Msg": fmt.Sprintf("save input file %s\n", err),
		})
		return
	}

	err = c.SaveUploadedFile(file.Output, outputDir)
	if err != nil {
		zap.L().Error("save output file error", zap.Error(err))
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"Msg": fmt.Sprintf("save output file %s\n", err),
		})
		return
	}
	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{
		"Msg": "OK",
	})
}

// GetFiles 获取问题文件列表
// @Summary 获取问题文件列表
// @Description 通过问题ID，获取指定问题的文件列表
// @Accept json
// @Produce json
// @Param id formData string true "要获取文件列表的问题ID"
// @Success 200 {object} models._ResponseGetData "成功获取问题文件列表"
// @Failure 403 {object} models._ResponseDataString "读取文件列表出错"
// @Router /problem/file/:id [post]
func GetFiles(c *gin.Context) {
	// 从请求中获取要获取文件列表的问题ID
	dir := Dir + "//" + c.PostForm("id")

	// 读取指定目录下的文件列表
	if files, err := ioutil.ReadDir(dir); err != nil {
		zap.L().Error("ReadDir error", zap.Error(err))
		c.JSON(http.StatusForbidden, gin.H{
			"data": fmt.Sprintf("%s", err),
		})
	} else {
		// 将文件列表转换为字符串切片
		result := make([]string, len(files))
		for k, info := range files {
			result[k] = info.Name()
		}
		// 返回成功响应
		c.JSON(http.StatusOK, gin.H{
			"data": result,
		})
	}
}
