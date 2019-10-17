package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/vivafoxdirector/gomon/common"
)

// You will test with insomnia tool
/* exe: # POST
$ curl -i -H 'Content-Type: application/json' \
-d '{"Code":"JP","Name":"Japan"}' http://127.0.0.1:8080/countries

HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
X-Powered-By: go-json-rest
Date: Tue, 18 Dec 2018 05:52:20 GMT
Content-Length: 37

{
 "Code": "JP",
 "Name": "Japan"
}
*/

type ServerInfoHolder struct {
	sync.RWMutex
	m map[string]*common.Record
}

var serverInfoHolder = ServerInfoHolder{m: make(map[string]*common.Record)}

func getServerAllInfo(w rest.ResponseWriter, r *rest.Request) {
	serverInfoHolder.RLock()
	var record []common.Record
	for _, v := range serverInfoHolder.m {
		record = append(record, *v)
	}
	serverInfoHolder.RUnlock()

	if record == nil {
		rest.NotFound(w, r)
		return
	}
	w.WriteJson(record)
}

// Server정보 취득
func getServerInfo(w rest.ResponseWriter, r *rest.Request) {
	// 파라메타 가져오기
	code := r.PathParam("code")

	// 정보 가져오기
	serverInfoHolder.RLock()
	var record *common.Record
	if serverInfoHolder.m[code] != nil {
		record = new(common.Record)
		*record = *serverInfoHolder.m[code]
	}
	serverInfoHolder.RUnlock()

	if record == nil {
		rest.NotFound(w, r)
		return
	}
	w.WriteJson(record)
}

// Server정보 저장
func setServerInfo(w rest.ResponseWriter, r *rest.Request) {
	record := new(common.Record)

	err := r.DecodeJsonPayload(&record)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if record.Code == "" {
		rest.Error(w, "record code required", 400)
		return
	}

	serverInfoHolder.Lock()
	serverInfoHolder.m[record.Code] = record
	serverInfoHolder.Unlock()
}

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/recv/server", getServerInfo),  // Read Agent Info
		rest.Post("/send/server", setServerInfo), // Write Agent Info
	)

	if err != nil {
		log.Fatal(err)
	}
	go task()

	// 등록
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}

func task() {
	for {
		// 1.
		fmt.Println("task...")
		time.Sleep(time.Second * 1)
	}
}
