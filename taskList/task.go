package tasklist

import (
	"database/sql"
	"fmt"
)

type Task struct {
	Id        int
	Name      string
	Status    string
	Priority  string
	CreatedAt string
	CreatedBy string
	DueDate   string
}

type TaskList struct {
	TasksDb *sql.DB
}

// Methods: create, update, get, getAll and delete

func (tl *TaskList) Create(t Task) error {
	_, err := tl.TasksDb.Exec(`INSERT INTO tasks
	(task_id, title, status, priority, created_at, created_by, due_date)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	`, t.Id, t.Name, t.Status, t.Priority, t.CreatedAt, t.CreatedBy, t.DueDate)

	return err
}

func (tl *TaskList) Update(t Task) error {
	res, err := tl.TasksDb.Exec(`UPDATE tasks 
			SET title = $1, status = $2, priority = $3, due_date = $4
			WHERE task_id = $5
			`, t.Name, t.Status, t.Priority, t.DueDate, t.Id)

	if err != nil {
		return err
	}

	if num, _ := res.RowsAffected(); num == 0 {
		return fmt.Errorf("task %d not exists", t.Id)
	}
	return nil
}

func (tl *TaskList) Get(id int) (Task, error) {
	row := tl.TasksDb.QueryRow("SELECT task_id, title, status, priority, created_at, created_by, due_date from tasks WHERE task_id = $1", id)

	var task, emptyTask Task
	row.Scan(&task.Id, &task.Name, &task.Status, &task.Priority, &task.CreatedAt, &task.CreatedBy, &task.DueDate)

	if task == emptyTask {
		return Task{}, fmt.Errorf("task %d not exists", id)
	}

	return task, nil
}

func (tl *TaskList) GetAll() ([]Task, error) {
	rows, err := tl.TasksDb.Query("SELECT task_id, title, status, priority, created_at, created_by, due_date from tasks ORDER BY task_id")
	if err != nil {
		return make([]Task, 0), err
	}
	var result []Task
	var tempTask Task

	for rows.Next() {
		rows.Scan(&tempTask.Id, &tempTask.Name, &tempTask.Status, &tempTask.Priority, &tempTask.CreatedAt, &tempTask.CreatedBy, &tempTask.DueDate)

		result = append(result, tempTask)
	}

	return result, nil

}

func (tl *TaskList) Delete(id int) {
	tl.TasksDb.QueryRow("DELETE FROM tasks WHERE task_id = $1", id)
}
