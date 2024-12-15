# Initialize up and down .sql files
migrate create -ext=sql -dir=migrations -seq init

# Command to up of our migrations
migrate -path=migrations -database "postgresql://book_shelf_user:book_shelf_pass@localhost:5480/book_shelf_db?sslmode=disable" -verbose up

# Command to down of our migrations
migrate -path=migrations -database "postgresql://book_shelf_user:book_shelf_pass@localhost:5480/book_shelf_db?sslmode=disable" -verbose down