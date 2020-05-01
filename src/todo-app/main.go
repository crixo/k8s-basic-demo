package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/crixo/k8s-basic-demo/config"
	"github.com/crixo/k8s-basic-demo/handlers"
	"github.com/crixo/k8s-basic-demo/model"

	flags "github.com/jessevdk/go-flags"
)

var (
	repository model.Repository
)

func main() {
	config := config.Config{}
	_, err := flags.Parse(&config)

	if err != nil {
		panic(err)
	}

	fmt.Printf("DbUser: %v\n", config.DbUser)
	fmt.Printf("DbPassword: %v\n", len(config.DbPassword))
	fmt.Printf("DbName: %v\n", config.DbName)
	fmt.Printf("DbHost: %v\n", config.DbHost)
	fmt.Printf("MigrationOnly: %v\n", config.MigrationOnly)
	fmt.Printf("CUnknown: %v\n", config.CUnknown)

	model.CreateMysqlDbIfNotExists(config.SqlConnString(), config.DbName)

	db := model.CreateDb(config.GormConnString())
	defer db.Close()

	if config.MigrationOnly {
		model.RunMigration(db)
		fmt.Printf("migration done. exiting")
		//os.Exit(1)
	} else {

		fmt.Printf("creating repository")
		repository = model.NewMysqlRepository(db)

		//httpClient := &http.Client{}
		//fs := http.FileServer(http.Dir("static"))

		handler := handlers.NewHandler(repository)
		//http.Handle("/", fs)
		http.Handle("/", http.HandlerFunc(handler.Index))
		http.Handle("/home", http.HandlerFunc(handler.Home))
		http.Handle("/send", http.HandlerFunc(handler.Send))
		http.Handle("/confirmation", http.HandlerFunc(handler.Confirmation))
		// http.Handle("/api/items", handler.NewGetItemsHandler(config, httpClient))
		// http.Handle("/api/pay", handler.NewPayHandler(config, httpClient))

		// c := make(chan os.Signal)
		// signal.Notify(c, os.Interrupt)

		// go func() {
		// 	select {
		// 	case sig := <-c:
		// 		fmt.Printf("Got %s signal. Aborting...\n", sig)
		// 		os.Exit(1)
		// 	}
		// }()

		log.Println("Listening on port 3000...")
		http.ListenAndServe(":3000", nil)
	}
}
