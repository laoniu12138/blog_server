package global

import (
	"blog_server/pkg/setting"
	"blog_server/pkg/logger"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger 			*logger.Logger
)
