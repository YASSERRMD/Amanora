CREATE TABLE IF NOT EXISTS lineage_nodes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    data_asset_id UUID REFERENCES data_assets(id) ON DELETE CASCADE,
    node_key TEXT NOT NULL,
    node_type TEXT NOT NULL,
    label TEXT NOT NULL,
    properties JSONB NOT NULL DEFAULT '{}'::jsonb,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT lineage_nodes_type_check CHECK (node_type IN ('source', 'asset', 'field', 'job', 'transformation', 'report')),
    UNIQUE (tenant_id, node_key)
);

CREATE INDEX IF NOT EXISTS lineage_nodes_tenant_type_idx ON lineage_nodes (tenant_id, node_type);
