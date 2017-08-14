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

// Components defined from schema:
// {
//   "title": "components",
//   "type": "array",
//   "items": {
//     "title": "Component",
//     "type": "object",
//     "properties": {
//       "assignee": {
//         "title": "User",
//         "type": "object",
//         "properties": {
//           "accountId": {
//             "type": "string"
//           },
//           "active": {
//             "type": "boolean"
//           },
//           "applicationRoles": {
//             "$ref": "#/definitions/simple-list-wrapper"
//           },
//           "avatarUrls": {
//             "type": "object",
//             "patternProperties": {
//               ".+": {
//                 "type": "string"
//               }
//             }
//           },
//           "displayName": {
//             "type": "string"
//           },
//           "emailAddress": {
//             "type": "string"
//           },
//           "expand": {
//             "type": "string"
//           },
//           "groups": {
//             "$ref": "#/definitions/simple-list-wrapper"
//           },
//           "key": {
//             "type": "string"
//           },
//           "locale": {
//             "type": "string"
//           },
//           "name": {
//             "type": "string"
//           },
//           "self": {
//             "type": "string"
//           },
//           "timeZone": {
//             "type": "string"
//           }
//         }
//       },
//       "assigneeType": {
//         "title": "assigneeType",
//         "type": "string"
//       },
//       "description": {
//         "title": "description",
//         "type": "string"
//       },
//       "id": {
//         "title": "id",
//         "type": "string"
//       },
//       "isAssigneeTypeValid": {
//         "title": "isAssigneeTypeValid",
//         "type": "boolean"
//       },
//       "lead": {
//         "title": "User",
//         "type": "object",
//         "properties": {
//           "accountId": {
//             "type": "string"
//           },
//           "active": {
//             "type": "boolean"
//           },
//           "applicationRoles": {
//             "$ref": "#/definitions/simple-list-wrapper"
//           },
//           "avatarUrls": {
//             "type": "object",
//             "patternProperties": {
//               ".+": {
//                 "type": "string"
//               }
//             }
//           },
//           "displayName": {
//             "type": "string"
//           },
//           "emailAddress": {
//             "type": "string"
//           },
//           "expand": {
//             "type": "string"
//           },
//           "groups": {
//             "$ref": "#/definitions/simple-list-wrapper"
//           },
//           "key": {
//             "type": "string"
//           },
//           "locale": {
//             "type": "string"
//           },
//           "name": {
//             "type": "string"
//           },
//           "self": {
//             "type": "string"
//           },
//           "timeZone": {
//             "type": "string"
//           }
//         }
//       },
//       "leadUserName": {
//         "title": "leadUserName",
//         "type": "string"
//       },
//       "name": {
//         "title": "name",
//         "type": "string"
//       },
//       "project": {
//         "title": "project",
//         "type": "string"
//       },
//       "projectId": {
//         "title": "projectId",
//         "type": "integer"
//       },
//       "realAssignee": {
//         "title": "User",
//         "type": "object",
//         "properties": {
//           "accountId": {
//             "type": "string"
//           },
//           "active": {
//             "type": "boolean"
//           },
//           "applicationRoles": {
//             "$ref": "#/definitions/simple-list-wrapper"
//           },
//           "avatarUrls": {
//             "type": "object",
//             "patternProperties": {
//               ".+": {
//                 "type": "string"
//               }
//             }
//           },
//           "displayName": {
//             "type": "string"
//           },
//           "emailAddress": {
//             "type": "string"
//           },
//           "expand": {
//             "type": "string"
//           },
//           "groups": {
//             "$ref": "#/definitions/simple-list-wrapper"
//           },
//           "key": {
//             "type": "string"
//           },
//           "locale": {
//             "type": "string"
//           },
//           "name": {
//             "type": "string"
//           },
//           "self": {
//             "type": "string"
//           },
//           "timeZone": {
//             "type": "string"
//           }
//         }
//       },
//       "realAssigneeType": {
//         "title": "realAssigneeType",
//         "type": "string"
//       },
//       "self": {
//         "title": "self",
//         "type": "string"
//       }
//     }
//   }
// }
type Components []*Component
