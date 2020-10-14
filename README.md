# GraphQL and Apollo Federation for Querying Microservices
Apollo GraphQL Federation allows you to stitch together the schemas of independent web services
to achieve a separation of concerns. But what happens when you have complex queries that involve
not just requesting data from multiple services but also filtering based on properties of that data
which may be owned by separate services? Even using Federation, each query/subquery is resolved
by a single service, so filtering items based on properties of an entity that are owned by different
services would require the resolver for the query to have direct access to the API or
datasource that contains each property.

This example shows how combining Federation with a composite pattern can enable these types of queries
in a declarative rather than programmatic way.

## Candidate Schema
Let's take a candidate schema which we want to break up across multiple federated microservices.

```graphql
type PageInfo {
  totalCount: Int!
}

type WidgetPage {
  items: [Widget!]!
  pageInfo: PageInfo!
}

enum WidgetColor {
  RED
  GREEN
  BLUE
  ORANGE
  BLACK
}

type Widget {
  id: Int!
  size: Int!
  color: WidgetColor!
  price: Int!
  inventory: Int!
}

input WidgetFilter {
  colors: [WidgetColor!]
  minSize: Int
  maxSize: Int
}

type Query {
  widgets(filter: WidgetFilter): WidgetSource!
}
```

Now what if we had `prices` and `inventory` services which knew about the price and inventory stock
of each widget. With Apollo Federation, we could split this graph into the federated schemas in 
[widgets](widgets/graph/schema.graphqls), [prices](prices/graph/schema.graphqls), and 
[inventory](inventory/graph/schema.graphqls).

The key to making filter queries work across the different services are the types with the pattern `*WidgetSource`
which encapsulate a page of realized results and a set of keys resolved by the query. Services that 
need to add queryable fields can extend the `*WidgetSource` from an external query which allows them
to be chained together in any order.

An example that filters by `color` and `inventory` and `price`:
```graphql
{
  widgets(filter: {colors: [BLACK, BLUE, GREEN]}) {
    widgetsByInventory(filter: {minInventory: 5}) {
      widgetsByPrice(filter: {minPrice: 0, maxPrice: 20}) {       
        widgets {
          pageInfo {
            totalCount
          }
          items {
            id
            price
            inventory
          }
      	}
      }
    }
  }
}
```

With this pattern, these filters could be chained in any order and still yield the same result.