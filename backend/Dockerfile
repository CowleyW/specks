# syntax=docker/dockerfile:1

FROM golang:1.19

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
COPY . ./

ENV DATABASE_ADDRESS=db-prod:3306

RUN CGO_ENABLED=0 GOOS=linux go build -o /go-backend

EXPOSE 4000

# Run
CMD ["/go-backend"]
