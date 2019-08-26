package schema

import (
	"github.com/graphql-go/graphql"
)

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"tasks": &ListTasks,
		"task":  &FindTask,
	},
})
