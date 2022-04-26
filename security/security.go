package security

import(
	"github.com/platsec-inspector-manager/clients"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

// GetAWSSessionToken gets a valid aws session token.
func GetAWSSessionToken(userCredentials *factory.UserInput, stsClient *sts.Client) error {
	sessionTokenResult, err := stsClient.GetSessionToken(userCredentials.UserContext, &sts.GetSessionTokenInput{
		TokenCode:       &userCredentials.MfaToken,
		DurationSeconds: &userCredentials.SessionDuration,
		SerialNumber:    userCredentials.GetSerialNumber(),
	}, func(options *sts.Options) {
		options.Region = userCredentials.Region
	})
	if err != nil {
		return err
	}
	userCredentials.TemporaryCredentials.AccessKeyId = *sessionTokenResult.Credentials.AccessKeyId
	userCredentials.TemporaryCredentials.SecretAccessKeyId = *sessionTokenResult.Credentials.SecretAccessKey
	userCredentials.TemporaryCredentials.SessionToken = *sessionTokenResult.Credentials.SessionToken
	return nil
}
