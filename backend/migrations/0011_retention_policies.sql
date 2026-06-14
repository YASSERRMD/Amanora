CREATE TABLE IF NOT EXISTS retention_policies (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    description TEXT,
    retention_days INTEGER NOT NULL,
    action TEXT NOT NULL DEFAULT 'review',
    legal_hold BOOLEAN NOT NULL DEFAULT false,
    criteria JSONB NOT NULL DEFAULT '{}'::jsonb,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT retention_policies_days_check CHECK (retention_days > 0),
    CONSTRAINT retention_policies_action_check CHECK (action IN ('review', 'archive', 'delete', 'anonymize')),
    UNIQUE (tenant_id, name)
);

CREATE TABLE IF NOT EXISTS data_asset_retention_policies (
    tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    data_asset_id UUID NOT NULL REFERENCES data_assets(id) ON DELETE CASCADE,
    retention_policy_id UUID NOT NULL REFERENCES retention_policies(id) ON DELETE CASCADE,
    assigned_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    PRIMARY KEY (tenant_id, data_asset_id, retention_policy_id)
);
