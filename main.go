package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

var (
	connectionString = flag.String("conn", "postgres://kukkomah:@127.0.0.1:5432/alfamt-messenger?sslmode=disable", "PostgreSQL connection string")
	listenAddr       = flag.String("addr", ":1234", "HTTP address to listen on")
	db               *sqlx.DB
)

type Messenger int

func connectToDB() {
	var err error
	db, err = sqlx.Open("pgx", *connectionString)
	if err != nil {
		log.Fatalf("Unable to establish connection: %v\n", err)
	}
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(100)
}

func createServer() {
	s := rpc.NewServer()
	s.RegisterCodec(json.NewCodec(), "application/json")
	s.RegisterCodec(json.NewCodec(), "application/json;charset=UTF-8")

	s.RegisterService(new(Messenger), "")

	r := mux.NewRouter()
	r.Handle("/rpc", s)
	http.ListenAndServe(*listenAddr, r)
}

func main() {
	flag.Parse()

	connectToDB()
	createServer()
}
