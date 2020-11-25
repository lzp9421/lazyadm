package api

import "lazyadm/controller"

type MenuController struct {
	controller.Controller
}

func (ctl *MenuController) SideAction() {

	data := []map[string]interface{}{
		{
			"title":"首页",
			"icon": "layui-icon-home",
			"href": "/",
		},
		{
			"title":"列表页",
			"icon": "layui-icon-unorderedlist",
			"childs":[]map[string]interface{}{
				{
					"title": "表格列表",
					"href":"/app/list/table/id=1",
				},
				{
					"title": "卡片列表",
					"href":"/app/list/card",
				},
			},
		},
		{
			"title":"详情页",
			"icon": "layui-icon-container",
			"childs":[]map[string]interface{}{
				{
					"title": "工作计划",
					"href":"/app/detail/plan",
				},
				{
					"title":"数据统计",
					"href": "/app/chart/index",
				},
			},
		},
		{
			"title":"表单页",
			"icon": "layui-icon-file-exception",
			"childs":[]map[string]interface{}{
				{
					"title": "表单元素",
					"href":"/app/form/basic",
				},
				{
					"title": "表单组合",
					"href":"/app/form/group",
				},
			},
		},
		{
			"title": "新增模块",
			"icon": "layui-icon-experiment",
			"notice":3,
			"childs":[]map[string]interface{}{
				{
					"title": "admin",
					"href":"/app/module/admin",
				},
				{
					"title": "helper",
					"href":"/app/module/helper",
				},
				{
					"title": "loadBar",
					"href":"/app/module/loadbar",
				},
			},
		},
		{
			"title": "多级导航",
			"icon": "layui-icon-apartment",
			"childs":[]map[string]interface{}{
				{
					"title": "Dota2",
					"childs":[]map[string]interface{}{
						{
							"title": "法力损毁",
						},
						{
							"title": "闪烁",
						},
						{
							"title": "法术护盾",
						},
						{
							"title": "法力虚空",
						},
					},
				},
				{
					"title": "帕吉",
					"childs":[]map[string]interface{}{
						{
							"title": "肉钩",
						},
						{
							"title": "腐烂",
						},
						{
							"title": "腐肉堆积",
						},
						{
							"title": "肢解",
						},
					},
				},
				{
					"title": "LOL",
					"childs":[]map[string]interface{}{},
				},
			},
		},
	}
	ctl.ApiSucc(map[string]interface{}{"menu": data}, "aaaaaaa")
}

func (ctl *MenuController) MainAction() {

	data := []map[string]interface{}{
		{
			"title": "开发者",
			"icon": "layui-icon-star",
			"href":"/develop/index",
		},
		{
			"title": "商品管理",
			"icon": "layui-icon-star",
			"href":"/app/index",
		},
		{
			"title": "用户",
			"icon": "layui-icon-star",
			"href":"/user/index",
		},
		{
			"title":"其它系统",
			"icon": "layui-icon-error",
			"childs":[]map[string]interface{}{
				{
					"title": "邮件管理",
					"href":"/others1/index",
				},
				{
					"title": "消息管理",
					"href":"/others2/index",
				},
				{
					"title": "授权管理",
					"href":"/others3/index",
				},
			},
		},
	}
	ctl.ApiSucc(map[string]interface{}{"menu": data}, "aaaa")
}

func (ctl *MenuController) List2Action() {

	data := []map[string]interface{}{
		{
			"title":"异常页",
			"icon": "layui-icon-error",
			"childs":[]map[string]interface{}{
				{
					"title": "403",
					"href":"/exception/403",
				},
				{
					"title": "404",
					"href":"/exception/404",
				},
				{
					"title": "500",
					"href":"/exception/500",
				},
			},
		},
		{
			"title": "图标",
			"icon": "layui-icon-star",
			"href":"/icon/index",
		},
	}
	ctl.ApiSucc(map[string]interface{}{"menu": data}, "aaaaaaa")
}
