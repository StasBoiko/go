package main

import (
	"context"

	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"work2/config"
	"work2/handlers"
	"work2/middlewares"
	"work2/repository/postgres"
	"work2/services/task"
	"work2/services/user"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"go.uber.org/zap"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	envconfig "github.com/sethvargo/go-envconfig"
)

func main() {
	var mc config.MyConfig
	ctx := context.Background()

	if err := envconfig.Process(ctx, &mc.Ps); err != nil {
		log.Fatal(err)
	}

	pr, err := postgres.NewRepo(mc.Ps)
	if err != nil {
		log.Fatal(err)
	}

	tSer := task.TaskNewService(pr)
	uSer := user.UserNewService(pr)
	logz, _ := zap.NewDevelopment()
	defer logz.Sync()
	//var logz *zap.Logger
	//router.Handle("/logout", middlewares.Authz(http.HandlerFunc(s.Logout))).Methods("POST")
	// http://localhost:3010/tasks?filterType=day&value=2012-01-01
	//s := middlewares.NewServer(tSer, uSer, log)
	s := handlers.NewServer(tSer, uSer, logz)
	m := middlewares.NewMiddleware(uSer, logz)

	//router := mux.NewRouter()
	//router.HandleFunc("/signup", s.Signup).Methods("POST")
	//router.HandleFunc("/login", s.Login).Methods("POST")
	//router.HandleFunc("/logout", s.Logout).Methods("POST")
	//router.Handle("/tasks", m.Authz(http.HandlerFunc(s.GetTasks))).Methods("GET")
	//router.Handle("/tasks", m.Authz(http.HandlerFunc(s.CreateTask))).Methods("POST")
	//router.Handle("/tasks/{id}", m.Authz(http.HandlerFunc(s.UpdateTask))).Methods("PUT")
	//router.Handle("/tasks/{id}", m.Authz(http.HandlerFunc(s.DeleteTask))).Methods("DELETE")
	////правильно ли я сделал? Это работает :)
	//routerMetrics := mux.NewRouter()
	//routerMetrics.Handle("/metrics", promhttp.Handler())
	//handlerMetrics := cors.Default().Handler(routerMetrics)
	//go http.ListenAndServe(":8000", handlerMetrics)
	//handler := cors.Default().Handler(router)
	//log.Fatal(http.ListenAndServe(":3000", handler))

	//prometheus.MustRegister(models.HttpReqs)
	//prometheus.MustRegister(models.UsersCreated)
	//models.UsersCreated.WithLabelValues("Vasya").Add(1)
	//models.HttpReqs.WithLabelValues("Vasya").Add(5)
	//models.UsersCreated.WithLabelValues("Petya").Add(1)
	//models.HttpReqs.WithLabelValues("Petya").Add(2)
	//models.UsersCreated.WithLabelValues("Kolya").Add(1)
	//models.HttpReqs.WithLabelValues("Kolya").Add(3)
	//models.HttpReqs.WithLabelValues("404", "POST").Add(42)
	// If you have to access the same set of labels very frequently, it
	// might be good to retrieve the metric only once and keep a handle to
	// it. But beware of deletion of that metric, see below!
	//foo := models.HttpReqs.WithLabelValues("200", "GET")
	//for i := 0; i < 1000000; i++ {
	//	foo.Inc()
	//}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		oscall := <-c
		log.Printf("system call:%+v", oscall)
		cancel()
	}()

	if err := serve(ctx, s, m); err != nil {
		log.Printf("failed to serve:+%v\n", err)
	}
}

func serve(ctx context.Context, s *handlers.Server, m *middlewares.Middleware) (err error) {
	router := mux.NewRouter()
	router.HandleFunc("/signup", s.Signup).Methods("POST")
	router.HandleFunc("/login", s.Login).Methods("POST")
	router.HandleFunc("/logout", s.Logout).Methods("POST")
	router.Handle("/tasks", m.Authz(http.HandlerFunc(s.GetTasks))).Methods("GET")
	router.Handle("/tasks", m.Authz(http.HandlerFunc(s.CreateTask))).Methods("POST")
	router.Handle("/tasks/{id}", m.Authz(http.HandlerFunc(s.UpdateTask))).Methods("PUT")
	router.Handle("/tasks/{id}", m.Authz(http.HandlerFunc(s.DeleteTask))).Methods("DELETE")

	routerMetrics := mux.NewRouter()
	routerMetrics.Handle("/metrics", promhttp.Handler())
	//routerMetrics.Handle("/metrics", promhttp.HandlerFor(s.GetTasks))
	handlerMetrics := cors.Default().Handler(routerMetrics)
	go http.ListenAndServe(":9091", handlerMetrics)
	handler := cors.Default().Handler(router)
	//log.Fatal(http.ListenAndServe(":3000", handler))

	srv := &http.Server{
		Addr:    ":2000",
		Handler: handler,
	}

	go func() {
		if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen:%+s\n", err)
		}
	}()

	log.Printf("server started")
	<-ctx.Done()
	log.Printf("server stopped")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err = srv.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("server Shutdown Failed:%+s", err)
	}

	log.Printf("server exited properly")
	if err == http.ErrServerClosed {
		err = nil
	}
	return
}

//func init() {
//	// must register counter on init
//	prometheus.MustRegister(models.RequestCounter)
//}

//func recordMetrics() {
//	go func() {
//		for {
//			opsProcessed.Inc()
//			time.Sleep(2 * time.Second)
//		}
//	}()
//}

//func GetTasks2(w http.ResponseWriter, r *http.Request) {
//	var mc config.MyConfig
//	ctx := context.Background()
//
//	if err := envconfig.Process(ctx, &mc.Ps); err != nil {
//		log.Fatal(err)
//	}
//
//	psqlconn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s ",
//		mc.Ps.Host, mc.Ps.Port, mc.Ps.User, mc.Ps.Dbname, mc.Ps.Password)
//
//	db, err := sql.Open("postgres", psqlconn)
//	CheckError(err)
//
//	// close database
//	defer db.Close()
//
//	// check db
//	err = db.Ping()
//	CheckError(err)
//
//	rows, err := db.Query("SELECT * FROM tasks")
//	CheckError(err)
//
//	var tasks []models.Task
//	for rows.Next() {
//		var id uint
//		var created_at time.Time
//		var updated_at time.Time
//		var deleted_at sql.NullTime
//		var name string
//		var event string
//		var date time.Time
//		var description string
//		err = rows.Scan(&id, &created_at, &updated_at, &deleted_at, &name, &event, &date, &description)
//		CheckError(err)
//
//		// var allTasks := []string{}
//		tasks = append(tasks, models.Task{
//			ID:          id,
//			CreatedAt:   created_at,
//			UpdatedAt:   updated_at,
//			DeletedAt:   deleted_at,
//			Name:        name,
//			Event:       event,
//			Date:        date,
//			Description: description,
//		})
//	}
//
//	json.NewEncoder(w).Encode(&tasks)
//}

//func CheckError(err error) {
//	if err != nil {
//		panic(err)
//	}
//}
