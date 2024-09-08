package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"

    "contacts-api/src/models"
    "contacts-api/src/services"
    "github.com/gorilla/mux"
)

type ContactHandler struct {
    Service *services.ContactService
}

func NewContactHandler(service *services.ContactService) *ContactHandler {
    return &ContactHandler{Service: service}
}

func (h *ContactHandler) GetContacts(w http.ResponseWriter, r *http.Request) {
    contacts, err := h.Service.GetAllContacts()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(contacts)
}

func (h *ContactHandler) CreateContact(w http.ResponseWriter, r *http.Request) {
    var contact models.Contact
    if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    newContact, err := h.Service.CreateContact(&contact)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(newContact)
}

func (h *ContactHandler) UpdateContact(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, err := strconv.Atoi(params["id"])
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    var contact models.Contact
    if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    contact.ID = id

    updatedContact, err := h.Service.UpdateContact(&contact)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(updatedContact)
}

func (h *ContactHandler) DeleteContact(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, err := strconv.Atoi(params["id"])
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    err = h.Service.DeleteContact(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
