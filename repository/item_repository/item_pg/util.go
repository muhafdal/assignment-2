package item_pg

import "fmt"

func (itemPG *itemPG) generateGetItemsByCodesQuery(itemAmount int) string {
	baseQuery := `SELECT "item_id", "item_code", "quantity", "description", "order_id", "created_at" FROM "items"
	WHERE "item_code" IN`

	statement := "("

	for i := 1; i <= itemAmount; i++ {

		if i == itemAmount {
			statement += fmt.Sprintf("$%d)", i)
			break
		}

		statement += fmt.Sprintf("$%d,", i)

	}

	return fmt.Sprintf("%s %s", baseQuery, statement)
}
