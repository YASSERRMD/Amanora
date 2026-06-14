INSERT INTO tenants (id, name, slug, status)
VALUES ('00000000-0000-0000-0000-000000000001', 'Amanora Demo', 'amanora-demo', 'active')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO data_sources (id, tenant_id, name, source_type, status, configuration)
VALUES (
    '00000000-0000-0000-0000-000000000101',
    '00000000-0000-0000-0000-000000000001',
    'Customer Warehouse',
    'postgres',
    'healthy',
    '{"host":"postgres","database":"warehouse"}'::jsonb
) ON CONFLICT (tenant_id, name) DO NOTHING;

INSERT INTO data_assets (id, tenant_id, data_source_id, name, schema_name, asset_type, fully_qualified_name, row_count, discovered_at)
VALUES (
    '00000000-0000-0000-0000-000000000201',
    '00000000-0000-0000-0000-000000000001',
    '00000000-0000-0000-0000-000000000101',
    'customers',
    'public',
    'table',
    'postgres.customer_warehouse.public.customers',
    125000,
    now()
) ON CONFLICT (tenant_id, fully_qualified_name) DO NOTHING;

INSERT INTO data_fields (id, tenant_id, data_asset_id, name, ordinal_position, data_type, nullable, sample_values)
VALUES
    ('00000000-0000-0000-0000-000000000301', '00000000-0000-0000-0000-000000000001', '00000000-0000-0000-0000-000000000201', 'customer_id', 1, 'uuid', false, '["cst_001"]'::jsonb),
    ('00000000-0000-0000-0000-000000000302', '00000000-0000-0000-0000-000000000001', '00000000-0000-0000-0000-000000000201', 'email', 2, 'text', false, '["sample@example.com"]'::jsonb),
    ('00000000-0000-0000-0000-000000000303', '00000000-0000-0000-0000-000000000001', '00000000-0000-0000-0000-000000000201', 'phone', 3, 'text', true, '["+971501234567"]'::jsonb)
ON CONFLICT (tenant_id, data_asset_id, name) DO NOTHING;

INSERT INTO data_classifications (tenant_id, data_field_id, classification, sensitivity_level, confidence, classifier, evidence)
VALUES
    ('00000000-0000-0000-0000-000000000001', '00000000-0000-0000-0000-000000000302', 'email_address', 'restricted', 0.9800, 'seed', '{"reason":"email sample"}'::jsonb),
    ('00000000-0000-0000-0000-000000000001', '00000000-0000-0000-0000-000000000303', 'phone_number', 'confidential', 0.9200, 'seed', '{"reason":"phone sample"}'::jsonb);

INSERT INTO data_owners (id, tenant_id, name, email, team)
VALUES ('00000000-0000-0000-0000-000000000401', '00000000-0000-0000-0000-000000000001', 'Data Platform Owner', 'owner@example.com', 'Data Platform')
ON CONFLICT (tenant_id, email) DO NOTHING;

INSERT INTO data_stewards (id, tenant_id, name, email, team)
VALUES ('00000000-0000-0000-0000-000000000402', '00000000-0000-0000-0000-000000000001', 'Privacy Steward', 'steward@example.com', 'Governance')
ON CONFLICT (tenant_id, email) DO NOTHING;

INSERT INTO data_asset_owners (tenant_id, data_asset_id, data_owner_id)
VALUES ('00000000-0000-0000-0000-000000000001', '00000000-0000-0000-0000-000000000201', '00000000-0000-0000-0000-000000000401')
ON CONFLICT DO NOTHING;

INSERT INTO data_asset_stewards (tenant_id, data_asset_id, data_steward_id)
VALUES ('00000000-0000-0000-0000-000000000001', '00000000-0000-0000-0000-000000000201', '00000000-0000-0000-0000-000000000402')
ON CONFLICT DO NOTHING;

INSERT INTO pii_findings (tenant_id, data_asset_id, data_field_id, pii_type, detector, confidence, match_count, evidence)
VALUES
    ('00000000-0000-0000-0000-000000000001', '00000000-0000-0000-0000-000000000201', '00000000-0000-0000-0000-000000000302', 'email', 'seed-email-detector', 0.9800, 250, '{"sample":"sample@example.com"}'::jsonb);

INSERT INTO compliance_frameworks (id, tenant_id, name, code, jurisdiction, description)
VALUES
    ('00000000-0000-0000-0000-000000000501', '00000000-0000-0000-0000-000000000001', 'GDPR-style Privacy Checks', 'GDPR_STYLE', 'EU', 'Baseline privacy governance checks'),
    ('00000000-0000-0000-0000-000000000502', '00000000-0000-0000-0000-000000000001', 'Internal Enterprise Policy', 'INTERNAL', 'Global', 'Internal data governance controls')
ON CONFLICT (tenant_id, code) DO NOTHING;

INSERT INTO policies (id, tenant_id, name, category, severity, document)
VALUES
    ('00000000-0000-0000-0000-000000000601', '00000000-0000-0000-0000-000000000001', 'PII data must have an owner', 'ownership', 'high', '{"if":{"classification":"pii"},"then":{"requires":"owner"}}'::jsonb),
    ('00000000-0000-0000-0000-000000000602', '00000000-0000-0000-0000-000000000001', 'High-risk data must have retention policy', 'retention', 'high', '{"if":{"risk":"high"},"then":{"requires":"retention_policy"}}'::jsonb)
ON CONFLICT (tenant_id, name) DO NOTHING;

INSERT INTO retention_policies (id, tenant_id, name, retention_days, action, criteria)
VALUES ('00000000-0000-0000-0000-000000000701', '00000000-0000-0000-0000-000000000001', 'Customer PII Review', 365, 'review', '{"classification":"pii"}'::jsonb)
ON CONFLICT (tenant_id, name) DO NOTHING;

INSERT INTO risk_scores (tenant_id, data_asset_id, score, level, factors)
VALUES ('00000000-0000-0000-0000-000000000001', '00000000-0000-0000-0000-000000000201', 72, 'high', '{"pii_findings":1,"owner_present":true}'::jsonb);

INSERT INTO audit_events (tenant_id, actor_type, actor_id, action, entity_type, entity_id, outcome, metadata)
VALUES ('00000000-0000-0000-0000-000000000001', 'system', 'seed', 'seed.loaded', 'tenant', '00000000-0000-0000-0000-000000000001', 'success', '{"phase":2}'::jsonb);
