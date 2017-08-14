package jiradata

/////////////////////////////////////////////////////////////////////////
// This Code is Generated by SlipScheme Project:
// https://github.com/coryb/slipscheme
//
// Generated with command:
// slipscheme -dir data -pkg jiradata -overwrite schemas/SearchResults.json
/////////////////////////////////////////////////////////////////////////
//                            DO NOT EDIT                              //
/////////////////////////////////////////////////////////////////////////

// ChangeItem defined from schema:
// {
//   "title": "Change Item",
//   "type": "object",
//   "properties": {
//     "field": {
//       "title": "field",
//       "type": "string"
//     },
//     "fieldId": {
//       "title": "fieldId",
//       "type": "string"
//     },
//     "fieldtype": {
//       "title": "fieldtype",
//       "type": "string"
//     },
//     "from": {
//       "title": "from",
//       "type": "string"
//     },
//     "fromString": {
//       "title": "fromString",
//       "type": "string"
//     },
//     "to": {
//       "title": "to",
//       "type": "string"
//     },
//     "toString": {
//       "title": "toString",
//       "type": "string"
//     }
//   }
// }
type ChangeItem struct {
	Field      string `json:"field,omitempty" yaml:"field,omitempty"`
	FieldID    string `json:"fieldId,omitempty" yaml:"fieldId,omitempty"`
	Fieldtype  string `json:"fieldtype,omitempty" yaml:"fieldtype,omitempty"`
	From       string `json:"from,omitempty" yaml:"from,omitempty"`
	FromString string `json:"fromString,omitempty" yaml:"fromString,omitempty"`
	To         string `json:"to,omitempty" yaml:"to,omitempty"`
	ToString   string `json:"toString,omitempty" yaml:"toString,omitempty"`
}
