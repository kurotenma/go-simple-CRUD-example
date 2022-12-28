new:
	- go mod init golang-ecommerce-example
	- go mod tidy
	- go get github.com/labstack/echo/v4
	- go get github.com/jackc/pgx/v5/pgxpool
	- go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	- go install github.com/cosmtrek/air@latest
	- go get github.com/go-playground/validator/v10
	- go get github.com/joho/godotenv

database:
	- migrate -database "postgres://root:root@localhost:5436/go-simple-crud-example?sslmode=disable" -path db/migrations down
	- migrate -database "postgres://root:root@localhost:5436/go-simple-crud-example?sslmode=disable" -path db/migrations up

docker:
	- sudo docker compose down
	- sudo docker volume rm go-simple-crud-example-db
	- sudo docker compose build
	- sudo docker compose up -d
