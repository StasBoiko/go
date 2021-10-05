// models/models.go

package models

import (
	//	"work2/repository/sqlitedb"

	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"gorm.io/gorm"
	//"github.com/jinzhu/gorm"
)

type Task struct {
	gorm.Model
	//ID          uint
	//UpdatedAt   time.Time
	//CreatedAt   time.Time
	//DeletedAt   sql.NullTime

	Name        string    `gorm:"column:name"`
	Event       string    `gorm:"column:event"`
	Date        time.Time `gorm:"column:date"`
	Description string    `gorm:"column:description"`
	UserId      uint      `gorm:"column:user_id"`
}

type Message struct {
	Message string
}

const (
	LayoutDay   = "2006-01-02"
	LayoutMonth = "2006-01"
	LayoutYear  = "2006"
)

type UserIdTaskId struct {
	UserID uint `json:"userId"`
	TaskID uint `json:"taskId"`
}

// User defines the user in db
type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	LoginResponse
}

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

var (
	OpsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	})
	TasksCreated = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_created_tasks_total",
		Help: "The total number of created tasks",
	})
	RequestCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "http_request_total_count", // metric name
		Help: "Total number of requests",
	})
	HttpReqs = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_user_requests",
			Help: "How many HTTP requests processed, partitioned by status code and HTTP method.",
		},
		[]string{"user"},
	)
	UsersCreated = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_user_created",
			Help: "How many HTTP requests processed, partitioned by status code and HTTP method.",
		},
		[]string{"user"},
	)
)
