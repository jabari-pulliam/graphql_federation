package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"prices/graph/generated"
	"prices/graph/model"
	"prices/interactor"
)

func (r *inventoryWidgetSourceResolver) WidgetsByPrice(ctx context.Context, obj *model.InventoryWidgetSource, filter model.PriceFilter) (*model.PriceWidgetSource, error) {
	return interactor.FindWidgetsByPrice(ctx, r.DbConn, obj.Ids, filter)
}

func (r *widgetSourceResolver) WidgetsByPrice(ctx context.Context, obj *model.WidgetSource, filter model.PriceFilter) (*model.PriceWidgetSource, error) {
	return interactor.FindWidgetsByPrice(ctx, r.DbConn, obj.Ids, filter)
}

// InventoryWidgetSource returns generated.InventoryWidgetSourceResolver implementation.
func (r *Resolver) InventoryWidgetSource() generated.InventoryWidgetSourceResolver {
	return &inventoryWidgetSourceResolver{r}
}

// WidgetSource returns generated.WidgetSourceResolver implementation.
func (r *Resolver) WidgetSource() generated.WidgetSourceResolver { return &widgetSourceResolver{r} }

type inventoryWidgetSourceResolver struct{ *Resolver }
type widgetSourceResolver struct{ *Resolver }
