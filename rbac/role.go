package rbac

import (
	"github.com/phpdi/ant/reflectutil"
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

//rbac权限服务
type roleService struct {
}


//角色管理
func (this *roleService) List(search RoleSearchParams) []Role {
	var roles []Role

	q := ConfigService.Orm.QueryTable(ConfigService.RoleTable)
	if search.Name != "" {
		q = q.Filter("name__contains", search.Name)
	}

	_, _ = q.All(&roles)

	return roles
}

//获取角色数据
func (this *roleService) GetRoles(roleIds ...int) []Role {
	var roles []Role
	roleIdStr := reflectutil.Join(roleIds, ",")
	if roleIdStr == "" {
		return roles
	}

	_, _ = ConfigService.Orm.QueryTable(ConfigService.RoleTable).Filter("id__in", roleIds).All(&roles)

	var res []Perm_RolePerm
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("p.id", "p.name", "p.perm_url", "p.module_id", "rp.role_id", "rp.perm_id").
		From(ConfigService.PermTable + " as p").
		InnerJoin(ConfigService.RolePermTable + " as rp").On("rp.perm_id = p.id").
		Where(fmt.Sprintf("rp.role_id in (%s)", roleIdStr))

	// 执行 SQL 语句
	_, _ = ConfigService.Orm.Raw(qb.String()).QueryRows(&res)

	for k, v := range roles {
		for _, v1 := range res {
			if v.Id == v1.RoleId {
				roles[k].Perms = append(roles[k].Perms, v1.Perm)
			}
		}
	}

	return roles

}

//获取单条角色信息
func (this *roleService) GetRoleByRoleId(roleId int) (Role, error) {
	role := Role{
		Id: roleId,
	}

	err := ConfigService.Orm.Read(&role)
	if err != nil {
		return role, err
	}

	role.Perms = PermService.GetPermsByRoleId(false, roleId)

	return role, nil

}

//通过类型获取角色数据
func (this *roleService) GetRolesByType(typeIds ...int) []Role {

	var roles []Role
	_, _ = ConfigService.Orm.QueryTable(ConfigService.RoleTable).Filter("TypeId__in", typeIds).Filter("status", RoleStatusEnable).All(&roles)

	return roles

}

func (this *roleService) SaveRole(role Role, permIds ...int) error {

	return Transaction(func(o orm.Ormer) error {
		//检查角色名重复
		if _, ok := Repeat(&role, ConfigService.RoleTable, "Id", "Name"); ok {
			return errors.New("角色名称重复")
		}

		var roleId int64
		var err error

		role.UpdateTime = time.Now()
		if role.Id == 0 {
			//添加
			role.Status = RoleStatusEnable
			role.CreateTime = time.Now()
			roleId, err =ConfigService.Orm.Insert(&role)
			if err != nil {
				return errors.New("添加角色失败")
			}
		} else {
			//编辑
			_, err = ConfigService.Orm.Update(&role, reflectutil.GetNotEmptyFields(role)...)
			if err != nil {
				return errors.New("编辑角色失败")
			}
			roleId = int64(role.Id)

			//刷新用户权限
			PermService.FlushUserPermsCache(int(roleId))
		}

		//删除绑定关系
		_, err = o.QueryTable(ConfigService.RolePermTable).Filter("role_id", roleId).Delete()
		if err != nil {
			return err
		}

		for _, permId := range permIds {
			item := RolePerm{
				RoleId: int(roleId),
				PermId: permId,
			}

			_, err := o.Insert(&item)

			if err != nil {
				return errors.New("绑定权限失败")
			}

		}

		return nil

	})

}

//禁用角色
func (this *roleService) RoleStatus(roleId int, status int) error {
	if roleId == 0 {
		return errors.New("角色ID缺失")
	}

	role := Role{
		Id:     roleId,
		Status: status,
	}

	_, err := ConfigService.Orm.Update(&role, reflectutil.GetNotEmptyFields(role)...)

	//刷新用户权限
	PermService.FlushUserPermsCache(roleId)

	return err
}

//通过用户Id获取角色信息
func (this *roleService) GetRolesByUserId(userId int) []Role {
	var res []Role
	//构建查询器
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("r.id", "r.name", "r.type_id", "r.desc", "r.status", "r.create_time", "r.update_time").
		From(ConfigService.RoleTable + " as r").
		InnerJoin(ConfigService.RoleUserTable + " as ru").On("r.id = ru.role_id").
		Where(fmt.Sprintf("ru.user_id=%d", userId))

	// 执行 SQL 语句
	_, _ = ConfigService.Orm.Raw(qb.String()).QueryRows(&res)

	return res

}

//绑定角色的权限
func (this *roleService) AddRolePerms(roleId int, permIds ...int) error {

	return Transaction(func(o orm.Ormer) error {
		//删除绑定关系
		_, err := o.QueryTable(ConfigService.RolePermTable).Filter("role_id", roleId).Delete()
		if err != nil {
			return err
		}

		for _, permId := range permIds {
			item := RolePerm{
				RoleId: roleId,
				PermId: permId,
			}

			_, err := o.Insert(&item)

			if err != nil {
				return errors.New("绑定权限失败")
			}

		}

		return nil

	})

}

//绑定用户角色,支持绑定多个角色
func (this *roleService) AddUserRoles(userId int, roleIds ...int) error {

	return Transaction(func(o orm.Ormer) error {
		//删除绑定关系
		_, err := o.QueryTable(ConfigService.RoleUserTable).Filter("user_id", userId).Delete()
		if err != nil {
			return nil
		}

		for _, roleId := range roleIds {
			item := RoleUser{
				RoleId: roleId,
				UserId: userId,
			}

			_, err := o.Insert(&item)

			if err != nil {
				return errors.New("绑定角色失败")
			}

		}

		return nil

	})

}
