package interactor

import (
	"context"
	"github.com/jackc/pgx/v4"
	"prices/graph/model"
)

func FindWidgetById(ctx context.Context, conn *pgx.Conn, id int) (*model.Widget, error) {
	var price int
	err := conn.QueryRow(ctx, "select price from prices where widget_id = $1", id).Scan(&price)
	if err != nil {
		return nil, err
	}

	return &model.Widget{
		ID:    id,
		Price: price,
	}, nil
}
