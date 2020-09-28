package web

import (
	"encoding/binary"
	"encoding/json"
	"net/http"
	"sync"

	"github.com/gofiber/fiber"
	"github.com/spaolacci/murmur3"
	"shortLink/base62"
)

// Web web服务
type Web struct {
	port        int               // 端口
	app         *fiber.App        // fiber应用
	urlMap      map[string]string // url映射
	urlMapMutex sync.RWMutex      // url映射读写锁
}

// Run 运行
func (object *Web) Run() (err error) {
	// 短连接重定向
	object.app.Get("/redirect/:code", func(ctx *fiber.Ctx) {
		code := ctx.Params("code", "")
		if 0 >= len(code) {
			ctx.SendStatus(http.StatusBadRequest)
			return
		}
		object.urlMapMutex.RLock()
		url := object.urlMap[code]
		object.urlMapMutex.RUnlock()
		if 0 >= len(url) {
			ctx.SendStatus(http.StatusBadRequest)
			return
		}
		ctx.Redirect(url, http.StatusFound)
	})
	// 添加链接
	object.app.Post("/add", func(ctx *fiber.Ctx) {
		url := ctx.FormValue("url")
		if 0 >= len(url) {
			ctx.SendStatus(http.StatusBadRequest)
			return
		}
		hash := murmur3.New32()
		hash.Write([]byte(url))
		code := base62.Inst.Encode(binary.BigEndian.Uint32(hash.Sum(nil)))
		object.urlMapMutex.Lock()
		defer object.urlMapMutex.Unlock()
		object.urlMap[string(code)] = url
		ctx.SendString("/redirect/" + string(code))
	})
	// 删除链接
	object.app.Post("/remove/:code", func(ctx *fiber.Ctx) {
		code := ctx.Params("code", "")
		if 0 >= len(code) {
			ctx.SendStatus(http.StatusBadRequest)
			return
		}
		object.urlMapMutex.Lock()
		url := object.urlMap[code]
		delete(object.urlMap, code)
		object.urlMapMutex.Unlock()
		raw, err := json.Marshal(&struct {
			Code string `json:"code"`
			Url  string `json:"url"`
		}{Code: code, Url: url})
		if nil != err {
			ctx.Next(err)
			return
		}
		ctx.Send(raw)
	})
	return object.app.Listen(object.port)
}

// Shutdown shutdown web
func (object *Web) Shutdown() (err error) {
	err = object.app.Shutdown()
	return
}

// New 工厂方法
func New(port int) *Web {
	return &Web{
		port:   port,
		app:    fiber.New(),
		urlMap: make(map[string]string),
	}
}
