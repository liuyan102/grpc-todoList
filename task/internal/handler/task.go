package handler

import (
	"context"
	"task/internal/repository"
	"task/internal/service"
	"task/pkg/e"
)

type TaskService struct {
}

func NewTaskService() *TaskService {
	return &TaskService{}
}

func (*TaskService) TaskCreate(ctx context.Context, req *service.TaskRequest) (*service.CommonResponse, error) {
	var task repository.Task
	resp := &service.CommonResponse{}
	err := task.TaskCreate(req)
	if err != nil {
		resp.Code = e.Error
		resp.Msg = e.GetMsg(e.Error)
		resp.Data = err.Error()
		return resp, err
	}
	resp.Code = e.Success
	resp.Msg = e.GetMsg(e.Success)
	return resp, nil
}

func (*TaskService) TaskUpdate(ctx context.Context, req *service.TaskRequest) (*service.CommonResponse, error) {
	var task repository.Task
	resp := &service.CommonResponse{}
	err := task.TaskUpdate(req)
	if err != nil {
		resp.Code = e.Error
		resp.Msg = e.GetMsg(e.Error)
		resp.Data = err.Error()
		return resp, err
	}
	resp.Code = e.Success
	resp.Msg = e.GetMsg(e.Success)
	return resp, nil
}

func (*TaskService) TaskShow(ctx context.Context, req *service.TaskRequest) (*service.TaskDetailResponse, error) {
	var task repository.Task
	resp := &service.TaskDetailResponse{}
	taskList, err := task.TaskShow(req)
	if err != nil {
		resp.Code = e.Error
		resp.TaskDetail = nil
		return resp, err
	}
	resp.Code = e.Success
	resp.TaskDetail = repository.BuildTaskList(taskList)
	return resp, nil

}

func (*TaskService) TaskDelete(ctx context.Context, req *service.TaskRequest) (*service.CommonResponse, error) {
	var task repository.Task
	resp := &service.CommonResponse{}
	err := task.TaskDelete(req)
	if err != nil {
		resp.Code = e.Error
		resp.Msg = e.GetMsg(e.Error)
		resp.Data = err.Error()
		return resp, err
	}
	resp.Code = e.Success
	resp.Msg = e.GetMsg(e.Success)
	return resp, nil
}
