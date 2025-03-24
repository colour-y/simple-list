package dao

import (
	"context"
	"simplelist/types"

	"gorm.io/gorm"
	"simplelist/res/db/model"
)

type TaskDao struct {
	*gorm.DB
}

func NewTaskdao(ctx context.Context) *TaskDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &TaskDao{NewDBClient(ctx)}
}

func (s *TaskDao) CreateTask(task *model.Task) error {
	return s.Model(&model.Task{}).Create(&task).Error
}

func (s *TaskDao) ListTask(start, limit int, userId uint) (r []*model.Task, total int64, err error) {
	err = s.Model(&model.Task{}).Preload("User").Where("uid = ?", userId).
		Count(&total).
		Limit(limit).Offset((start - 1) * limit).
		Find(&r).Error

	return
}
func (s *TaskDao) FindTaskByIdAndUserId(uId, id uint) (r *model.Task, err error) {
	err = s.Model(&model.Task{}).Where("id = ? AND uid = ?", id, uId).First(&r).Error

	return
}

func (s *TaskDao) UpdataTask(uid uint, req *types.UpdateTaskReq) error {
	t := new(model.Task)
	err := s.Model(&model.Task{}).Where("id = ? AND uid=?", req.ID, uid).First(&t).Error
	if err != nil {
		return nil
	}

	if req.Status != 0 {
		t.Status = req.Status
	}

	if req.Title != "" {
		t.Title = req.Title
	}

	if req.Content != "" {
		t.Content = req.Content
	}
	return s.Save(t).Error
}

func (s *TaskDao) SearchTask(uid uint, info string) (tasks []*model.Task, err error) {
	err = s.Where("uid=?", uid).Preload("User").First(&tasks).Error
	if err != nil {
		return
	}
	err = s.Model(&model.Task{}).Where("title LIKE ? OR content LIKE ?", "%"+info+"%", "%"+info+"%").Find(&tasks).Error

	return
}

func (s *TaskDao) DeleteTaskById(uid, tId uint) error {
	r, err := s.FindTaskByIdAndUserId(uid, tId)
	if err != nil {
		return err
	}
	return s.Delete(&r).Error
}
