package schema

import (
	"github.com/graphql-go/graphql"
)

var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"createTask": &CreateTask,
		"deleteTask": &DeleteTask,
		"updateTask": &UpdateTask,
	},
})
