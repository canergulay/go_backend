package env

import (
	"os"
)

func SetEnvorinmentVariables() {
	os.Setenv("SERVER", "localhost")
	os.Setenv("USER_NAME", "postgres")
	os.Setenv("PASSWORD", "canerrenac1")
	os.Setenv("DB_NAME", "ressurection")
	os.Setenv("DB_PORT", "5432")
}
