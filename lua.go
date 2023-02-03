package request

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"github.com/vela-ssoc/vela-kit/vela"
	"github.com/vela-ssoc/vela-kit/lua"
	"net/http"
	"strings"
)

var xEnv vela.Environment

func ByHeader(L *lua.LState) int {
	r := New().R()
	r.H(L)
	L.Push(r)
	return 1
}

func ByParam(L *lua.LState) int {
	r := New().R()
	SetQueryParam(L, r, L.Get(1))
	L.Push(r)
	return 1
}

func ByBody(L *lua.LState) int {
	var buf bytes.Buffer
	r := New().R()
	n := L.GetTop()
	if n == 0 {
		goto done
	}

	for i := 1; i <= n; i++ {
		buf.WriteString(L.Get(i).String())
	}

	r.SetBody(buf.Bytes())
done:
	L.Push(r)
	return 1

}

func ByInsecure(L *lua.LState) int {
	cli := New()
	cli.SetTLSClientConfig(&tls.Config{
		InsecureSkipVerify: true,
	})

	r := cli.R()
	r.H(L)
	L.Push(r)
	return 1
}

func rawL(L *lua.LState) int {
	raw := L.CheckString(1)

	reader := bufio.NewReader(strings.NewReader(raw))
	req, err := http.ReadRequest(reader)
	if err != nil {
		L.Push(NewRespE(nil, err))
		return 1
	}

	cli := New()
	r := cli.NewRequest()
	r.RawRequest = req

	resp, err := cli.execute(r)
	if err != nil {
		L.Push(NewRespE(r, err))
		return 1
	}

	L.Push(resp)
	return 1
}

func indexL(L *lua.LState, key string) lua.LValue {
	switch key {

	case "client":
		return New()

	case "H":
		return lua.NewFunction(ByHeader)
	case "param":
		return lua.NewFunction(ByParam)
	case "k":
		return lua.NewFunction(ByInsecure)
	case "body":
		return lua.NewFunction(ByBody)

	case "GET", "POST", "PUT", "HEAD", "OPTIONS", "PATCH", "DELETE", "TRACE":
		r := New().R()
		r.Method = key
		return L.NewFunction(r.exec)

	case "TLS":
		return L.NewFunction(newLuaTlsInfo)

	case "raw":
		return L.NewFunction(rawL)

	case "save":
		r := New().R()
		return L.NewFunction(r.save)
	}
	return lua.LNil
}

/*
	local r = http.H("cookie:12312312312313111").GET("http://www.baidu.com").case("code = 200").pipe(print)
	local r = http.H("cookie:12312312312313111").H().H().H().GET("http://www.baidu.com").case("code = 200").pipe(print)

	http.k(true)
		.H("Host:www.baidu.com")
		.H("Content-Type:123")
		.P("a=123").body("123")
		.GET("http://www.baidu.com")
		.case("code = 200")
		.pipe(function(r)
			local v = vela.json(r.body)
			print(v["name"])
		end)
*/

func WithEnv(env vela.Environment) {
	xEnv = env
	env.Global("http", lua.NewExport("vela.http.export", lua.WithIndex(indexL)))
}
