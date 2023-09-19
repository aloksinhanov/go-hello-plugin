package models

type Association struct {
	SrcID    string
	DestID   string
	Relation string
}

type Policy struct {
	ID   string
	Name string
}

type EntitlementFeature struct {
	ID   string
	Name string
}
