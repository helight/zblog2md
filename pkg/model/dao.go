package model

import (
	"zblog2md/pkg/config"
)

func init() {
	DefaultZbpPostConnection = ZbpPostConnection(GetDB2DBConnect(config.GetDBConnect("zblog")))
	DefaultZbpCategoryConnection = ZbpCategoryConnection(GetDB2DBConnect(config.GetDBConnect("zblog")))
	DefaultZbpTagConnection = ZbpTagConnection(GetDB2DBConnect(config.GetDBConnect("zblog")))
}