package catalog

import (
	"testing"

	"github.com/YASSERRMD/Amanora/backend/internal/datasource"
)

func TestCatalogFiltersAndTags(t *testing.T) {
	service := NewService()
	asset := service.UpsertAsset(Asset{
		AssetMetadata: datasource.AssetMetadata{
			Name:               "customers",
			FullyQualifiedName: "postgres.crm.public.customers",
			Fields: []datasource.FieldMetadata{
				{Name: "email", DataType: "text"},
			},
		},
		Owner:           "privacy",
		RiskLevel:       "high",
		Classifications: []string{"email"},
	})

	if got := service.ListAssets(SearchFilter{Query: "customer", Classification: "email", Owner: "privacy", Risk: "high"}); len(got) != 1 {
		t.Fatalf("expected filtered asset, got %d", len(got))
	}
	updated, err := service.AssignTags(asset.ID, []string{"pii", "pii", "customer"})
	if err != nil {
		t.Fatal(err)
	}
	if len(updated.Tags) != 2 {
		t.Fatalf("expected unique tags, got %v", updated.Tags)
	}
	if _, ok := service.GetField(asset.ID + "_field_001"); !ok {
		t.Fatal("expected field detail")
	}
}
