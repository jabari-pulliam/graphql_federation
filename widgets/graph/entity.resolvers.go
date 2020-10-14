package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"widgets/graph/generated"
	"widgets/graph/model"
	"widgets/interactor"
)

func (r *entityResolver) FindWidgetByID(ctx context.Context, id int) (*model.Widget, error) {
	return interactor.FindWidgetById(ctx, r.DbConn, id)
}

func (r *entityResolver) FindWidgetSourceByIds(ctx context.Context, ids []int) (*model.WidgetSource, error) {
	panic(fmt.Errorf("not implemented"))
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
