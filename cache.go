package goproxy


type respData struct {
	Data []byte
	Finished bool
	Connect bool
	Md5  string
}

func (res *respData) M5() {
	res.Md5 = Md5(res.Data)
}

var Data  map[string]*respData

func init() {
	Data = make(map[string]*respData)
}