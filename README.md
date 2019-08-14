# Topics
* [Dependencies](#dependencies)
* [Settings](#settings)
* [Endpoints](#endpoints)
    * [POST /token/auth](#post-tokenauth)
    * [POST /token](#post-token)
    * [GET /token/{id}](#get-tokenid)
    * [DELETE /token/{id}](#delete-tokenid)
    * [POST /token/{id}/disable](#post-tokeniddisable)
    * [GET /token/list](#get-tokenlist)
* [Service Endpoints](#service-endpoints)
    * [GET /health](#get-health)
* [Errors](#errors)

## Dependencies
* Cassandra
* GO

## Settings
 Env                     | Description
-------------------------|-----------------------------------
 AUTH_HTTP_ADDRESS       | HTTP listener address
 AUTH_CASSANDRA_HOST     | cassandra host
 AUTH_CASSANDRA_PORT     | cassandra port
 AUTH_CASSANDRA_LOGIN    | cassandra login
 AUTH_CASSANDRA_PASSWORD | cassandra password
 AUTH_CASSANDRA_KEYSPACE | cassandra keyspace

## Endpoints
### POST /token/auth
Performs auth operation by basic token.

**Request info**

```
HTTP Request method    POST
Request URL            /v1/token/auth HTTP/2.0
Headers                Authorization: {token}
```

- `Authorization` header should be taken [POST /token](#post-token).

**Response info**
```
HTTP/2.0 200 OK
Content-Type: application/json
Account-Id: {account_id}
```

### POST /v1/token
Performs create operation for basic tokens.

**Request info**

```
HTTP Request method    POST
Request URL            /token HTTP/2.0
Headers                Account-Id: {account_id}
```

- `Account-Id` - user id

**Response info**
```
HTTP/2.0 200 OK
Content-Type: application/json
```

**Response**
```json
{
    "token": {token{64}},
}
```

### DELETE /v1/token/{id}
Performs token delete operation.

**Request info**

```
HTTP Request method    DELETE
Request URL            /token/{id} HTTP/2.0
Headers                Account-Id: {account_id}
```

- `Account-Id` - user id

**Response info**
```
HTTP/1.1 200 OK
```

### PUT /v1/token/{id}
Performs token update operation.

**Request info**

```
HTTP Request method    PUT
Request URL            /v1/token/{id} HTTP/2.0
Headers                Account-Id: {account_id}
```

**Request body**
```json
{
    "is_enabled": false,
}
```

- `Account-Id` - user id

**Response info**
```
HTTP/1.1 200 OK
```

### GET /v1/token/{id}
Performs token get operation.

**Request info**

```
HTTP Request method    GET
Request URL            /v1/token/{id} HTTP/2.0
Headers                Account-Id: {account_id}
```

- `Account-Id` - user id

**Response info**
```
HTTP/1.1 200 OK
```

**Response body**
```json
{
    "token": {token{64}}
}
```

### GET /v1/token/list
Performs get token list by account.

**Request info**

```
HTTP Request method    GET
Request URL            /v1/token/list HTTP/2.0
Headers                Account-Id: {account_id}
```

- `Account-Id` - user id

**Response info**
```
HTTP/1.1 200 OK
```

**Response body**
```json
[{
    "id": {id},
    "created_at": timestamp,
    "is_active": true,
},
{
    "id": {id},
    "created_at": timestamp,
    "is_active": false,
},
{
    "id": {id},
    "created_at": timestamp,
    "is_active": true,
}]
```

# Service Endpoints
## GET /health
Internal endpoint to support livenes probes for kubernetes 
**Request info**

```
HTTP Request method    GET
Request URL            /health
```

**Response info**
```
HTTP/1.1 200 OK
```

##Errors

 Code     | Message
----------|-------------------------------------------------------------------------------------------
 4001     | invalid json
 4002     | invalid base64
 4003     | 'Account-Id' header is not set
 4004     | 'Authorization' header is not set
 4040     | token not found
 4005     | token id is invalid
 4006     | token can't be removed, invalid 'Account-Id' header"
