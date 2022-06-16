package clients

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/inspector2"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"testing"
)

// TestUserInput_GetSerialNumber Tests the formatting of serial numbers
func TestUserInput_GetSerialNumber(t *testing.T) {
	testCases := []struct{
		name string
		input UserInput
		expected string
	}{
		{name: "SerialNumberValid",input: UserInput{AwsAccount: "13423425",Username: "mark.teasdale"},expected: "arn:aws:iam::13423425:mfa/mark.teasdale"},
	}
	for _, tc := range testCases {
		actual := tc.input.GetSerialNumber()
		if *actual != tc.expected {
			t.Errorf("Test Failed expected %s got %s", *actual,tc.expected)
		}
	}
}

// TestSetRole tests setting the role
func TestSetRole(t *testing.T) {
	userData := UserInput{}
	testCases := []struct {
		name string
		account string
		roleName string
		expected string
	}{
		{name: "SetRole",  
		account: "118949222011", 
		roleName: "RoleSecurityAdministrator",
		expected: "arn:aws:iam::118949222011:role/RoleSecurityAdministrator"},
	}
	for _, tc := range testCases {
		userData.SetRole(tc.account, tc.roleName)
		if userData.ServiceCredentials.AssumedRole != tc.expected {
			t.Errorf("Test Failed expected %s got %s", userData.ServiceCredentials.AssumedRole, tc.expected)
		}
	}

}

// TestNewInspectorClientFactory test if a factory is being returned
func TestNewInspectorClientFactory(t *testing.T){
   testCases := []struct {
   	  name string
   	  expected bool
   }{
   	{
   		name: "ValidFactory",
		expected: true,
	},
   }
	for _, tc := range testCases {
		actual := isType(NewInspectorClientFactory())
        if actual != tc.expected {
			t.Errorf("Expected %v got %v",tc.expected,actual)
		}
	}
}

// TestNewSTSClientFactory test if a factory is being returned
func TestNewSTSClientFactory(t *testing.T){
	testCases := []struct {
		name string
		expected bool
	}{
		{
			name: "ValidFactory",
			expected: true,
		},
	}
	for _, tc := range testCases {
		actual := isType(NewSTSClientFactory())
		if actual != tc.expected {
			t.Errorf("Expected %v got %v",tc.expected,actual)
		}
	}
}
// isType checks for type
func isType(t interface{}) bool{
	switch t.(type) {
	case func(cfg aws.Config, stsCredentials UserInput) *inspector2.Client:
		return true
	case func(cfg aws.Config) *sts.Client:
		return true
	default:
		return false
	}
}