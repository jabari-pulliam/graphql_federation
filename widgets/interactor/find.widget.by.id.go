package interactor

import (
	"context"
	"github.com/jackc/pgx/v4"
	"widgets/graph/model"
)

func FindWidgetById(context context.Context, conn *pgx.Conn, id int) (*model.Widget, error) {
	var color string
	var size int
	err := conn.QueryRow(context, "select color, size from widgets where id = $1", id).Scan(&color, &size)
	if err == pgx.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &model.Widget{
		ID:    id,
		Size:  size,
		Color: model.WidgetColor(color),
	}, nil
}
