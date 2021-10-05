package handlers

import (
	"context"
	"encoding/json"

	"io/ioutil"

	"github.com/golang/mock/gomock"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"time"

	"net/http"
	"net/http/httptest"
	"testing"
	"work2/middlewares"
	"work2/models"

	"github.com/stretchr/testify/assert"
)

type MockRepo struct{}
type UserMockRepo struct{}

var timeNow = time.Now()

func TestGetTasks(t *testing.T) {
	var m *MockRepo
	var um *UserMockRepo
	var tasks []models.Task

	tasks = append(tasks, models.Task{
		//ID:          123,
		//CreatedAt:   time.Date(timeNow.Year(), timeNow.Month(), 1, 0, 0, 0, 0, time.UTC),
		//UpdatedAt:   time.Date(timeNow.Year(), timeNow.Month(), 1, 0, 0, 0, 0, time.UTC),
		//DeletedAt:   sql.NullTime{},
		Name:        "name",
		Event:       "event",
		Date:        time.Date(timeNow.Year(), timeNow.Month(), 1, 0, 0, 0, 0, time.UTC),
		Description: "description",
	})
	var result []models.Task

	req := httptest.NewRequest(http.MethodGet, "/tasks?filterType=day&value=2005-05-07", nil)

	w := httptest.NewRecorder()
	c := gomock.NewController(t)
	defer c.Finish()
	s := &Server{ts: m, us: um}
	s.GetTasks(w, req)

	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("error: %v", err)
		return
	}
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Errorf("error: %v", err)
		return
	}
	assert.Equal(t, w.Code, 200)
	assert.Equal(t, result, tasks)
}

func NewContextWithRequestID(ctx context.Context, r *http.Request) context.Context {
	return context.WithValue(ctx, "reqId", "1234")
}

func AddContextWithRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var ctx = context.Background()
		ctx = NewContextWithRequestID(ctx, r)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func TestAuthz(t *testing.T) {
	// create a handler to use as "next" which will verify the request
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		val := r.Context().Value("reqId")
		if val == nil {
			t.Error("reqId not present")
		}
		valStr, ok := val.(string)
		if !ok {
			t.Error("not string")
		}
		if valStr != "1234" {
			t.Error("wrong reqId")
		}
	})
	// create the handler to test, using our custom "next" handler
	handlerToTest := AddContextWithRequestID(nextHandler)
	// create a mock request to use
	req := httptest.NewRequest("GET", "http://localhost:2000/tasks", nil)
	// call the handler using a mock response recorder (we'll not use that anyway)
	handlerToTest.ServeHTTP(httptest.NewRecorder(), req)
}

func TestAuthzNoHeader(t *testing.T) {
	var m *MockRepo
	var um *UserMockRepo

	req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
	w := httptest.NewRecorder()
	s := &Server{ts: m, us: um}
	//uSer := user.UserNewService(um)
	//m := middlewares.NewMiddleware(uSer, logz)
	middlewares.Authz(http.HandlerFunc(s.GetTasks))
	s.GetTasks(w, req)
	assert.Equal(t, 403, w.Code)
}

func TestAuthzInvalidTokenFormat(t *testing.T) {
	var m TaskService
	var um UserService

	req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
	w := httptest.NewRecorder()
	s := &Server{ts: m, us: um}
	req.Header.Set("Authorization", "test")
	middlewares.Authz(http.HandlerFunc(s.GetTasks))
	s.GetTasks(w, req)
	assert.Equal(t, 403, w.Code)
}

func TestAuthzInvalidToken(t *testing.T) {
	invalidToken := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
	var m TaskService
	var um UserService
	req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
	w := httptest.NewRecorder()
	s := &Server{ts: m, us: um}
	req.Header.Set("Authorization", invalidToken)
	middlewares.Authz(http.HandlerFunc(s.GetTasks))
	s.GetTasks(w, req)
	assert.Equal(t, 403, w.Code)
}

func (mock *MockRepo) GetTasks(ctx context.Context, f string, v string, userId uint) (tasks []models.Task, err error) {
	tasks = append(tasks, models.Task{
		//ID:          123,
		//CreatedAt:   time.Date(timeNow.Year(), timeNow.Month(), 1, 0, 0, 0, 0, time.UTC),
		//UpdatedAt:   time.Date(timeNow.Year(), timeNow.Month(), 1, 0, 0, 0, 0, time.UTC),
		//DeletedAt:   sql.NullTime{},
		Name:        "name",
		Event:       "event",
		Date:        time.Date(timeNow.Year(), timeNow.Month(), 1, 0, 0, 0, 0, time.UTC),
		Description: "description",
	})
	return tasks, err
}
func (mock *MockRepo) CreateTask(ctx context.Context, task models.Task) (tasks models.Task, err error) {
	return tasks, err
}
func (mock *MockRepo) SaveTask(ctx context.Context, task models.Task) (tasks models.Task, err error) {
	return tasks, err
}
func (mock *MockRepo) FirstTask(ctx context.Context, id string, task *models.Task) (err error) {
	err, _ = ctx.Value("error").(error)
	return err
}
func (mock *MockRepo) DeleteTask(ctx context.Context, id string) (tasks []models.Task, err error) {
	return tasks, err
}

func (u *UserMockRepo) CreateUserRecord(ctx context.Context, users models.User) (user models.User, err error) {
	return user, err
}

func (u *UserMockRepo) SaveUserToken(ctx context.Context, users models.User) (user models.User, err error) {
	return user, err
}

// HashPassword encrypts user password
func (u *UserMockRepo) HashPassword(ctx context.Context, user *models.User) error {
	err, _ := ctx.Value("error").(error)
	return err
}

// CheckPassword checks user password
func (u *UserMockRepo) CheckPassword(ctx context.Context, userPass string, providedPass string) error {
	err, _ := ctx.Value("error").(error)
	return err
}

func (u *UserMockRepo) GetFirstUser(ctx context.Context, email string, token string, users models.User) (user models.User, err error) {
	return user, err
}

func (u *UserMockRepo) GetUser(ctx context.Context, email string, users models.User) (user models.User, err error) {
	return user, err
}
