package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/graph-gophers/dataloader"
	"github.com/jackc/pgx/v4"
	"log"
	"prices/graph/model"
	"strconv"
	"strings"
	"time"
)

type WidgetLoaderKey string

const loaderKey = WidgetLoaderKey("widgetloaderkey")

func WidgetLoaderMiddleware(conn *pgx.Conn) gin.HandlerFunc {
	loadWidgets := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var results []*dataloader.Result

		sb := strings.Builder{}
		sb.WriteString("select prices.widget_id, prices.price from prices inner join unnest(ARRAY[")
		keysLen := len(keys)
		log.Printf("Loading %d keys", keysLen)
		for i := 0; i < keysLen; i++ {
			sb.WriteString(keys[i].String())

			if i < keysLen-1 {
				sb.WriteString(", ")
			}
		}
		sb.WriteString("]) as pid on pid = prices.widget_id")

		rows, err := conn.Query(ctx, sb.String())
		if err != nil {
			for range keys {
				results = append(results, &dataloader.Result{
					Error: err,
				})
			}
		} else {
			data := make(map[int]*dataloader.Result)
			for rows.Next() {
				var id int
				var price int
				err = rows.Scan(&id, &price)
				data[id] = &dataloader.Result{
					Data: &model.Widget{
						ID:    id,
						Price: price,
					},
					Error: err,
				}
			}

			for _, k := range keys.Keys() {
				i, _ := strconv.Atoi(k)
				results = append(results, data[i])
			}
		}

		return results
	}

	return func(c *gin.Context) {
		loader := dataloader.NewBatchedLoader(loadWidgets, dataloader.WithWait(50*time.Millisecond))
		ctx := context.WithValue(c.Request.Context(), loaderKey, loader)
		c.Request = c.Request.WithContext(ctx)
		log.Println("Created loader")

		c.Next()
	}
}

func GetWidgetLoader(ctx context.Context) *dataloader.Loader {
	return ctx.Value(loaderKey).(*dataloader.Loader)
}
