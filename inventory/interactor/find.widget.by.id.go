package interactor

import (
	"context"
	"github.com/jackc/pgx/v4"
	"inventory/graph/model"
)

func FindWidgetById(ctx context.Context, conn *pgx.Conn, id int) (*model.Widget, error) {
	var inventory int
	err := conn.QueryRow(ctx, "select inventory from inventory where widget_id = $1", id).Scan(&inventory)
	if err != nil {
		return nil, err
	}
	return &model.Widget{
		ID:        id,
		Inventory: inventory,
	}, nil
}
