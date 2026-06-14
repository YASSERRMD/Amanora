CREATE TABLE IF NOT EXISTS policies (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    description TEXT,
    category TEXT NOT NULL,
    severity TEXT NOT NULL DEFAULT 'medium',
    enabled BOOLEAN NOT NULL DEFAULT true,
    document JSONB NOT NULL DEFAULT '{}'::jsonb,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT policies_category_check CHECK (category IN ('privacy', 'ownership', 'retention', 'lineage', 'security', 'quality')),
    CONSTRAINT policies_severity_check CHECK (severity IN ('low', 'medium', 'high', 'critical')),
    UNIQUE (tenant_id, name)
);

CREATE TABLE IF NOT EXISTS policy_rules (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    policy_id UUID NOT NULL REFERENCES policies(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    condition JSONB NOT NULL,
    effect TEXT NOT NULL,
    message TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT policy_rules_effect_check CHECK (effect IN ('allow', 'deny', 'warn', 'require_review'))
);

CREATE INDEX IF NOT EXISTS policies_tenant_category_idx ON policies (tenant_id, category, enabled);
