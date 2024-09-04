# Makefile for managing Docker containers

# Target to start Docker containers
# Usage: make up
up:
	docker compose up

# Target to stop Docker containers
# Usage: make down
down:
	docker compose down
