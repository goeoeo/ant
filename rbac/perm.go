package rbac

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)


//权限服务
type permService struct {
	perms               []Perm     //所有的权限数据
	permTree            []PermTree //权限树结构

}

//构造函数
func NewPerm() *permService {
	this := &permService{}
	this.InitPerms()

	return this
}


//初始化权限数据
func (this *permService) InitPerms() *permService {
	_, _ = ConfigService.Orm.QueryTable(ConfigService.PermTable).Limit(-1).All(&this.perms)
	return this
}

//初始化权限树结构
func (this *permService) InitPermTree(filePath string) *permService {
	//配置文件中基础权限树
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		beego.Error(err.Error())
	}
	err = json.Unmarshal(content, &this.permTree)
	if err != nil {
		beego.Error(err.Error())
	}

	//动态的树结构
	for k, _ := range this.permTree {
		for k1, v1 := range this.permTree[k].SubPermTree {
			this.permTree[k].SubPermTree[k1].SubPermTree = this.GetPermTreeByModuleId(v1.Id)
		}
	}

	return this
}

//获取权限树结构
func (this *permService) GetPermTree() []PermTree {
	return this.permTree
}

//通过module获取树结构
func (this *permService) GetPermTreeByModuleId(moduleId int) []PermTree {
	var subTree []PermTree

	var perms []Perm
	_, _ = ConfigService.Orm.QueryTable(ConfigService.PermTable).Filter("module_id", moduleId).All(&perms)

	for _, v := range perms {
		subTree = append(subTree, PermTree{
			Id:   v.Id,
			Name: v.Name,
		})
	}

	return subTree
}

//通过角色获取权限数据,支持多角色
func (this *permService) GetPermsByRoleId(statusOn bool, roleIds ...int) []Perm {
	var res []Perm
	var roles []Role
	var roleIdStr string

	//过滤禁用状态的角色id
	q := ConfigService.Orm.QueryTable(ConfigService.RoleTable).Filter("id__in", roleIds)
	if statusOn {
		q = q.Filter("status", RoleStatusEnable)
	}

	_, _ = q.All(&roles)

	for _, v := range roles {
		roleIdStr += strconv.Itoa(v.Id) + ","
	}
	roleIdStr = strings.TrimRight(roleIdStr, ",")
	if roleIdStr == "" {
		return res
	}

	//构建查询器
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("p.id", "p.name", "p.perm_url", "p.module_id").
		From(ConfigService.PermTable + " as p").
		InnerJoin(ConfigService.RolePermTable + " as rp").On("rp.perm_id = p.id").
		Where(fmt.Sprintf("rp.role_id in (%s)", roleIdStr)).
		GroupBy("p.id")

	// 执行 SQL 语句
	_, _ = ConfigService.Orm.Raw(qb.String()).QueryRows(&res)

	return res
}

//通过用户Id获取权限信息
func (this *permService) GetPermsByUserId(userId int) []Perm {
	var res []Perm
	roles := RoleService.GetRolesByUserId(userId)

	var roleIds []int
	for _, v := range roles {
		roleIds = append(roleIds, v.Id)
	}

	if len(roleIds) == 0 {
		return res
	}

	return this.GetPermsByRoleId(true, roleIds...)

}

//通过用户id获取权限数据
func (this *permService) GetPermUrlsByUserId(userId int) []string {
	//缓存前缀
	permUrls := []string{}
	if userId == 0 {
		return permUrls
	}

	//从缓存中获取权限数据
	permUrls, ok := CacheService.Get(ConfigService.UserPermCachePreFix + strconv.Itoa(userId)).([]string)
	if ok {
		return permUrls
	}

	//从数据库获取权限数据
	perms := this.GetPermsByUserId(userId)
	for _, v := range perms {
		permUrls = append(permUrls, v.PermUrl)
	}

	//将数据存入缓存
	_ = CacheService.Put(ConfigService.UserPermCachePreFix+strconv.Itoa(userId), permUrls, 48*time.Hour)

	return permUrls

}

//通过角色id刷新用户缓存
func (this *permService) FlushUserPermsCache(roleId int) {
	var roleUsers []RoleUser
	_, _ = ConfigService.Orm.QueryTable(ConfigService.RoleUserTable).Filter("role_id", roleId).All(&roleUsers)

	for _, v := range roleUsers {
		_ = CacheService.Delete(ConfigService.UserPermCachePreFix + strconv.Itoa(v.UserId))
	}

}

////判定用户是否拥有权限
//func HasPerm(user User, permUrl string) bool {
//
//	//超级用户
//	if user.UserType == User_Type_Admin {
//		return true
//	}
//
//	//在权限表中无数据,判定为有权限访问
//	var isPermUrl = false
//	for _, v := range this.perms {
//		if v.PermUrl == permUrl {
//			isPermUrl = true
//			break
//		}
//
//	}
//	if !isPermUrl {
//		return true
//	}
//
//	//判定权限
//	return utils.InArray(permUrl, this.GetPermUrlsByUserId(user.Id))
//}
