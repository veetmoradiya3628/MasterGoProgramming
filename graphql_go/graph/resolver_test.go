package graph

import (
	"context"
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/veetmoradiya3628/graph_go/graph/model"
	"github.com/veetmoradiya3628/graph_go/models"
)

func TestCreateTodo(t *testing.T) {

	store := models.NewTodoStore()
	resolver := &Resolver{TodoStore: store}

	input := model.NewTodo{Text: "Test todo"}

	todo, err := resolver.Mutation().CreateTodo(context.Background(), input)

	assert.NoError(t, err)
	assert.Equal(t, "Test todo", todo.Text)
	assert.False(t, todo.Done)
	assert.NotEmpty(t, todo.ID)

}

func TestCreateTodo_Validation_Errors(t *testing.T) {
	store := models.NewTodoStore()

	resolver := &Resolver{TodoStore: store}

	input := model.NewTodo{Text: ""}
	todo, err := resolver.Mutation().CreateTodo(context.Background(), input)

	assert.Nil(t, todo)
	assert.Error(t, err)
	assert.Equal(t, "validation failed for text: text cannot be empty", err.Error())

	input = model.NewTodo{Text: generateString(256)}
	todo, err = resolver.Mutation().CreateTodo(context.Background(), input)

	assert.Nil(t, todo)
	assert.Error(t, err)
	assert.Equal(t, "validation failed for text: text cannot exceed 255 characters", err.Error())

}

func TestUpdateTodo_Not_FoundWithID(t *testing.T) {
	store := models.NewTodoStore()

	resolver := &Resolver{TodoStore: store}

	input := model.UpdateTodo{Text: nil, Done: nil}
	todo, err := resolver.Mutation().UpdateTodo(context.Background(), "1", input)

	assert.Nil(t, todo)
	assert.Error(t, err)
	assert.Equal(t, "Todo with ID 1 not found", err.Error())
}

func TestToggleTodo_Not_FoundWithID(t *testing.T) {
	store := models.NewTodoStore()

	id := "1"
	resolver := &Resolver{TodoStore: store}
	todo, err := resolver.Mutation().ToggleTodo(context.Background(), id)

	assert.Nil(t, todo)
	assert.Error(t, err)
	assert.Equal(t, fmt.Sprintf("Todo with ID %v not found", id), err.Error())
}

func TestGetTodos(t *testing.T) {
	store := models.NewTodoStore()

	resolver := &Resolver{TodoStore: store}

	store.Create("Todo 1")
	store.Create("Todo 2")

	todos, err := resolver.Query().Todos(context.Background())
	assert.NoError(t, err)

	assert.Len(t, todos, 2)
}

func generateString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = 'a' + byte(rand.Intn(26))
	}
	return string(b)
}
