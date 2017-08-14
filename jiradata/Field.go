package jiradata

/////////////////////////////////////////////////////////////////////////
// This Code is Generated by SlipScheme Project:
// https://github.com/coryb/slipscheme
//
// Generated with command:
// slipscheme -pkg jiradata -dir data schemas/Field.json
/////////////////////////////////////////////////////////////////////////
//                            DO NOT EDIT                              //
/////////////////////////////////////////////////////////////////////////

// Field defined from schema:
// {
//   "title": "Field",
//   "id": "https://docs.atlassian.com/jira/REST/schema/field#",
//   "type": "object",
//   "properties": {
//     "clauseNames": {
//       "title": "clauseNames",
//       "type": "array",
//       "items": {
//         "type": "string"
//       }
//     },
//     "custom": {
//       "title": "custom",
//       "type": "boolean"
//     },
//     "id": {
//       "title": "id",
//       "type": "string"
//     },
//     "key": {
//       "title": "key",
//       "type": "string"
//     },
//     "name": {
//       "title": "name",
//       "type": "string"
//     },
//     "navigable": {
//       "title": "navigable",
//       "type": "boolean"
//     },
//     "orderable": {
//       "title": "orderable",
//       "type": "boolean"
//     },
//     "schema": {
//       "title": "Json Type",
//       "type": "object",
//       "properties": {
//         "custom": {
//           "title": "custom",
//           "type": "string"
//         },
//         "customId": {
//           "title": "customId",
//           "type": "integer"
//         },
//         "items": {
//           "title": "items",
//           "type": "string"
//         },
//         "system": {
//           "title": "system",
//           "type": "string"
//         },
//         "type": {
//           "title": "type",
//           "type": "string"
//         }
//       }
//     },
//     "searchable": {
//       "title": "searchable",
//       "type": "boolean"
//     }
//   }
// }
type Field struct {
	ClauseNames ClauseNames `json:"clauseNames,omitempty" yaml:"clauseNames,omitempty"`
	Custom      bool        `json:"custom,omitempty" yaml:"custom,omitempty"`
	ID          string      `json:"id,omitempty" yaml:"id,omitempty"`
	Key         string      `json:"key,omitempty" yaml:"key,omitempty"`
	Name        string      `json:"name,omitempty" yaml:"name,omitempty"`
	Navigable   bool        `json:"navigable,omitempty" yaml:"navigable,omitempty"`
	Orderable   bool        `json:"orderable,omitempty" yaml:"orderable,omitempty"`
	Schema      *JSONType   `json:"schema,omitempty" yaml:"schema,omitempty"`
	Searchable  bool        `json:"searchable,omitempty" yaml:"searchable,omitempty"`
}
