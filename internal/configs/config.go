package configs

import (
	"fmt"

	"github.com/gookit/goutil/strutil"
)

const (
	//运行环境
	RUN_RELEASED = "prod"
	RUN_TEST     = "test"
)

var configuration *Configuration = &Configuration{Options: Options{}}

type Configuration struct {
	Port    int
	Options Options
}

type Options struct {
	RunEnvironment string
}

func InitializateCallback() {
}

func GetConfiguration() *Configuration {
	return configuration
}

func SetConfiguration(c *Configuration) {
	fmt.Printf("loading start config: %v", &c)
	configuration = c
}

func (c *Configuration) IsReleased() bool {
	return strutil.Equal(RUN_RELEASED, c.Options.RunEnvironment)
}
