package main

import (
	"net/http"

	_ "github.com/dogntao/rymz/controllers"
	"github.com/dogntao/rymz/routers"
)

func main() {
	http.HandleFunc("/", routers.Router.Router)
	http.ListenAndServe("127.0.0.1:8088", nil)
}
