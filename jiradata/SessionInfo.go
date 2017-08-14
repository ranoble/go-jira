package jiradata

/////////////////////////////////////////////////////////////////////////
// This Code is Generated by SlipScheme Project:
// https://github.com/coryb/slipscheme
//
// Generated with command:
// slipscheme -pkg jiradata -dir data schemas/AuthSuccess.json
/////////////////////////////////////////////////////////////////////////
//                            DO NOT EDIT                              //
/////////////////////////////////////////////////////////////////////////

// SessionInfo defined from schema:
// {
//   "title": "Session Info",
//   "type": "object",
//   "properties": {
//     "name": {
//       "title": "name",
//       "type": "string"
//     },
//     "value": {
//       "title": "value",
//       "type": "string"
//     }
//   }
// }
type SessionInfo struct {
	Name  string `json:"name,omitempty" yaml:"name,omitempty"`
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
