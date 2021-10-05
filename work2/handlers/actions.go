//go:generate mockgen -source actions.go -destination mock/actions_mock.go -package mock
package handlers

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"work2/auth"
	"work2/middlewares"
	"work2/models"

	"go.uber.org/zap"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/gorilla/mux"
)

type TaskService interface {
	GetTasks(ctx context.Context, f string, v string, userId uint) (tasks []models.Task, err error)
	CreateTask(ctx context.Context, task models.Task) (tasks models.Task, err error)
	SaveTask(ctx context.Context, task models.Task) (tasks models.Task, err error)
	FirstTask(ctx context.Context, id string, task *models.Task) error
	DeleteTask(ctx context.Context, id string) (tasks []models.Task, err error)
}

type UserService interface {
	CreateUserRecord(ctx context.Context, user models.User) (models.User, error)
	SaveUserToken(ctx context.Context, user models.User) (models.User, error)
	HashPassword(ctx context.Context, user *models.User) error
	CheckPassword(ctx context.Context, userPass string, providedPass string) error
	GetFirstUser(ctx context.Context, email string, token string, user models.User) (models.User, error)
	GetUser(ctx context.Context, email string, user models.User) (models.User, error)
}

type Server struct {
	ts  TaskService
	us  UserService
	log *zap.Logger
}

func NewServer(ts TaskService, us UserService, log *zap.Logger) *Server {
	return &Server{ts: ts, us: us, log: log}
}

// Req: http://localhost:3000/tasks?filterType=day&value=2012-01-01
func (s *Server) GetTasks(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user_id")
	userIdUint, ok := userId.(uint)
	if !ok {
		w.WriteHeader(http.StatusForbidden)
		s.log.Error("user Id Str", zap.String("key1", "value1"))
		return
	}
	tasks, err := s.ts.GetTasks(r.Context(), r.FormValue("filterType"), r.FormValue("value"), userIdUint)
	if err != nil {
		s.error(w, err)
		return
	}
	s.response(w, &tasks)
}

func (s *Server) CreateTask(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user_id")
	userIdUint, ok := userId.(uint)
	if !ok {
		w.WriteHeader(http.StatusForbidden)
		s.log.Error("user Id Str", zap.String("key1", "value1"))
		return
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		s.error(w, err)
		return
	}
	var task models.Task
	//если мы хотим сохранить нормальную дату,
	//то нужно использовать строку формата: 2005-05-07T21:16:32.082049+03:00
	//либо в этом месте написать код, который будет преобразовывать строку в time:
	//2005-05-07 -> 2005-05-07T21:16:32.082049+03:00
	//сейчас просто в запросе: 2005-05-07T21:16:32.082049+03:00
	err = json.Unmarshal(reqBody, &task)
	if err != nil {
		s.error(w, err)
		return
	}
	task.UserId = userIdUint
	tasks, err := s.ts.CreateTask(r.Context(), task)
	if err != nil {
		s.error(w, err)
		return
	}
	s.response(w, &tasks)
}

func (s *Server) UpdateTask(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		s.error(w, err)
		return
	}
	params := mux.Vars(r)
	var task models.Task
	err = s.ts.FirstTask(r.Context(), params["id"], &task)
	if err != nil {
		s.error(w, err)
		return
	}
	err = json.Unmarshal(reqBody, &task)
	if err != nil {
		s.error(w, err)
		return
	}
	task, err = s.ts.SaveTask(r.Context(), task)
	if err != nil {
		s.error(w, err)
		return
	}
	s.response(w, &task)
}

func (s *Server) DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tasks, err := s.ts.DeleteTask(r.Context(), params["id"])
	if err != nil {
		s.error(w, err)
		return
	}
	s.response(w, &tasks)
}

func (s *Server) Signup(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		s.error(w, err)
		return
	}
	var user models.User
	err = json.Unmarshal(reqBody, &user)
	if err != nil {
		s.error(w, err)
		return
	}
	err = s.us.HashPassword(r.Context(), &user)
	if err != nil {
		s.error(w, err)
		return
	}
	user, err = s.us.CreateUserRecord(r.Context(), user)
	if err != nil {
		s.error(w, err)
		return
	}
	s.response(w, &user)
}

func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		s.error(w, err)
		return
	}
	var payload models.User
	err = json.Unmarshal(reqBody, &payload)

	var user models.User
	user, err = s.us.GetUser(r.Context(), payload.Email, user)
	if err != nil {
		s.error(w, err)
		return
	}
	err = s.us.CheckPassword(r.Context(), user.Password, payload.Password)
	if err != nil {
		s.error(w, err)
		return
	}
	jwtWrapper := auth.JwtWrapper{
		SecretKey:       "verysecretkey",
		Issuer:          "AuthService",
		ExpirationHours: 24,
	}
	signedToken, err := jwtWrapper.GenerateToken(user.Email)
	if err != nil {
		s.error(w, err)
		return
	}
	tokenResponse := models.LoginResponse{
		Token: signedToken,
	}
	user.LoginResponse = tokenResponse
	user, err = s.us.SaveUserToken(r.Context(), user)
	if err != nil {
		s.error(w, err)
		return
	}
	s.response(w, &tokenResponse)
}

func (s *Server) Logout(w http.ResponseWriter, r *http.Request) {
	email := r.Context().Value("email_and_token").(middlewares.Values).Get("email")
	var user models.User
	user, err := s.us.GetUser(r.Context(), email, user)
	if err != nil {
		s.error(w, err)
		return
	}
	//var user models.User
	tokenResponse := models.LoginResponse{
		Token: "",
	}
	user.LoginResponse = tokenResponse
	user, err = s.us.SaveUserToken(r.Context(), user)
	if err != nil {
		s.error(w, err)
		return
	}
	w.Write([]byte("Logged out!\n"))
}

func (s *Server) error(w http.ResponseWriter, err error) {
	log.Print(err)
	//s.log.Error("user Id Str", zap.String("key1", "value1"))
	s.log.Error("error: ", zap.Error(err))
	s.log.Sugar().Errorw("error", err)
	s.log.Sugar().Errorf("%+v", err)
	w.WriteHeader(http.StatusBadRequest)
}

func (s *Server) response(w http.ResponseWriter, tasks interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&tasks)
}
