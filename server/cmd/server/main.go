package main

import (
	"github.com/hermant2/angelventureserver/pkg/applogger"
	"github.com/hermant2/angelventureserver/pkg/routes"
	"net/http"
)

func main() {
	applogger.Instance().Fatal("listen_and_serve_http", http.ListenAndServe(":8080", routes.Router()))
}
