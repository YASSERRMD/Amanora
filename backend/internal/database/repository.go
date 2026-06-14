package database

import (
	"context"
	"time"
)

type ID string

type Tenant struct {
	ID        ID
	Name      string
	Slug      string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type DataSource struct {
	ID            ID
	TenantID      ID
	Name          string
	SourceType    string
	ConnectionURI string
	Status        string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type DataAsset struct {
	ID                 ID
	TenantID           ID
	DataSourceID       ID
	Name               string
	SchemaName         string
	AssetType          string
	FullyQualifiedName string
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

type DataField struct {
	ID              ID
	TenantID        ID
	DataAssetID     ID
	Name            string
	OrdinalPosition int
	DataType        string
	Nullable        bool
}

type DataClassification struct {
	ID               ID
	TenantID         ID
	DataAssetID      ID
	DataFieldID      ID
	Classification   string
	SensitivityLevel string
	Confidence       float64
	Classifier       string
	ClassifiedAt     time.Time
}

type PiiFinding struct {
	ID          ID
	TenantID    ID
	DataAssetID ID
	DataFieldID ID
	PiiType     string
	Detector    string
	Confidence  float64
	MatchCount  int
	Status      string
	DetectedAt  time.Time
}

type LineageNode struct {
	ID       ID
	TenantID ID
	NodeKey  string
	NodeType string
	Label    string
}

type LineageEdge struct {
	ID         ID
	TenantID   ID
	FromNodeID ID
	ToNodeID   ID
	EdgeType   string
	Confidence float64
	ObservedAt time.Time
}

type Policy struct {
	ID       ID
	TenantID ID
	Name     string
	Category string
	Severity string
	Enabled  bool
}

type RetentionPolicy struct {
	ID            ID
	TenantID      ID
	Name          string
	RetentionDays int
	Action        string
	LegalHold     bool
}

type ComplianceFramework struct {
	ID           ID
	TenantID     ID
	Name         string
	Code         string
	Jurisdiction string
	Enabled      bool
}

type ComplianceCheck struct {
	ID       ID
	TenantID ID
	CheckKey string
	Title    string
	Status   string
	Severity string
}

type RiskScore struct {
	ID           ID
	TenantID     ID
	DataAssetID  ID
	Score        int
	Level        string
	CalculatedAt time.Time
}

type AuditEvent struct {
	ID         ID
	TenantID   ID
	ActorType  string
	ActorID    string
	Action     string
	EntityType string
	EntityID   ID
	Outcome    string
	OccurredAt time.Time
}

type TenantRepository interface {
	CreateTenant(ctx context.Context, tenant Tenant) (Tenant, error)
	GetTenant(ctx context.Context, id ID) (Tenant, error)
	GetTenantBySlug(ctx context.Context, slug string) (Tenant, error)
}

type DataSourceRepository interface {
	CreateDataSource(ctx context.Context, source DataSource) (DataSource, error)
	ListDataSources(ctx context.Context, tenantID ID) ([]DataSource, error)
	GetDataSource(ctx context.Context, tenantID ID, id ID) (DataSource, error)
}

type MetadataRepository interface {
	UpsertDataAsset(ctx context.Context, asset DataAsset) (DataAsset, error)
	UpsertDataField(ctx context.Context, field DataField) (DataField, error)
	ListAssets(ctx context.Context, tenantID ID) ([]DataAsset, error)
	ListFields(ctx context.Context, tenantID ID, assetID ID) ([]DataField, error)
}

type ClassificationRepository interface {
	SaveClassification(ctx context.Context, classification DataClassification) (DataClassification, error)
	SavePiiFinding(ctx context.Context, finding PiiFinding) (PiiFinding, error)
	ListPiiFindings(ctx context.Context, tenantID ID) ([]PiiFinding, error)
}

type LineageRepository interface {
	UpsertLineageNode(ctx context.Context, node LineageNode) (LineageNode, error)
	UpsertLineageEdge(ctx context.Context, edge LineageEdge) (LineageEdge, error)
	ListUpstream(ctx context.Context, tenantID ID, nodeID ID) ([]LineageEdge, error)
	ListDownstream(ctx context.Context, tenantID ID, nodeID ID) ([]LineageEdge, error)
}

type GovernanceRepository interface {
	ListPolicies(ctx context.Context, tenantID ID) ([]Policy, error)
	ListRetentionPolicies(ctx context.Context, tenantID ID) ([]RetentionPolicy, error)
	ListComplianceFrameworks(ctx context.Context, tenantID ID) ([]ComplianceFramework, error)
	SaveComplianceCheck(ctx context.Context, check ComplianceCheck) (ComplianceCheck, error)
	SaveRiskScore(ctx context.Context, score RiskScore) (RiskScore, error)
}

type AuditRepository interface {
	RecordAuditEvent(ctx context.Context, event AuditEvent) (AuditEvent, error)
	ListAuditEvents(ctx context.Context, tenantID ID) ([]AuditEvent, error)
}
