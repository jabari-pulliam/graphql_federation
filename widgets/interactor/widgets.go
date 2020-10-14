package interactor

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"log"
	"strings"
	"widgets/graph/model"
)

func FindWidgets(ctx context.Context, conn *pgx.Conn, filter *model.WidgetFilter) (*model.WidgetSource, error) {
	query := buildQuery(filter)

	log.Printf("Query: %s", query)

	rows, err := conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var widgets []*model.Widget
	var ids []int
	for rows.Next() {
		var id int
		var color string
		var size int
		err = rows.Scan(&id, &color, &size)
		if err != nil {
			return nil, err
		}
		widgets = append(widgets, &model.Widget{
			ID:    id,
			Color: model.WidgetColor(color),
			Size:  size,
		})
		ids = append(ids, id)
	}

	return &model.WidgetSource{
		Ids: ids,
		Widgets: &model.WidgetPage{
			Items: widgets,
			PageInfo: &model.PageInfo{
				TotalCount: len(widgets),
			},
		},
	}, nil
}

func buildQuery(filter *model.WidgetFilter) string {
	sb := strings.Builder{}
	sb.WriteString("select * from widgets")

	clauses := make([]string, 0)

	if len(filter.Colors) > 0 {
		cb := strings.Builder{}

		cb.WriteString("color in (")
		colors := make([]string, 0)
		for _, c := range filter.Colors {
			colors = append(colors, fmt.Sprintf("'%s'", c))
		}
		cb.WriteString(strings.Join(colors, ","))
		cb.WriteString(")")

		clauses = append(clauses, cb.String())
	}

	if filter.MinSize != nil {
		clauses = append(clauses, fmt.Sprintf("size >= %d", *filter.MinSize))
	}

	if filter.MaxSize != nil {
		clauses = append(clauses, fmt.Sprintf("size <= %d", *filter.MaxSize))
	}

	if len(clauses) > 0 {
		sb.WriteString(" where ")
		sb.WriteString(strings.Join(clauses, " and "))
	}

	return sb.String()
}
