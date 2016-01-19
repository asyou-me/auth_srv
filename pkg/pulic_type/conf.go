package pulic_type

type (
	ConfType struct {
		Name     string                  `json:"name"`
		MicroSer map[string]MicroSerType `json:"micro_ser"`
		Version  string                  `json:"version"`
	}

	MicroSerType struct {
		Addr   string            `json:"addr"`
		AppId  string            `json:"app_id"`
		Appkey string            `json:"app_key"`
		Attr   map[string]string `json:"attr"`
	}
)
