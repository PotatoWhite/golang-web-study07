package main

import (
	"encoding/json"
	"github.com/gorilla/pat"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
	"net/http"
	"time"
)

var rd *render.Render

type User struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func getUser(writer http.ResponseWriter, request *http.Request) {
	aUser := User{
		Name:  "potato",
		Email: "potato@example.com",
	}

	rd.JSON(writer, http.StatusOK, aUser)
}

func createUser(writer http.ResponseWriter, request *http.Request) {
	aUser := new(User)
	err := json.NewDecoder(request.Body).Decode(&aUser)
	if err != nil {
		rd.JSON(writer, http.StatusBadRequest, err.Error())
		return
	}

	aUser.CreatedAt = time.Now()
	rd.JSON(writer, http.StatusCreated, aUser)
}

func main() {
	n := NewHandler()
	http.ListenAndServe(":8080", n)
}

func NewHandler() *negroni.Negroni {
	rd = render.New(render.Options{
		Directory:  "templates",
		Extensions: []string{".html", ".tmpl"},
		Layout:     "hello",
	})
	mux := pat.New()
	mux.Get("/users", getUser)
	mux.Post("/users", createUser)
	mux.Get("/hello", sayHello)

	//mux.Handle("/", http.FileServer(http.Dir("public")))
	// apply negroni

	n := negroni.Classic()
	n.UseHandler(mux)
	return n
}

func sayHello(writer http.ResponseWriter, request *http.Request) {
	rd.HTML(writer, http.StatusOK, "body", User{Name: "Potato", Email: "potato@example.com"})
}
