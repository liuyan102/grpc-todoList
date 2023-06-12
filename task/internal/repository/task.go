package repository

import "task/internal/service"

type Task struct {
	TaskID    uint `gorm:"primaryKey"`
	UserID    uint `gorm:"index"`
	Status    int  `gorm:"default:0"`
	Title     string
	Content   string `gorm:"type:longtext"`
	StartTime int64
	EndTime   int64
}

func (*Task) TaskCreate(request *service.TaskRequest) error {
	task := &Task{
		UserID:    uint(request.UserID),
		Title:     request.Title,
		Content:   request.Content,
		StartTime: int64(request.StartTime),
		EndTime:   int64(request.EndTime),
	}
	return DB.Create(&task).Error
}

func (*Task) TaskUpdate(request *service.TaskRequest) error {
	var task Task
	err := DB.Model(&Task{}).Where("task_id=?", request.TaskID).First(&task).Error
	if err != nil {
		return err
	}
	task.UserID = uint(request.UserID)
	task.Status = int(request.Status)
	task.Title = request.Title
	task.Content = request.Content
	task.StartTime = int64(request.StartTime)
	task.EndTime = int64(request.EndTime)

	return DB.Save(&task).Error
}

func (*Task) TaskShow(request *service.TaskRequest) ([]Task, error) {
	var taskList []Task
	err := DB.Model(&Task{}).Where("user_id=?", request.UserID).Find(&taskList).Error
	if err != nil {
		return nil, err
	}
	return taskList, err
}

func (*Task) TaskDelete(request *service.TaskRequest) error {
	return DB.Model(&Task{}).Where("task_id=?", request.TaskID).Delete(&Task{}).Error
}

func BuildTask(task Task) *service.TaskModel {
	taskModel := &service.TaskModel{
		TaskID:    uint32(task.TaskID),
		UserID:    uint32(task.UserID),
		Status:    uint32(task.Status),
		Title:     task.Title,
		Content:   task.Content,
		StartTime: uint32(task.StartTime),
		EndTime:   uint32(task.EndTime),
	}
	return taskModel
}

func BuildTaskList(tasks []Task) []*service.TaskModel {
	var taskList []*service.TaskModel
	for _, item := range tasks {
		task := BuildTask(item)
		taskList = append(taskList, task)
	}
	return taskList
}
