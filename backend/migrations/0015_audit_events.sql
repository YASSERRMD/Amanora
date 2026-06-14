CREATE TABLE IF NOT EXISTS audit_events (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    actor_type TEXT NOT NULL,
    actor_id TEXT,
    action TEXT NOT NULL,
    entity_type TEXT NOT NULL,
    entity_id UUID,
    outcome TEXT NOT NULL,
    metadata JSONB NOT NULL DEFAULT '{}'::jsonb,
    occurred_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT audit_events_actor_type_check CHECK (actor_type IN ('user', 'api_key', 'worker', 'system')),
    CONSTRAINT audit_events_outcome_check CHECK (outcome IN ('success', 'failure', 'warning'))
);

CREATE INDEX IF NOT EXISTS audit_events_tenant_time_idx ON audit_events (tenant_id, occurred_at DESC);
CREATE INDEX IF NOT EXISTS audit_events_entity_idx ON audit_events (tenant_id, entity_type, entity_id);
