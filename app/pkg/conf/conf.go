package conf

import (
	"github.com/spf13/viper"
)

var (
	Conf config // holds the global app config.
)

type config struct {
	Image map[string]image
}

type image struct {
	Path string
}

func Init() error {
	err := viper.Unmarshal(&Conf)
	return err
}
