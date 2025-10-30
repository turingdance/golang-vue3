package types

import (
	"database/sql/driver"
	"encoding/json"
)

type KeyValue map[string]interface{}
type KeyValueArr []KeyValue
type DataItem struct {
	Value string `json:"value"`
	Label string `json:"label"`
	Score int    `json:"score"` //此项得分
}

func (ds *KeyValue) Scan(val interface{}) error {
	if val == nil || *&val == "" {
		*ds = make(map[string]interface{}, 0)
		return nil
	} else {
		b, _ := val.([]byte)
		return json.Unmarshal(b, ds)
	}
}

func (ds KeyValue) Value() (driver.Value, error) {
	return json.Marshal(ds)
}

func (ds *KeyValueArr) Scan(val interface{}) error {
	if val == nil || *&val == "" {
		*ds = make(KeyValueArr, 0)
		return nil
	} else {
		b, _ := val.([]byte)
		return json.Unmarshal(b, ds)
	}
}

func (ds KeyValueArr) Value() (driver.Value, error) {
	return json.Marshal(ds)
}
