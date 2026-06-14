# API Examples

Use `X-API-Key: dev-api-key` for local requests.

```bash
curl -H "X-API-Key: dev-api-key" http://localhost:8080/api/v1/datasources
```

```bash
curl -X POST http://localhost:8080/api/v1/policies \
  -H "X-API-Key: dev-api-key" \
  -H "Content-Type: application/json" \
  -d '{"name":"PII must have owner","conditions":{"classification":"pii","owner":"present"},"effect":"warn"}'
```

```bash
curl -X POST http://localhost:8080/api/v1/retention/policies \
  -H "X-API-Key: dev-api-key" \
  -H "Content-Type: application/json" \
  -d '{"name":"Customer PII Review","retentionDays":365,"action":"review"}'
```
