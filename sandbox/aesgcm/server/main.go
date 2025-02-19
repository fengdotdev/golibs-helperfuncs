package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"

	"github.com/fengdotdev/golibs-helperfuncs/data"
	"github.com/fengdotdev/golibs-helperfuncs/secret"
)

func ENCODED(w http.ResponseWriter, r *http.Request) {

	///----ONLY FOR TESTING PURPOSES
	queryValue := r.URL.Query().Get("key")
	fmt.Println("key64: ", queryValue)
	key, err := data.Decode64Bytes(queryValue)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	///----ONLY FOR TESTING PURPOSES

	var payload Payload

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(body, &payload)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("payload: ", payload)
	//decode
	fmt.Println("payload.Cypher64: ", payload.Cypher64)
	dataCypher, err := payload.GetCypher()
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("iv64: ", payload.Additionaldata.IV64)
	iv, err := payload.Additionaldata.GetIV()
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("additionaldata: ", payload.Additionaldata)

	addjson, err := payload.GetAdditionalDataAsBinary()
	fmt.Println("additionaldata as binary64: ", data.Encode64Bytes(addjson))

	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//decrypt

	plainText, err := secret.DecryptAESGCM(key, iv, dataCypher, addjson)
	if err != nil {

		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(string(plainText))
	w.Write([]byte(plainText))
}

func Ping(w http.ResponseWriter, r *http.Request) {

	fmt.Println("connected client at " + r.RemoteAddr + " to /ping")

	fmt.Println("sending pong")
	w.Write([]byte("pong"))
}

type Foo struct {
	Bar string `json:"bar"`
}

func FooResource(w http.ResponseWriter, r *http.Request) {

	fmt.Println("connected client at " + r.RemoteAddr + " to /foo")

	foo := Foo{Bar: "baz"}
	fooJSON, err := json.Marshal(foo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("sending foo" + string(fooJSON))

	w.Write(fooJSON)
}

type operation struct {
	Operation string `json:"operation"`
	N1        int    `json:"n1"`
	N2        int    `json:"n2"`
}

type result struct {
	Result int `json:"result"`
}

func Add(w http.ResponseWriter, r *http.Request) {
	var op operation

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(body, &op)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("connected client at " + r.RemoteAddr + " to /add")
	fmt.Println("operation: ", op)

	if op.Operation != "add" {
		http.Error(w, "Invalid operation", http.StatusBadRequest)
		return
	}

	fmt.Println("n1: ", op.N1)
	fmt.Println("n2: ", op.N2)
	fmt.Println("result: ", op.N1+op.N2)

	res := result{Result: op.N1 + op.N2}
	resJSON, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resJSON)
}

func main() {
	// Create a new instance of the server
	server := NewServer()

	// Start the server
	server.Run("8080")

}

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run(port string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	http.HandleFunc("/404", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("404 Not Found")
		http.Error(w, "404 Not Found", http.StatusNotFound)
	})
	http.HandleFunc("/add", Add)
	http.HandleFunc("/foo", FooResource)
	http.HandleFunc("/encoded", ENCODED)
	http.HandleFunc("/ping", Ping)

	if !Assert4charsAndNumbeable(port) {
		panic("Invalid port")
	}
	

	localip,err  := getIP()
	if err != nil {
		fmt.Println("Could not obtain the IP address")
	}
	ip := localip + ":" + port

	fmt.Println("Server running on " + ip)
	http.ListenAndServe(":"+port, nil)
}

func getIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ip4 := ipNet.IP.To4(); ip4 != nil {
				return ip4.String(), nil
			}
		}
	}

	return "", fmt.Errorf("could not obtain the IP address")
}

func Assert4charsAndNumbeable(s string) bool {
	l := len(s)
	if !(l == 4) {
		return false
	}

	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}

	return true
}

type AdditionalData struct {
	Algorithm string `json:"algorithm"` // AES
	Mode      string `json:"mode"`      // GCM
	Strength  int    `json:"strength"`  // 256
	IV64      string `json:"iv64"`      // ex: 32bVr0KW+Cj5pPLB
}

func (a *AdditionalData) GetIV() ([]byte, error) {
	return base64.StdEncoding.DecodeString(a.IV64)
}

type Payload struct {
	Cypher64       string         `json:"cypher64"`
	Additionaldata AdditionalData `json:"additionaldata"`
}

func (p *Payload) GetCypher() ([]byte, error) {
	return base64.StdEncoding.DecodeString(p.Cypher64)
}

func (p *Payload) GetAdditionalDataAsBinary() ([]byte, error) {
	jsonData, err := json.Marshal(p.Additionaldata)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

type PayloadWithKey struct {
	Key64   string  `json:"key64"`
	Payload Payload `json:"payload"`
}

func (p *PayloadWithKey) GetKey() ([]byte, error) {
	return base64.StdEncoding.DecodeString(p.Key64)
}
