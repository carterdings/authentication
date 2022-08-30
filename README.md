# authentication
A simple authentication and authorization service.

With time limit, some cases are tested manually, such as delete user and then check\_role, invalidate token and then check, expire token and then check, etc.

## APIs

### /user/create

Create user.

usage:

```
POST /user/create
```

example:

```
curl -v 'http://127.0.0.1:8080/user/create' -H 'Content-Type: application/json' -d '{"user_name":"cat","password":"test"}' -X POST
```

return when success:

```
{"code":0,"msg":""}
```

### /user/delete

Delete user.

usage:

```
POST /user/delete
```

example:

```
curl -v 'http://127.0.0.1:8080/user/delete' -H 'Content-Type: application/json' -d '{"user_name":"carter"}' -X POST
```

return when success:

```
{"code":0,"msg":""}
```

### /role/create

Create role.

usage:

```
POST /role/create
```

example:

```
curl -v 'http://127.0.0.1:8080/role/create' -H 'Content-Type: application/json' -d '{"role_name":"root"}' -X POST
```

return when success:

```
{"code":0,"msg":""}
```

### /role/delete

Delete role.

usage:

```
POST /role/delete
```

example:

```
curl -v 'http://127.0.0.1:8080/role/delete' -H 'Content-Type: application/json' -d '{"role_name":"mqq"}' -X POST
```

return when success:

```
{"code":0,"msg":""}
```

### /user/add_role

Add role to user.

usage:

```
POST /user/add_role
```

example:

```
curl -v 'http://127.0.0.1:8080/user/add_role' -H 'Content-Type: application/json' -d '{"user_name":"cat","password":"test","role_name":"root"}' -X POST
```

return when success:

```
{"code":0,"msg":""}
```

### /auth/authenticate

Authenticate.

usage:

```
POST /auth/authenticate
```

example:

```
curl -v 'http://127.0.0.1:8080/auth/authenticate' -H 'Content-Type: application/json' -d '{"user_name":"cat","password":"test"}' -X POST
```

return when success:

```
// \u0026 means "&"
{"code":0,"msg":"","token":"expire=7200\u0026rand=HhYNOq9lFsbu9vDemPwFQA\u0026token=hVPsXOiVQ11zDmEZ1KlLoFQUbZGlFEeOyGhVKABu3MPuS-Ids2tB4K-RQ_FbQj4InFjdwpxuyQTfKpy1r3aQjvTfRx49K4mEYSmzZA3LMCoKI2zDBV1y9eIXrHRL5iMKdz6w2mLyybac7F0M0MM0LUq61y5gfj6HR5iZt3dkKBU\u0026ts=1661846597\u0026user=cat"}
```

Token should be used in interfaces as follow.

### /auth/invalidate

Invalidate.

usage:

```
POST /auth/invalidate
```

example:

```
curl -v 'http://127.0.0.1:8080/auth/invalidate' -H 'Content-Type: application/json' -d '{"user_name":"cat","password":"test","token":"expire=7200&rand=ZrCV4SLd0euk4lcHHH2cHA&token=FKunVi5yiLpGOnt5CplnT7rWtzdp-eJ4w_l9T9Yx_eUHkqBOP-ZxDHKi6nqn33JjCeSetuGlEsQ8thBU9Y5ZXG__lvBcwFhRWbWLHR_fiXQgyobrtM4bxvzXTZpGNX5Jf9ssL2YoHqeihGuHWq4DyJnqZkiVz51P5Kqh3-2WVqA&ts=1661841402&user=cat"}' -X POST
```

return when success:

```
{"code":0,"msg":""}
```

### /user/check_role

Check role.

usage:

```
POST /user/check_role
```

example:

```
curl -v 'http://127.0.0.1:8080/user/check_role' -H 'Content-Type: application/json' -d '{"user_name":"cat","password":"test","role_name":"root","token":"expire=7200&rand=ZrCV4SLd0euk4lcHHH2cHA&token=FKunVi5yiLpGOnt5CplnT7rWtzdp-eJ4w_l9T9Yx_eUHkqBOP-ZxDHKi6nqn33JjCeSetuGlEsQ8thBU9Y5ZXG__lvBcwFhRWbWLHR_fiXQgyobrtM4bxvzXTZpGNX5Jf9ssL2YoHqeihGuHWq4DyJnqZkiVz51P5Kqh3-2WVqA&ts=1661841402&user=cat"}' -X POST
```

return when success:

```
{"check_result":true,"code":0,"msg":""}
```

### /user/all_roles

All roles.

usage:

```
POST /user/all_roles
```

example:

```
curl -v 'http://127.0.0.1:8080/user/all_roles' -H 'Content-Type: application/json' -d '{"user_name":"cat","password":"test","token":"expire=7200&rand=ZrCV4SLd0euk4lcHHH2cHA&token=FKunVi5yiLpGOnt5CplnT7rWtzdp-eJ4w_l9T9Yx_eUHkqBOP-ZxDHKi6nqn33JjCeSetuGlEsQ8thBU9Y5ZXG__lvBcwFhRWbWLHR_fiXQgyobrtM4bxvzXTZpGNX5Jf9ssL2YoHqeihGuHWq4DyJnqZkiVz51P5Kqh3-2WVqA&ts=1661841402&user=cat"}' -X POST
```

return when success:

```
{"code":0,"msg":"","roles":["root"]}
```

