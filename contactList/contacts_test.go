package contacts

import (
	"fmt"
	"postgre_task/db"
	"testing"
)

var (
	testContacts = []Contact{
		{
			Id:        1000,
			FirstName: "Sam",
			LastName:  "Smith",
			Phone:     "(695)-175-4661",
			Email:     "sam@local.com",
		},
		{
			Id:        1100,
			FirstName: "Eugene",
			LastName:  "Williamson",
			Phone:     "(139)-191-0039",
			Email:     "eugene@local.com",
		},
		{
			Id:        1200,
			FirstName: "Brian",
			LastName:  "Robinson",
			Phone:     "(045)-207-9455",
			Email:     "brian.robinson@example.com",
		},
	}
	updContact = Contact{
		Id:        1200,
		FirstName: "Lee",
		LastName:  "Wright",
		Phone:     "(215)-511-9272",
		Email:     "lee.wright@example.com",
	}
	clt = new(ContactList)
)

func TestCreateContact(t *testing.T) {
	clt.ContactsDb = db.Db
	var del string

	for i, v := range testContacts {
		err := clt.Create(v)
		if err != nil {
			t.Error(err)
		}

		if i == len(testContacts)-1 {
			del += fmt.Sprintf("%d", v.Id)
		} else {
			del += fmt.Sprintf("%d, ", v.Id)
		}
	}

	t.Cleanup(func() {
		delQuery := fmt.Sprintf("DELETE FROM contacts WHERE contact_id IN (%s)", del)
		_, err := clt.ContactsDb.Exec(delQuery)
		clt = new(ContactList)
		if err != nil {
			t.Error(err)
		}
	})
}

func TestUpdateContact(t *testing.T) {
	TestCreateContact(t)

	err := clt.Update(updContact)
	if err != nil {
		t.Error(err)
	}
}

func TestGetContact(t *testing.T) {
	TestCreateContact(t)

	for _, v := range testContacts {
		temp, err := clt.Get(v.Id)
		if err != nil {
			t.Error(err)
		}
		if temp != v {
			t.Error("method get failed")
		}
	}
}

func TestGetAllContact(t *testing.T) {
	TestCreateContact(t)

	temp, err := clt.GetAll()
	if err != nil {
		t.Error(err)
	}
	for i := range testContacts {
		if testContacts[i] != temp[i] {
			t.Error("failed getall method")
		}
	}
}

func TestGetDeleteContact(t *testing.T) {
	TestCreateContact(t)

	for _, v := range testContacts {
		err := clt.Delete(v.Id)
		if err != nil {
			t.Error(err)
		}
	}
}
