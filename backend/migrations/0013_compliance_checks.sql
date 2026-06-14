CREATE TABLE IF NOT EXISTS compliance_checks (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    compliance_framework_id UUID NOT NULL REFERENCES compliance_frameworks(id) ON DELETE CASCADE,
    data_asset_id UUID REFERENCES data_assets(id) ON DELETE CASCADE,
    policy_id UUID REFERENCES policies(id) ON DELETE SET NULL,
    check_key TEXT NOT NULL,
    title TEXT NOT NULL,
    status TEXT NOT NULL,
    severity TEXT NOT NULL DEFAULT 'medium',
    evidence JSONB NOT NULL DEFAULT '{}'::jsonb,
    evaluated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT compliance_checks_status_check CHECK (status IN ('pass', 'fail', 'warning', 'not_applicable')),
    CONSTRAINT compliance_checks_severity_check CHECK (severity IN ('low', 'medium', 'high', 'critical'))
);

CREATE INDEX IF NOT EXISTS compliance_checks_framework_idx ON compliance_checks (tenant_id, compliance_framework_id, status);
CREATE INDEX IF NOT EXISTS compliance_checks_asset_idx ON compliance_checks (tenant_id, data_asset_id);
