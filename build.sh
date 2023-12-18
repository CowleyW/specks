#! /usr/bin/bash

if [ "$1" == "--prod" ]; then
	docker compose --project-name specks up --build
elif [ "$1" == "--dev" ] || [ "$1" == "" ]; then
	docker compose -f docker-compose.dev.yml --project-name specks-dev up --build
else
	echo "Invalid build type: $1"
	exit 1;
fi
