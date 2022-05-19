package configmanagement

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type InspectorConfig struct {
	Account string
	RoleName string 
}


func InitConfig() InspectorConfig {

	viper.SetConfigName("config") //name for config file
	viper.SetConfigType("yaml") //file extension type 
	viper.AddConfigPath("../")  
	viper.ReadInConfig()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("error config file not found: default \n", err)
		}
		
	if err != nil { // Handle errors reading the config file
		fmt.Println("fatal error config file: default \n", err)
        os.Exit(1)
	}
} 

return InspectorConfig{
	Account: viper.GetString("aws.account"),
	RoleName: viper.GetString("aws.rolename"),
}

}