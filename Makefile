up:
	docker-compose -f deploy/docker-compose.yml up -d postgres redis
ps:
	docker-compose -f deploy/docker-compose.yml ps -a
down:
	docker-compose -f deploy/docker-compose.yml down
rebuild:
	docker-compose -f deploy/docker-compose.yml up --build --force-recreate
reload:
	cd web/backend && air -c air.toml
test:
	cd web/backend && go test ./internal/service
dev:
	cd web/frontend && npm i && npm run dev
