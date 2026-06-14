CREATE TABLE IF NOT EXISTS risk_scores (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    data_asset_id UUID NOT NULL REFERENCES data_assets(id) ON DELETE CASCADE,
    score INTEGER NOT NULL,
    level TEXT NOT NULL,
    factors JSONB NOT NULL DEFAULT '{}'::jsonb,
    calculated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT risk_scores_score_check CHECK (score >= 0 AND score <= 100),
    CONSTRAINT risk_scores_level_check CHECK (level IN ('low', 'medium', 'high', 'critical'))
);

CREATE INDEX IF NOT EXISTS risk_scores_asset_idx ON risk_scores (tenant_id, data_asset_id, calculated_at DESC);
CREATE INDEX IF NOT EXISTS risk_scores_level_idx ON risk_scores (tenant_id, level, score DESC);
