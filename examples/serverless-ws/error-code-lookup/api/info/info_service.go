package info

import (
	"error-code-lookup/config/app"

	"github.com/rs/zerolog/log"
)

type AppInfoService struct {
}

func NewService() *AppInfoService {
	return &AppInfoService{}
}

func (service *AppInfoService) getAppInfo() (*AppInfo, error) {
	// this block make GetAppInfo() not testable.
	info := &AppInfo{
		AppName:     app.Config.AppInfo.Name,
		Description: app.Config.AppInfo.Description,
		Version:     app.Config.AppInfo.Version,
		Env:         app.Config.AppInfo.Env,
	}

	log.Info().Any("info", *info)
	return info, nil
}
