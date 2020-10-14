package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"inventory/graph/generated"
	"inventory/graph/model"
	"inventory/interactor"
)

func (r *priceWidgetSourceResolver) WidgetsByInventory(ctx context.Context, obj *model.PriceWidgetSource, filter model.InventoryFilter) (*model.InventoryWidgetSource, error) {
	return interactor.FindWidgetsByInventory(ctx, r.DbConn, obj.Ids, filter)
}

func (r *widgetSourceResolver) WidgetsByInventory(ctx context.Context, obj *model.WidgetSource, filter model.InventoryFilter) (*model.InventoryWidgetSource, error) {
	return interactor.FindWidgetsByInventory(ctx, r.DbConn, obj.Ids, filter)
}

// PriceWidgetSource returns generated.PriceWidgetSourceResolver implementation.
func (r *Resolver) PriceWidgetSource() generated.PriceWidgetSourceResolver {
	return &priceWidgetSourceResolver{r}
}

// WidgetSource returns generated.WidgetSourceResolver implementation.
func (r *Resolver) WidgetSource() generated.WidgetSourceResolver { return &widgetSourceResolver{r} }

type priceWidgetSourceResolver struct{ *Resolver }
type widgetSourceResolver struct{ *Resolver }
