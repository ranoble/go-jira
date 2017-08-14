package jiradata

/////////////////////////////////////////////////////////////////////////
// This Code is Generated by SlipScheme Project:
// https://github.com/coryb/slipscheme
//
// Generated with command:
// slipscheme -pkg jiradata -dir data schemas/WorklogWithPagination.json
/////////////////////////////////////////////////////////////////////////
//                            DO NOT EDIT                              //
/////////////////////////////////////////////////////////////////////////

// Worklogs defined from schema:
// {
//   "title": "worklogs",
//   "type": "array",
//   "items": {
//     "title": "Worklog",
//     "type": "object",
//     "properties": {
//       "author": {
//         "title": "User",
//         "type": "object",
//         "properties": {
//           "accountId": {
//             "type": "string"
//           },
//           "active": {
//             "type": "boolean"
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
//           "key": {
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
//       "comment": {
//         "title": "comment",
//         "type": "string"
//       },
//       "created": {
//         "title": "created",
//         "type": "string"
//       },
//       "id": {
//         "title": "id",
//         "type": "string"
//       },
//       "issueId": {
//         "title": "issueId",
//         "type": "string"
//       },
//       "properties": {
//         "title": "properties",
//         "type": "array",
//         "items": {
//           "title": "Entity Property",
//           "type": "object",
//           "properties": {
//             "key": {
//               "title": "key",
//               "type": "string"
//             },
//             "value": {
//               "title": "value"
//             }
//           }
//         }
//       },
//       "self": {
//         "title": "self",
//         "type": "string"
//       },
//       "started": {
//         "title": "started",
//         "type": "string"
//       },
//       "timeSpent": {
//         "title": "timeSpent",
//         "type": "string"
//       },
//       "timeSpentSeconds": {
//         "title": "timeSpentSeconds",
//         "type": "integer"
//       },
//       "updateAuthor": {
//         "title": "User",
//         "type": "object",
//         "properties": {
//           "accountId": {
//             "type": "string"
//           },
//           "active": {
//             "type": "boolean"
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
//           "key": {
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
//       "updated": {
//         "title": "updated",
//         "type": "string"
//       },
//       "visibility": {
//         "title": "Visibility",
//         "type": "object",
//         "properties": {
//           "type": {
//             "title": "type",
//             "type": "string"
//           },
//           "value": {
//             "title": "value",
//             "type": "string"
//           }
//         }
//       }
//     }
//   }
// }
type Worklogs []*Worklog
