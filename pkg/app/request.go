package app

import (
	"github.com/astaxie/beego/validation"

	"hub/src/app/pkg/logging"
)

// MarkErrors logs error logs
func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		logging.Info(err.Key, err.Message)
	}

	return
}
