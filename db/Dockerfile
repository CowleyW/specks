# syntax=docker/dockerfile:1

FROM python:3.10-alpine AS build

WORKDIR /app

COPY . /app

RUN python3 builder.py

FROM mysql:8

WORKDIR /docker-entrypoint-initdb.d/
COPY --from=build /app/*.sql .

EXPOSE 3306

CMD ["mysqld"]
