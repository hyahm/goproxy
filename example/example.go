package main

import (
	"fmt"
	"github.com/hyahm/goproxy"
	"github.com/hyahm/xmux"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func test(w http.ResponseWriter,r *http.Request) {
	url,err := url.Parse("http://18.139.189.181:80")
	if err != nil {
		fmt.Println(err.Error())
	}
	md5 := r.Header.Get("Md5")
	if md5 == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if md5 != goproxy.Data[r.URL.Path].Md5 {
		goproxy.Data[r.URL.Path]= nil

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
	b , _:= ioutil.ReadFile("/root/下载/SignTest.ipa")
	m := goproxy.Md5(b)
	fmt.Println(m)
	router := xmux.NewRouter()
	router.SetHeader("Access-Control-Allow-Headers", "Md5")
	router.Pattern("/{string:name}").Get(test)
	fmt.Println("listen on :8000")
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatal(err)
	}
}