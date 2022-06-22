package inspector

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/inspector2"
	"github.com/aws/aws-sdk-go-v2/service/inspector2/types"
	"testing"
)

// mockCreateFilterAPI is the function definition that implements the interface
type mockCreateFilterAPI func(ctx context.Context, params *inspector2.CreateFilterInput, optFns ...func(*inspector2.Options)) (*inspector2.CreateFilterOutput, error)


func (m mockCreateFilterAPI) CreateFilter(ctx context.Context, params *inspector2.CreateFilterInput, optFns ...func(*inspector2.Options)) (*inspector2.CreateFilterOutput, error) {
	testARN := "mytestARN"
	output := &inspector2.CreateFilterOutput{
		Arn: &testARN,
	}
	return output, nil
}

func TestGetFilterOnTypeCategory (t *testing.T){
	expectedValue := "Package Vulnerability"
	testCases:= []struct{
		name string
		typeCategory string
		typeCategoryComparisonOperator string
		expected types.StringFilter
	}{
		{
			name: "TestEqualsSuccess",
			typeCategory: "Package Vulnerability",
			typeCategoryComparisonOperator: "EQUALS",
			expected: types.StringFilter{Comparison: "EQUALS",Value: &expectedValue},
		},
	}

	for _, tc := range testCases{
		actual := getFilterOnTypeCategory(tc.typeCategory,tc.typeCategoryComparisonOperator)
		if *actual.Value != *tc.expected.Value {
			t.Errorf("Error expected %s but got %s",*tc.expected.Value,*actual.Value)
		}
	}
}


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

// TesAWSCreateFilter tests the API call to the Inspector Clients CreateFilter function
func TestAWSCreateFilter (t *testing.T){
    expectedARN := "mytestARN"
	// Create Test Pipeline
	testPipeline := InspectorFilterPipeline{}

	// Create the mockClient and context
	var mockClient mockCreateFilterAPI
    ctx:=context.TODO()

    // Callout to the SUT passing in the mock client
	testPipeline.ProcessFilterRequest(mockClient,ctx)
    actualResult := *testPipeline.FilterResponse.Arn
    if actualResult != expectedARN {
    	t.Errorf("Error unexpected ARN expecting %s got %s", expectedARN,*testPipeline.FilterResponse.Arn)
	}
}

// TestCreateTypeCategoryFilterRequest test creation of

func TestCreateTypeCategoryFilterRequest(t *testing.T){
	categoryType := getFilterOnTypeCategory("","EQUALS")
	categoryTypes := []types.StringFilter{}
	categoryTypes = append(categoryTypes, categoryType)
    testCases := []struct {
    	name string
    	input InspectorFilterPipeline
    	want inspector2.CreateFilterInput
	}{
		{
			name: "Test",
			input: InspectorFilterPipeline{FilterName: "TestCategoryTypeFilter",Action: "SUPPRESS",TypeCategoryFilters: categoryTypes},
			want: inspector2.CreateFilterInput{},
		},
	}

	for _, tc := range testCases {
		actual := tc.input.CreateTypeCategoryFilterRequest()
		if actual.Action != tc.input.Action {
			t.Errorf("Expected %s, but got %s", actual.Action, tc.input.Action)
		}

		if len(actual.TypeCategoryFilters) == 0 {
			t.Errorf("Expected %d got %d", len(tc.input.TypeCategoryFilters), len(actual.TypeCategoryFilters))
		}
	}
}


func TestCreateVulnerabilityIdFilterRequest(t *testing.T){
	testCases := []struct {
		name string
		input InspectorFilterPipeline
		want inspector2.CreateFilterInput
	}{
		{
			name: "Test",
			input: InspectorFilterPipeline{FilterName: "TestCategoryTypeFilter",Action: "SUPPRESS",CVETitles: []string{"CVE12345"}},
			want: inspector2.CreateFilterInput{Action: types.FilterActionSuppress},
		},
	}

	for _, tc := range testCases {
		actual := tc.input.CreateVulnerabilityIdFilterRequest()
		if actual.Action != tc.want.Action{
			t.Errorf("Wanted %s got %s",tc.want.Action,actual.Action)
		}
	}
}