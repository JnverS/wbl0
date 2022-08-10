package apiserver

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"strings"
)

type Server struct {
	http.Server
	router *mux.Router
}

func (a *Api) StartHttp(addr string) *Server {
	router := mux.NewRouter()

	server := Server {
			Server: http.Server{
					Addr:    addr,
					Handler: router,
			},
	}
	router.HandleFunc("/", a.homeHandler)
	router.HandleFunc("/order/{id:[0-9]+}", a.orderHandler)
	return &server
}

func (a Api) homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		log.Println("StatusMethodNotAllowed")
		return
	}
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Println("Can not expand template")
		return
	}
	tmpl.Execute(w, make(map[int]struct{}))
}

func (a *Api) orderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		log.Println("Error method")
		return
	}
	stringId := strings.Replace(r.URL.Path, "/order/", "", -1)
	id, er := strconv.Atoi(stringId)
	if er != nil {
		log.Println("Invalid id")
		return
	}
	model, ok := a.Get(id)
	if !ok {
		log.Println("Couldn't get the order")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(model)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

}
