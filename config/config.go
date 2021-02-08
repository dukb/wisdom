package config

import "github.com/spf13/viper"

type Setting struct {
	vp *viper.Viper
}

func NewSetting() (*Setting, error) {
	var err error
	vp := viper.New()
	vp.SetConfigName("setting")
	vp.AddConfigPath("./")
	vp.SetConfigType("yaml")
	if err = vp.ReadInConfig(); err != nil {
		return nil, err
	}
	return &Setting{vp}, nil
}
