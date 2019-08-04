package main

import "fmt"
import "bufio"
import "net"
import "net/http"
import "net/http/httputil"
import "io/ioutil"
import "strings"

func main() {
	listen, _ := net.Listen("tcp", "localhost:9999")
	fmt.Println("서버기동@http://localhost:9999")
	for {
		conn, _ := listen.Accept()
		go func() {
			fmt.Println("Remote Addr: ", conn.RemoteAddr())

			request, _ := http.ReadRequest(bufio.NewReader(conn))
			dump, _ := httputil.DumpRequest(request, true)
			fmt.Println(string(dump))

			response := http.Response{
				StatusCode: 200,
				ProtoMajor: 1,
				ProtoMinor: 0,
				Body:       ioutil.NopCloser(strings.NewReader("Hello World.\n")),
			}
			response.Write(conn)
			conn.Close()
		}()
	}
}
