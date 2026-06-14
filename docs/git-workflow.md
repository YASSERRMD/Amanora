# Git Workflow

## Required Author Identity

Before any commit, configure git with the project author identity:

```bash
git config user.name "YASSERRMD"
git config user.email "arafath.yasser@gmail.com"
```

All commits must be authored as:

```text
YASSERRMD <arafath.yasser@gmail.com>
```

Do not use automation, assistant, bot, or vendor names as the commit author.

## Phase Branches

Each phase starts from the latest `main`:

```bash
git checkout main
git pull origin main
git checkout -b phase-XX-short-name
```

Use short, descriptive phase branch names such as:

```text
phase-00-planning
phase-01-foundation
phase-02-data-model
```

## Atomic Commits

Commit after each atomic task. Commit messages use this format:

```text
phase <number>: <small completed task>
```

Examples:

```text
phase 1: initialize go backend module
phase 1: add backend health endpoint
phase 2: add data source schema
phase 5: add email detector
```

## Validation

Run validation commands after meaningful changes and before phase completion.
Use the commands relevant to the files changed in the phase.

Backend:

```bash
go fmt ./...
go test ./...
go vet ./...
```

Frontend:

```bash
npm run lint
npm run build
```

Docker:

```bash
docker compose -f deploy/docker-compose.yml config
```

Documentation-only phases should at minimum verify repository status and inspect
the resulting files.

## Author Verification

Before pushing a phase branch, verify recent commit authors:

```bash
git log --format='%an <%ae>' -n 20
```

The expected author for project commits is:

```text
YASSERRMD <arafath.yasser@gmail.com>
```

## Phase Completion

Push once at the end of the phase:

```bash
git push -u origin phase-XX-short-name
```

Then:

1. Create a pull request.
2. Merge the pull request into `main`.
3. Delete the completed phase branch.
4. Pull the latest `main`.
5. Start the next phase from updated `main`.

## Reporting

At the end of each phase, report:

- Branch name
- Commits created
- Validation commands executed
- Push status
- Pull request status
- Merge status
- Branch deletion status
- Any incomplete items
