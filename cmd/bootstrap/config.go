package bootstrap

import (
	"github.com/XelaMP/getIP.git/model"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"strings"
)

const defaultPort = "8080"

func newConfiguration() (model.Configuration, error) {
	// we ignore the error because it can take the system environments
	_ = godotenv.Load()

	portHttp, err := strconv.Atoi(os.Getenv("PORT_HTTP"))
	if err != nil {
		return model.Configuration{}, err
	}

	return model.Configuration{
		PortHttp:       portHttp,
		AllowedOrigins: strings.Split(os.Getenv("ALLOWED_ORIGINS"), ","),
		AllowedMethods: strings.Split(os.Getenv("ALLOWED_METHODS"), ","),
	}, nil
}
