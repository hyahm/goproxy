package main

import (
	"fmt"
	"github.com/hyahm/goproxy"
	"github.com/hyahm/xmux"
	"log"
	"net/http"
	"net/url"
)

func test(w http.ResponseWriter,r *http.Request) {
	url,err := url.Parse("http://18.139.189.181:80")
	if err != nil {
		fmt.Println(err.Error())
	}
	if v, ok := goproxy.Data[r.URL.Path]; ok {
		if v.Finished {
			w.Write(v.Data)
			return
		}
	}
	proxy := goproxy.NewSingleHostReverseProxy(url)
	proxy.CacheData = true
	proxy.ServeHTTP(w, r)


}

func main() {
	fmt.Println(32<<10)
	router := xmux.NewRouter()
	router.Pattern("/{string:name}").Get(test)
	fmt.Println("listen on :8000")
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatal(err)
	}
}