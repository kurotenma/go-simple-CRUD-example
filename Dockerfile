# Start from golang base image
#FROM golang:alpine as base
#FROM base as dev
#
#RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
#
#WORKDIR /golang-ecommerce-example
#CMD ["air"]
#
#
#COPY . .
#EXPOSE 3000
#CMD ["migrate", "database", "postgres://root:root@localhost:5434/golang-ecommerce-example-sslmode=disable", "-path db/migrations", "down"]
#CMD ["migrate", "database", "postgres://root:root@localhost:5434/golang-ecommerce-example-sslmode=disable", "-path db/migrations", "up"]
#CMD ["/golang-ecommerce-example"]