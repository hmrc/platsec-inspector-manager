package clients

import (
	"testing"
)

// Tests the formatting of serial numbers
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

// tests setting the role 
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
