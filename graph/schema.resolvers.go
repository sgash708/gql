package graph

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/sgash708/gql/graph/generated"
	"github.com/sgash708/gql/graph/model"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	// TODO: applicationを呼び出して処理
	// fmt.Println("CreateTodo")
	// return nil, nil

	todo := &model.Todo{
		ID:   fmt.Sprintf("T%d", rand.Int()),
		Text: input.Text,
		Done: false,
		User: &model.User{
			ID:   input.UserID,
			Name: "user:" + input.UserID,
		},
	}
	r.todos = append(r.todos, todo)

	return todo, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	// TODO: applicationを呼び出して処理
	// fmt.Println("Todos")
	// return nil, nil

	return r.todos, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
