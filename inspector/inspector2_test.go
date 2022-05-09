package inspector

import (
	"testing"
)

func TestInspectorFilterPipeline_PopulateAccountFilters(t *testing.T) {
	filterPipeline := InspectorFilterPipeline{
		AWSAccounts: []string{"12345348035"},
	}
	filterPipeline.PopulateAccountFilters("EQUALS")
	if filterPipeline.FilterError != nil {
		t.Errorf("Error populating account filters %s", filterPipeline.FilterError.Error())
	}
	if len(filterPipeline.AccountFilters) != 1 {
		t.Errorf("Error populating account filters expecting 1 got %d", len(filterPipeline.AWSAccounts))
	}
	if *filterPipeline.AccountFilters[0].Value != "12345348035" {
		t.Errorf("Error populating account filters account expecting 12345348035 got %s",
			*filterPipeline.AccountFilters[0].Value)
	}
}

func TestInspectorFilterPipeline_PopulateTitleFilters(t *testing.T) {
	filterPipeline := InspectorFilterPipeline{
		CVETitles: []string{"CVE-12345"},
	}
	filterPipeline.PopulateTitleFilters("EQUALS")
	if filterPipeline.FilterError != nil {
		t.Errorf("Error populating title filters %s", filterPipeline.FilterError.Error())
	}
	if len(filterPipeline.CVETitleFilters) != 1 {
		t.Errorf("Error populating account filters expecting 1 got %d", len(filterPipeline.CVETitles))
	}
	if *filterPipeline.CVETitleFilters[0].Value != "CVE-12345" {
		t.Errorf("Error populating account filters account expecting CVE-12345 got %s",
			*filterPipeline.CVETitleFilters[0].Value)
	}
}