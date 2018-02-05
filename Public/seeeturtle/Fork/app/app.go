package app

import (
	"fmt"
	"log"
	"net/http"

	"database/sql"

	"github.com/gorilla/mux"
	"github.com/seeeturtle/Fork/app/handler"
	"github.com/seeeturtle/Fork/app/model"
	"github.com/seeeturtle/Fork/config"
	_ "github.com/lib/pq"
)

// App has router and db instances
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Initialize initializes the app with predefined configuration
func (a *App) Initialize(config *config.Config) {
	dbinfo := fmt.Sprintf("user=%s dbname=%s sslmode=disable password=%s",
		config.DB.Username,
		config.DB.Name,
		config.DB.Password,
	)

	db, err := sql.Open(config.DB.Dialect, dbinfo)
	if err != nil {
		log.Fatal("Could not connect database")
	}

	a.DB = db
	a.Router = mux.NewRouter()
	a.setRouters()
	a.setModels()
}

// setRouters sets the all required routers
func (a *App) setRouters() {
	a.Get("/keyboard", a.GetKeyboard)
	a.Post("/message", a.CreateMessage)
}

func (a *App) setModels() {
	model.Lunches = model.Lunches.New(a.DB)
}

// Get wraps the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Post wraps the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Put wraps the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

/*
** Handlers
 */
func (a *App) GetKeyboard(w http.ResponseWriter, r *http.Request) {
	handler.GetKeyboard(w, r)
}

func (a *App) CreateMessage(w http.ResponseWriter, r *http.Request) {
	handler.CreateMessage(w, r)
}

// Run the app and log request
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
	defer a.DB.Close()
}
