package provider

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/amiraliio/avn-promotion/config"
	"github.com/amiraliio/avn-promotion/handler"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Init() {
	app := new(config.App)
	app.Init()

	err := make(chan error, 4)

	go func() {
		grpcListener(app)
	}()

	go func() {
		httpListener(app, err)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		err <- fmt.Errorf("%s", <-c)
	}()

	fmt.Printf("terminated %s", <-err)
}

func httpListener(app *config.App, err chan<- error) {
	router := mux.NewRouter()

	handler.HTTP(app, router)

	router.Use(mux.CORSMethodMiddleware(router))

	fmt.Println("Listening HTTP on port " + config.AppConfig.GetString("APP.HTTP_PORT"))

	err <- http.ListenAndServe(":"+config.AppConfig.GetString("APP.HTTP_PORT"), router)
}

func grpcListener(app *config.App) {
	listener, _ := net.Listen("tcp", ":"+config.AppConfig.GetString("APP.TCP_PORT"))

	server := grpc.NewServer()

	handler.GRPC(app, server)

	reflection.Register(server)

	fmt.Println("Listening TCP on port " + config.AppConfig.GetString("APP.TCP_PORT"))

	_= server.Serve(listener)
}
