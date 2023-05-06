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
  "name":"user_srv"
}
```
_dove/auth_srv_

```json
{
  "name":"auth_srv"
}
```


