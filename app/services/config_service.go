package services

import (
	"os"

	"github.com/tsuru/config"
)

type configService struct {
}

func ConfigService() *configService {
	return &configService{}
}

func (this *configService) LoadConfigFile(path string) {
	err := config.ReadConfigFile(os.Args[1])
	if err != nil {
		panic(err)
	}
}

func (this *configService) GetConfig(key string, defaultValue string) {

}