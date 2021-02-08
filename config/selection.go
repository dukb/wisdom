package config

type ServerSetting struct {
	Port string
}
type DatabaseSetting struct {
	Port     string
	Host     string
	User     string
	DBname   string
	Password string
	LogMode  bool
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	if err := s.vp.UnmarshalKey(k, v); err != nil {
		return err
	}
	return nil
}
