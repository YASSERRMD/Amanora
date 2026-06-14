CREATE TABLE IF NOT EXISTS lineage_edges (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    from_node_id UUID NOT NULL REFERENCES lineage_nodes(id) ON DELETE CASCADE,
    to_node_id UUID NOT NULL REFERENCES lineage_nodes(id) ON DELETE CASCADE,
    edge_type TEXT NOT NULL,
    operation TEXT,
    confidence NUMERIC(5,4) NOT NULL DEFAULT 1,
    properties JSONB NOT NULL DEFAULT '{}'::jsonb,
    observed_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT lineage_edges_type_check CHECK (edge_type IN ('ingests', 'transforms', 'depends_on', 'contains', 'produces', 'reads', 'writes')),
    CONSTRAINT lineage_edges_confidence_check CHECK (confidence >= 0 AND confidence <= 1),
    UNIQUE (tenant_id, from_node_id, to_node_id, edge_type)
);

CREATE INDEX IF NOT EXISTS lineage_edges_from_idx ON lineage_edges (tenant_id, from_node_id);
CREATE INDEX IF NOT EXISTS lineage_edges_to_idx ON lineage_edges (tenant_id, to_node_id);
