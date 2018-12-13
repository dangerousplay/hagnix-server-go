``
Hagnix Rotmg Server - Go
``

An implementation of **ROTMG** Server (Delivery Server) in Golang using Iris, Xorm...

<br/><br/><br/><br/>

#### Environment Variables Need to Set to application Work:

At the momento only MySQL (Percona) database is supported.

DB_HOST = Host of Database
DB_DATABASE = Name of the Schema
DB_USER = Name of the user
DB_PASSWORD = Password of the user

PORT = Port which will used to server the HTTP server

WSERVER_HOST = Wserver Host which are running gRPC

TLS = Use TLS certificates

TLS_EMAIL = Email of TLS

DOMAIN = Domain to generate TLS

REDIS_HOST = Redis host

REDIS_PORT = Redis Port

REDIS_PASSWORD = Redis Password

<br/><br/><br/>

#### Config Files:

If **not setted SETTINGS_VARIABLE** you will need to provide a config file in
_/app/server.json_ 

<br/>

###### Example:

```JSON
{
"VerifyEmail": false,
  "ServerDomain": "",
  "Servers": [{
    "Name": "localhost",
    "Address": "localhost",
    "Location": ""
  }]
}
```

