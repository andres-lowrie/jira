package jiradata

/////////////////////////////////////////////////////////////////////////
// This Code is Generated by SlipScheme Project:
// https://github.com/coryb/slipscheme
//
// Generated with command:
// slipscheme -dir jiradata -pkg jiradata -overwrite schemas/Project.json
/////////////////////////////////////////////////////////////////////////
//                            DO NOT EDIT                              //
/////////////////////////////////////////////////////////////////////////

// RemoteEntityLink defined from schema:
// {
//   "title": "Remote Entity Link",
//   "type": "object",
//   "properties": {
//     "link": {
//       "title": "link"
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
type RemoteEntityLink struct {
	Link interface{} `json:"link,omitempty" yaml:"link,omitempty"`
	Name string      `json:"name,omitempty" yaml:"name,omitempty"`
	Self string      `json:"self,omitempty" yaml:"self,omitempty"`
}