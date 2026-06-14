# Repository Structure

## Current State

Amanora starts from a greenfield repository with only the root license committed.
There is no application code, documentation set, deployment configuration, or
test harness yet.

## Existing Files

```text
.
├── .git/
└── LICENSE
```

## Planned Top-Level Layout

The platform will use a monorepo layout with separate backend, frontend,
deployment, documentation, and example data areas:

```text
.
├── backend/
├── frontend/
├── deploy/
├── docs/
├── examples/
├── LICENSE
└── README.md
```

## Initial Implementation Notes

- Backend services will be implemented in Go with command entry points for the
  API and worker processes.
- Frontend application code will live under a Next.js TypeScript app.
- Infrastructure definitions will live under `deploy/` and target local Docker
  Compose development first.
- Documentation will be built incrementally per phase before the system surface
  expands.
