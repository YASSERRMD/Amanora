CREATE TABLE IF NOT EXISTS data_classifications (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    data_asset_id UUID REFERENCES data_assets(id) ON DELETE CASCADE,
    data_field_id UUID REFERENCES data_fields(id) ON DELETE CASCADE,
    classification TEXT NOT NULL,
    sensitivity_level TEXT NOT NULL,
    confidence NUMERIC(5,4) NOT NULL DEFAULT 0,
    classifier TEXT NOT NULL,
    evidence JSONB NOT NULL DEFAULT '{}'::jsonb,
    classified_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT data_classifications_target_check CHECK (
        (data_asset_id IS NOT NULL AND data_field_id IS NULL)
        OR (data_asset_id IS NULL AND data_field_id IS NOT NULL)
    ),
    CONSTRAINT data_classifications_sensitivity_check CHECK (sensitivity_level IN ('public', 'internal', 'confidential', 'restricted')),
    CONSTRAINT data_classifications_confidence_check CHECK (confidence >= 0 AND confidence <= 1)
);

CREATE INDEX IF NOT EXISTS data_classifications_tenant_idx ON data_classifications (tenant_id, classification);
CREATE INDEX IF NOT EXISTS data_classifications_field_idx ON data_classifications (tenant_id, data_field_id);
