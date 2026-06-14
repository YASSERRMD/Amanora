CREATE TABLE IF NOT EXISTS data_owners (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    email TEXT NOT NULL,
    team TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE (tenant_id, email)
);

CREATE TABLE IF NOT EXISTS data_stewards (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    email TEXT NOT NULL,
    team TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE (tenant_id, email)
);

CREATE TABLE IF NOT EXISTS data_asset_owners (
    tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    data_asset_id UUID NOT NULL REFERENCES data_assets(id) ON DELETE CASCADE,
    data_owner_id UUID NOT NULL REFERENCES data_owners(id) ON DELETE CASCADE,
    assigned_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    PRIMARY KEY (tenant_id, data_asset_id, data_owner_id)
);

CREATE TABLE IF NOT EXISTS data_asset_stewards (
    tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    data_asset_id UUID NOT NULL REFERENCES data_assets(id) ON DELETE CASCADE,
    data_steward_id UUID NOT NULL REFERENCES data_stewards(id) ON DELETE CASCADE,
    assigned_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    PRIMARY KEY (tenant_id, data_asset_id, data_steward_id)
);
