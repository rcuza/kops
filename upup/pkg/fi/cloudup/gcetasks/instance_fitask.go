// Code generated by ""fitask" -type=Instance"; DO NOT EDIT

package gcetasks

import (
	"encoding/json"

	"k8s.io/kops/upup/pkg/fi"
)

// Instance

// JSON marshalling boilerplate
type realInstance Instance

func (o *Instance) UnmarshalJSON(data []byte) error {
	var jsonName string
	if err := json.Unmarshal(data, &jsonName); err == nil {
		o.Name = &jsonName
		return nil
	}

	var r realInstance
	if err := json.Unmarshal(data, &r); err != nil {
		return err
	}
	*o = Instance(r)
	return nil
}

var _ fi.HasName = &Instance{}

func (e *Instance) GetName() *string {
	return e.Name
}

func (e *Instance) SetName(name string) {
	e.Name = &name
}

func (e *Instance) String() string {
	return fi.TaskAsString(e)
}