package main
import (
	"flag" 
	"github.com/platsec-inspector-manager/clients"
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

	myUserInput := factory.UserInput{
		AwsAccount: *awsAccount,
		Usernname: *username,
		Region: *region,
		Profile: *profile,
		Action: *action,
		FilterName: *filterName,
		FilterType: *filterType,
		ComparissionOperator: *comparissonOperator,
		MfaToken: *mfaToken,
		SessionDuration: 3600,

	}

	myUserInput.SetDefaultConfig()
	// Get Session Token
	stsFactory := factory.NewSTSClientFactory()
	stsClient := stsFactory(myUserInput.UserConfig)
}