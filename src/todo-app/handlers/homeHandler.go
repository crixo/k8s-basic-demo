package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
	"text/template"

	"github.com/crixo/k8s-basic-demo/model"
)

type handler struct {
	repository model.Repository
}

func NewHandler(repository model.Repository) *handler {
	return &handler{
		repository: repository,
	}
}

func (h handler) Index(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadFile("/tmp/k8s-basic-demo-config/my-file.txt")
	var output string
	if err != nil {
		log.Fatal(err)
		output = "[error]" + err.Error()
	} else {
		output = string(content)
	}

	//fmt.Printf("File contents: %s", content)
	render(w, "templates/index.html", output)
}

func (h handler) Home(w http.ResponseWriter, r *http.Request) {
	render(w, "templates/home.html", h.repository.GetList())
}

func (h handler) Send(w http.ResponseWriter, r *http.Request) {
	// Step 1: Validate form
	msg := &model.Message{
		Email:   r.PostFormValue("email"),
		Content: r.PostFormValue("content"),
	}

	if msg.Validate() == false {
		render(w, "templates/contact.html", msg)
		return
	}

	// Step 2: Store message into the repository

	if err := h.repository.Store(msg); err != nil {
		log.Println(err)
		http.Error(w, "Sorry, something went wrong", http.StatusInternalServerError)
		return
	}

	// Step 3: Redirect to confirmation page
	http.Redirect(w, r, "/confirmation", http.StatusSeeOther)
}

func (h handler) Confirmation(w http.ResponseWriter, r *http.Request) {
	render(w, "templates/confirmation.html", h.repository.GetList())
}

func render(w http.ResponseWriter, filename string, data interface{}) {
	tmpl, err := template.ParseFiles(filename)
	if err != nil {
		log.Println(err)
		http.Error(w, "Sorry, something went wrong", http.StatusInternalServerError)
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Println(err)
		http.Error(w, "Sorry, something went wrong", http.StatusInternalServerError)
	}
}
