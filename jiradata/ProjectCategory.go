package jiradata

/////////////////////////////////////////////////////////////////////////
// This Code is Generated by SlipScheme Project:
// https://github.com/coryb/slipscheme
//
// Generated with command:
// slipscheme -dir jiradata -pkg jiradata schemas/Project.json
/////////////////////////////////////////////////////////////////////////
//                            DO NOT EDIT                              //
/////////////////////////////////////////////////////////////////////////

// ProjectCategory defined from schema:
// {
//   "title": "Project Category",
//   "type": "object",
//   "properties": {
//     "description": {
//       "title": "description",
//       "type": "string"
//     },
//     "id": {
//       "title": "id",
//       "type": "string"
//     },
//     "name": {
//       "title": "name",
//       "type": "string"
//     },
//     "self": {
//       "title": "self",
//       "type": "string"
//     }
//   }
// }
type ProjectCategory struct {
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	ID          string `json:"id,omitempty" yaml:"id,omitempty"`
	Name        string `json:"name,omitempty" yaml:"name,omitempty"`
	Self        string `json:"self,omitempty" yaml:"self,omitempty"`
}
