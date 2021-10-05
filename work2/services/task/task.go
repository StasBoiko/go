//go:generate mockgen -source task.go -destination mock/task_mock.go -package mock
package task

import (
	"context"
	"fmt"
	"time"
	"work2/models"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type TaskRepository interface {
	GetDayTasks(ctx context.Context, t time.Time, userId uint) (tasks []models.Task, err error)
	GetMonthTasks(ctx context.Context, t time.Time, userId uint) (tasks []models.Task, err error)
	GetYearTasks(ctx context.Context, t time.Time, userId uint) (tasks []models.Task, err error)
	GetAllTasks(ctx context.Context, tasks *[]models.Task) (err error)
	GetAllUserTasks(ctx context.Context, tasks *[]models.Task, userId uint) (err error)
	CreateOneTask(ctx context.Context, task *models.Task) (err error)
	SaveTask(ctx context.Context, task *models.Task) (err error)
	FirstTask(ctx context.Context, task *models.Task, id string) (err error)
	DeleteOneTask(ctx context.Context, task *models.Task) (err error)
}

type TaskService struct {
	db TaskRepository
}

func TaskNewService(db TaskRepository) *TaskService {
	return &TaskService{db: db}
}

func (s *TaskService) GetTasks(ctx context.Context, filterType string, value string, userId uint) (tasks []models.Task, err error) {
	if filterType != "" && value != "" {
		switch filterType {
		case "day":
			t, err := time.Parse(models.LayoutDay, value)
			if err != nil {
				return nil, err
			}
			tasks, err = s.db.GetDayTasks(ctx, t, userId)
			if err != nil {
				return nil, err
			}
			return tasks, nil
		case "month":
			t, err := time.Parse(models.LayoutMonth, value)
			if err != nil {
				return nil, err
			}
			tasks, err = s.db.GetMonthTasks(ctx, t, userId)
			if err != nil {
				return nil, err
			}
			return tasks, nil
		case "year":
			t, err := time.Parse(models.LayoutYear, value)
			if err != nil {
				return nil, err
			}
			tasks, err = s.db.GetYearTasks(ctx, t, userId)
			if err != nil {
				return nil, err
			}
			return tasks, nil
		default:
			err = fmt.Errorf("not valid filterType")
			return nil, err
		}
	} else {
		//fmt.Println(1111)
		var tasks []models.Task
		//err = s.db.GetAllTasks(ctx, &tasks)
		err = s.db.GetAllUserTasks(ctx, &tasks, userId)
		if err != nil {
			return nil, err
		}
		return tasks, nil
	}
}

func (s *TaskService) CreateTask(ctx context.Context, task models.Task) (models.Task, error) {
	err := s.db.CreateOneTask(ctx, &task)
	if err != nil {
		return task, err
	}
	return task, err
}
func (s *TaskService) FirstTask(ctx context.Context, id string, task *models.Task) (err error) {
	err = s.db.FirstTask(ctx, task, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *TaskService) SaveTask(ctx context.Context, task models.Task) (models.Task, error) {
	err := s.db.SaveTask(ctx, &task)
	if err != nil {
		return task, err
	}
	return task, nil
}

func (s *TaskService) DeleteTask(ctx context.Context, id string) (tasks []models.Task, err error) {
	var task models.Task
	err = s.db.FirstTask(ctx, &task, id)
	if err != nil {
		return nil, err
	}
	err = s.db.DeleteOneTask(ctx, &task)
	if err != nil {
		return nil, err
	}
	err = s.db.GetAllTasks(ctx, &tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
