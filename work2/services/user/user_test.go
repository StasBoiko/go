package user

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"work2/config"
	"work2/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	envconfig "github.com/sethvargo/go-envconfig"
	"github.com/stretchr/testify/assert"
)

type UserMockRepo struct{}

func TestHashPassword(t *testing.T) {
	user := models.User{
		Password: "secret",
	}
	var um *UserMockRepo
	s := UserNewService(um)
	err := s.HashPassword(context.TODO(), &user)
	assert.NoError(t, err)
	os.Setenv("passwordHash", user.Password)
}

func TestCreateUserRecord(t *testing.T) {
	var userResult models.User
	var mc config.MyConfig
	ctx := context.Background()
	if err := envconfig.Process(ctx, &mc.Ps); err != nil {
		log.Fatal(err)
	}
	ps := mc.Ps
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s ",
		ps.Host, ps.Port, ps.User, ps.Dbname, ps.Password)
	db, err := gorm.Open("postgres", psqlconn)

	assert.NoError(t, err)
	user := models.User{
		Name:     "Test User",
		Email:    "test@email.com",
		Password: os.Getenv("passwordHash"),
	}
	var um *UserMockRepo
	s := UserNewService(um)
	user, err = s.CreateUserRecord(context.TODO(), user)
	assert.NoError(t, err)
	fmt.Println(err)
	db.Where("email = ?", user.Email).Find(&userResult)
	db.Unscoped().Delete(&user)
	assert.Equal(t, "Test User", userResult.Name)
	assert.Equal(t, "test@email.com", userResult.Email)
}

func TestCheckPassword(t *testing.T) {
	var um *UserMockRepo
	s := UserNewService(um)
	err := s.CheckPassword(context.TODO(), "$2a$14$BuN8OevOoGTYi/UQUc2kj.q0Tqg0oA6p2Vqlh.OI6KoZXwvNuv9g6", "secret888")
	assert.NoError(t, err)
}

func (p *UserMockRepo) CreateUser(ctx context.Context, user *models.User) (err error) {
	var mc config.MyConfig
	ctx = context.Background()
	if err := envconfig.Process(ctx, &mc.Ps); err != nil {
		log.Fatal(err)
	}
	ps := mc.Ps
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s ",
		ps.Host, ps.Port, ps.User, ps.Dbname, ps.Password)
	db, err := gorm.Open("postgres", psqlconn)

	result := db.Create(&user)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *UserMockRepo) SaveUser(ctx context.Context, user *models.User) (err error) {
	return err
}

func (u *UserMockRepo) FirstUser(ctx context.Context, user *models.User, email string, token string) (err error) {
	return err
}

func (u *UserMockRepo) GetUser(ctx context.Context, user *models.User, email string) (err error) {
	return err
}
