package main

import (
	"bufio"
	"bytes"
	"context"
	"flag"
	"fmt"
	"github.com/go-kit/kit/log"
	_ "github.com/lib/pq"
	"strings"

	"github.com/go-kit/kit/log/level"

	"net/http"
	"os"
	"os/signal"
	"syscall"

	"balance-service-go/balance"
)

func main() {
	var httpAddr = flag.String("http", ":8080", "http listen address")
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "account",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	dbHost, dbUser, dbPassword, dbName :=
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME")
	database, err := Initialize(dbHost, dbUser, dbPassword, dbName)
	if err != nil {
		level.Error(logger).Log("exit", "Could not set up database")
		os.Exit(-1)
	}
	defer database.Conn.Close()

	script := strings.Split(readFile(), ";")
	for _, s := range script {
		database.Conn.Exec(s)
	}

	flag.Parse()
	ctx := context.Background()
	var srv balance.Service
	{
		repository := balance.NewRepo(database.Conn)

		srv = balance.NewService(repository)
	}

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	endpoints := balance.MakeEndpoints(srv)

	go func() {
		fmt.Println("listening on port", *httpAddr)
		handler := balance.NewHTTPServer(ctx, endpoints)
		errs <- http.ListenAndServe(*httpAddr, handler)
	}()

	level.Error(logger).Log("exit", <-errs)
}

func readFile() string {
	file, err := os.Open("start.sql")
	if err != nil {
		fmt.Println(err)
	}

	wr := bytes.Buffer{}
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		wr.WriteString(sc.Text())
	}

	return wr.String()
}
