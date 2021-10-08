package main

import (
	"fmt"
	contacts "postgre_task/contactList"
	"postgre_task/db"
	tasklist "postgre_task/taskList"
	"time"
)

func main() {

	taskL := tasklist.TaskList{TasksDb: db.Db}

	task1 := tasklist.Task{
		Id:        12,
		Name:      "Task 122",
		Status:    "testing",
		Priority:  "impoortant",
		CreatedAt: time.Now().Format(time.UnixDate),
		CreatedBy: "Sam",
		DueDate:   "08.10.2021",
	}

	taskL.Create(task1)

	// taskL.Delete(12)

	/*cont1 := contacts.Contact{
		Id:        10,
		FirstName: "Sam",
		LastName:  "Smith",
		Phone:     "(695)-175-4661",
		Email:     "sam@local.com",
	}

	cont2 := contacts.Contact{
		Id:        11,
		FirstName: "Eugene",
		LastName:  "Williamson",
		Phone:     "(139)-191-0039",
		Email:     "eugene@local.com",
	}*/

	contL := contacts.ContactList{ContactsDb: db.Db}

	// contL.Create(cont1)
	// contL.Create(cont2)

	contL.Delete(11)
	fmt.Println(contL.GetAll())

	fmt.Println("Finish!")

	defer db.Db.Close()
}
