package configmanagement

import (
	"testing"
)

type InspectorConfigTest struct {
	Account string
	RoleName string 
}

// TestInitConfig tests that config is being read
func TestInitConfig (t *testing.T){
	
	testCases:= []struct{
		name string
		account string
		roleName string
		expected string
		configError error

	}{
		{
			name: "TestInitConfigMissingFile",
			account: "",
			roleName: "",
		},
	}

	for _, tc := range testCases{
		actual, _ := InitConfig()

		if actual.Account != tc.account {
				t.Errorf("Error expected %s but got %s",tc.account,actual.Account)
			}

			if actual.RoleName != tc.roleName {
				t.Errorf("Error expected %s but got %s", tc.roleName, actual.RoleName)
			}
	}
}

/*
	1) Create a InspectorConfig object this will be your expected
    2) Set the Account and role name to be of values that you would expect
    3) In the test call the function to get your actual
    4) Then you can test if expected.Account == actual.Account
    5) Also test for Rolename
 */