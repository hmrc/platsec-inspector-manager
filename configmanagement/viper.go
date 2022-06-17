package configmanagement

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

// InspectorConfig holds configuration items loaded from file
type InspectorConfig struct {
	Account string
	RoleName string 
}

// InitConfig loads the configuration from a file.
func InitConfig() (InspectorConfig, error){
    configDetails := InspectorConfig{}
	viper.SetConfigName("config") //name for config file
	viper.SetConfigType("yaml") //file extension type 
	viper.AddConfigPath("../")  
	err := viper.ReadInConfig()

	if err != nil {
		os.Exit(1)
	}

	if err = viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return configDetails, err
		}
		
		if err != nil { // Handle errors reading the config file
			fmt.Println("fatal error config file: default \n", err)
        	os.Exit(1)
		}
	}

	configDetails.Account = viper.GetString("aws.account")
	configDetails.RoleName = viper.GetString("aws.rolename")
	return configDetails, nil
}