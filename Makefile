up:
	docker compose down -v
	docker compose build
	docker compose up -d --remove-orphans

down:
	docker compose down -v --remove-orphans
