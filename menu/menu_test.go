package services

import (
	"fmt"
	"github.com/astaxie/beego/context"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"testing"
	"yungengxin2019/lwyweb/utils"
)

var curUser = NewAuth()

var ctx = context.NewContext()

func init() {
	ctx.Request = &http.Request{
		Form: url.Values{"mid": {"33810"}, "cid": {"%7b1AF9A1A1-FF46-41B7-BC7A-D1FAE3E25B01%7d"}, "page": {"3"}},
		URL:  &url.URL{Path: "/netbar/list", RawPath: "menuType=网吧管理"},
	}

}

//菜单服务实例化
func TestNewMenuService(t *testing.T) {

	menuService := NewMenuService(curUser, ctx, "../conf/menu.json")

	fmt.Printf("%+v", menuService)
}

//根据Url获取当前菜单列表
func TestMenuService_GetMenuList(t *testing.T) {
	menuService := NewMenuService(curUser, ctx, "../conf/menu.json")

	_, _ = menuService.GetMenus()

	//utils.JsonPrint(menuTop)
	//utils.JsonPrint(menuLeft)
	//utils.JsonPrint(menuService.curUrl)
	//utils.JsonPrint(menuService.menuUrls)

}

//根据菜单类型获取第一个子菜单数据
func TestMenuService_GetFirstMenu(t *testing.T) {
	menuService := NewMenuService(curUser, ctx, "../conf/menu.json")

	menu := menuService.GetFirstMenu("网吧管理")

	utils.JsonPrint(menu)
	utils.JsonPrint(menuService.curUrl)
}

func TestCallFun(t *testing.T) {
	menuService := NewMenuService(curUser, ctx, "../conf/menu.json")
	objV := reflect.ValueOf(menuService)
	f := objV.MethodByName("Gett")

	params := make([]reflect.Value, 1)
	menu := &Menu{Id: 2}

	params[0] = reflect.ValueOf(menu)
	f.Call(params)

	fmt.Println(menu.Id)
}

func TestUrlDecode(t *testing.T) {
	res := url.QueryEscape("网吧管理test")

	log.Println("编码后:", res)

	res, _ = url.QueryUnescape(res)

	log.Println("解码后:", res)

}
