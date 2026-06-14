package discovery

import "github.com/YASSERRMD/Amanora/backend/internal/datasource"

type SamplingService struct{}

func NewSamplingService() SamplingService {
	return SamplingService{}
}

func (SamplingService) BuildSamples(assets []datasource.AssetMetadata) []SampleSet {
	return BuildSampleSets(assets)
}
