package entity

const (
	SaltLen     = 32
	TokenExpire = 2 * 60 * 60
)

// UserReq create or delete user request
type UserReq struct {
	UserName string `json:"user_name,omitempty"`
	Password string `json:"password,omitempty"`
}

// RoleReq create or delete role request
type RoleReq struct {
	RoleName string `json:"role_name,omitempty"`
}

// CommRsp common response
type CommRsp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// UserRoleReq add role to user, check role request
type UserRoleReq struct {
	UserName string `json:"user_name,omitempty"`
	Password string `json:"password,omitempty"`
	RoleName string `json:"role_name,omitempty"`
	Token    string `json:"token,omitempty"`
}

// Key user-role key
func (ur *UserRoleReq) Key() string {
	return "bind_" + ur.UserName + "_" + ur.RoleName
}

// UserRoleRsp authenticate, invalidate, check role, all roles response
type UserRoleRsp struct {
	Code        int      `json:"code"`
	Msg         string   `json:"msg"`
	Token       string   `json:"token,omitempty"`
	CheckResult bool     `json:"check_result,omitempty"`
	Roles       []string `json:"roles,omitempty"`
}

// User user
type User struct {
	UserName   string `json:"user_name,omitempty"`
	Salt       []byte `json:"salt,omitempty"`
	Password   []byte `json:"password,omitempty"`
	CreateTime int64  `json:"create_time,omitempty"`
}

// Key for searching
func (u *User) Key() string {
	return "user_" + u.UserName
}

// UserRolesKey for searching all roles
func (u *User) UserRolesKey() string {
	return "user_roles_" + u.UserName
}

// Role role
type Role struct {
	RoleName   string `json:"role_name,omitempty"`
	CreateTime int64  `json:"create_time,omitempty"`
}

// Key for searching
func (r *Role) Key() string {
	return "role_" + r.RoleName
}

// Token token
type Token struct {
	UserName  string `json:"user_name,omitempty"`
	Password  []byte `json:"password,omitempty"`
	Rand      string `json:"rand,omitempty"`
	Timestamp int64  `json:"ts,omitempty"`
	Expire    int64  `json:"expire,omitempty"`
}
