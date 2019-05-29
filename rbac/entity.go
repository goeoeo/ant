package rbac

import "time"

//权限
type Perm struct {
	Id       int
	Name     string `field:"权限名称"`
	PermUrl  string `field:"权限名称"`
	ModuleId int    `field:"模块名称"` //用于树形结构展示
}

//权限数结构
type PermTree struct {
	Id          int
	Name        string
	SubPermTree []PermTree
}


//角色
type Role struct {
	Id     int
	Name   string `field:"角色名称" valid:"ChnAlphaNumeric;MinSize(4);MaxSize(20)"`
	TypeId int    `field:"角色类型"`
	Desc   string `field:"角色描述"`
	Status int    `field:"角色状态"` //1=启用,2=禁用

	CreateTime time.Time `field:"创建时间"`
	UpdateTime time.Time `field:"更新时间"`

	Perms []Perm `orm:"-"` //角色的权限信息
}

//角色状态
const (
	RoleStatusEnable  = 1 //启用
	RoleStatusDisable = 2 //禁用
)

//角色-权限 N:N
type RolePerm struct {
	Id     int
	RoleId int
	PermId int
}

//角色-用户 N:N
type RoleUser struct {
	Id     int
	RoleId int
	UserId int
}

type Perm_RolePerm struct {
	Perm
	RolePerm
}

//用户搜索参数
type RoleSearchParams struct {
	Name string
}
