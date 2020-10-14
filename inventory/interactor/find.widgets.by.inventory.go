package interactor

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"inventory/graph/model"
	"strconv"
	"strings"
)

func FindWidgetsByInventory(ctx context.Context, conn *pgx.Conn, ids []int, filter model.InventoryFilter) (*model.InventoryWidgetSource, error) {
	rows, err := conn.Query(ctx, buildQuery(ids, filter))
	if err != nil {
		return nil, err
	}

	var widgets []*model.Widget
	var destIds []int
	for rows.Next() {
		var id int
		var inventory int

		err = rows.Scan(&id, &inventory)
		if err != nil {
			return nil, err
		}

		widgets = append(widgets, &model.Widget{ID: id, Inventory: inventory})
		destIds = append(destIds, id)
	}

	return &model.InventoryWidgetSource{
		Widgets: &model.WidgetPage{
			Items: widgets,
			PageInfo: &model.PageInfo{
				TotalCount: len(widgets),
			},
		},
		Ids: destIds,
	}, nil
}

func buildQuery(ids []int, filter model.InventoryFilter) string {
	sb := strings.Builder{}

	sb.WriteString("select inventory.widget_id, inventory.inventory from inventory inner join unnest(ARRAY[")
	idb := strings.Builder{}
	idLen := len(ids)
	for i := 0; i < idLen; i++ {
		idb.WriteString(strconv.Itoa(ids[i]))

		if i < idLen-1 {
			idb.WriteString(",")
		}
	}
	sb.WriteString(idb.String())
	sb.WriteString("]) as wid on wid = inventory.widget_id where ")
	sb.WriteString(fmt.Sprintf("inventory.inventory >= %d", filter.MinInventory))

	return sb.String()
}
