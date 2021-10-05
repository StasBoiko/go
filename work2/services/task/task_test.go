package task

import (
	"context"
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"testing"
	"time"
	"work2/models"

	//	"work2/services"
	"work2/services/task/mock"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type MockRepo struct{}

func TestServiceGetTasks(t *testing.T) {
	type fields struct {
		db TaskRepository
	}
	type args struct {
		ctx        context.Context
		filterType string
		value      string
		userId     uint
	}
	var m *MockRepo
	var tasks []models.Task
	ctxVe := context.TODO()
	err := fmt.Errorf("error!")
	ctxEr := context.WithValue(context.TODO(), "error", err)
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantTasks []models.Task
		wantErr   bool
	}{
		{name: "error", fields: fields{m}, args: args{ctxEr, "", "", 1}, wantTasks: tasks, wantErr: true},
		{name: "ok", fields: fields{m}, args: args{ctxVe, "", "", 1}, wantTasks: tasks, wantErr: false},
		{name: "ok", fields: fields{m}, args: args{ctxVe, "day", "2005-05-07", 1}, wantTasks: tasks, wantErr: false},
		{name: "ok", fields: fields{m}, args: args{ctxVe, "month", "2005-05", 1}, wantTasks: tasks, wantErr: false},
		{name: "ok", fields: fields{m}, args: args{ctxVe, "year", "2005", 1}, wantTasks: tasks, wantErr: false},
		{name: "error", fields: fields{m}, args: args{ctxEr, "day", "2005-05-07", 1}, wantTasks: tasks, wantErr: true},
		{name: "error", fields: fields{m}, args: args{ctxEr, "month", "2005-05", 1}, wantTasks: tasks, wantErr: true},
		{name: "error", fields: fields{m}, args: args{ctxEr, "year", "2005", 1}, wantTasks: tasks, wantErr: true},
		//{name: "error", fields: fields{m}, args: args{task}, wantTasks: tasks, wantErr: false, c: ctxVe},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &TaskService{
				db: tt.fields.db,
			}
			//gotTasks, err := s.GetTasks(tt.args.ctx, tt.args.filterType, tt.args.value)
			gotTasks, err := s.GetTasks(tt.args.ctx, tt.args.filterType, tt.args.value, tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTasks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTasks, tt.wantTasks) {
				t.Errorf("GetTasks() gotTasks = %v, want %v", gotTasks, tt.wantTasks)
			}
		})
	}
}

func TestServiceCreateTask(t *testing.T) {
	var m *mock.MockTaskRepository
	err := fmt.Errorf("error!")
	ctxVe := context.TODO()
	ctxEr := context.WithValue(context.TODO(), "error", err)
	type mockBehavior func(r *mock.MockTaskRepository, ctx, task interface{})
	type fields struct {
		db TaskRepository
	}
	type args struct {
		task models.Task
	}
	var task models.Task
	var tasks models.Task
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantTasks    models.Task
		wantErr      bool
		c            context.Context
		mockBehavior func(r *mock.MockTaskRepository, task interface{})
	}{
		{
			name: "error", fields: fields{m}, args: args{task},
			wantTasks: tasks, wantErr: true, c: ctxEr,
			mockBehavior: func(r *mock.MockTaskRepository, task interface{}) {
				r.EXPECT().CreateOneTask(gomock.Any(), &tasks).Return(err)
			},
		},
		{
			name: "ok", fields: fields{m}, args: args{task},
			wantTasks: tasks, wantErr: false, c: ctxVe,
			mockBehavior: func(r *mock.MockTaskRepository, task interface{}) {
				r.EXPECT().CreateOneTask(gomock.Any(), &tasks)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			repo := mock.NewMockTaskRepository(c)
			tt.mockBehavior(repo, task)
			s := &TaskService{db: repo}
			gotTasks, err := s.CreateTask(ctxVe, tt.args.task)

			if tt.wantErr {
				require.NotNil(t, err)
				return
			}
			require.Nil(t, err)
			assert.EqualValues(t, gotTasks, tt.wantTasks)
		})
	}
}

func TestServiceGetFirstTask(t *testing.T) {
	type fields struct {
		db TaskRepository
	}
	type args struct {
		task *models.Task
	}
	var task *models.Task
	var tasks models.Task
	var m *MockRepo
	err := fmt.Errorf("error!")
	ctxVe := context.TODO()
	ctxEr := context.WithValue(context.TODO(), "error", err)

	tests := []struct {
		name      string
		fields    fields
		args      args
		wantTasks models.Task
		wantErr   bool
		c         context.Context
		id        string
	}{
		// {"test case 1", fields{pr}, args{task}, tasks, true},
		{name: "error", fields: fields{m}, args: args{task},
			wantTasks: tasks, wantErr: true, c: ctxEr, id: strconv.Itoa(rand.Intn(100))},
		{name: "ok", fields: fields{m}, args: args{task},
			wantTasks: tasks, wantErr: false, c: ctxVe, id: strconv.Itoa(rand.Intn(100))},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &TaskService{
				db: tt.fields.db,
			}
			err := s.FirstTask(tt.c, tt.id, tt.args.task)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetFirstTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(gotTasks, tt.wantTasks) {
			//	t.Errorf("Service.GetFirstTask() = %v, want %v", gotTasks, tt.wantTasks)
			//}
		})
	}
}

