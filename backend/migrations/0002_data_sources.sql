CREATE TABLE IF NOT EXISTS data_sources (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    source_type TEXT NOT NULL,
    connection_uri TEXT,
    configuration JSONB NOT NULL DEFAULT '{}'::jsonb,
    status TEXT NOT NULL DEFAULT 'registered',
    last_tested_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT data_sources_type_check CHECK (source_type IN ('postgres', 'mysql', 'oracle', 'csv', 'rest')),
    CONSTRAINT data_sources_status_check CHECK (status IN ('registered', 'healthy', 'degraded', 'disabled')),
    UNIQUE (tenant_id, name)
);

CREATE INDEX IF NOT EXISTS data_sources_tenant_type_idx ON data_sources (tenant_id, source_type);
