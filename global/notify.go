package global

type Notify struct {
	Id      int64
	Code    string
	Target  []string
	Message map[string]string
	Title   string
}

var _ch chan Notify

func init() {
	_ch = make(chan Notify, 1000)
}

func SendNotify(item Notify) {
	_ch <- item
}

func GetChannel() chan Notify {
	return _ch
}
