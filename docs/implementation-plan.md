# Implementation Plan

## Phase 0: Repository Inspection and Planning

Goal: document the starting state and establish the architecture, implementation
plan, and git workflow.

Atomic tasks:

1. Inspect repository structure.
2. Document current state.
3. Add project overview to README.
4. Add architecture documentation.
5. Add implementation plan.
6. Add git workflow documentation.
7. Add Codex attribution to README.

## Phase 1: Project Foundation

Goal: create the backend, frontend, and local infrastructure foundation.

Atomic tasks:

1. Initialize the Go backend module.
2. Add command structure for API and workers.
3. Add API, discovery, classifier, policy, and retention command entry points.
4. Initialize the Next.js frontend.
5. Add backend and frontend Dockerfiles.
6. Add Docker Compose for PostgreSQL, Redis, Neo4j, and Kafka-compatible
   eventing.
7. Add environment template and Makefile.
8. Add backend health endpoint and frontend landing page.

## Phase 2: Core Governance Data Model

Goal: create the metadata schema and repository contracts.

Entities:

- Tenant
- DataSource
- DataAsset
- DataField
- DataClassification
- PiiFinding
- DataOwner
- DataSteward
- LineageNode
- LineageEdge
- Policy
- PolicyRule
- RetentionPolicy
- ComplianceFramework
- ComplianceCheck
- RiskScore
- AuditEvent

## Phase 3: Data Source and Connector Framework

Goal: build connector abstractions for data discovery.

Initial connectors:

- PostgreSQL
- CSV files
- MySQL placeholder
- Oracle placeholder
- REST API placeholder

Initial endpoints:

- `POST /api/v1/datasources`
- `GET /api/v1/datasources`
- `GET /api/v1/datasources/{id}`
- `POST /api/v1/datasources/{id}/test`

## Phase 4: Data Discovery Engine

Goal: discover schemas, tables, columns, files, and sample values.

Initial endpoints:

- `POST /api/v1/discovery/jobs`
- `GET /api/v1/discovery/jobs`
- `GET /api/v1/discovery/jobs/{id}`
- `POST /api/v1/discovery/jobs/{id}/run`

## Phase 5: PII Detection and Classification

Goal: detect and classify sensitive data.

Initial detectors:

- Email
- Phone number
- Credit card placeholder
- Passport placeholder
- National ID placeholder
- Name placeholder
- Address placeholder
- Free-text sensitive pattern placeholder

Initial endpoints:

- `POST /api/v1/classification/jobs`
- `GET /api/v1/classification/jobs`
- `GET /api/v1/classification/findings`
- `POST /api/v1/classification/run`

## Phase 6: Data Catalog

Goal: build a searchable catalog of governed data assets.

Initial endpoints:

- `GET /api/v1/catalog/assets`
- `GET /api/v1/catalog/assets/{id}`
- `GET /api/v1/catalog/fields/{id}`
- `GET /api/v1/catalog/search`
- `POST /api/v1/catalog/assets/{id}/tags`

## Phase 7: Metadata Graph and Lineage

Goal: build graph-based lineage and impact analysis.

Initial endpoints:

- `GET /api/v1/lineage/assets/{id}`
- `GET /api/v1/lineage/assets/{id}/upstream`
- `GET /api/v1/lineage/assets/{id}/downstream`
- `POST /api/v1/lineage/edges`

## Phase 8: Governance Policy Engine

Goal: implement policy-based governance decisions.

Policy examples:

- PII data must have an owner.
- High-risk data must have a retention policy.
- Public data cannot contain passport numbers.
- Sensitive data must not be stored in non-approved sources.
- Assets without a steward are non-compliant.

Initial endpoints:

- `POST /api/v1/policies`
- `GET /api/v1/policies`
- `POST /api/v1/policies/{id}/evaluate`
- `POST /api/v1/policies/evaluate-all`

## Phase 9: Retention Policy Engine

Goal: manage data lifecycle and retention rules.

Initial endpoints:

- `POST /api/v1/retention/policies`
- `GET /api/v1/retention/policies`
- `POST /api/v1/retention/evaluate`
- `POST /api/v1/retention/simulate`

## Phase 10: Compliance Checks and Risk Scoring

Goal: create compliance framework mapping and asset risk scores.

Initial frameworks:

- GDPR-style privacy checks
- UAE PDPL-style privacy checks placeholder
- India DPDP-style privacy checks placeholder
- Internal enterprise policy checks

Initial endpoints:

- `GET /api/v1/compliance/frameworks`
- `POST /api/v1/compliance/run`
- `GET /api/v1/compliance/results`
- `GET /api/v1/risk/assets`

## Phase 11: Frontend Foundation

Goal: build the Amanora admin console shell.

Pages:

- `/dashboard`
- `/datasources`
- `/discovery`
- `/catalog`
- `/classification`
- `/lineage`
- `/policies`
- `/retention`
- `/compliance`
- `/risk`
- `/audit`
- `/settings`

## Phase 12: Frontend Governance Screens

Goal: build product screens for governance workflows.

Screens:

- Data source registration
- Discovery jobs
- Catalog search
- Asset detail
- PII findings
- Lineage viewer placeholder
- Policy management
- Retention
- Compliance dashboard
- Risk dashboard
- Audit log

## Phase 13: Security and Multi-Tenancy

Goal: harden tenant isolation and governance access.

Scope:

- Tenant middleware
- API key authentication
- Role placeholders
- Tenant-scoped queries
- Secure headers
- Rate limiting
- Audit middleware
- Sensitive field redaction
- Security tests

## Phase 14: Sample Data and Demo Flow

Goal: provide a realistic demo experience.

Scope:

- Sample customer, HR, and finance datasets
- CSV connector examples
- Sample PII findings
- Sample lineage graph
- Sample policies
- Demo setup command
- Demo walkthrough documentation

## Phase 15: Final Documentation and Polish

Goal: make the repository complete and easy to evaluate.

Scope:

- README update
- Architecture diagram
- Discovery flow diagram
- PII classification flow
- Metadata graph diagram
- Policy engine diagram
- Local setup guide
- API examples
- Roadmap
- Known limitations
- Final verification notes

## Validation Strategy

Backend phases will run:

```bash
go fmt ./...
go test ./...
go vet ./...
```

Frontend phases will run:

```bash
npm run lint
npm run build
```

Infrastructure phases will run:

```bash
docker compose -f deploy/docker-compose.yml config
```

Full local environment validation will use Docker Compose once all required
services and images exist.
