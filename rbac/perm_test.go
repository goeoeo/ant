package rbac

import (
	"testing"
	"yungengxin2019/lwyweb/utils"
)

func TestNewPerm(t *testing.T) {

	perm := NewPerm("../conf/permtree.json")

	if len(perm.perms) == 0 {
		t.Error("初始化数据失败")
	}
	utils.JsonPrint(perm.permTree)

}

func TestPermService_GetPermTreeByModuleId(t *testing.T) {
	perm := NewPerm("../conf/permtree.json")

	res := perm.GetPermTreeByModuleId(4)
	utils.JsonPrint(res)

}

func TestPermService_GetPermUrlsByUserId(t *testing.T) {
	perm := NewPerm("../conf/permtree.json")

	_ = perm.GetPermUrlsByUserId(14)

	t.Log("缓存数据")
	res := CacheService.Get(perm.userPermCachePreFix + "14")
	utils.JsonPrint(res)

	t.Log("清除缓存数据")
	_ = CacheService.Delete(perm.userPermCachePreFix + "14")

	res = CacheService.Get(perm.userPermCachePreFix + "14")
	utils.JsonPrint(res)

}
