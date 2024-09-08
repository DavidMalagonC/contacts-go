package repositories

import (
    "contacts-api/src/models"
    "database/sql"
)

type ContactRepository struct {
    DB *sql.DB
}

func NewContactRepository(db *sql.DB) *ContactRepository {
    return &ContactRepository{DB: db}
}

func (r *ContactRepository) GetAll() ([]models.Contact, error) {
    rows, err := r.DB.Query("SELECT id, name, email, phone FROM contacts")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var contacts []models.Contact
    for rows.Next() {
        var contact models.Contact
        if err := rows.Scan(&contact.ID, &contact.Name, &contact.Email, &contact.Phone); err != nil {
            return nil, err
        }
        contacts = append(contacts, contact)
    }

    return contacts, nil
}

func (r *ContactRepository) Create(contact *models.Contact) (*models.Contact, error) {
    result, err := r.DB.Exec("INSERT INTO contacts (name, email, phone) VALUES (?, ?, ?)", contact.Name, contact.Email, contact.Phone)
    if err != nil {
        return nil, err
    }

    id, err := result.LastInsertId()
    if err != nil {
        return nil, err
    }
    contact.ID = int(id)

    return contact, nil
}

func (r *ContactRepository) Update(contact *models.Contact) (*models.Contact, error) {
    _, err := r.DB.Exec("UPDATE contacts SET name=?, email=?, phone=? WHERE id=?", contact.Name, contact.Email, contact.Phone, contact.ID)
    if err != nil {
        return nil, err
    }

    return contact, nil
}

func (r *ContactRepository) Delete(id int) error {
    _, err := r.DB.Exec("DELETE FROM contacts WHERE id=?", id)
    return err
}
