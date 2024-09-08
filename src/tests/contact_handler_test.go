package tests

import (
    "bytes"
    "contacts-api/src/handlers"
    "contacts-api/src/models"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "strconv"
    "testing"
    "github.com/gorilla/mux"
    "github.com/stretchr/testify/assert"
)

func TestGetContacts(t *testing.T) {
    service := &MockContactService{
        Contacts: []models.Contact{
            {ID: 1, Name: "John Doe", Email: "john@example.com", Phone: "123456789"},
        },
    }
    handler := handlers.NewContactHandler(service)

    req, err := http.NewRequest("GET", "/contacts", nil)
    assert.NoError(t, err)

    rr := httptest.NewRecorder()
    handler.GetContacts(rr, req)

    assert.Equal(t, http.StatusOK, rr.Code)
    assert.Contains(t, rr.Body.String(), "John Doe")
}

func TestCreateContact(t *testing.T) {
    service := &MockContactService{}
    handler := handlers.NewContactHandler(service)

    contact := models.Contact{Name: "Jane Doe", Email: "jane@example.com", Phone: "987654321"}
    body, _ := json.Marshal(contact)

    req, err := http.NewRequest("POST", "/contacts", bytes.NewBuffer(body))
    assert.NoError(t, err)

    rr := httptest.NewRecorder()
    handler.CreateContact(rr, req)

    assert.Equal(t, http.StatusOK, rr.Code)
    assert.Contains(t, rr.Body.String(), "Jane Doe")
}

func TestUpdateContact(t *testing.T) {
    service := &MockContactService{
        Contacts: []models.Contact{
            {ID: 1, Name: "John Doe", Email: "john@example.com", Phone: "123456789"},
        },
    }
    handler := handlers.NewContactHandler(service)

    updatedContact := models.Contact{ID: 1, Name: "John Smith", Email: "john.smith@example.com", Phone: "987654321"}
    body, _ := json.Marshal(updatedContact)

    req, err := http.NewRequest("PUT", "/contacts/1", bytes.NewBuffer(body))
    assert.NoError(t, err)

    req = mux.SetURLVars(req, map[string]string{"id": strconv.Itoa(updatedContact.ID)})

    rr := httptest.NewRecorder()
    handler.UpdateContact(rr, req)

    assert.Equal(t, http.StatusOK, rr.Code)
    assert.Contains(t, rr.Body.String(), "John Smith")
}

func TestDeleteContact(t *testing.T) {
    service := &MockContactService{
        Contacts: []models.Contact{
            {ID: 1, Name: "John Doe", Email: "john@example.com", Phone: "123456789"},
        },
    }
    handler := handlers.NewContactHandler(service)

    req, err := http.NewRequest("DELETE", "/contacts/1", nil)
    assert.NoError(t, err)

    req = mux.SetURLVars(req, map[string]string{"id": "1"})

    rr := httptest.NewRecorder()
    handler.DeleteContact(rr, req)

    assert.Equal(t, http.StatusNoContent, rr.Code)
    assert.Empty(t, rr.Body.String())
}
