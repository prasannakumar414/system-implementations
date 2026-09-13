package main

import (
	"context"

	"github.com/prasannakumar414/system-implementations/http"
)

func main() {
	conn, err := InitConn()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	err = Init(context.Background(), "test-service", conn)
	if err != nil {
		panic(err)
	}

	logger := InitZapLogger()
	defer logger.Sync() // Flush any buffered log entries before exiting.
	router := http.NewRouter(logger)
	router.Run(":8085")
}
