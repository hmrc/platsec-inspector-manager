package main
import (
	"context"
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/inspector2/types"
	"github.com/platsec-inspector-manager/auditing"
	"github.com/platsec-inspector-manager/clients"
	"github.com/platsec-inspector-manager/configmanagement"
	"github.com/platsec-inspector-manager/inspector"
	"github.com/platsec-inspector-manager/security"
	"os"
)

func main() {
	awsAccount := flag.String("account", "", "AWS account")
	username := flag.String("username", "", "AWS username")
	region := flag.String("region", "eu-west-2", "AWS region")
	profile := flag.String("profile", "", "AWS profile")
	action := flag.String("action", "SUPPRESS", "action to apply")
	filterName := flag.String("filter-name", "", "filter name")
	filterType := flag.String("filter-type", "cve", "type of filter")
	comparisonOperator := flag.String("comparison-operator", "EQUALS", "comparison operator")
	mfaToken := flag.String("mfa-token", "", "MFA token")
	vulnerabilityId := flag.String("vulnerability-id", "", "vulnerability ID (CVE-2021-3711)")
	flag.Parse()

	// Load in config from file
	configValues,err := configmanagement.InitConfig()
	if err != nil {
		os.Exit(1)
	}

	myUserInput := clients.UserInput{
		AwsAccount: *awsAccount,
		Username: *username,
		UserContext: context.TODO(),
		Region: *region,
		Account: configValues.Account,
		RoleName: configValues.RoleName,
		Profile: *profile,
		Action: *action,
		FilterName: *filterName,
		FilterType: *filterType,
		ComparisonOperator: *comparisonOperator,
		MfaToken: *mfaToken,
		VulnerabilityId: *vulnerabilityId,
		SessionDuration: 3600,
		SessionName: "inspector",
	}

	myUserInput.SetDefaultConfig()
	// Get Session Token
	stsFactory := clients.NewSTSClientFactory()
	stsClient := stsFactory(myUserInput.UserConfig)

	err = security.GetAWSSessionToken(&myUserInput, stsClient)

	if err != nil {
		auditing.Log(err.Error())
		os.Exit(1)
	}

	stsServiceFactory := clients.NewSTSClientSessionConfig()
    err = security.AssumeAccountRole(&myUserInput, stsServiceFactory, myUserInput.AwsAccount)
	if err!= nil {
		auditing.Log(err.Error())
		os.Exit(1)
	}

	inspectorFactory := clients.NewInspectorClientFactory()
    inspectorClient := inspectorFactory(myUserInput.UserConfig, *&myUserInput)

	filterPipeline := inspector.InspectorFilterPipeline{
        AWSAccounts: []string{myUserInput.AwsAccount},
        Action:      types.FilterAction(*action),
        FilterName:  myUserInput.FilterName,
        CVETitles:   []string{myUserInput.VulnerabilityId},
    }
	filterPipeline.PopulateAccountFilters(myUserInput.ComparisonOperator).CreateAccountFilterRequest().ProcessFilterRequest(inspectorClient, myUserInput.UserContext)
	if filterPipeline.FilterError != nil {
		fmt.Printf("Error processing pipeline %s", filterPipeline.FilterError.Error())
		auditing.Log(filterPipeline.FilterError.Error())
		fmt.Printf("Filter Output %s", *filterPipeline.FilterResponse.Arn)
	}	
}
	
// TODO: refactor below and add switch statement 

/* 	if logonCredentials.FilterType == "cve" {
		filterPipeline.PopulateTitleFilters(logonCredentials.ComparisonOperator).CreateVulnerabilityIdFilterRequest().ProcessFilterRequest(inspectorClient, logonCredentials.UserContext)
	} 
	
	elif logonCredentials.FilterType == "awsAccounts" {
		filterPipeline.PopulateAccountFilters(logonCredentials.ComparisonOperator).CreateAccountFilterRequest().ProcessFilterRequest(inspectorClient, logonCredentials.UserContext)
	}

	elif logonCredentials.FilterType == "typeCategory" {
		filterPipeline.PopulateTypeCategoryFilters(logonCredentials.ComparisonOperator).CreateTypeCategoryFilterRequest().ProcessFilterRequest(inspectorClient, logonCredentials.UserContext)
	}
	else filterPipeline.FilterError != nil {
		fmt.Printf("Error processing pipeline %s", filterPipeline.FilterError.Error())
		auditing.Log(filterPipeline.FilterError.Error())
	}

	fmt.Printf("Filter Output %s", *filterPipeline.FilterResponse.Arn)
} */

/*
switch (logonCredentials.FilterType){
case "cve":
	filterPipeline.PopulateTitleFilters(logonCredentials.ComparisonOperator).CreateVulnerabilityIdFilterRequest().ProcessFilterRequest(inspectorClient, logonCredentials.UserContext)
case "awsAccounts":
	filterPipeline.PopulateAccountFilters(logonCredentials.ComparisonOperator).CreateAccountFilterRequest().ProcessFilterRequest(inspectorClient, logonCredentials.UserContext)
case "typeCategory":
	filterPipeline.PopulateTypeCategoryFilters(logonCredentials.ComparisonOperator).CreateTypeCategoryFilterRequest().ProcessFilterRequest(inspectorClient, logonCredentials.UserContext)
default: 
	fmt.Println("no filter types matched")
*/