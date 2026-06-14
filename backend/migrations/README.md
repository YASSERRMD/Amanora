# Database Migrations

Amanora migrations are plain SQL files applied in lexical order. Each migration
uses an `NNNN_description.sql` file name and should be safe to run against a new
PostgreSQL database during local development.

Phase 2 introduces the core governance metadata model. Later phases can add
indexes, constraints, and workflow-specific tables without rewriting migration
history.
