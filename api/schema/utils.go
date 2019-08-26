package schema

import (
	"github.com/graphql-go/graphql"
)

type ListMetadata struct {
	total int
}

var ListMetadataType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ListMetadata",
	Fields: graphql.Fields{
		"total": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
})

func NewListNonNull(of *graphql.Object) *graphql.NonNull {
	return graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(of)))
}

func NewListType(of *graphql.Object) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: of.Name() + "List",
		Fields: graphql.Fields{
			"data": &graphql.Field{
				Type:        NewListNonNull(of),
				Description: "List of " + of.Name(),
			},
			"metadata": &graphql.Field{
				Type: graphql.NewNonNull(ListMetadataType),
			},
		},
	})
}
