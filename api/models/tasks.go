package models

type Task struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    int    `json:"priority"`
	Done        bool   `json:"done"`
}

func InsertTask(task *Task) (*Task, error) {
	err := db.QueryRow(
		"INSERT INTO tasks (title, description, priority) VALUES ($1, $2, $3) RETURNING id",
		task.Title, task.Description, task.Priority,
	).Scan(&task.Id)
	return task, err
}

func UpdateTask(task *Task) (*Task, error) {
	_, err := db.Exec(
		"UPDATE tasks set title=$1, description=$2, priority=$3, done=$4  WHERE id = $5",
		task.Title, task.Description, task.Priority, task.Done, task.Id,
	)
	return task, err
}

func DeleteTask(id int) error {
	_, err := db.Exec(
		"DELETE FROM tasks WHERE id = $1", id,
	)
	return err
}

func FindTask(id int) (*Task, error) {
	task := new(Task)
	err := db.QueryRow(
		"SELECT id, title, description, priority FROM tasks WHERE id = $1", id,
	).Scan(&task.Id, &task.Title, &task.Description, &task.Priority)
	return task, err
}

func AllTasks() ([]*Task, error) {
	rows, err := db.Query("SELECT id, title, description, priority, done FROM tasks")
	tasks := make([]*Task, 0)
	if err != nil {
		return tasks, err
	}

	for rows.Next() {
		task := new(Task)
		err := rows.Scan(
			&task.Id,
			&task.Title,
			&task.Description,
			&task.Priority,
			&task.Done,
		)

		if err != nil {
			return []*Task{}, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}
