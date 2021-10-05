package main

import (
	// "gorm.io/driver/postgres"
	// "gorm.io/gorm"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/rs/cors"
)

const (
	layoutDay   = "2006-01-03"
	layoutMonth = "2006-01"
	layoutYear  = "2006"
)

//TODO: нужно ли реализовывать RequestHandler, Logger interface и Sevice, который реализует Logger?
type RequestHandler interface {
	GetTasks()
	CreateTask()
	UpdateTask()
	DeleteTask()
	GetDayTasks()
	GetMonthTasks()
	GetYearTasks()
}
type Logger interface {
	Println(v ...interface{})
	Fatalf(format string, v ...interface{})
}
type Service struct {
	log Logger
}

func NewService(log Logger) *Service {
	return &Service{log: log}
}

type Task struct {
	gorm.Model

	Name        string    `gorm:"column:name"`
	Event       string    `gorm:"column:event"`
	Date        time.Time `gorm:"column:date"`
	Description string    `gorm:"column:description"`
}

var db *gorm.DB

var err error

var t = time.Now()

var (
	tasks = []Task{
		{Name: "some task1", Event: "event", Date: time.Now(), Description: "some description1"},
		{Name: "some task2", Event: "event", Date: t.AddDate(0, 0, -1), Description: "some description2"},
		{Name: "some task3", Event: "event", Date: t.AddDate(0, 0, -2), Description: "some description3"},
	}
)

type PsConfig struct {
	Dbname   string `env:"POSTGRES_DBNAME"`
	Host     string `env:"POSTGRES_HOST"`
	Port     string `env:"POSTGRES_PORT"`
	User     string `env:"POSTGRES_USER"`
	Password string `env:"POSTGRES_PASSWORD"`
}

func main() {

	// ctx := context.Background()

	// var c MyConfig
	// if err := envconfig.Process(ctx, &c); err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("!!!")
	// fmt.Println(c.Port)

	router := mux.NewRouter()

	// const (
	// 	host     = "localhost"
	// 	port     = 5432
	// 	user     = "sboik4"
	// 	password = ""
	// 	dbname   = "postgres"
	// )

	// dsn := "host=localhost user=sboik4 password= dbname=postgres port=5432 sslmode=disable"
	// db, err := gorm.Open(postgres.Open(dsn))

	// fmt.Println("postgres", "host=localhost port=5432 user=sboik4 dbname=postgres sslmode=disable password=")
	// psqlconn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable password=%s ", host, port, user, dbname, password)
	// fmt.Println(psqlconn)
	// fmt.Println("postgres", psqlconn)
	// db, err := gorm.Open("postgres", psqlconn)

	// var c PsConfig
	// ctx := context.Background()

	// if err := envconfig.Process(ctx, &c); err != nil {
	// 	log.Fatal(err)
	// }
	// psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", c.Host, c.Port, c.User, c.Password, c.Dbname)
	// db, err := gorm.Open("postgres", psqlconn)
	// fmt.Println("postgres", psqlconn)
	// fmt.Println("postgres", "host=localhost port=5432 user=sboik4 password= dbname=postgres sslmode=disable")
	// if psqlconn != "host=localhost port=5432 user=sboik4 password= dbname=postgres sslmode=disable" {
	// 	fmt.Println("diff!!!")
	// }
	// fmt.Println(strings.Compare(psqlconn, "host=localhost port=5432 user=sboik4 password= dbname=postgres sslmode=disable"))

	db, err = gorm.Open("postgres", "host=localhost port=5432 user=sboik4 dbname=postgres sslmode=disable password=")
	CheckError(err)

	defer db.Close()

	db.AutoMigrate(&Task{})

	for index := range tasks {

		db.Create(&tasks[index])

	}

	// database.Init()

	router.HandleFunc("/tasks", GetTasks).Methods("GET")

	router.HandleFunc("/tasks", CreateTask).Methods("POST")

	router.HandleFunc("/tasks/{id}", UpdateTask).Methods("PUT")

	router.HandleFunc("/tasks/{id}", DeleteTask).Methods("DELETE")

	// окончательный вариант:
	// http://localhost:3000/tasks?filterType=day&value=2012-01-01
	// http://localhost:3000/tasks?filterType=month&value=2012-12
	// http://localhost:3000/tasks?filterType=year&value=2012

	handler := cors.Default().Handler(router)

	//TODO: нужно ли?
	// logger := log.New(os.Stdout, "INFO: ", 0)
	// var rH RequestHandler
	// rH = NewService(logger)

	log.Fatal(http.ListenAndServe(":3010", handler))

}

func GetTasks(w http.ResponseWriter, r *http.Request) {

	if r.FormValue("filterType") != "" && r.FormValue("value") != "" {

		filterType := r.FormValue("filterType")

		value := r.FormValue("value")

		switch filterType {
		case "day":
			GetDayTasks(w, value)
		case "month":
			GetMonthTasks(w, value)
		case "year":
			GetYearTasks(w, value)
		default:
			fmt.Println("NOT Valid filterType !!!")
		}

	} else {

		fmt.Println(1)
		var tasks []Task
		fmt.Println(tasks)
		db.Find(&tasks)
		fmt.Println(2)
		json.NewEncoder(w).Encode(&tasks)

	}

}

func CreateTask(w http.ResponseWriter, r *http.Request) {

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {

		fmt.Fprintf(w, "NOT Valid Task !!!")

	}

	var task Task

	json.Unmarshal(reqBody, &task)

	db.Create(&task)

	var tasks []Task

	db.Find(&tasks)

	json.NewEncoder(w).Encode(&tasks)

}

func UpdateTask(w http.ResponseWriter, r *http.Request) {

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {

		fmt.Fprintf(w, "NOT Valid Task !!!")

	}

	params := mux.Vars(r)

	var task Task

	db.First(&task, params["id"])

	json.Unmarshal(reqBody, &task)

	db.Save(&task)

	var tasks []Task

	db.Find(&tasks)

	json.NewEncoder(w).Encode(&tasks)

}

func DeleteTask(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	var task Task

	db.First(&task, params["id"])

	db.Delete(&task)

	var tasks []Task

	db.Find(&tasks)

	json.NewEncoder(w).Encode(&tasks)

}

func GetDayTasks(w http.ResponseWriter, v string) {

	t, err := time.Parse(layoutDay, v)

	CheckError(err)

	var tasks []Task

	db.Where("date BETWEEN ? AND ?", t, t.Add(time.Hour*24)).Find(&tasks)

	json.NewEncoder(w).Encode(&tasks)

}

func GetMonthTasks(w http.ResponseWriter, v string) {

	t, err := time.Parse(layoutMonth, v)

	CheckError(err)

	var tasks []Task

	db.Where("date BETWEEN ? AND ?", t, t.Add(time.Hour*24*30)).Find(&tasks)

	json.NewEncoder(w).Encode(&tasks)

}

func GetYearTasks(w http.ResponseWriter, v string) {

	t, err := time.Parse(layoutYear, v)

	CheckError(err)

	var tasks []Task

	db.Where("date BETWEEN ? AND ?", t, t.Add(time.Hour*24*365)).Find(&tasks)

	json.NewEncoder(w).Encode(&tasks)

}

func CheckError(err error) {

	if err != nil {

		log.Fatal(err)

	}

}
