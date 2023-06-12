package handler

import (
	"api-gateway/internal/service"
	"api-gateway/pkg/e"
	"api-gateway/pkg/res"
	"api-gateway/pkg/util"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TaskCreate(ctx *gin.Context) {
	var taskReq service.TaskRequest
	// 绑定数据
	PanicIfTaskError(ctx.Bind(&taskReq))
	// 获取并解析token
	claim, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	// 获取token中的userID
	taskReq.UserID = uint32(claim.UserId)
	// gin.key中获取服务实例
	taskService := ctx.Keys["task"].(service.TaskServiceClient)
	commonResp, err := taskService.TaskCreate(context.Background(), &taskReq)
	PanicIfTaskError(err)
	r := res.Response{
		Data:   commonResp,
		Status: uint(commonResp.Code),
		Msg:    e.GetMsg(uint(commonResp.Code)),
	}
	ctx.JSON(http.StatusOK, r)
}

func TaskShow(ctx *gin.Context) {
	var taskReq service.TaskRequest
	PanicIfTaskError(ctx.Bind(&taskReq))
	// 获取并解析token
	claim, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	// 获取token中的userID
	taskReq.UserID = uint32(claim.UserId)
	taskService := ctx.Keys["task"].(service.TaskServiceClient)
	taskDetailResp, err := taskService.TaskShow(context.Background(), &taskReq)
	PanicIfTaskError(err)
	r := res.Response{
		Data:   taskDetailResp,
		Status: uint(taskDetailResp.Code),
		Msg:    e.GetMsg(uint(taskDetailResp.Code)),
	}
	ctx.JSON(http.StatusOK, r)
}

func TaskUpdate(ctx *gin.Context) {
	var taskReq service.TaskRequest
	PanicIfTaskError(ctx.Bind(&taskReq))
	// 获取并解析token
	claim, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	// 获取token中的userID
	taskReq.UserID = uint32(claim.UserId)
	taskService := ctx.Keys["task"].(service.TaskServiceClient)
	commonResp, err := taskService.TaskUpdate(context.Background(), &taskReq)
	PanicIfTaskError(err)
	r := res.Response{
		Data:   commonResp,
		Status: uint(commonResp.Code),
		Msg:    e.GetMsg(uint(commonResp.Code)),
	}
	ctx.JSON(http.StatusOK, r)
}

func TaskDelete(ctx *gin.Context) {
	var taskReq service.TaskRequest
	PanicIfTaskError(ctx.Bind(&taskReq))
	// 获取并解析token
	claim, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	// 获取token中的userID
	taskReq.UserID = uint32(claim.UserId)
	taskService := ctx.Keys["task"].(service.TaskServiceClient)
	commonResp, err := taskService.TaskDelete(context.Background(), &taskReq)
	PanicIfTaskError(err)
	r := res.Response{
		Data:   commonResp,
		Status: uint(commonResp.Code),
		Msg:    e.GetMsg(uint(commonResp.Code)),
	}
	ctx.JSON(http.StatusOK, r)
}
