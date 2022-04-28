package inspector

import (
	"github.com/aws/aws-sdk-go-v2/service/inspector2"
	"github.com/aws/aws-sdk-go-v2/service/inspector2/types"
)

type InspectorFilterPipeline struct {
    AWSAccounts     []string
    CVETitles       []string
    AccountFilters  []types.StringFilter
    CVETitleFilters []types.StringFilter
    FilterRequest   *inspector2.CreateFilterInput
    FilterResponse  *inspector2.CreateFilterOutput
    Action          types.FilterAction
    FilterName      string
    FilterError     error
}

// getFilterOnAWSAccount creates a filter on account
func getFilterOnAWSAccount(awsAccount string, accountComparisonOperator string) types.StringFilter {
    filterType := types.StringFilter{
        Comparison: types.StringComparison(accountComparisonOperator),
        Value:      &awsAccount,
    }
    return filterType
}

func (i *InspectorFilterPipeline) PopulateTitleFilters(comparisonOperator string) *InspectorFilterPipeline {
    var titleFilters []types.StringFilter
    if len(i.CVETitles) > 0 {
        for _, cveTitle := range i.CVETitles {
            accountFilter := getFilterOnAWSAccount(cveTitle, comparisonOperator)
            titleFilters = append(titleFilters, accountFilter)
        }
    }
    i.CVETitleFilters = titleFilters
    return i
}

func (i *InspectorFilterPipeline) PopulateAccountFilters(comparisonOperator string) *InspectorFilterPipeline {
    var accountFilters []types.StringFilter
    if len(i.AWSAccounts) > 0 {
        for _, awsAccount := range i.AWSAccounts {
            accountFilter := getFilterOnAWSAccount(awsAccount, comparisonOperator)
            accountFilters = append(accountFilters, accountFilter)
        }
    }
    i.AccountFilters = accountFilters
    return i
}