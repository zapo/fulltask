package schema

import (
	"github.com/graphql-go/graphql"
	"todoapi/models"
)

type TaskList struct {
	Data     []*models.Task
	Metadata ListMetadata
}

func NewTaskList(data []*models.Task, total int) TaskList {
	return TaskList{
		data,
		ListMetadata{total},
	}
}

var TaskType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Task",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "Identifier of the task",
		},
		"title": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "Short title of the task",
		},
		"description": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "Long description of the task",
		},
		"priority": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "Priority",
		},
		"done": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Boolean),
			Description: "Priority",
		},
	},
})

var CreateTask = graphql.Field{
	Type:        TaskType,
	Description: "Create a new Task",
	Args: graphql.FieldConfigArgument{
		"title": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"description": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"priority": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		title, _ := params.Args["title"].(string)
		description, _ := params.Args["description"].(string)
		priority, _ := params.Args["priority"].(int)

		return models.InsertTask(&models.Task{
			Title:       title,
			Description: description,
			Priority:    priority,
			Done:        false,
		})
	},
}

var DeleteTask = graphql.Field{
	Type:        TaskType,
	Description: "Delete a Task",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		id := params.Args["id"].(int)
		return models.Task{}, models.DeleteTask(id)
	},
}

var UpdateTask = graphql.Field{
	Type:        TaskType,
	Description: "Update a Task",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"title": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"description": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"priority": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		id := params.Args["id"].(int)
		title := params.Args["title"].(string)
		description, _ := params.Args["description"].(string)
		priority, _ := params.Args["priority"].(int)
		return models.UpdateTask(&models.Task{
			Id:          id,
			Title:       title,
			Description: description,
			Priority:    priority,
			Done:        false,
		})
	},
}

var ListTasks = graphql.Field{
	Type:        NewListType(TaskType),
	Description: "List of tasks",
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		tasks, err := models.AllTasks()
		return NewTaskList(tasks, 0), err
	},
}

var FindTask = graphql.Field{
	Type:        TaskType,
	Description: "Find a task",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		id := params.Args["id"].(int)
		return models.FindTask(id)
	},
}
