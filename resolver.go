package smartmei_backend

import (
	"context"
)

type Resolver struct{}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) TarifaAtual(ctx context.Context) (*Tarifa, error) {

}
