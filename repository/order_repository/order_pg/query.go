package order_pg

const (
	createOrderQuery = `
		INSERT INTO "orders"
		("ordered_at", "customer_name")
		VALUES ($1, $2)

		RETURNING "order_id"
	`

	createItemQuery = `
		INSERT INTO "items"
		("item_code", "description", "quantity", "order_id")
		VALUES ($1, $2, $3, $4)
	`

	getOrdersWithItemsQuery = `
		SELECT "o"."order_id", "o"."customer_name", "o"."ordered_at", "o"."created_at", "o"."updated_at",
		"i"."item_id", "i"."item_code", "i"."quantity", "i"."description", "i"."order_id", "i"."created_at", "i"."updated_at"
		from "orders" as "o"
		LEFT JOIN "items" as "i" ON "o"."order_id" = "i"."order_id"
		ORDER BY "o"."order_id" ASC
	`

	getOrderById = `
		SELECT "order_id", "customer_name", "ordered_at", "created_at", "updated_at" FROM "orders"
		WHERE "order_id" = $1

	`

	updateOrderByIdQuery = `
		UPDATE "orders"
		SET "ordered_at" = $2,
		"customer_name" = $3
		WHERE "order_id" = $1
	`

	updateItemByCodeQuery = `
		UPDATE "items"
		SET "description" = $2,
		"quantity" = $3
		WHERE "item_code" = $1
	`

	deleteOrderQuery = `
	DELETE FROM "orders"
	WHERE "order_id" = $1
	`
)
