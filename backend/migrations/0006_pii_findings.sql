CREATE TABLE IF NOT EXISTS pii_findings (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    data_asset_id UUID NOT NULL REFERENCES data_assets(id) ON DELETE CASCADE,
    data_field_id UUID REFERENCES data_fields(id) ON DELETE SET NULL,
    pii_type TEXT NOT NULL,
    detector TEXT NOT NULL,
    confidence NUMERIC(5,4) NOT NULL,
    match_count INTEGER NOT NULL DEFAULT 0,
    sample_hashes JSONB NOT NULL DEFAULT '[]'::jsonb,
    evidence JSONB NOT NULL DEFAULT '{}'::jsonb,
    status TEXT NOT NULL DEFAULT 'open',
    detected_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    reviewed_at TIMESTAMPTZ,
    CONSTRAINT pii_findings_confidence_check CHECK (confidence >= 0 AND confidence <= 1),
    CONSTRAINT pii_findings_status_check CHECK (status IN ('open', 'confirmed', 'false_positive', 'remediated'))
);

CREATE INDEX IF NOT EXISTS pii_findings_asset_idx ON pii_findings (tenant_id, data_asset_id);
CREATE INDEX IF NOT EXISTS pii_findings_type_idx ON pii_findings (tenant_id, pii_type, status);
