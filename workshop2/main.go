package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// 	"strconv"

// 	"github.com/gorilla/mux"
// )

//TODO: - postgres!
// type task struct {
// 	ID      int    `json:"ID"`
// 	Name    string `json:"Name"`
// 	Type    string `json:"Type"`
// 	Date    string `json:"Date"`
// 	Content string `json:"Content"`
// }

// type allTasks []task

// var tasks = allTasks{
// 	{
// 		ID:      1,
// 		Name:    "Task One",
// 		Type:    "Event",
// 		Date:    "2012-01-01",
// 		Content: "Some Content1",
// 	},
// 	{
// 		ID:      2,
// 		Name:    "Task Two",
// 		Type:    "Reminder",
// 		Date:    "2012-01-02",
// 		Content: "Some Content2",
// 	},
// }

// func indexRoute(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello!")
// }

// func createTask(w http.ResponseWriter, r *http.Request) {
// 	var newTask task
// 	reqBody, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		fmt.Fprintf(w, "Insert a Valid Task Data")
// 	}

// 	json.Unmarshal(reqBody, &newTask)
// 	newTask.ID = len(tasks) + 1
// 	tasks = append(tasks, newTask)

// 	// w.Header().Set("Content-Type", "application/json")
// 	// w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(newTask)

// }

// func getTasks(w http.ResponseWriter, r *http.Request) {
// 	// fmt.Println(mux.Vars(r))
// 	// fmt.Println(r.FormValue("period"))
// 	// fmt.Println(r.FormValue("filterType"))
// 	// fmt.Println(r.FormValue("value"))
// 	if r.FormValue("filterType") != "" && r.FormValue("value") != "" {
// 		filterType := r.FormValue("filterType")
// 		value := r.FormValue("value")
// 		fmt.Println(filterType, value)
// 		// TODO: Some sql request

// 	} else {
// 		// json.NewEncoder(w).Encode(tasks)
// 	}

// 	// w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(tasks)
// }

// // func getOneTask(w http.ResponseWriter, r *http.Request) {
// // 	vars := mux.Vars(r)
// // 	taskID, err := strconv.Atoi(vars["id"])
// // 	if err != nil {
// // 		return
// // 	}

// // 	for _, task := range tasks {
// // 		if task.ID == taskID {
// // 			// w.Header().Set("Content-Type", "application/json")
// // 			json.NewEncoder(w).Encode(task)
// // 		}
// // 	}
// // }

// func updateTask(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	taskID, err := strconv.Atoi(vars["id"])
// 	var updatedTask task

// 	if err != nil {
// 		fmt.Fprintf(w, "Invalid ID")
// 	}

// 	reqBody, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		fmt.Fprintf(w, "Please Enter Valid Data")
// 	}
// 	json.Unmarshal(reqBody, &updatedTask)

// 	for i, t := range tasks {
// 		if t.ID == taskID {
// 			tasks = append(tasks[:i], tasks[i+1:]...)

// 			updatedTask.ID = t.ID
// 			tasks = append(tasks, updatedTask)

// 			fmt.Fprintf(w, "The task with ID %v has been updated successfully", taskID)
// 		}
// 	}
// }

// func deleteTask(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	taskID, err := strconv.Atoi(vars["id"])

// 	if err != nil {
// 		fmt.Fprintf(w, "Invalid User ID")
// 		return
// 	}

// 	for i, t := range tasks {
// 		if t.ID == taskID {
// 			tasks = append(tasks[:i], tasks[i+1:]...)
// 			fmt.Fprintf(w, "The task with ID %v has been removed successfully", taskID)
// 		}
// 	}
// }

// func main() {
// 	router := mux.NewRouter().StrictSlash(true)

// 	router.HandleFunc("/", indexRoute)
// router.HandleFunc("/tasks", createTask).Methods("POST")
// router.HandleFunc("/tasks", getTasks).Methods("GET")
// router.HandleFunc("/tasks", getTasks).Methods("GET")

// // router.HandleFunc("/tasks/{id}", getOneTask).Methods("GET")
// router.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")
// router.HandleFunc("/tasks/{id}", updateTask).Methods("PUT")

// 	log.Fatal(http.ListenAndServe(":3000", router))
// }

// // http://localhost:3000/tasks?filterType=date&value=2012-01-01
// // http://localhost:3000/tasks?period=month&filterType=name&value=

// // окончательный вариант:
// // http://localhost:3000/tasks?filterType=day&value=2012-01-01
// // http://localhost:3000/tasks?filterType=month&value=2012-12
// // http://localhost:3000/tasks?filterType=year&value=2012

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"workshop2/config"

	_ "github.com/lib/pq"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	envconfig "github.com/sethvargo/go-envconfig"
)

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "sboik4"
// 	password = ""
// 	dbname   = "postgres"
// )

// func main() {
// 	// connection string
// 	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

// 	// open database
// 	db, err := sql.Open("postgres", psqlconn)
// 	CheckError(err)

// 	// close database
// 	defer db.Close()

// 	// check db
// 	err = db.Ping()
// 	CheckError(err)

// 	fmt.Println("Connected!")
// }

