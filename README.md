# prueba REST Go + CockroachDB


### Requisitos
- Go 1.22+
- CockroachDB 24.x (o Cloud)
- (Opcional) golang-migrate


### Pasos rápidos
1. `cp .env.example .env`
2. `docker compose up --build -d`
3. `make migrate-up`
4. `make run`


### Endpoints
- GET /health
- GET /v1/users
- POST /v1/users
- GET /v1/users/{id}
- GET /v1/jobs
- POST /v1/jobs
- POST /v1/jobs/{id}/run