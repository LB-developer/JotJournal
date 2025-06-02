.PHONY: setup migrate run dev frontend docker-dev lint test valkey-start

# === Setup ===

setup:
	cp server/.env.local.example server/.env
	cp client/.env.local.example client/.env
	@bash scripts/local_db_setup.sh

migrate:
	make -C server migrate-up

# === Local Dev ===

valkey-start:
	valkey-server &

backend:
	make -C server run

frontend:
	cd client && npm install && npm run dev

dev:
	@echo "Starting backend and frontend in dev mode..."
	make valkey-start & make backend & make frontend

# === Docker ===

docker-dev:
	cp server/.env.docker.example server/.env
	cp client/.env.docker.example client/.env
	docker compose --profile dev up --build -d
