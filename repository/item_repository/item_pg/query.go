package item_pg

const (
	getItemsByCodes = `
		SELECT * from "items" WHERE "item_code"
	`
)
