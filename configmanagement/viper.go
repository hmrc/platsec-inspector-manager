package configmanagement

import (
	"github.com/spf13/viper"
)

// InspectorConfig holds configuration items loaded from file
type InspectorConfig struct {
	Account string
	RoleName string 
}

// InitConfig loads the configuration from a file.
func InitConfig(fileName string, fileType string, fileLocation string) (InspectorConfig, error){
    configDetails := InspectorConfig{}
	viper.SetConfigName(fileName) //name for config file
	viper.SetConfigType(fileType) //file extension type
	viper.AddConfigPath(fileLocation)
	err := viper.ReadInConfig()

	if err != nil {
		return configDetails, err
	}

	configDetails.Account = viper.GetString("aws.account")
	configDetails.RoleName = viper.GetString("aws.rolename")
	return configDetails, nil
}