package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/StasBoiko/test-work/config"
	"github.com/StasBoiko/test-work/gqlgen"
	"github.com/StasBoiko/test-work/postgres"

	envconfig "github.com/sethvargo/go-envconfig"
)

func main() {
	var mc config.MyConfig
	ctx := context.Background()
	if err := envconfig.Process(ctx, &mc.Ps); err != nil {
		log.Fatal(err)
	}

	var ps = mc.Ps
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s ",
		ps.Host, ps.Port, ps.User, ps.Dbname, ps.Password)
	db, err := postgres.Open(psqlconn)

	// initialize the db
	defer check(db.Close)
	if err != nil {
		panic(err)
	}

	// initialize the repository
	repo := postgres.NewRepository(db)

	// configure the server
	mux := http.NewServeMux()
	mux.Handle("/", gqlgen.NewPlaygroundHandler("/query"))
	mux.Handle("/query", gqlgen.NewHandler(repo))

	// run the server
	port := ":8080"
	fmt.Fprintf(os.Stdout, "🚀 Server ready at http://localhost%s\n", port)
	fmt.Fprintln(os.Stderr, http.ListenAndServe(port, mux))
}

func check(f func() error) {
	if err := f(); err != nil {
		panic(err)
	}
}
