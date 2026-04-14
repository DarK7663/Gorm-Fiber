package main

import (
	"time"

	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) GetTaskById(id uint) (*Task, error) {
	var task Task
	if err := r.db.Where("id = ?", id).First(&task).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *TaskRepository) GetAllTasks() ([]Task, error) {
	var tasks []Task
	if err := r.db.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *TaskRepository) CreateTask(title string, t time.Time, date time.Time) (*Task, error) {
	t, err := time.Parse(time.RFC3339, "2024-03-28")
	if err != nil {

	}
	task := &Task{Title: title, Time: t.UTC(), Date: date}
	if err := r.db.Create(task).Error; err != nil {
		return nil, err
	}
	return task, nil
}

func (r *TaskRepository) UpdateTask(id uint, title string) error {
	if _, err := r.GetTaskById(id); err != nil {
		return err
	}
	if err := r.db.Model(&Task{}).Where("id = ?", id).Update("title", title).Error; err != nil {
		return err
	}
	return nil
}

func (r *TaskRepository) DeleteTask(id uint) error {
	if _, err := r.GetTaskById(id); err != nil {
		return err
	}
	if err := r.db.Delete(&Task{}, id).Error; err != nil {
		return err
	}
	return nil
}
