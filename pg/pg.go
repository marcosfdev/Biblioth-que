package pg

import (
	"context"

	_ "github.com/lib/pq" // required
)

// Repository is the application's data layer functionality.
type Repository interface {
	// agent queries
	CreateAgent(ctx context.Context, arg CreateAgentParams) (Agent, error)
	DeleteAgent(ctx context.Context, id int64) (Agent, error)
	GetAgent(ctx context.Context, id int64) (Agent, error)
	ListAgents(ctx context.Context) ([]Agent, error)
	UpdateAgent(ctx context.Context, arg UpdateAgentParams) (Agent, error)

	// author queries
	CreateAuthor(ctx context.Context, arg CreateAuthorParams) (Author, error)
	DeleteAuthor(ctx context.Context, id int64) (Author, error)
	GetAuthor(ctx context.Context, id int64) (Author, error)
	ListAuthors(ctx context.Context) ([]Author, error)
	UpdateAuthor(ctx context.Context, arg UpdateAuthorParams) (Author, error)

	// book queries
	CreateBook(ctx context.Context, bookArg CreateBookParams, authorIDs []int64) (*Book, error)
	UpdateBook(ctx context.Context, bookArg UpdateBookParams, authorIDs []int64) (*Book, error)
	DeleteBook(ctx context.Context, id int64) (Book, error)
	GetBook(ctx context.Context, id int64) (Book, error)
	ListBooks(ctx context.Context) ([]Book, error)
}
