# request
> 磐石开发框架默认http请求模块

## 内置信息
- [http.client](#client) &emsp;生成client对象
- [http.H](#请求头设置) &emsp;请求头设置
- [http.param](#请求参数) &emsp;请求参数
- [http.body](#请求内容) &emsp;请求内容
- [http.save](#结果保存) &emsp;保存结果设置

## 直接请求
- [http.GET](#)
- [http.POST](#)
- [http.PUT](#)
- [http.HEAD](#)
- [http.OPTIONS](#)
- [http.PATCH](#)
- [http.DELETE](#)
- [http.TRACE](#)



## client
> cli = http.client <br />
> 参数为空默认创建一个请求的客户端,可以复用

内部函数:

- [before(pipe)](#)&emsp;设置请求之前的动作,参数:[request](#request)
- [after(pipe)](#)&emsp;设置请求之后的动作,参数:[response](#response)
- [auth(string)](#)&emsp;设置认证信息
- [proxy(string)](#)&emsp;设置代理信息

```lua
    local cli = http.client --这里会新建
```

## request
>支持GET,HEAD,POST,PUT,PATCH,DELETE,CONNECT,OPTIONS,TRACE,POST <br />
> 返回: [response](#response)结构体 <br />

内置方法:

- [H(string)](#设置请求头)设置请求头
- [save(string)](#结果保存)结果保存
- [param 或者 P](#设置参数)设置参数
- [body(v)](#请求体)请求体
>

```lua
    local resp = http.H("cookie:aa").H"cookie:123".GET("http://www.vela.com")
    print(resp.ERR)
    print(resp.code)
    print(resp.body)

    local r = http.P("a=123").H("cookie:3")
    r.body("hello") 
    r.GET("http://www.vela.com")
```


## 设置参数
> [request](#request) = http.param("a=123")
```lua
    local r = http.param("a=124&b=456")
```

## 请求头设置
>  [request](#request) = http.H(string) <br />
> 设置请求头  冒号分割
```lua
    local r = http.H("cookie:xxxxx").H("host:ww.vela.com")
```

## 请求体
> [request](#request) = http.body(string) <br />
> 请求内容
```lua
    local r = http.body("hello")
```

## 结果保存
> 这是个一个实现下载的库,可以用保存请求结果到文件
```lua
   http.save("1.html").H("cookie:www.baidu.com").body("helo").POST("http://www.baidu.com")
```

## response
> 返回结果 对象内容

内置方法:
* response.ERR
* response.code
* response.body
* response.header
* [save(string)](#) &emsp;保存到本地路径



```lua
    local r = http.GET("http://www.a.com/1.html")
    print(r.ERR)
    print(r.code)
    print(r.body)
```
