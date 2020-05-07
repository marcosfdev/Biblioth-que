package gqlgen

import (
	"context"

	"github.com/marcosfdev/bibliotheque/pg"
)

type agentResolver struct{ *Resolver }

func (r *agentResolver) Authors(ctx context.Context, obj *pg.Agent) ([]pg.Author, error) {
	panic("not implemented")
}

// Resolver connects individual resolvers with the datalayer.
type Resolver struct {
	Repository pg.Repository
}
