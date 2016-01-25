package handler

import (
	"github.com/asyoume/Auth_ser/pkg/models"
	"github.com/asyoume/lib/pulic_type"
)

var conf *pulic_type.ConfType

func Init(conf_path string) error {
	conf, err := ConfigInit(conf_path)
	if err != nil {
		return err
	}
	models.Init(conf.MicroSer["db"], conf.MicroSer["dblog"])
	if err != nil {
		return err
	}
	return nil
}
