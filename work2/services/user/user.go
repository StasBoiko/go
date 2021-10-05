//go:generate mockgen -source user.go -destination mock/user_mock.go -package mock
package user

import (
	"context"
	"work2/models"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) (err error)
	FirstUser(ctx context.Context, user *models.User, email string, token string) (err error)
	GetUser(ctx context.Context, user *models.User, email string) (err error)
	SaveUser(ctx context.Context, user *models.User) (err error)
}

type UserService struct {
	db UserRepository
}

func UserNewService(db UserRepository) *UserService {
	return &UserService{db: db}
}

//CreateUserRecord creates a user record in the database
func (u *UserService) CreateUserRecord(ctx context.Context, user models.User) (models.User, error) {
	err := u.db.CreateUser(ctx, &user)
	if err != nil {
		return user, err
	}
	return user, err
}

func (u *UserService) SaveUserToken(ctx context.Context, user models.User) (models.User, error) {
	err := u.db.SaveUser(ctx, &user)
	if err != nil {
		return user, err
	}
	return user, nil
}

// HashPassword encrypts user password
func (u *UserService) HashPassword(ctx context.Context, user *models.User) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

// CheckPassword checks user password
func (u *UserService) CheckPassword(ctx context.Context, userPass string, providedPass string) error {
	err := bcrypt.CompareHashAndPassword([]byte(userPass), []byte(providedPass))
	if err != nil {
		return err
	}
	return nil
}

func (u *UserService) GetFirstUser(ctx context.Context, email string, token string, user models.User) (models.User, error) {
	err := u.db.FirstUser(ctx, &user, email, token)
	if err != nil {
		return user, err
	}
	return user, err
}

func (u *UserService) GetUser(ctx context.Context, email string, user models.User) (models.User, error) {
	err := u.db.GetUser(ctx, &user, email)
	if err != nil {
		return user, err
	}
	return user, err
}
