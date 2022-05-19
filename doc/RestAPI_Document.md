# Mercury RestAPI Document

## 1. Overview

This document the introduction for cloud agent RESTful API.

**Note**: The startup parameter ```--enable-pack```, all messages will be encrypt and packed to envelope .



## 2. Rest API List

| Name                                                       | Method | URL                              | Description                 |
| ---------------------------------------------------------- | ------ | -------------------------------- | --------------------------- |
| [invitation](#21-invitation)                              | POST   | /api/v1/invitation               | create a invitation         |
| [connection request](#22-connection-request)              | POST   | /api/v1/connectionrequest        | request connection          |
| [send message](#23-send-message)                          | POST   | /api/v1/sendbasicmsg             | send a basic message      |
| [send proposal credential](#24-send-proposal-credential)  | POST   | /api/v1/sendproposalcredential   | send a proposal credential  |
| [send request credential](#25-send-request-credential)    | POST   | /api/v1/sendrequestcredential    | send a request credential   |
| [send request presentation](#26-send-request-presentation)| POST   | /api/v1/sendrequestpresentation  | send a request presentation |
| [query credential](#27-query-credential)                  | POST   | /api/v1/querycredential          | query a credential          |
| [query presentation](#28-query-presentation)              | POST   | /api/v1/querypresentation        | query a presentation        |
| [query basic message](#29-query-basic-message)            | POST   | /api/v1/querybasicmsg            | query basic message       |
| [send disconnect](#210-send-disconnect)                   | POST   | /api/v1/senddisconnect           | send disconnect request     |
| [proposal credential](#211-proposal-credential)           | POST   | /api/v1/proposalcredential       | proposal credential request     |
| [offer credential](#212-offer-credential)                 | POST   | /api/v1/offercredential          | offer credential     |
| [credential ack](#213-credential-ack)                     | POST   | /api/v1/credentialack            | credential ack     |
| [del credential](#214-del-credential)                     | POST   | /api/v1/delcredential            | del credential     |
| [connection response](#215-connection-response)           | POST   | /api/v1/connectionresponse       | connection response     |
| [connection ack](#216-connection-ack)                     | POST   | /api/v1/connectionack            | connection ack     |
| [disconnect](#217-disconnect)                             | POST   | /api/v1/disconnect               | disconnect     |
| [receive basic msg](#218-receive-basic-msg)               | POST   | /api/v1/receivebasicmsg          | receive basic msg     |
| [send disconnect](#219-send-disconnect)                   | POST   | /api/v1/senddisconnect           | send disconnect request     |
| [query connections](#220-query-connections)               | POST   | /api/v1/queryconnections         | query connections     |
| [request presentation](#221-request-presentation)         | POST   | /api/v1/requestpresentation      | request presentation     |
| [present proof](#222-present-proof)                       | POST   | /api/v1/presentproof             | present proof     |
| [presentation ack](#223-presentation-ack)                 | POST   | /api/v1/presentationack          | presentation ack     |
| [delete presentation](#224-delete-presentation)           | POST   | /api/v1/deletepresentation       | delete presentation     |


### 2.1 invitation

POST

```
/api/v1/invitation
```

Request body example:

```json
{
  "@type":"spec/connection/1.0/invitation",
  "@id":"A000000020",
  "label":"alice",
  "did":"did:ont:TQAiaefkdypSBiCSV9h9MfBJ2Ypy9fa7LY",
  "router":[
    "did:ont:TKgH6JiYWSLxWpCyoDZuky6rpNrG79zedz#1"
  ]
}
```

Response:

```json
{
  "code":0,
  "msg":"",
  "data":{
    "message_type":0,
    "content":{
      "@type":"spec/connection/1.0/invitation",
      "@id":"A000000020",
      "label":"alice",
      "did":"did:ont:TQAiaefkdypSBiCSV9h9MfBJ2Ypy9fa7LY",
      "router":[
        "did:ont:TKgH6JiYWSLxWpCyoDZuky6rpNrG79zedz#1"
      ]
    }
  }
}
```



### 2.2 connection request

POST

```
/api/v1/connectionrequest
```

Request body example:

```json
{
  "@id":"000019",
  "@type":"spec/connections/1.0/request",
  "label":"bob",
  "connection":{
    "my_did":"did:ont:TGA8YWpqwxe9LDQCdTGC7wmxTmumEQ9Gjx",
    "my_router":[
      "did:ont:TKgH6JiYWSLxWpCyoDZuky6rpNrG79zedz#1"
    ],
    "their_did":"did:ont:TQAiaefkdypSBiCSV9h9MfBJ2Ypy9fa7LY",
    "their_router":[
      "did:ont:TKgH6JiYWSLxWpCyoDZuky6rpNrG79zedz#1"
    ]
  },
  "invitation_id":"A000000019"
}

```

Response:

```json
{
  "code":0,
  "msg":""
}
```



### 2.3 send message

POST

```
/api/v1/sendbasicmsg
```

Request body example:

```json
{
  "content":"hello world",
  "connection":{
    "my_did":"did:ont:TQAiaefkdypSBiCSV9h9MfBJ2Ypy9fa7LY",
    "their_did":"did:ont:TGA8YWpqwxe9LDQCdTGC7wmxTmumEQ9Gjx",
    "my_router":[
      "did:ont:TKgH6JiYWSLxWpCyoDZuky6rpNrG79zedz#1"
    ],
    "their_router":[
      "did:ont:TKgH6JiYWSLxWpCyoDZuky6rpNrG79zedz#1"
    ]
  }
}
```

Response :

```json
{
  "code":0,
  "msg":""
}
```

### 2.4 send proposal credential

POST

```
/api/v1/sendproposalcredential
```

Request body example:

```json
{
  "@type":"spec/issue-credential/1.0/propose-credential",
  "@id":"P000002",
  "comment":"proposal1",
  "connection":{
    "my_did":"did:ont:TQAiaefkdypSBiCSV9h9MfBJ2Ypy9fa7LY",
    "their_did":"did:ont:TGA8YWpqwxe9LDQCdTGC7wmxTmumEQ9Gjx",
    "my_router":[
      "did:ont:TKgH6JiYWSLxWpCyoDZuky6rpNrG79zedz#1"
    ],
    "their_router":[
      "did:ont:TKgH6JiYWSLxWpCyoDZuky6rpNrG79zedz#1"
    ]
  }
}
```

Response:

```json
{
  "code":0,
  "msg":""
}
```

### 2.5 send request credential

POST

```
/api/v1/sendrequestcredential
```

Request body example:

```json
{
  "@type":"spec/issue-credential/1.0/request-credential",
  "@id":"RC00000031",
  "comment":"request 020",
  "connection":{
    "my_did":"did:ont:TQAiaefkdypSBiCSV9h9MfBJ2Ypy9fa7LY",
    "their_did":"did:ont:TGA8YWpqwxe9LDQCdTGC7wmxTmumEQ9Gjx",
    "my_router":[
      "did:ont:TKgH6JiYWSLxWpCyoDZuky6rpNrG79zedz#1"
    ],
    "their_router":[
      "did:ont:TKgH6JiYWSLxWpCyoDZuky6rpNrG79zedz#1"
    ]
  },
  "formats":[
    {
      "attach_id":"1",
      "format":"string"
    }
  ],
  "requests_attach":[
    {
      "@id":"1",
      "data":{
        "json":{
          "name":"age",
          "value":"greater than 18"
        }
      }
    }
  ]
}
```

Response:

```json
{
  "code":0,
  "msg":""
}
```

### 2.6 send request presentation

POST

```
/api/v1/sendrequestpresentation
```

Request body example

```json
{
  "@type":"spec/present-proof/1.0/request-presentation",
  "@id":"RP00000019",
  "comment":"test0001",
  "connection":{
    "my_did":"did:ont:TQAiaefkdypSBiCSV9h9MfBJ2Ypy9fa7LY",
    "their_did":"did:ont:TGA8YWpqwxe9LDQCdTGC7wmxTmumEQ9Gjx",
    "my_router":[
      "did:ont:TKgH6JiYWSLxWpCyoDZuky6rpNrG79zedz#1"
    ],
    "their_router":[
      "did:ont:TKgH6JiYWSLxWpCyoDZuky6rpNrG79zedz#1"
    ]
  },
  "formats":[
    {
      "attach_id":"1",
      "format":"base64"
    }
  ],
  "request_presentation_attach":[
    {
      "@id":"1",
      "data":{
        "base64":"UkMwMDAwMDAzMQ=="
      }
    }
  ]
}
```

Response

```json
{
  "code":0,
  "msg":""
}
```

### 2.7 query credential

POST

```
/api/v1/querycredential
```

Request body example:

```json
{
  "did":"did:ont:TQAiaefkdypSBiCSV9h9MfBJ2Ypy9fa7LY",
  "id":"RC00000030"
}
```

Response

```json
{
  "code":0,
  "msg":"",
  "data":{
    "message_type":14,
    "content":{
      "formats":[
        {
          "attach_id":"1",
          "format":"base64"
        }
      ],
      "credentials~attach":[
        {
          "@id":"1",
          "lastmod_time":"2020-07-15T16:54:58.511900118+08:00",
          "data":{
            "base64":"eyJhbGciOiJFUzI1NiIsImtpZCI6ImRpZDpvbnQ6VEdBOFlXcHF3eGU5TERRQ2RUR0M3d214VG11bUVROUdqeCNrZXlzLTEiLCJ0eXAiOiJKV1QifQ==.eyJpc3MiOiJkaWQ6b250OlRHQThZV3Bxd3hlOUxEUUNkVEdDN3dteFRtdW1FUTlHangiLCJleHAiOjE1OTQ4ODk2OTcsIm5iZiI6MTU5NDgwMzI5OCwiaWF0IjoxNTk0ODAzMjk4LCJqdGkiOiJ1cm46dXVpZDoxYWY3ZGJmZC1mODRjLTQ3NzctOTgzZC1iNTIzZGZlYTA0NmUiLCJ2YyI6eyJAY29udGV4dCI6WyJodHRwczovL3d3dy53My5vcmcvMjAxOC9jcmVkZW50aWFscy92MSIsImh0dHBzOi8vb250aWQub250LmlvL2NyZWRlbnRpYWxzL3YxIiwiY29udGV4dDEiLCJjb250ZXh0MiJdLCJ0eXBlIjpbIlZlcmlmaWFibGVDcmVkZW50aWFsIiwib3RmIl0sImNyZWRlbnRpYWxTdWJqZWN0IjpbeyJuYW1lIjoiYWdlIiwidmFsdWUiOiJncmVhdGVyIHRoYW4gMTgifV0sImNyZWRlbnRpYWxTdGF0dXMiOnsiaWQiOiIwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwIiwidHlwZSI6IkF0dGVzdENvbnRyYWN0In0sInByb29mIjp7ImNyZWF0ZWQiOiIyMDIwLTA3LTE1VDA4OjU0OjU4WiIsInByb29mUHVycG9zZSI6ImFzc2VydGlvbk1ldGhvZCJ9fX0=.qrgt72U95rby3L+Ox+ZV0rr7doJ8T/yv1OYZcvx7BH63oJY2npxWl9X9sGOTfOHskAUyauwAvzGOIi53oKHlhA=="
          }
        }
      ]
    }
  }
}
```

### 2.8 query presentation

POST

```
/api/v1/querypresentation
```

Request body example:

```json
{
  "did":"did:ont:TQAiaefkdypSBiCSV9h9MfBJ2Ypy9fa7LY",
  "id":"RP00000019"
}
```

Response

```json
{
  "code":0,
  "msg":"",
  "data":{
    "message_type":20,
    "content":{
      "formats":[
        {
          "attach_id":"1",
          "format":"base64"
        }
      ],
      "presentations~attach":[
        {
          "@id":"1",
          "lastmod_time":"2020-07-15T16:54:58.511900118+08:00",
          "data":{
            "base64":"eyJhbGciOiJFUzI1NiIsImtpZCI6ImRpZDpvbnQ6VEdBOFlXcHF3eGU5TERRQ2RUR0M3d214VG11bUVROUdqeCNrZXlzLTEiLCJ0eXAiOiJKV1QifQ==.eyJpc3MiOiJkaWQ6b250OlRHQThZV3Bxd3hlOUxEUUNkVEdDN3dteFRtdW1FUTlHangiLCJleHAiOjE1OTQ4ODk2OTcsIm5iZiI6MTU5NDgwMzI5OCwiaWF0IjoxNTk0ODAzMjk4LCJqdGkiOiJ1cm46dXVpZDoxYWY3ZGJmZC1mODRjLTQ3NzctOTgzZC1iNTIzZGZlYTA0NmUiLCJ2YyI6eyJAY29udGV4dCI6WyJodHRwczovL3d3dy53My5vcmcvMjAxOC9jcmVkZW50aWFscy92MSIsImh0dHBzOi8vb250aWQub250LmlvL2NyZWRlbnRpYWxzL3YxIiwiY29udGV4dDEiLCJjb250ZXh0MiJdLCJ0eXBlIjpbIlZlcmlmaWFibGVDcmVkZW50aWFsIiwib3RmIl0sImNyZWRlbnRpYWxTdWJqZWN0IjpbeyJuYW1lIjoiYWdlIiwidmFsdWUiOiJncmVhdGVyIHRoYW4gMTgifV0sImNyZWRlbnRpYWxTdGF0dXMiOnsiaWQiOiIwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwIiwidHlwZSI6IkF0dGVzdENvbnRyYWN0In0sInByb29mIjp7ImNyZWF0ZWQiOiIyMDIwLTA3LTE1VDA4OjU0OjU4WiIsInByb29mUHVycG9zZSI6ImFzc2VydGlvbk1ldGhvZCJ9fX0=.qrgt72U95rby3L+Ox+ZV0rr7doJ8T/yv1OYZcvx7BH63oJY2npxWl9X9sGOTfOHskAUyauwAvzGOIi53oKHlhA=="
          }
        }
      ]
    }
  }
}
```

### 2.9 query basic message

POST

```
 /api/v1/querybasicmsg
```

Request body example:

```json
{
  "did":"did:ont:TQAiaefkdypSBiCSV9h9MfBJ2Ypy9fa7LY",
  "latest":false,
  "remove_after_read":false
}
```
```
**latest**: true : return the latest message, false:return all messages.

**remove_after_read**:true :remove the message in storage.
```
Response

```json
{
  "code":0,
  "msg":"",
  "data":{
    "message_type":24,
    "content":[
      {
        "@type":"spec/didcomm/1.0/basicmessage",
        "@id":"6ff22592-3476-42a8-8e50-2d76cf771cb7",
        "send_time":"0001-01-01T00:00:00Z",
        "content":"124 adfasdfasefa",
        "~I10n":{
          "locale":"en"
        },
        "connection":{
          "my_did":"did:ont:TQAiaefkdypSBiCSV9h9MfBJ2Ypy9fa7LY",
          "my_router":[
            "did:ont:TKgH6JiYWSLxWpCyoDZuky6rpNrG79zedz#1"
          ],
          "their_did":"did:ont:TGA8YWpqwxe9LDQCdTGC7wmxTmumEQ9Gjx",
          "their_router":[
            "did:ont:TKgH6JiYWSLxWpCyoDZuky6rpNrG79zedz#1"
          ]
        }
      }
    ]
  }
}
```

### 2.10 send disconnect

Delete connection

POST

```
/api/v1/senddisconnect
```

Request body example:

```json
{
  "@type":"disconnect",
  "@id":"someid",
  "connection":{
    "my_did":"did:ont:TGA8YWpqwxe9LDQCdTGC7wmxTmumEQ9Gjx",
    "my_router":[
      "did:ont:TKgH6JiYWSLxWpCyoDZuky6rpNrG79zedz#1"
    ],
    "their_did":"did:ont:TQAiaefkdypSBiCSV9h9MfBJ2Ypy9fa7LY",
    "their_router":[
      "did:ont:TKgH6JiYWSLxWpCyoDZuky6rpNrG79zedz#1"
    ]
  }
}
```

Response

```json
{
  "code":0,
  "msg":""
}
```

### 2.11 proposal credential

Proposal credential

POST

```
/api/v1/proposalcredential
```

Request body example

```json
{
  "@type":"spec/connection/1.0/proposal-credential",
  "@id":"someid",
  "comment":"test0001",
  "connection":{
    "my_did":"did:ont:TGA8YWpqwxe9LDQCdTGC7wmxTmumEQ9Gjx",
    "my_router":[
      "did:ont:TKgH6JiYWSLxWpCyoDZuky6rpNrG79zedz#1"
    ],
    "their_did":"did:ont:TQAiaefkdypSBiCSV9h9MfBJ2Ypy9fa7LY",
    "their_router":[
      "did:ont:TKgH6JiYWSLxWpCyoDZuky6rpNrG79zedz#1"
    ]
  }
}
```

Response

```json
{
  "code":0,
  "msg":""
}
```

### 2.12 offer credential

Offer credential

POST

```
 /api/v1/offercredential
```

Request body example:

```json
{
  "@type":"spec/connection/1.0/offer-credential",
  "@id":"someid",
  "comment":"test0001",
  "connection":{
    "my_did":"did:ont:TGA8YWpqwxe9LDQCdTGC7wmxTmumEQ9Gjx",
    "my_router":[
      "did:ont:TKgH6JiYWSLxWpCyoDZuky6rpNrG79zedz#1"
    ],
    "their_did":"did:ont:TQAiaefkdypSBiCSV9h9MfBJ2Ypy9fa7LY",
    "their_router":[
      "did:ont:TKgH6JiYWSLxWpCyoDZuky6rpNrG79zedz#1"
    ]
  }
}
```

Response

```json
{
  "code":0,
  "msg":""
}
```

### 2.13 credential ack

Credential Ack

POST

```
 /api/v1/credentialack
```

Request body example:

```json
{
  "@type":"spec/connection/1.0/credential-ack",
  "@id":"someid",
  "status":"",
  "connection":{
    "my_did":"did:ont:TGA8YWpqwxe9LDQCdTGC7wmxTmumEQ9Gjx",
    "my_router":[
      "did:ont:TKgH6JiYWSLxWpCyoDZuky6rpNrG79zedz#1"
    ],
    "their_did":"did:ont:TQAiaefkdypSBiCSV9h9MfBJ2Ypy9fa7LY",
    "their_router":[
      "did:ont:TKgH6JiYWSLxWpCyoDZuky6rpNrG79zedz#1"
    ]
  }
}
```

Response

```json
{
  "code":0,
  "msg":""
}
```

### 2.14 del credential

Del Credential

POST

```
 /api/v1/delcredential
```

Request body example:

```json
{
  "@did":"did:ont:TGA8YWpqwxe9LDQCdTGC7wmxTmumEQ9Gjx",
  "@id":"someid"
}
```

Response

```json
{
  "code":0,
  "msg":""
}
```

### 2.15 connection response

Connection Response

POST

```
 /api/v1/connectionresponse
```

Request body example:

```json
{
  "@type":"spec/connection/1.0/connection-response",
  "@id":"someid",
  "connection":{
    "my_did":"did:ont:TGA8YWpqwxe9LDQCdTGC7wmxTmumEQ9Gjx",
    "my_router":[
      "did:ont:TKgH6JiYWSLxWpCyoDZuky6rpNrG79zedz#1"
    ],
    "their_did":"did:ont:TQAiaefkdypSBiCSV9h9MfBJ2Ypy9fa7LY",
    "their_router":[
      "did:ont:TKgH6JiYWSLxWpCyoDZuky6rpNrG79zedz#1"
    ]
  }
}
```

Response

```json
{
  "code":0,
  "msg":""
}
```


### 2.16 connection ack

Connection Ack

POST

```
 /api/v1/connectionack
```

Request body example:

```json
{
  "@type":"spec/connection/1.0/connection-ack",
  "@id":"someid",
  "status":"",
  "connection":{
    "my_did":"did:ont:TGA8YWpqwxe9LDQCdTGC7wmxTmumEQ9Gjx",
    "my_router":[
      "did:ont:TKgH6JiYWSLxWpCyoDZuky6rpNrG79zedz#1"
    ],
    "their_did":"did:ont:TQAiaefkdypSBiCSV9h9MfBJ2Ypy9fa7LY",
    "their_router":[
      "did:ont:TKgH6JiYWSLxWpCyoDZuky6rpNrG79zedz#1"
    ]
  }
}
```

Response

```json
{
  "code":0,
  "msg":""
}
```

### 2.17 disconnect

Disconnect

POST

```
 /api/v1/disconnect
```

Request body example:

```json
{
  "@type":"spec/connection/1.0/disconnect",
  "@id":"someid",
  "connection":{
    "my_did":"did:ont:TGA8YWpqwxe9LDQCdTGC7wmxTmumEQ9Gjx",
    "my_router":[
      "did:ont:TKgH6JiYWSLxWpCyoDZuky6rpNrG79zedz#1"
    ],
    "their_did":"did:ont:TQAiaefkdypSBiCSV9h9MfBJ2Ypy9fa7LY",
    "their_router":[
      "did:ont:TKgH6JiYWSLxWpCyoDZuky6rpNrG79zedz#1"
    ]
  }
}
```

Response

```json
{
  "code":0,
  "msg":""
}
```


### 2.18 receive basic msg

Receive BasicMsg

POST

```
 /api/v1/receivebasicmsg
```

Request body example:

```json
{
  "@type":"spec/didcomm/1.0/receive-basic-message",
  "@id":"6ff22592-3476-42a8-8e50-2d76cf771cb7",
  "send_time":"0001-01-01T00:00:00Z",
  "content":"124 adfasdfasefa",
  "~I10n":{
    "locale":"en"
  },
  "connection":{
    "my_did":"did:ont:TGA8YWpqwxe9LDQCdTGC7wmxTmumEQ9Gjx",
    "my_router":[
      "did:ont:TKgH6JiYWSLxWpCyoDZuky6rpNrG79zedz#1"
    ],
    "their_did":"did:ont:TQAiaefkdypSBiCSV9h9MfBJ2Ypy9fa7LY",
    "their_router":[
      "did:ont:TKgH6JiYWSLxWpCyoDZuky6rpNrG79zedz#1"
    ]
  }
}
```

Response

```json
{
  "code":0,
  "msg":""
}
```

### 2.19 send disconnect

Send Disconnect

POST

```
 /api/v1/senddisconnect
```

Request body example:

```json
{
  "@type":"spec/connection/1.0/send-disconnect",
  "@id":"someid",
  "connection":{
    "my_did":"did:ont:TGA8YWpqwxe9LDQCdTGC7wmxTmumEQ9Gjx",
    "my_router":[
      "did:ont:TKgH6JiYWSLxWpCyoDZuky6rpNrG79zedz#1"
    ],
    "their_did":"did:ont:TQAiaefkdypSBiCSV9h9MfBJ2Ypy9fa7LY",
    "their_router":[
      "did:ont:TKgH6JiYWSLxWpCyoDZuky6rpNrG79zedz#1"
    ]
  }
}
```

Response

```json
{
  "code":0,
  "msg":""
}
```


### 2.20 query connections

Query Connections

POST

```
 /api/v1/queryconnections
```

Request body example:

```json
{
  "@id":"someid"
}
```

Response

```json
{
  "code":0,
  "msg":""
}
```


### 2.21 request presentation

Request Presentation

POST

```
 /api/v1/requestpresentation
```

Request body example:

```json
{
  "@type":"spec/connection/1.0/request-presentation",
  "@id":"someid",
  "connection":{
    "my_did":"did:ont:TGA8YWpqwxe9LDQCdTGC7wmxTmumEQ9Gjx",
    "my_router":[
      "did:ont:TKgH6JiYWSLxWpCyoDZuky6rpNrG79zedz#1"
    ],
    "their_did":"did:ont:TQAiaefkdypSBiCSV9h9MfBJ2Ypy9fa7LY",
    "their_router":[
      "did:ont:TKgH6JiYWSLxWpCyoDZuky6rpNrG79zedz#1"
    ]
  }
}
```

Response

```json
{
  "code":0,
  "msg":""
}
```


### 2.22 present proof

Present Proof

POST

```
 /api/v1/presentproof
```

Request body example:

```json
{
  "@type":"spec/present-proof/1.0/present-proof",
  "@id":"RP00000019",
  "comment":"test0001",
  "formats":[
    {
      "attach_id":"1",
      "format":"base64"
    }
  ],
  "presentation_attach":[
    {
      "@id":"1",
      "data":{
        "base64":"UkMwMDAwMDAzMQ=="
      }
    }
  ],
  "connection":{
    "my_did":"did:ont:TQAiaefkdypSBiCSV9h9MfBJ2Ypy9fa7LY",
    "their_did":"did:ont:TGA8YWpqwxe9LDQCdTGC7wmxTmumEQ9Gjx",
    "my_router":[
      "did:ont:TKgH6JiYWSLxWpCyoDZuky6rpNrG79zedz#1"
    ],
    "their_router":[
      "did:ont:TKgH6JiYWSLxWpCyoDZuky6rpNrG79zedz#1"
    ]
  }
}
```

Response

```json
{
  "code":0,
  "msg":""
}
```

### 2.23 presentation ack

Presentation Ack

POST

```
 /api/v1/presentationack
```

Request body example:

```json
{
  "@type":"spec/present-proof/1.0/presentation-ack",
  "@id":"someid",
  "@status":"",
  "connection":{
    "my_did":"did:ont:TGA8YWpqwxe9LDQCdTGC7wmxTmumEQ9Gjx",
    "my_router":[
      "did:ont:TKgH6JiYWSLxWpCyoDZuky6rpNrG79zedz#1"
    ],
    "their_did":"did:ont:TQAiaefkdypSBiCSV9h9MfBJ2Ypy9fa7LY",
    "their_router":[
      "did:ont:TKgH6JiYWSLxWpCyoDZuky6rpNrG79zedz#1"
    ]
  }
}
```

Response

```json
{
  "code":0,
  "msg":""
}
```

### 2.24 delete presentation

Delete Presentation

POST

```
 /api/v1/deletepresentation
```

Request body example:

```json
{
  "@did":"did:ont:TGA8YWpqwxe9LDQCdTGC7wmxTmumEQ9Gjx",
  "@id":"someid"
}
```

Response

```json
{
  "code":0,
  "msg":""
}
```
