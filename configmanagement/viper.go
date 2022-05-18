package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"errors"
)

func main() {

	viper.SetConfigName("config") //name for config file
	viper.SetConfigType("yaml") //file extension type 
	viper.AddConfigPath("./")  
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

	account := viper.GetString("aws.account")
	roleName := viper.GetString("aws.rolename")

	fmt.Println("---------- Example ----------")
	fmt.Println("aws.account :",account)
	fmt.Println("aws.rolename :",roleName)
}