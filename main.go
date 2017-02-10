package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/CardInfoLink/log"
	"github.com/wonsikin/queuing/goconf"
)

func main() {
	log.SetLevel(log.DebugLevel)

	addr := ":7002"
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/seq", handler)

	log.Infof("seq http is listening, addr = %s", addr)
	log.Error(http.ListenAndServe(addr, nil))
}

// 每一次请求都会生成一个新的序号
// 序号格式为 YMMDD%d。 %d为至少两位的数字
func handler(w http.ResponseWriter, r *http.Request) {
	s := goconf.SeqTick()
	d := time.Now().Format("060102")
	result := fmt.Sprintf("%s%02d", d[1:], s)

	msg := fmt.Sprintf(`{"seq": "%s"}`, result)
	w.Write([]byte(msg))
}
