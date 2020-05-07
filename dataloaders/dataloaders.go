package dataloaders

//go:generate go run github.com/vektah/dataloaden AgentLoader int64 *github.com/marcosfdev/bibliotheque/pg.Agent
//go:generate dataloaden AuthorSliceLoader int64 []github.com/marcosfdev/bibliotheque/pg.Author

import (
	"context"
	"time"

	"github.com/marcosfdev/bibliotheque/pg"
)

type contextKey string

const key = contextKey("dataloaders")

// Loaders holds references to the individual dataloaders.
type Loaders struct {
	AgentByAuthorID *AgentLoader
}

func newLoaders(ctx context.Context, repo pg.Repository) *Loaders {
	return &Loaders{
		// individual loaders will be initialized here
	}
}

func newAgentByAuthorID(ctx context.Context, repo pg.Repository) *AgentLoader {
	return NewAgentLoader(AgentLoaderConfig{
		MaxBatch: 100,
		Wait:     5 * time.Millisecond,
		Fetch: func(authorIDs []int64) ([]*pg.Agent, []error) {
			// db query
			res, err := repo.ListAgentsByAuthorIDs(ctx, authorIDs)
			if err != nil {
				return nil, []error{err}
			}
			// map
			groupByAuthorID := make(map[int64]*pg.Agent, len(authorIDs))
			for _, r := range res {
				groupByAuthorID[r.AuthorID] = &pg.Agent{
					ID:    r.ID,
					Name:  r.Name,
					Email: r.Email,
				}
			}
			// order
			result := make([]*pg.Agent, len(authorIDs))
			for i, authorID := range authorIDs {
				result[i] = groupByAuthorID[authorID]
			}
			return result, nil
		},
	})
}
