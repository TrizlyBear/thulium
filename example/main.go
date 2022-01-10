package main

import (
	th "github.com/TrizlyBear/thulium"
	"github.com/gorilla/mux"
	"net/http"
)

var (
	comp = &th.P{Text: "Hello world!"}
)

func main()  {
	r := mux.NewRouter()
	r.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		th.Render(writer, comp)
	})

	http.ListenAndServe("localhost:6969",r)
}
