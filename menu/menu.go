package services

//
//import (
//	"context"
//	"github.com/phpdi/ant/stringutil"
//	"encoding/json"
//	"fmt"
//	"io/ioutil"
//	"log"
//	"net/url"
//	"reflect"
//	"strings"
//)
//
//type Menu struct {
//	Id         int    //菜单id,用于处理不同菜单的处理情况
//	Name       string //菜单名称
//	Action     string //链接地址(例GameController.List)
//	Url        string //完整的url地址
//	Icon       string //图标参看bootstrap
//	Active     bool   //是否是激活状态
//	Submenu    []Menu //子菜单数组
//	SubmenuFun string //子菜单为空时调用的函数
//}
//
////存储menu的cookie键
//const (
//	CookieKeyCurUrl   = "curUrl"   //菜单url
//	CookieKeyMenuType = "menuType" //菜单类型
//)
//
////菜单服务
//type MenuService struct {
//	ctx         *context.Context //用于操作cookie
//	curUrl      string           //当前url
//	curMenuType string           //当前菜单类型
//	menuUrls    []string         //所有的菜单Url集合
//
//	breadCrumbs []Menu //面包屑
//	menus       []Menu //菜单结构
//
//}
//
////初始化函数
//func NewMenuService(ctx *context.Context, filePath string) *MenuService {
//	c := &MenuService{
//		////curUser: entity.User{
//		////	Id:       1,
//		////	UserType: entity.User_Type_Admin,
//		////},
//		ctx: ctx,
//	}
//
//	c.initMenu(filePath)
//
//	return c
//}
//
////初始化菜单数据
//func (this *MenuService) initMenu(filePath string) {
//
//	//配置文件中基础菜单数据
//	content, err := ioutil.ReadFile(filePath)
//	if err != nil {
//		log.Println(err.Error())
//	}
//	err = json.Unmarshal(content, &this.menus)
//	if err != nil {
//		log.Println(err.Error())
//	}
//
//}
//
////获取二级菜单信息
//func (this *MenuService) GetMenus() (one []Menu, two Menu) {
//
//	//从url或者cookie获取出当前菜单类型
//	//this.setCurMenuType()
//
//	//生成菜单树
//	this.makeMenuByType(this.curMenuType)
//
//	//采集所有的菜单url
//	this.recursiveCollectionMenuUrl(this.menus)
//
//	//设置当前激活的菜单url
//	//this.setCurUrl()
//
//	//激活菜单
//	this.recursiveActiveMenus(this.menus, this.curUrl)
//
//	////根据权限置空菜单
//	//this.recursiveDropMenu(this.menus)
//
//
//	return
//}
//
////根据菜单类型递归生成子菜单
//func (this *MenuService) makeMenuByType(menuType string) {
//
//	for k, v := range this.menus {
//		if v.Name == menuType {
//			this.callSubMenuFun(&this.menus[k], reflect.ValueOf(this))
//		}
//	}
//
//}
//
////获取当前菜单类型的第一个可进入菜单
//func (this *MenuService) GetFirstMenu(menuType string) Menu {
//	//从url或者cookie获取出当前菜单类型
//	//this.setCurMenuType()
//
//	//生成菜单树
//	this.makeMenuByType(this.curMenuType)
//
//	//根据权限置空菜单
//	//this.recursiveDropMenu(this.menus)
//
//	var curMenu Menu
//	for _, v := range this.menus {
//		if v.Name == menuType {
//			curMenu = v
//		}
//	}
//
//	return recursiveMenuGetFirstAvailableChildMenu(curMenu)
//
//}
//
////设置菜单类型
////1从url中获取,2.从cookie中获取
////func (this *MenuService) setCurMenuType() *MenuService {
////	var menuType string
////
////	//从url中获取
////	if menuType == "" {
////		menuType = this.ctx.Request.Form.Get(CookieKeyMenuType)
////	}
////
////	//从cookie中获取
////	if menuType == "" {
////		menuType = this.GetCookie(CookieKeyMenuType)
////	}
////
////	//设置当前菜单类型
////	this.curMenuType = menuType
////
////	this.SetCookie(CookieKeyMenuType, this.curMenuType)
////
////	return this
////
////}
//
////设置当前菜单url,依赖this.menuUrls
////func (this *MenuService) setCurUrl() *MenuService {
////	//当前url
////	browserUrl, _ := url.QueryUnescape(this.ctx.Request.URL.String())
////
////	if this.isMenuUrl(browserUrl) {
////		//浏览器url是菜单url
////		this.curUrl = browserUrl
////
////		//设置当前菜单到cookie
////		this.SetCookie(CookieKeyCurUrl, this.curUrl)
////	} else {
////		//从cookie中获取菜单
////		this.curUrl = this.GetCookie(CookieKeyCurUrl)
////	}
////
////	return this
////}
//
////判定url是否为菜单的url
//func (this *MenuService) isMenuUrl(browserUrl string) bool {
//	return stringutil.inArray(browserUrl, this.menuUrls)
//}
//
//
//
//
//
////递归一个菜单返回第一个可用(可点击)的菜单
//func recursiveMenuGetFirstAvailableChildMenu(menu Menu) Menu {
//	var res Menu
//
//	for _, v := range menu.Submenu {
//		if v.Url != "" {
//			res = v
//			break
//		}
//		tmp := recursiveMenuGetFirstAvailableChildMenu(v)
//
//		if tmp.Url != "" {
//			res = tmp
//			break
//		}
//	}
//	return res
//}
//
////递归激活菜单
//func (this *MenuService) recursiveActiveMenus(menu []Menu, curUrl string) bool {
//
//	for k, _ := range menu {
//
//		if menu[k].Url == curUrl {
//			menu[k].Active = true
//
//			this.breadCrumbs = append([]Menu{menu[k]}, this.breadCrumbs...)
//			return true
//		}
//
//		if len(menu[k].Submenu) == 0 {
//			continue
//		}
//
//		if this.recursiveActiveMenus(menu[k].Submenu, curUrl) {
//			menu[k].Active = true
//			this.breadCrumbs = append([]Menu{menu[k]}, this.breadCrumbs...)
//			return true
//		}
//	}
//
//	return false
//}
//
////调用生成子菜单的函数
//func (this *MenuService) callSubMenuFun(menu *Menu, objV reflect.Value) {
//
//	if len(menu.Submenu) > 0 {
//		for k, _ := range menu.Submenu {
//			this.callSubMenuFun(&menu.Submenu[k], objV)
//		}
//
//	} else {
//
//		//根据SubmenuFun去生成子菜单
//		if menu.SubmenuFun == "" {
//			return
//		}
//
//		f := objV.MethodByName(menu.SubmenuFun)
//		if (f == reflect.Value{}) {
//			return
//		}
//
//		params := make([]reflect.Value, 1)
//		params[0] = reflect.ValueOf(menu)
//		f.Call(params)
//
//	}
//
//}
//
////递归采集菜单url
//func (this *MenuService) recursiveCollectionMenuUrl(menu []Menu) {
//
//	for k, _ := range menu {
//
//		if menu[k].Url != "" {
//			this.menuUrls = append(this.menuUrls, menu[k].Url)
//		}
//
//		this.recursiveCollectionMenuUrl(menu[k].Submenu)
//	}
//
//}
//
////递归删除没有权限的菜单
////func (this *MenuService) recursiveDropMenu(menu []Menu) {
////	for k, v := range menu {
////
////		route := strings.Split(v.Action, ".")
////		if len(route) == 2 && !this.auth.HasAccessPerm(route[0], route[1]) {
////			menu[k] = Menu{}
////			continue
////		} else {
////			if len(menu) > k && len(menu[k].Submenu) > 0 {
////				this.recursiveDropMenu(menu[k].Submenu)
////			}
////		}
////
////	}
////}
//
//
//
////生成菜单url
////生成菜单url的目的有两个
////1.将生成的url装入this.MenuUrls中,2.beego.URLFor 生成的菜单的url参数是无序的,这会导致后面的激活菜单失效
//func (this *MenuService) createUrl(url string, params ...interface{}) string {
//
//	if len(params)%2 != 0 {
//		return ""
//	}
//
//	out := beego.URLFor(url)
//
//	pathName := ""
//	for k, v := range params {
//		if k%2 == 0 {
//			pathName += fmt.Sprintf("&%v=", v)
//		} else {
//			pathName += fmt.Sprintf("%v", v)
//		}
//	}
//
//	if pathName != "" {
//		out = out + "?" + strings.TrimLeft(pathName, "&")
//	}
//
//	return out
//
//}
//
////设置cookie
//func (this *MenuService) SetCookie(key string, value string) {
//
//	this.ctx.SetCookie(key, url.QueryEscape(value), 86400)
//}
//
////获取cookie
//func (this *MenuService) GetCookie(key string) string {
//
//	res := this.ctx.GetCookie(key)
//
//	res, _ = url.QueryUnescape(res)
//
//	return res
//}
//
////判定当前菜单是否有子菜单
//func HaveSubMenu(menu []Menu) bool {
//
//	var name bool
//
//	for _, v := range menu {
//		if v.Name != "" {
//			name = true
//			break
//		}
//	}
//
//	if len(menu) > 0 && name {
//		return true
//	}
//
//	return false
//
//}
