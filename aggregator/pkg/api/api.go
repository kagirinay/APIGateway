package api

import (
	"APIGateway/aggregator/pkg/storage"
	"APIGateway/aggregator/pkg/storage/postgres"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

const limit = 10

// API Программный интерфейс сервера.
type API struct {
	db     storage.Interface // БД.
	router *mux.Router       // Маршрутизатор.
}

// New Конструктор объекта API.
func New(db *postgres.Store) *API {
	api := API{}
	api.router = mux.NewRouter()
	api.endpoints()

	return &api
}

// Router Получение маршрутизатора запросов.
// Требуется для передачи маршрутизатора веб-серверу.
func (api *API) Router() *mux.Router {

	return api.Router()
}

// Регистрация методов API в маршрутизаторе запросов.
func (api *API) endpoints() {
	api.router.HandleFunc("/news", api.postsHandler).Methods(http.MethodGet, http.MethodOptions)
	api.router.HandleFunc("/news/latest", api.newsLatestHandler).Methods(http.MethodGet, http.MethodOptions)
	api.router.HandleFunc("/news/search", api.newsDetailedHandler).Methods(http.MethodGet, http.MethodOptions)
}

// postsHandler Возвращает все публикации.
func (api *API) postsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	pageParam := r.URL.Query().Get("page")
	if pageParam == "" {
		pageParam = "1"
	}
	sParam := r.URL.Query().Get("s")
	page, err := strconv.Atoi(pageParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	posts, pagination, err := api.db.PostSearchILIKE(sParam, limit, (page-1)*limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response := struct {
		Posts      []storage.Post
		Pagination storage.Pagination
	}{
		Posts:      posts,
		Pagination: pagination,
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Получение страницы с определенным номером
func (api *API) newsLatestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	pageParam := r.URL.Query().Get("page")
	if pageParam == "" {
		pageParam = "1"
	}
	page, err := strconv.Atoi(pageParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	posts, err := api.db.Posts(limit, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(posts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Получение публикаций по id
func (api *API) newsDetailedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	idParam := r.URL.Query().Get("id")
	log.Println(idParam)
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	post, err := api.db.PostDetail(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
