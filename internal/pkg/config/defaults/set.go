package defaults

import "github.com/spf13/viper"

// Set applies the default values to the configuration instance.
func Set(conf *viper.Viper, defaults ...map[string]interface{}) {
	for _, m := range defaults {
		for k, v := range m {
			conf.SetDefault(k, v)
		}
	}
}
