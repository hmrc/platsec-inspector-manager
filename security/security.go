package security

import (
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/platsec-inspector-manager/clients"
)

// GetAWSSessionToken gets a valid aws session token.
func GetAWSSessionToken(userCredentials *clients.UserInput, stsClient *sts.Client) error {
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

// AssumeAccountRole returns an assumed role
func AssumeAccountRole(userCredentials *clients.UserInput, factory func(stsCredentials *clients.UserInput) *sts.Client,
	targetAccount string) error {
	stsClient := factory(userCredentials)
	userCredentials.SetRole(userCredentials.Account, userCredentials.RoleName)
	assumeRoleResult, err := stsClient.
		AssumeRole(userCredentials.UserContext, &sts.AssumeRoleInput{
			DurationSeconds: &userCredentials.SessionDuration,
			RoleArn:         &userCredentials.ServiceCredentials.AssumedRole,
			RoleSessionName: &userCredentials.SessionName,
		})
	if err != nil {
		return err
	}
	userCredentials.ServiceCredentials.AccessKeyId = *assumeRoleResult.Credentials.AccessKeyId
	userCredentials.ServiceCredentials.SecretAccessKeyId = *assumeRoleResult.Credentials.SecretAccessKey
	userCredentials.ServiceCredentials.SessionToken = *assumeRoleResult.Credentials.SessionToken
	return nil
}