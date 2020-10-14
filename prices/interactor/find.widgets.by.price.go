package interactor

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"prices/graph/model"
	"strconv"
	"strings"
)

func FindWidgetsByPrice(ctx context.Context, conn *pgx.Conn, ids []int, filter model.PriceFilter) (*model.PriceWidgetSource, error) {
	rows, err := conn.Query(ctx, buildQuery(ids, filter))
	if err != nil {
		return nil, err
	}

	var widgets []*model.Widget
	var destIds []int
	for rows.Next() {
		var id int
		var price int

		err = rows.Scan(&id, &price)
		if err != nil {
			return nil, err
		}

		widgets = append(widgets, &model.Widget{ID: id, Price: price})
		destIds = append(destIds, id)
	}

	return &model.PriceWidgetSource{
		Widgets: &model.WidgetPage{
			Items: widgets,
			PageInfo: &model.PageInfo{
				TotalCount: len(widgets),
			},
		},
		Ids: destIds,
	}, nil
}

func buildQuery(ids []int, filter model.PriceFilter) string {
	sb := strings.Builder{}

	sb.WriteString("select prices.widget_id, prices.price from prices inner join unnest(ARRAY[")
	idb := strings.Builder{}
	idLen := len(ids)
	for i := 0; i < idLen; i++ {
		idb.WriteString(strconv.Itoa(ids[i]))

		if i < idLen-1 {
			idb.WriteString(",")
		}
	}
	sb.WriteString(idb.String())
	sb.WriteString("]) as wid on wid = prices.widget_id where ")

	var clauses []string

	if filter.MinPrice != nil {
		clauses = append(clauses, fmt.Sprintf("prices.price >= %d", *filter.MinPrice))
	}

	if filter.MaxPrice != nil {
		clauses = append(clauses, fmt.Sprintf("prices.price <= %d", *filter.MaxPrice))
	}

	sb.WriteString(strings.Join(clauses, " and "))

	return sb.String()
}
