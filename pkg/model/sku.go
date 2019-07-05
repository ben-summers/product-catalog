package model

type Sku struct {
	Persistence
	Attributes []*KeyValue `bson:"Attributes,omitempty"`
}
