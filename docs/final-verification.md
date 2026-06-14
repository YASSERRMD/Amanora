# Final Verification

Validation commands used across the project:

```bash
cd backend && go fmt ./... && go test ./... && go vet ./...
cd frontend && npm run lint && npm run build && npm audit --audit-level=moderate
docker compose -f deploy/docker-compose.yml config
```

Final expected local entry points:

- Frontend: `http://localhost:3000`
- API health: `http://localhost:8080/healthz`
- Docker Compose: `deploy/docker-compose.yml`

The README includes the required attribution:

```text
Built with OpenAI Codex.
```
