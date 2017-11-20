package main

import (
	"Ksana/controller"
	"Ksana/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	app := new(router.Router)

	app.Get("/", func(p controller.Context) {
		fmt.Fprintf(p.Res, "Hello World")
	})
	app.Post("/post/add", controller.AddPost)

	if err := http.ListenAndServe(":9090", app); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}