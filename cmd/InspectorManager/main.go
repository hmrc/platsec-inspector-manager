package main
import (
	"os"
	"flag" 
	"github.com/platsec-inspector-manager/clients"
	"github.com/platsec-inspector-manager/security"
	"github.com/platsec-inspector-manager/auditing"
)

func main() {
	awsAccount := flag.String("account", "", "AWS account")
	username := flag.String("username", "", "AWS username")
	region := flag.String("region", "eu-west-2", "AWS region")
	profile := flag.String("profile", "", "AWS profile")
	action := flag.String("action", "SUPRESS", "action to apply")
	filterName := flag.String("filter-name", "", "filter name")
	filterType := flag.String("filter-type", "cve", "type of filter")
	comparissonOperator := flag.String("comparisson-opertaor", "EQUALS", "comaprsison operator")
	mfaToken := flag.String("mfa-token", "", "MFA token")
	flag.Parse()

	myUserInput := clients.UserInput{
		AwsAccount: *awsAccount,
		Username: *username,
		Region: *region,
		Profile: *profile,
		Action: *action,
		FilterName: *filterName,
		FilterType: *filterType,
		ComparissionOperator: *comparissonOperator,
		MfaToken: *mfaToken,
		SessionDuration: 3600,
		SessionName: "inspector",
	}

	myUserInput.SetDefaultConfig()
	// Get Session Token
	stsFactory := clients.NewSTSClientFactory()
	stsClient := stsFactory(myUserInput.UserConfig)
	err := security.GetAWSSessionToken(&myUserInput, stsClient)
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
}