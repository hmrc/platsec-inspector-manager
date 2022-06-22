package configmanagement_test

import (
	"github.com/platsec-inspector-manager/configmanagement"
	"testing"
)

// InspectorConfigTest test struct for holding testing data
type InspectorConfigTest struct {
	Account string
	RoleName string 
}

// TestInitConfigErrorCases tests that config is being read
func TestInitConfigErrorCases (t *testing.T){
	
	testCases:= []struct{
		name string
		fileName string
		fileType string
		fileLocation string
	}{
		{
			name: "TestInitConfigMissingFile",
			fileName: "invalidname",
			fileType: "yaml",
			fileLocation: "../",
		},
		{
			name: "TestInitConfigBadFileFormat",
			fileName: "config",
			fileType: "invalidformat",
			fileLocation: "../",
		},
	}

	for _, tc := range testCases{
		_ , actual := configmanagement.InitConfig(tc.fileName,tc.fileType,tc.fileLocation)

		if actual == nil {
			t.Errorf("Test %s failed expecting an error but got %v", tc.name,actual)
		}
	}
}

// TestInitConfigCases tests that config is being read
func TestInitConfigCases (t *testing.T){

	testCases:= []struct{
		name string
		fileName string
		fileType string
		fileLocation string
	}{
		{
			name: "TestInitConfigValidFile",
			fileName: "config",
			fileType: "yaml",
			fileLocation: "../",
		},
	}

	for _, tc := range testCases{
		_ , actual := configmanagement.InitConfig(tc.fileName,tc.fileType,tc.fileLocation)

		if actual != nil {
			t.Errorf("Test %s failed expecting an no error but got %v", tc.name,actual)
		}
	}
}