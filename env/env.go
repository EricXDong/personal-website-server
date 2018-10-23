package env

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Env struct {
	EmailUsername string
	EmailPassword string
	ContactEmail  string
	Logger        *log.Logger
}

func GetEnv() *Env {
	env := &Env{}

	env.Logger = log.New(os.Stdout, "| SERVER | ", log.Lshortfile)

	viper.AddConfigPath("config")
	viper.SetConfigName("secrets")

	err := viper.ReadInConfig()
	if err != nil {
		env.Logger.Fatalf("Error reading config: %v", err)
	}

	env.EmailUsername = viper.GetString("emailUsername")
	env.EmailPassword = viper.GetString("emailPassword")
	env.ContactEmail = viper.GetString("contactEmail")

	return env
}
