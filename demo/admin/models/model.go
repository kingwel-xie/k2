package models

import (
	"github.com/kingwel-xie/k2/core/logger"
)

var log = logger.Logger("model")

type ModelTable interface {
	TableName() string
}