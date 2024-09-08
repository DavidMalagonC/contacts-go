package dependencies

import (
	"contacts-api/config"
	"contacts-api/src/handlers"
	"contacts-api/src/repositories"
	"contacts-api/src/services"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func Initialize() error {
	db, err := config.NewConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	contactRepo := repositories.NewContactRepository(db)
	contactService := services.NewContactService(contactRepo)
	contactHandler := handlers.NewContactHandler(contactService)

	router := mux.NewRouter()
	router.HandleFunc("/contacts", contactHandler.GetContacts).Methods("GET")
	router.HandleFunc("/contacts", contactHandler.CreateContact).Methods("POST")
	router.HandleFunc("/contacts/{id}", contactHandler.UpdateContact).Methods("PUT")
	router.HandleFunc("/contacts/{id}", contactHandler.DeleteContact).Methods("DELETE")

	log.Println("Server running on port 8080")
	return http.ListenAndServe(":8080", router)
}
