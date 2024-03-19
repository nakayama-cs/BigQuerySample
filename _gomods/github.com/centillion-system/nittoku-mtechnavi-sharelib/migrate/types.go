package migrate

import (
	"encoding/json"
	"time"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type DataContainer struct {
	CollectionPath string

	DocumentID string

	Value interface{}

	CreateTime time.Time

	UpdateTime time.Time
}

func (v *DataContainer) Clone() DataContainer {
	clone := *v
	if msg, ok := v.Value.(proto.Message); ok {
		clone.Value = proto.Clone(msg)
	}
	return clone
}

func (v *DataContainer) MarshalJSON() ([]byte, error) {
	raw := map[string]interface{}{
		"collectionPath": v.CollectionPath,
		"documentId":     v.DocumentID,
		"createTime":     v.CreateTime.Format(time.RFC3339Nano),
		"updateTime":     v.UpdateTime.Format(time.RFC3339Nano),
	}
	if msg, ok := v.Value.(proto.Message); ok {
		b, err := protojson.Marshal(msg)
		if err != nil {
			return nil, err
		}
		raw["value"] = map[string]interface{}{
			"@type": string(msg.ProtoReflect().Descriptor().FullName()),
			"data":  json.RawMessage(b),
		}
	} else {
		raw["value"] = map[string]interface{}{
			"data": v.Value,
		}
	}
	return json.Marshal(raw)
}

func (v *DataContainer) UnmarshalJSON(data []byte) error {
	var raw struct {
		CollectionPath string                     `json:"collectionPath"`
		DocumentID     string                     `json:"documentId"`
		Value          map[string]json.RawMessage `json:"value"`
		CreateTime     time.Time                  `json:"createTime"`
		UpdateTime     time.Time                  `json:"updateTime"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	v.CollectionPath = raw.CollectionPath
	v.DocumentID = raw.DocumentID
	v.CreateTime = raw.CreateTime
	v.UpdateTime = raw.UpdateTime
	if typeNameJSON, ok := raw.Value["@type"]; ok {
		var typeName string
		if err := json.Unmarshal(typeNameJSON, &typeName); err != nil {
			return err
		}
		msg := newMessage(protoreflect.FullName(typeName))
		if err := protojson.Unmarshal(raw.Value["data"], msg); err != nil {
			return err
		}
		v.Value = msg
	} else {
		var value interface{}
		if err := json.Unmarshal(raw.Value["data"], value); err != nil {
			return err
		}
		v.Value = value
	}
	return nil
}

type planItem struct {
	From DataContainer `json:"from"`

	To DataContainer `json:"to"`

	DiffOption string `json:"diffOption"`
}
