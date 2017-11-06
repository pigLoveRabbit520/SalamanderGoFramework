package main

import (

	"github.com/go-chi/chi"
	"github.com/salamander/SalamanderGo/controls"
)


func InitAPIRouter(checkUserToken bool) chi.Router {
	api := chi.NewRouter()

	api.Get("/", controls.ShowHome)

	return api
}
