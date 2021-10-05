//go:generate mockgen -source repository.go -destination mock/repository_mock.go -package mock
//go:generate mockgen -destination=repository_mock.go -package=events -build_flags=-mod=mod work2/services Repository
package postgres

import (
	"context"
	"fmt"
	"time"
	"work2/config"
	"work2/models"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/jinzhu/gorm"
)

type PostgresRepo struct {
	db *gorm.DB
}

func NewRepo(ps config.PsConfig) (p *PostgresRepo, err error) {
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s ",
		ps.Host, ps.Port, ps.User, ps.Dbname, ps.Password)
	db, err := gorm.Open("postgres", psqlconn)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.Task{})
	db.AutoMigrate(&models.User{})

	return &PostgresRepo{db}, nil
}

func (p *PostgresRepo) GetDayTasks(ctx context.Context, t time.Time, userId uint) (tasks []models.Task, err error) {
	err = p.db.Where("user_id = ? AND date BETWEEN ? AND ?", userId, t, t.Add(time.Hour*24)).Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (p *PostgresRepo) GetMonthTasks(ctx context.Context, t time.Time, userId uint) (tasks []models.Task, err error) {
	tl := time.Date(t.Year(), t.Month()+1, 1, 0, 0, 0, 0, time.UTC)
	err = p.db.Where("user_id = ? AND date BETWEEN ? AND ?", userId, t, tl).Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (p *PostgresRepo) GetYearTasks(ctx context.Context, t time.Time, userId uint) (tasks []models.Task, err error) {
	tl := time.Date(t.Year()+1, 1, 1, 0, 0, 0, 0, time.UTC)
	err = p.db.Where("user_id = ? AND date BETWEEN ? AND ?", userId, t, tl).Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (p *PostgresRepo) GetAllUserTasks(ctx context.Context, tasks *[]models.Task, userId uint) (err error) {
	err = p.db.Where("user_id = ?", userId).Find(&tasks).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *PostgresRepo) GetAllTasks(ctx context.Context, tasks *[]models.Task) (err error) {
	err = p.db.Find(&tasks).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *PostgresRepo) CreateOneTask(ctx context.Context, task *models.Task) (err error) {
	err = p.db.Create(&task).Error
	if err != nil {
		return err
	}
	models.TasksCreated.Add(1)
	return nil
}

func (p *PostgresRepo) FirstTask(ctx context.Context, task *models.Task, id string) (err error) {
	err = p.db.First(&task, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *PostgresRepo) SaveTask(ctx context.Context, task *models.Task) (err error) {
	err = p.db.Save(&task).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *PostgresRepo) DeleteOneTask(ctx context.Context, task *models.Task) (err error) {
	err = p.db.Delete(&task).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *PostgresRepo) CreateUser(ctx context.Context, user *models.User) (err error) {
	result := p.db.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	models.UsersCreated.WithLabelValues(user.Name).Add(2)
	//models.UsersCreated.Add(1)
	return nil
}

func (p *PostgresRepo) SaveUser(ctx context.Context, user *models.User) (err error) {
	err = p.db.Save(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *PostgresRepo) FirstUser(ctx context.Context, user *models.User, email string, token string) (err error) {
	result := p.db.Where("email = ? AND token = ?", email, token).First(&user)
	//возможно стоит использовать Find вместо First?
	//result := p.db.Where("email = ? AND token = ?", email, token).Find(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (p *PostgresRepo) GetUser(ctx context.Context, user *models.User, email string) (err error) {
	result := p.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
