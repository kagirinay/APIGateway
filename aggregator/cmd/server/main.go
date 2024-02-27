package main

// Сервер.
type server struct {
	db  storage.Interface
	api *api.API
}
