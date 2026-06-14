# Known Limitations

- Runtime services use in-memory stores for API workflows.
- SQL migrations define the production model, but repositories are interfaces or memory-backed implementations.
- Graph lineage is memory-backed; Neo4j is available in Compose but not wired as a repository.
- Kafka-compatible infrastructure is present, but event publishing is represented by in-memory hooks.
- Authentication is API-key based for local development.
- Frontend screens are functional console views with static/demo rows, not fully connected to all API endpoints.
