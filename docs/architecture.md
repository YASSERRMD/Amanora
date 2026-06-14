# Amanora Architecture

## Architecture Goals

Amanora is an autonomous data governance platform that discovers enterprise
data, classifies sensitive content, catalogs metadata, tracks lineage, evaluates
governance policies, and records audit evidence.

The architecture is organized around independently testable domain modules and
worker processes so discovery, classification, policy evaluation, retention, and
audit workflows can scale separately.

## System Context

```text
Data sources
  ├─ PostgreSQL
  ├─ CSV files
  ├─ MySQL placeholder
  ├─ Oracle placeholder
  └─ REST API placeholder
        │
        ▼
Discovery workers ── events ── Classification workers
        │                         │
        ▼                         ▼
PostgreSQL metadata store   PII findings and risk scores
        │                         │
        ├──────── Metadata graph ─┤
        │                         │
        ▼                         ▼
Policy, retention, and compliance engines
        │
        ▼
REST API and governance dashboard
```

## Backend

The backend will be written in Go and split into command entry points:

- `api`: REST API for the governance console and integrations
- `discovery-worker`: source scanning, metadata extraction, and sampling
- `classifier-worker`: PII detection and classification execution
- `policy-worker`: policy evaluation and compliance decision processing
- `retention-worker`: lifecycle evaluation and retention violation detection

Domain code will live under `backend/internal/` with package boundaries for
configuration, HTTP transport, database access, connectors, discovery,
classification, catalog, lineage, policy, retention, compliance, risk, audit,
and tenant isolation.

## Data Stores

- PostgreSQL stores tenants, data sources, assets, fields, classifications,
  policies, compliance results, risk scores, retention rules, and audit events.
- Redis stores cache entries and short-lived policy evaluation state.
- A graph-ready lineage abstraction represents metadata relationships and can be
  backed by Neo4j or another graph implementation.
- Kafka-compatible eventing coordinates discovery, classification, policy, and
  retention workflows.

## Frontend

The frontend will be a Next.js TypeScript admin console with Tailwind CSS. It
will provide operational views for:

- Dashboard
- Data sources
- Discovery jobs
- Catalog search and asset details
- Classification findings
- Lineage
- Policies
- Retention
- Compliance
- Risk
- Audit
- Settings

Charts and operational summaries will use Recharts or ECharts.

## Event Flow

1. A user registers a data source through the API.
2. A discovery job is created and published to the event stream.
3. Discovery workers inspect schemas, tables, fields, files, and samples.
4. Discovery results are stored in PostgreSQL and lineage edges are synced to
   the graph abstraction.
5. Classification jobs evaluate samples and metadata for PII.
6. Risk, compliance, retention, and policy checks consume metadata and findings.
7. Every governance action writes audit events.
8. The dashboard reads curated API views for operational visibility.

## Multi-Tenancy and Security

Amanora will enforce tenant scoping at the API, service, repository, and query
layers. Authentication starts with API keys and role placeholders, then expands
to finer-grained governance access controls. Audit logging, sensitive field
redaction, rate limiting, and secure headers will be added during security
hardening.

## Deployment Model

Local development will use Docker Compose with PostgreSQL, Redis, Neo4j, and a
Kafka-compatible broker. Backend and frontend Dockerfiles will support the same
service boundaries used in local development and CI.
