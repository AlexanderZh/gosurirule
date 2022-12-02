package model

import(
	"time"
	"gorm.io/gorm"
)

type RawRule struct{
	Source string `json: Source,omitempty`
	Date time.Time `json: Date,omitempty`
	User string `json: User,omitempty`
	Tags string `json: Tags,omitempty`
	Version int `json: Version,omitempty`
	
	// The raw rule string.
	Raw string `json: Raw,omitempty`
}

type RuleOption struct {
	Option string `json:"option"`
	Args   string `json:"args"`
}

// Rule is a struct representing an IDS rule.
type Rule struct {
	gorm.Model `json:"-"`
	Source string `json: Source,omitempty`
	Date time.Time `json: Date,omitempty`
	User string `json: User,omitempty`
	Tags string `json: Tags,omitempty`
	Version int `json: Version,omitempty`
	
	// The raw rule string.
	Raw string `json: Raw,omitempty`
	
	Enabled bool `json: Enabled,omitempty`

	// Header components.
	Action     string `json: Action,omitempty`
	Proto      string `json: Proto,omitempty`
	SourceAddr string `json: SourceAddr,omitempty`
	SourcePort string `json: SourcePort,omitempty`
	Direction  string `json: Direction,omitempty`
	DestAddr   string `json: DestAddr,omitempty`
	DestPort   string `json: DestPort,omitempty`

	// List of options in order.
	Options string `json: Options,omitempty`

	// Some options are also pulled out for easy access.
	Msg string `json: Msg,omitempty`
	Sid uint64 `json: Sid,omitempty`
	Gid uint64 `json: Gid,omitempty`
}
