package rbac

import (
	"testing"
	"yungengxin2019/lwyweb/entity"
	"yungengxin2019/lwyweb/utils"
)

func TestRbacService_AddRole(t *testing.T) {
	role := entity.Role{
		Name:   "员工类型-B",
		TypeId: entity.User_Type_Employee,
		Desc:   "员工类型",
	}

	t.Log("添加角色")
	err := roleService.SaveRole(role, 1, 2, 3, 4, 5)
	if err != nil {
		t.Error(err)
	}

	t.Log("名称重复性")
	err = roleService.SaveRole(role, 1, 2, 3, 4, 5)
	if err == nil {
		t.Error("名称重复性测试失败")
	}

}

func TestRbacService_GetRoles(t *testing.T) {
	t.Log("获取角色信息")
	res := roleService.GetRoles(1, 2)

	utils.JsonPrint(res)
}

func TestRbacService_AddRolePerms(t *testing.T) {
	t.Log("绑定角色的权限")
	err := roleService.AddRolePerms(1, 2, 4, 5)
	if err != nil {
		t.Error(err)
	}

}

func TestRbacService_AddUserRoles(t *testing.T) {
	t.Log("绑定用户角色")
	err := roleService.AddUserRoles(14, 1, 2)
	if err != nil {
		t.Error(err)
	}

}

func TestRbacService_RoleStatus(t *testing.T) {
	t.Log("用户状态")
	err := roleService.RoleStatus(1, entity.RoleStatusDisable)
	if err != nil {
		t.Error(err)
	}

}

func TestRbacService_EditRole(t *testing.T) {
	role := entity.Role{
		Id:   1,
		Desc: "备注" + utils.RandomString(5),
	}

	//t.Log("编辑备注")
	//err:=roleService.EditRole(role,1,2,3,4,5)
	//if err != nil {
	//	t.Error(err)
	//}
	//
	//t.Log("名称重复性")
	//role.Name="网吧管理员"
	//err=roleService.EditRole(role,1,2,3,4,5)
	//if err!= nil {
	//	t.Error("名称重复性测试失败")
	//}

	err := roleService.SaveRole(role, 1, 2, 3)
	if err != nil {
		t.Error(err)
	}

}

func TestRbacService_GetPermsByRoleId(t *testing.T) {
	roleId := 1
	t.Log("获取权限信息")
	t.Log("禁用用户")
	err := roleService.RoleStatus(roleId, entity.RoleStatusDisable)
	if err != nil {
		t.Error(err)
	}

	res := PermService.GetPermsByRoleId(true, roleId)
	if len(res) != 0 {
		t.Error("测试失败")
	}

	t.Log("未禁用用户")
	err = roleService.RoleStatus(roleId, entity.RoleStatusEnable)
	if err != nil {
		t.Error(err)
	}

	res = PermService.GetPermsByRoleId(true, roleId)
	if len(res) == 0 {
		t.Error("测试失败")
	}

	utils.JsonPrint(res)
}

func TestRbacService_GetRolesByUserId(t *testing.T) {
	userId := 14
	t.Log("获取用户角色信息")

	res := roleService.GetRolesByUserId(userId)

	utils.JsonPrint(res)
}

func TestRbacService_GetPermsByUserId(t *testing.T) {
	userId := 14
	t.Log("获取权限信息")

	res := PermService.GetPermsByUserId(userId)

	utils.JsonPrint(res)

}

func TestRbacService_HasPerm(t *testing.T) {
	t.Log("测试权限判定函数")
	user := entity.User{
		Id: 14,
	}
	permUrl := "main/login"

	t.Log("超级用户")
	user.UserType = entity.User_Type_Admin
	hasPerm := PermService.HasPerm(user, permUrl)
	if !hasPerm {
		t.Error("失败")
	}

	t.Log("权限表中无数据")
	user.UserType = entity.User_Type_Netbar
	hasPerm = PermService.HasPerm(user, permUrl)
	if !hasPerm {
		t.Error("失败")
	}

	t.Log("用户有权限")
	permUrl = "netbar/list"
	hasPerm = PermService.HasPerm(user, permUrl)
	if !hasPerm {
		t.Error("失败")
	}

	t.Log("用户无权限")
	permUrl = "client/list"
	hasPerm = PermService.HasPerm(user, permUrl)
	if hasPerm {
		t.Error("失败")
	}

}

func TestRbacService_GetRolesByType(t *testing.T) {
	t.Log("通过typeid获取角色")
	res := roleService.GetRolesByType(4)
	utils.JsonPrint(res)
}

func TestGetRoles(t *testing.T) {
	tmpRoles := roleService.GetRolesByType(4)
	//tmpRoles=tmpRoles[:1]

	var roles [][]entity.Role
	for i := 0; i < len(tmpRoles); i++ {

		if i+1 < len(tmpRoles) {
			roles = append(roles, []entity.Role{
				tmpRoles[i], tmpRoles[i+1],
			})
		} else {
			roles = append(roles, []entity.Role{
				tmpRoles[i],
			})
		}

		i++
	}

	utils.JsonPrint(roles)

}
