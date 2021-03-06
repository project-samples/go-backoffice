package app

import (
	"github.com/core-go/log"
	mid "github.com/core-go/log/middleware"
	"github.com/core-go/mongo"
	sv "github.com/core-go/service"
	"github.com/core-go/service/builder"
	"github.com/core-go/service/cors"
)

type Config struct {
	Server     sv.ServerConf          `mapstructure:"server"`
	Allow      cors.AllowConfig       `mapstructure:"allow"`
	Mongo      mongo.MongoConfig      `mapstructure:"mongo"`
	Log        log.Config             `mapstructure:"log"`
	MiddleWare mid.LogConfig          `mapstructure:"middleware"`
	Tracking   builder.TrackingConfig `mapstructure:"action"`
	Status     *sv.StatusConfig       `mapstructure:"status"`
	Action     *sv.ActionConfig       `mapstructure:"action"`
}
