package tests

import (
    "contacts-api/src/models"
)

type MockContactService struct {
    Contacts []models.Contact
}

func (m *MockContactService) GetAllContacts() ([]models.Contact, error) {
    return m.Contacts, nil
}

func (m *MockContactService) CreateContact(contact *models.Contact) (*models.Contact, error) {
    contact.ID = len(m.Contacts) + 1
    m.Contacts = append(m.Contacts, *contact)
    return contact, nil
}

func (m *MockContactService) UpdateContact(contact *models.Contact) (*models.Contact, error) {
    for i, c := range m.Contacts {
        if c.ID == contact.ID {
            m.Contacts[i] = *contact
            return contact, nil
        }
    }
    return nil, nil
}

func (m *MockContactService) DeleteContact(id int) error {
    for i, c := range m.Contacts {
        if c.ID == id {
            m.Contacts = append(m.Contacts[:i], m.Contacts[i+1:]...)
            return nil
        }
    }
    return nil
}
