package main

import (
	core "BankGW-Gin/Core"

	_ "BankGW-Gin/docs"

	"github.com/spf13/viper"
)

// @title TQR Api Gin
// @version 1.0
// @description.markdown
// @termsOfService http://somewhere.com/

// @contact.name API Support
// @contact.url http://somewhere.com/support
// @contact.email support@somewhere.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @schemes https http

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// viper.SetConfigFile("configs.json")
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.ReadInConfig()

	r := core.Router()

	r.Run(":7979")
}
