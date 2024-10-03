package config

type AppConf struct {
	AppName     string `env:"APP_NAME"`
	AppListener string `env:"APP_LISTENER"`
}
