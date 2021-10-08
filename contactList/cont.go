package contacts

import (
	"database/sql"
	"fmt"
)

type Contact struct {
	Id        int
	FirstName string
	LastName  string
	Phone     string
	Email     string
}

type ContactList struct {
	ContactsDb *sql.DB
}

// Methods: create, update, get, getAll and delete

func (cl *ContactList) Create(con Contact) error {
	_, err := cl.ContactsDb.Exec(`INSERT INTO contacts
	(contact_id, first_name, last_name, phone, email)
	VALUES ($1, $2, $3, $4, $5)
	`, con.Id, con.FirstName, con.LastName, con.Phone, con.Email)
	return err
}

func (cl *ContactList) Update(con Contact) error {
	res, err := cl.ContactsDb.Exec(`UPDATE tasks 
			SET first_name = $1, last_name = $2, phone = $3, email = $4
			WHERE contact_id = $5
			`, con.FirstName, con.LastName, con.Phone, con.Phone, con.Id)

	if err != nil {
		return err
	}

	if num, _ := res.RowsAffected(); num == 0 {
		return fmt.Errorf("task %d not exists", con.Id)
	}
	return nil
}

func (cl *ContactList) Get(id int) (Contact, error) {
	row := cl.ContactsDb.QueryRow("SELECT contact_id, first_name, last_name, phone, email from contacts WHERE contact_id = $1", id)

	var contact, emptyContact Contact
	row.Scan(&contact.Id, &contact.FirstName, &contact.LastName, &contact.Phone, &contact.Email)

	if contact == emptyContact {
		return Contact{}, fmt.Errorf("task %d not exists", id)
	}

	return contact, nil
}

func (cl *ContactList) GetAll() ([]Contact, error) {
	rows, err := cl.ContactsDb.Query("SELECT contact_id, first_name, last_name, phone, email from contacts ORDER BY contact_id")
	if err != nil {
		return make([]Contact, 0), err
	}
	var result []Contact
	var tempContact Contact

	for rows.Next() {
		rows.Scan(&tempContact.Id, &tempContact.FirstName, &tempContact.LastName, &tempContact.Phone, &tempContact.Email)

		result = append(result, tempContact)
	}

	return result, nil
}

func (cl *ContactList) Delete(id int) {
	cl.ContactsDb.QueryRow("DELETE FROM contacts WHERE contact_id = $1", id)
}
