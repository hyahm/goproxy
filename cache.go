package goproxy


type respData struct {
	Data []byte
	Finished bool
	Connect bool
}



var Data  map[string]*respData

func init() {
	Data = make(map[string]*respData)
}