package services

import (
    "contacts-api/src/models"
    "contacts-api/src/repositories"
)

type ContactService struct {
    Repo *repositories.ContactRepository
}

func NewContactService(repo *repositories.ContactRepository) *ContactService {
    return &ContactService{Repo: repo}
}

func (s *ContactService) GetAllContacts() ([]models.Contact, error) {
    return s.Repo.GetAll()
}

func (s *ContactService) CreateContact(contact *models.Contact) (*models.Contact, error) {
    return s.Repo.Create(contact)
}

func (s *ContactService) UpdateContact(contact *models.Contact) (*models.Contact, error) {
    return s.Repo.Update(contact)
}

func (s *ContactService) DeleteContact(id int) error {
    return s.Repo.Delete(id)
}
