package tasklist

import (
	"fmt"
	"postgre_task/db"
	"testing"
)

var (
	testTasks = []Task{
		{
			Id:        10,
			Name:      "Task 10",
			Status:    "done",
			Priority:  "Important",
			CreatedAt: "04.10.2021",
			CreatedBy: "05.10.2021",
			DueDate:   "08.10.2021",
		},
		{
			Id:        11,
			Name:      "Task 11",
			Status:    "testing",
			Priority:  "medium",
			CreatedAt: "04.10.2021",
			CreatedBy: "",
			DueDate:   "09.10.2021",
		},
		{
			Id:        12,
			Name:      "Task 12",
			Status:    "Initial",
			Priority:  "medium",
			CreatedAt: "04.10.2021",
			CreatedBy: "",
			DueDate:   "10.10.2021",
		},
	}

	updTask = Task{
		Id:        11,
		Name:      "Task 10",
		Status:    "done",
		Priority:  "Important",
		CreatedAt: "04.10.2021",
		CreatedBy: "06.10.2021",
		DueDate:   "09.10.2021",
	}
	tlt = new(TaskList)
)

func TestCreateTask(t *testing.T) {
	tlt.TasksDb = db.Db
	var del string

	for i, v := range testTasks {
		err := tlt.Create(v)
		if err != nil {
			t.Error(err)
		}

		if i == len(testTasks)-1 {
			del += fmt.Sprintf("%d", v.Id)
		} else {
			del += fmt.Sprintf("%d, ", v.Id)
		}
	}

	t.Cleanup(func() {
		delQuery := fmt.Sprintf("DELETE FROM tasks WHERE task_id IN (%s)", del)
		_, err := tlt.TasksDb.Exec(delQuery)
		tlt = new(TaskList)
		if err != nil {
			t.Error(err)
		}
	})
}

func TestUpdateTask(t *testing.T) {
	TestCreateTask(t)

	err := tlt.Update(updTask)
	if err != nil {
		t.Error(err)
	}
}

func TestGetTask(t *testing.T) {
	TestCreateTask(t)

	for _, v := range testTasks {
		temp, err := tlt.Get(v.Id)
		if err != nil {
			t.Error(err)
		}
		if temp.Id != v.Id {
			t.Error("method get failed")
		}
	}
}

func TestGetAllTask(t *testing.T) {
	TestCreateTask(t)

	temp, err := tlt.GetAll()
	if err != nil {
		t.Error(err)
	}
	for i := range testTasks {
		if testTasks[i].Id != temp[i].Id {
			t.Error("failed getall method")
		}
	}
}

func TestGetDeleteTask(t *testing.T) {
	TestCreateTask(t)

	for _, v := range testTasks {
		err := tlt.Delete(v.Id)
		if err != nil {
			t.Error(err)
		}
	}
}
