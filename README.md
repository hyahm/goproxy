# goproxy
go 代理

基于httputil 包反代的加强版， 原生包不支持获取数据， 此包支持获取传输的数据， 不用话请删除数据

### 开启方法
```
proxy := goproxy.NewSingleHostReverseProxy(url)
proxy.CacheData = true  // 开启缓存数据， Data   key 为url
proxy.ServeHTTP(w, r)
```
