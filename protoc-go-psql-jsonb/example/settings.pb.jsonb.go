// Code generated by protoc-gen-psql-json. DO NOT EDIT.
package example

import (
	"database/sql/driver"

	"google.golang.org/protobuf/encoding/protojson"
)

func (t *Settings) Scan(val interface{}) error {
	return protojson.Unmarshal(val.([]byte), t)
}

func (t *Settings) Value() (driver.Value, error) {
	return protojson.Marshal(t)
}
