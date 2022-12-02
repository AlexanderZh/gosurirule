package model


type Ruleset struct {
	ID        uint           
	Name             string `json:"Name"`
	Url          string `json:"Url"`
	AutoUpdate  bool `json:"AutoUpdate"`
	UpdateInterval int `json:"UpdateInterval"`	
}