// func CheckError(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "sboik4"
// 	password = ""
// 	dbname   = "postgres"
// )

func main() {
	// 	// connection string
	// psqlconn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable password=%s ", host, port, user, dbname, password)
	// db, err := sql.Open("postgres", psqlconn)

	var mc config.MyConfig
	// var c config.PsConfig
	// var mc MyConfig
	ctx := context.Background()

	if err := envconfig.Process(ctx, &mc.Ps); err != nil {
		log.Fatal(err)
	}
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s ", mc.Ps.Host, mc.Ps.Port, mc.Ps.User, mc.Ps.Dbname, mc.Ps.Password)
	db, err := sql.Open("postgres", psqlconn)

	// fmt.Println(mc.Ps.Host)
	// fmt.Println(mc.Ps.Port)
	// fmt.Println(mc.Ps.User)
	// fmt.Println(mc.Ps.Password)
	// fmt.Println(mc.Ps.Dbname)

	// fmt.Println(host)
	// fmt.Println(port)
	// fmt.Println(user)
	// fmt.Println(password)
	// fmt.Println(dbname)

	// fmt.Println(mc.Ps.Host == host)
	// fmt.Println(mc.Ps.Port == port)
	// fmt.Println(mc.Ps.User == user)
	// fmt.Println(mc.Ps.Password == password)
	// fmt.Println(mc.Ps.Dbname == dbname)

	CheckError(err)

	// close database
	defer db.Close()

	// check db
	err = db.Ping()
	CheckError(err)

	// // db, err := sql.Open("postgres", "host=localhost port=5432 user=sboik4 dbname=postgres sslmode=disable password=")

	fmt.Println("Connected!")

	fmt.Println("# Querying")
	rows, err := db.Query("SELECT * FROM taskinfo")
	fmt.Println(rows)
	CheckError(err)

	// fmt.Println("# Inserting values")

	// var lastInsertId int
	// err = db.QueryRow("INSERT INTO taskinfo(name,event_type,date,description) VALUES($1,$2,$3,$4) returning id;", "some name1", "event", time.Now(), "some description").Scan(&lastInsertId)
	// CheckError(err)
	// fmt.Println("last inserted id =", lastInsertId)

	// fmt.Println("# Updating")
	// stmt, err := db.Prepare("update taskinfo set name=$1 where id=$2")
	// CheckError(err)

	// res, err := stmt.Exec("astaxieupdate", lastInsertId)
	// CheckError(err)

	// affect, err := res.RowsAffected()
	// CheckError(err)

	// fmt.Println(affect, "rows changed")

	// fmt.Println("# Querying")
	// rows, err := db.Query("SELECT * FROM taskinfo")
	// fmt.Println(rows)
	// CheckError(err)

	// for rows.Next() {
	// 	var id int
	// 	var name string
	// 	var event_type string
	// 	var date time.Time
	// 	var description string
	// 	err = rows.Scan(&id, &name, &event_type, &date, &description)
	// 	CheckError(err)
	// 	fmt.Println("id | name | type | date | description ")
	// 	fmt.Printf("%3v | %8v | %6v | %6v | %8v \n", id, name, event_type, date, description)
	// }

	// fmt.Println("# Deleting")
	// stmt, err = db.Prepare("delete from taskinfo where id=$1")
	// CheckError(err)

	// res, err = stmt.Exec(5)
	// CheckError(err)

	// affect, err = res.RowsAffected()
	// CheckError(err)

	// fmt.Println(affect, "rows changed")

	//////////////////////////
	// psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// db, err := gorm.Open("postgres", psqlconn)
	// CheckError(err)

	// defer db.Close()
	///////////////////////////////
	// insert
	// hardcoded
	// insertStmt := `insert into "tasks"("id", "name", "type", "date", "description") values(3, 'name2', 'type2', 'date2', 'description2')`
	// _, e := db.Exec(insertStmt)
	// CheckError(e)

	// insert dynamic
	// insertDynStmt := `insert into "tasks"("id", "name", "type", "date", "description") values($1, $2, $3, $4, $5)`
	// _, e := db.Exec(insertDynStmt, 10, "name4", "type4", "date4", "description4")
	// CheckError(e)

	// update
	// updateStmt := `update "tasks" set "name"=$1, "type"=$2, "date"=$3, "description"=$4  where "id"=$5`
	// _, e := db.Exec(updateStmt, "name10", "type10", "date10", "description2", 10)
	// CheckError(e)

	// deleteStmt := `delete from "tasks" where id=$1`
	// _, e := db.Exec(deleteStmt, 10)
	// CheckError(e)

	// rows, err := db.Query(`SELECT "id", "name", "type", "date", "description" FROM "tasks"`)
	// CheckError(err)

	// defer rows.Close()
	// for rows.Next() {
	// 	var id int
	// 	var name string
	// 	var Type string
	// 	var date string
	// 	var description string

	// 	err = rows.Scan(&id, &name, &Type, &date, &description)
	// 	CheckError(err)

	// 	fmt.Println(id, name, Type, date, description)
	// }

	// CheckError(err)

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
