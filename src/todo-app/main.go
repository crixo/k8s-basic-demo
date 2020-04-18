package main

import (
	"html/template"
	"log"
	"net/http"
	// "github.com/gianarb/shopmany/frontend/config"
	// "github.com/gianarb/shopmany/frontend/handler"
	// flags "github.com/jessevdk/go-flags"
)

func main() {
	// config := config.Config{}
	// _, err := flags.Parse(&config)

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("Item Host: %v\n", config.ItemHost)
	// fmt.Printf("Pay Host: %v\n", config.PayHost)
	// fmt.Printf("Discount Host: %v\n", config.DiscountHost)

	//httpClient := &http.Client{}
	fs := http.FileServer(http.Dir("static"))

	http.Handle("/", fs)
	http.Handle("/home", http.HandlerFunc(home))
	http.Handle("/send", http.HandlerFunc(send))
	http.Handle("/confirmation", http.HandlerFunc(confirmation))
	// http.Handle("/api/items", handler.NewGetItemsHandler(config, httpClient))
	// http.Handle("/api/pay", handler.NewPayHandler(config, httpClient))

	log.Println("Listening on port 3000...")
	http.ListenAndServe(":3000", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	msgs := messages
	render(w, "templates/home.html", msgs)
}

func send(w http.ResponseWriter, r *http.Request) {
	// Step 1: Validate form
	msg := &Message{
		Email:   r.PostFormValue("email"),
		Content: r.PostFormValue("content"),
	}

	if msg.Validate() == false {
		render(w, "templates/contact.html", msg)
		return
	}

	// Step 2: Store message into the repository
	if err := msg.Store(); err != nil {
		log.Println(err)
		http.Error(w, "Sorry, something went wrong", http.StatusInternalServerError)
		return
	}

	// Step 3: Redirect to confirmation page
	http.Redirect(w, r, "/confirmation", http.StatusSeeOther)
}

func confirmation(w http.ResponseWriter, r *http.Request) {
	msgs := messages
	render(w, "templates/confirmation.html", msgs)
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
