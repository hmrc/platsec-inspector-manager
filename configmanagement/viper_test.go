package configmanagement

import (
	"testing"
)

type InspectorConfigTest struct {
	Account string
	RoleName string 
}



func TestInitConfig (t *testing.T){
	
	testCases:= []struct{
		name string
		Account string
		RoleName string
		expected string
	}{
		{
			name: "TestInitConfig",
			Account: "118949222011",
			RoleName: "RoleSecurityAdministrator",
			expected: ,
		},
	}

	for _, tc := range testCases{
		actual := InitConfig()
		if *actual.Value != *tc.expected.Value {
			t.Errorf("Error expected %s but got %s",*tc.expected.Value,*actual.Value)
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