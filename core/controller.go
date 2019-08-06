package core

import (
	"net/http"
	"encoding/json"
	"reflect"
	"lazyadm/core/library"
	"strings"
)

type IControllerRun interface {
	Run(action string)
}

type Controller struct {
	actions map[string]func()
	Res     http.ResponseWriter
	Req     *http.Request
	name    string
	Self    IControllerRun
	tpl     *Template
}

var actionsMap = map[string]map[string]int{}

func (c *Controller) Init(name string, self IControllerRun, res http.ResponseWriter, req *http.Request) {
	c.Res = res
	c.Req = req
	c.name = name
	c.Self = self
	c.tpl = NewTemplate(res, req, "templates/", ".html")

	if _, ok := actionsMap[name]; !ok {
		actionsMap[name] = map[string]int{}
		getType := reflect.TypeOf(self)
		// 获取方法
		for i := 0; i < getType.NumMethod(); i++ {
			m := getType.Method(i)
			methodName := m.Name
			if strings.HasSuffix(methodName, "Action") {
				methodName = methodName[0: len(methodName)-6]
				actionsMap[name][methodName] = 1
			}
		}
	}
}

func (c *Controller) GetString(name string) string {
	return library.ToString(c.Req.URL.Query().Get(name))
}

func (c *Controller) GetInt(name string) int {
	return library.ToInt(c.Req.URL.Query().Get(name))
}

func (c *Controller) PostString(name string) string {
	return library.ToString(c.Req.FormValue(name))
}

func (c *Controller) PostInt(name string) int {
	return library.ToInt(c.Req.FormValue(name))
}

func (c *Controller) Run(action string) {

	if action == "" {
		action = "Index"
	}
	action = library.ToCamelString(action)
	if c.actions == nil {
		c.actions = map[string]func(){}
	}
	if _, ok := c.actions[action]; !ok {
		if _, ok := actionsMap[c.name][action]; !ok {
			c.Abort(404, "404 page not found")
			return
		}
		getValue := reflect.ValueOf(c.Self)
		methodValue := getValue.MethodByName(action + "Action")
		if !methodValue.IsValid() {
			c.Abort(404, "404 page not found")
			return
		}
		method := methodValue.Interface().(func())
		c.actions[action] = method
	}
	c.actions[action]()
}

func (c *Controller) Abort(httpCode int, message string) {
	c.Res.WriteHeader(httpCode)
	c.Res.Write([]byte(message))
}

func (c *Controller) SetHeader(headerOptions map[string]string) *Controller {
	if len(headerOptions) > 0 {
		for field, val := range headerOptions {
			c.Res.Header().Set(field, val)
		}
	}
	return c
}

func (c *Controller) Response(code int, message string, data map[string]interface{}) {
	result := map[string]interface{}{
		"code":    code,
		"message": message,
		"data":    data,
	}
	jsonData, err := json.Marshal(result)
	if err != nil {
		jsonData = nil
	}
	c.Res.Header().Set("Content-Type", "application/json; charset=utf-8")
	c.Res.Write(jsonData)
}

func (c *Controller) AssignMap(data map[string]interface{}) *Template {
	return c.AssignMap(data)
}

func (c *Controller) Assign(key string, data interface{}) *Template {
	return c.Assign(key, data)
}

func (c *Controller) Display(tpl string) {
	c.tpl.Display(tpl)
}
