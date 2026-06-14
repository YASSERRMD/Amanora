CREATE TABLE IF NOT EXISTS data_fields (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    data_asset_id UUID NOT NULL REFERENCES data_assets(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    ordinal_position INTEGER,
    data_type TEXT NOT NULL,
    nullable BOOLEAN NOT NULL DEFAULT true,
    description TEXT,
    sample_values JSONB NOT NULL DEFAULT '[]'::jsonb,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE (tenant_id, data_asset_id, name)
);

CREATE INDEX IF NOT EXISTS data_fields_asset_idx ON data_fields (tenant_id, data_asset_id);
CREATE INDEX IF NOT EXISTS data_fields_name_idx ON data_fields (tenant_id, name);
