package database

import (
	"database/sql"
	"fmt"
	"h8-assignment-2/infra/config"
	"log"

	_ "github.com/lib/pq"
)

var (
	db  *sql.DB
	err error
)

func handleDatabaseConnection() {
	appConfig := config.GetAppConfig()
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		appConfig.DBHost, appConfig.DBPort, appConfig.DBUser, appConfig.DBPassword, appConfig.DBName,
	)

	db, err = sql.Open(appConfig.Dialect, psqlInfo)

	if err != nil {
		log.Panic("error occured while trying to validate database arguments:", err)
	}

	err = db.Ping()

	if err != nil {
		log.Panic("error occured while trying to connect to database:", err)
	}

}

func handleCreateRequiredTables() {
	orderTable := `
		CREATE TABLE IF NOT EXISTS "orders" (
			order_id SERIAL PRIMARY KEY,
			customer_name VARCHAR(255) NOT NULL,
			ordered_at timestamptz DEFAULT now(),
			created_at timestamptz DEFAULT now(),
			updated_at timestamptz DEFAULT now()
		);
	`

	itemTable := `
		CREATE TABLE IF NOT EXISTS "items" (
			item_id SERIAL PRIMARY KEY,
			item_code VARCHAR(191) NOT NULL,
			quantity int NOT NULL,
			description TEXT NOT NULL,
			order_id int NOT NULL,
			created_at timestamptz DEFAULT now(),
			updated_at timestamptz DEFAULT now(),
			CONSTRAINT items_order_id_fk
            FOREIGN KEY(order_id) 
                REFERENCES orders(order_id)
                    ON DELETE CASCADE
		);	
	`

	_, err = db.Exec(orderTable)

	if err != nil {
		log.Panic("error occured while trying to create order table:", err)
	}

	_, err = db.Exec(itemTable)

	if err != nil {
		log.Panic("error occured while trying to create item table:", err)
	}
}

func InitiliazeDatabase() {
	handleDatabaseConnection()
	handleCreateRequiredTables()
}

func GetDatabaseInstance() *sql.DB {
	if db == nil {
		log.Panic("database instance is still nill!!!")
	}

	return db
}
