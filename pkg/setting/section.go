package setting

import (
	"time"
)

// ServerSettingS ...
type ServerSettingS struct {
	RunMode      string
	HTTPPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// AppSettingS ...
type AppSettingS struct {
	DefaultPageSize int
	MaxPageSize     int
	LogSavePath     string
	logFileName     string
	logFileExt      string
}

// DatabaseSettingS ...
type DatabaseSettingS struct {
	DBType       string
	Username     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

// ReadSection ...
func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}
