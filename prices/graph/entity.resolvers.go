package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"prices/graph/generated"
	"prices/graph/model"
	"prices/interactor"
)

func (r *entityResolver) FindInventoryWidgetSourceByIds(ctx context.Context, ids []int) (*model.InventoryWidgetSource, error) {
	return &model.InventoryWidgetSource{
		Ids: ids,
	}, nil
}

func (r *entityResolver) FindPriceWidgetSourceByIds(ctx context.Context, ids []int) (*model.PriceWidgetSource, error) {
	return &model.PriceWidgetSource{
		Ids: ids,
	}, nil
}

func (r *entityResolver) FindWidgetByID(ctx context.Context, id int) (*model.Widget, error) {
	return interactor.FindWidgetById(ctx, r.DbConn, id)
}

func (r *entityResolver) FindWidgetSourceByIds(ctx context.Context, ids []int) (*model.WidgetSource, error) {
	return &model.WidgetSource{
		Ids: ids,
	}, nil
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
