package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"widgets/graph/generated"
	"widgets/graph/model"
	"widgets/interactor"
)

func (r *queryResolver) Widgets(ctx context.Context, filter *model.WidgetFilter) (*model.WidgetSource, error) {
	return interactor.FindWidgets(ctx, r.DbConn, filter)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
