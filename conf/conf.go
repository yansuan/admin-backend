package conf

import (
	"github.com/yansuan/go-admin-backend/pkg/httpx"
	"github.com/yansuan/go-admin-backend/pkg/logx"
	"github.com/yansuan/go-admin-backend/pkg/ormx"
	"github.com/yansuan/go-admin-backend/storage"
)

type ConfigType struct {
	Global GlobalConfig
	Log    logx.Config
	HTTP   httpx.Config
	DB     ormx.DBConfig
	Redis  storage.RedisConfig
}

type GlobalConfig struct {
	RunMode string
}

func InitConfig(configDir, cryptoKey string) (*ConfigType, error) {
	// var config = new(ConfigType)

	// if err := cfg.LoadConfigByDir(configDir, config); err != nil {
	// 	return nil, fmt.Errorf("failed to load configs of directory: %s error: %s", configDir, err)
	// }

	// return config, err

	return nil, nil
}
