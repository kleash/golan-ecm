package objectstore

import (
	"cjavellana.me/ecm/golan/internal/cfg"
	"cjavellana.me/ecm/golan/internal/core/ce"
	"cjavellana.me/ecm/golan/internal/core/engine/aws"
)

func Get(config cfg.AppConfig) ce.ObjectStore {

	switch config.StoreType {
	case cfg.StoreTypeAWS:
		return aws.GetObjectStore(&config)
	default:
		return nil
	}

}
