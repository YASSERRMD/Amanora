# Amanora

Amanora is an autonomous data governance platform for data discovery, PII classification, metadata lineage, retention governance, compliance checks, and policy-based data control.

## Purpose

Amanora is designed as data governance infrastructure for enterprise teams that
need to discover, classify, catalog, govern, and audit data across many systems.

## Core Capabilities

- Data discovery across databases, files, and API-backed sources
- PII detection and classification workflows
- Searchable data catalog with ownership metadata
- Metadata graph and lineage analysis
- Retention policy management
- Compliance checks and evidence capture
- Policy evaluation for governance decisions
- Risk scoring for assets and findings
- Audit trail for governance activity
- Governance dashboard for operators and stewards

## Planned Stack

- Backend: Go, REST APIs, PostgreSQL, Redis, graph-ready lineage abstraction, and Kafka-compatible eventing
- Frontend: Next.js, TypeScript, Tailwind CSS, and charting with Recharts or ECharts
- DevOps: Docker, Docker Compose, and GitHub pull request workflow

## Repository Status

This repository contains the Amanora platform foundation, including Go backend
services, a Next.js admin console, SQL migrations, Docker Compose infrastructure,
demo data, and governance documentation.

## Local Setup

```bash
cp .env.example .env
make frontend-install
make validate
make compose-up
```

Frontend: `http://localhost:3000`

API health: `http://localhost:8080/healthz`

Protected API requests use:

```text
X-API-Key: dev-api-key
X-Tenant-ID: default
```

## Documentation

- [Architecture](docs/architecture.md)
- [Implementation Plan](docs/implementation-plan.md)
- [Git Workflow](docs/git-workflow.md)
- [Demo Walkthrough](docs/demo-walkthrough.md)
- [API Examples](docs/api-examples.md)
- [Roadmap](docs/roadmap.md)
- [Known Limitations](docs/known-limitations.md)
- [Final Verification](docs/final-verification.md)

Built with OpenAI Codex.
