// Code generated by gopkg.in/reform.v1. DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type channelTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *channelTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("ia_channels").
func (v *channelTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *channelTableType) Columns() []string {
	return []string{
		"id",
		"summary",
		"type",
		"email_config",
		"pagerduty_config",
		"slack_config",
		"webhook_config",
		"disabled",
		"created_at",
		"updated_at",
	}
}

// NewStruct makes a new struct for that view or table.
func (v *channelTableType) NewStruct() reform.Struct {
	return new(Channel)
}

// NewRecord makes a new record for that table.
func (v *channelTableType) NewRecord() reform.Record {
	return new(Channel)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *channelTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// ChannelTable represents ia_channels view or table in SQL database.
var ChannelTable = &channelTableType{
	s: parse.StructInfo{
		Type:    "Channel",
		SQLName: "ia_channels",
		Fields: []parse.FieldInfo{
			{Name: "ID", Type: "string", Column: "id"},
			{Name: "Summary", Type: "string", Column: "summary"},
			{Name: "Type", Type: "ChannelType", Column: "type"},
			{Name: "EmailConfig", Type: "*EmailConfig", Column: "email_config"},
			{Name: "PagerDutyConfig", Type: "*PagerDutyConfig", Column: "pagerduty_config"},
			{Name: "SlackConfig", Type: "*SlackConfig", Column: "slack_config"},
			{Name: "WebHookConfig", Type: "*WebHookConfig", Column: "webhook_config"},
			{Name: "Disabled", Type: "bool", Column: "disabled"},
			{Name: "CreatedAt", Type: "time.Time", Column: "created_at"},
			{Name: "UpdatedAt", Type: "time.Time", Column: "updated_at"},
		},
		PKFieldIndex: 0,
	},
	z: new(Channel).Values(),
}

// String returns a string representation of this struct or record.
func (s Channel) String() string {
	res := make([]string, 10)
	res[0] = "ID: " + reform.Inspect(s.ID, true)
	res[1] = "Summary: " + reform.Inspect(s.Summary, true)
	res[2] = "Type: " + reform.Inspect(s.Type, true)
	res[3] = "EmailConfig: " + reform.Inspect(s.EmailConfig, true)
	res[4] = "PagerDutyConfig: " + reform.Inspect(s.PagerDutyConfig, true)
	res[5] = "SlackConfig: " + reform.Inspect(s.SlackConfig, true)
	res[6] = "WebHookConfig: " + reform.Inspect(s.WebHookConfig, true)
	res[7] = "Disabled: " + reform.Inspect(s.Disabled, true)
	res[8] = "CreatedAt: " + reform.Inspect(s.CreatedAt, true)
	res[9] = "UpdatedAt: " + reform.Inspect(s.UpdatedAt, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Channel) Values() []interface{} {
	return []interface{}{
		s.ID,
		s.Summary,
		s.Type,
		s.EmailConfig,
		s.PagerDutyConfig,
		s.SlackConfig,
		s.WebHookConfig,
		s.Disabled,
		s.CreatedAt,
		s.UpdatedAt,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Channel) Pointers() []interface{} {
	return []interface{}{
		&s.ID,
		&s.Summary,
		&s.Type,
		&s.EmailConfig,
		&s.PagerDutyConfig,
		&s.SlackConfig,
		&s.WebHookConfig,
		&s.Disabled,
		&s.CreatedAt,
		&s.UpdatedAt,
	}
}

// View returns View object for that struct.
func (s *Channel) View() reform.View {
	return ChannelTable
}

// Table returns Table object for that record.
func (s *Channel) Table() reform.Table {
	return ChannelTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Channel) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Channel) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Channel) HasPK() bool {
	return s.ID != ChannelTable.z[ChannelTable.s.PKFieldIndex]
}

// SetPK sets record primary key, if possible.
//
// Deprecated: prefer direct field assignment where possible: s.ID = pk.
func (s *Channel) SetPK(pk interface{}) {
	reform.SetPK(s, pk)
}

// check interfaces
var (
	_ reform.View   = ChannelTable
	_ reform.Struct = (*Channel)(nil)
	_ reform.Table  = ChannelTable
	_ reform.Record = (*Channel)(nil)
	_ fmt.Stringer  = (*Channel)(nil)
)

func init() {
	parse.AssertUpToDate(&ChannelTable.s, new(Channel))
}
