package commonv1

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

func NotNullString(value string) *String {
	return &String{
		Data:   value,
		IsNull: false,
	}
}
func NullString() *String {
	return &String{
		Data:   "",
		IsNull: true,
	}
}

func (m *String) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bsonx.String(m.Data).MarshalBSONValue()
}

func (m *String) UnmarshalBSONValue(t bsontype.Type, b []byte) error {
	if t == bsontype.String {
		m.Data = string(b)
		return nil
	} else {
		return fmt.Errorf("common.v1.String cannot unmarshal, require string but receive %s", t.String())
	}
}
