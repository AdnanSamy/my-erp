migrate_up:
	migrate -path=migrations -database "postgresql://${MY_ERP_DB_USER}:${MY_ERP_DB_PASS}@${MY_ERP_DB_HOST}/my_erp?sslmode=disable" -verbose up
migrate_down:
	migrate -path=migrations -database "postgresql://${MY_ERP_DB_USER}:${MY_ERP_DB_PASS}@${MY_ERP_DB_HOST}/my_erp?sslmode=disable" -verbose down

