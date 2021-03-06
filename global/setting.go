package global

import (
	"go-travel-blog/pkg/logger"
	"go-travel-blog/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS

	Logger *logger.Logger
)