func TestServiceSaveTask(t *testing.T) {
	type fields struct {
		db TaskRepository
	}
	type args struct {
		task models.Task
	}
	var task models.Task
	var tasks models.Task
	var m *MockRepo
	err := fmt.Errorf("error!")
	ctxVe := context.TODO()
	ctxEr := context.WithValue(context.TODO(), "error", err)

	tests := []struct {
		name      string
		fields    fields
		args      args
		wantTasks models.Task
		wantErr   bool
		c         context.Context
	}{
		// {"test case 1", fields{pr}, args{task}, tasks, true},
		{name: "error", fields: fields{m}, args: args{task},
			wantTasks: tasks, wantErr: true, c: ctxEr},
		{name: "ok", fields: fields{m}, args: args{task},
			wantTasks: tasks, wantErr: false, c: ctxVe},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &TaskService{
				db: tt.fields.db,
			}
			gotTasks, err := s.SaveTask(tt.c, tt.args.task)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.SaveTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTasks, tt.wantTasks) {
				t.Errorf("Service.SaveTask() = %v, want %v", gotTasks, tt.wantTasks)
			}
		})
	}
}

func TestServiceDeleteTask(t *testing.T) {
	type fields struct {
		db TaskRepository
	}
	type args struct {
		task models.Task
	}
	var task models.Task
	var tasks []models.Task
	var m *MockRepo
	err := fmt.Errorf("error!")
	ctxVe := context.TODO()
	ctxEr := context.WithValue(context.TODO(), "error", err)

	tests := []struct {
		name      string
		fields    fields
		args      args
		wantTasks []models.Task
		wantErr   bool
		c         context.Context
		id        string
	}{
		// {"test case 1", fields{pr}, args{task}, tasks, true},
		{name: "error", fields: fields{m}, args: args{task},
			wantTasks: tasks, wantErr: true, c: ctxEr, id: strconv.Itoa(rand.Intn(100))},
		{name: "ok", fields: fields{m}, args: args{task},
			wantTasks: tasks, wantErr: false, c: ctxVe, id: strconv.Itoa(rand.Intn(100))},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &TaskService{
				db: tt.fields.db,
			}
			gotTasks, err := s.DeleteTask(tt.c, tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.DeleteTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTasks, tt.wantTasks) {
				t.Errorf("Service.DeleteTask() = %v, want %v", gotTasks, tt.wantTasks)
			}
		})
	}
}

func (mock *MockRepo) GetDayTasks(ctx context.Context, t time.Time, userId uint) (tasks []models.Task, err error) {
	err, _ = ctx.Value("error").(error)
	return nil, err
}

func (mock *MockRepo) GetMonthTasks(ctx context.Context, t time.Time, userId uint) (tasks []models.Task, err error) {
	err, _ = ctx.Value("error").(error)
	return nil, err
}

func (mock *MockRepo) GetYearTasks(ctx context.Context, t time.Time, userId uint) (tasks []models.Task, err error) {
	err, _ = ctx.Value("error").(error)
	return nil, err
}

func (mock *MockRepo) GetAllTasks(ctx context.Context, tasks *[]models.Task) (err error) {
	err, _ = ctx.Value("error").(error)
	return err
}

func (mock *MockRepo) GetAllUserTasks(ctx context.Context, tasks *[]models.Task, userId uint) (err error) {
	err, _ = ctx.Value("error").(error)
	return err
}

func (mock *MockRepo) CreateOneTask(ctx context.Context, task *models.Task) (err error) {
	err, _ = ctx.Value("error").(error)
	return err
}
func (mock *MockRepo) FirstTask(ctx context.Context, task *models.Task, id string) (err error) {
	err, _ = ctx.Value("error").(error)
	return err
}
func (mock *MockRepo) SaveTask(ctx context.Context, task *models.Task) (err error) {
	err, _ = ctx.Value("error").(error)
	return err
}
func (mock *MockRepo) DeleteOneTask(ctx context.Context, task *models.Task) (err error) {
	err, _ = ctx.Value("error").(error)
	return err
}

// go:generate mockgen -destination=repository_mock.go -package=events -build_flags=-mod=mod work2/services Repository
