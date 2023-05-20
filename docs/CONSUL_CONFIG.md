## Consul Config

_dove/api_srv_

```json
{
  "name": "api",
  "port": 8080,

  "user_srv":
  {
    "name": "user_srv"
  },
  "auth_srv":
  {
    "name": "auth_srv"
  }
}
```

_dove/user_srv_

```json
{
  "name":"user_srv",
  "mysql": {
    "host": "127.0.0.1",
    "port": 3306,
    "user": "YOUR_USER",
    "password": "YOUR_PASSWORD",
    "db": "dove",
    "salt": "YOUR_SALT"
  }
}
```
_dove/auth_srv_

```json
{
  "name":"auth_srv"
}
```


