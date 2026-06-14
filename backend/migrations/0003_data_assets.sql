CREATE TABLE IF NOT EXISTS data_assets (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    data_source_id UUID NOT NULL REFERENCES data_sources(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    schema_name TEXT,
    asset_type TEXT NOT NULL,
    fully_qualified_name TEXT NOT NULL,
    description TEXT,
    row_count BIGINT,
    size_bytes BIGINT,
    discovered_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT data_assets_type_check CHECK (asset_type IN ('database', 'schema', 'table', 'view', 'file', 'api_resource')),
    UNIQUE (tenant_id, fully_qualified_name)
);

CREATE INDEX IF NOT EXISTS data_assets_tenant_source_idx ON data_assets (tenant_id, data_source_id);
CREATE INDEX IF NOT EXISTS data_assets_type_idx ON data_assets (tenant_id, asset_type);
