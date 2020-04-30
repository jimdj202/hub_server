package util

import "hub/src/app/pkg/setting"

// Setup Initialize the util
func Setup() {
	jwtSecret = []byte(setting.AppSetting.JwtSecret)
}