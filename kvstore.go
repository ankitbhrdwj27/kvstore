package main

import (
	"net/http" 
	"io"
	"fmt"
	"sync"
	"strings"
	)

var mutex sync.Mutex
var hash map[string]string




func main() {
mutex.Lock()
hash=make(map[string]string)
hash["first"]="demo"
mutex.Unlock()
http.HandleFunc("/index", index)
go http.HandleFunc("/insert/", insert)
go http.HandleFunc("/deletekv/", deletekv)
go http.HandleFunc("/check/", check)
http.ListenAndServe(":8080", nil)
}




func index(res http.ResponseWriter, req *http.Request) {
res.Header().Set("Content-Type","text/html")
io.WriteString(res,
`<doctype html>
<html>
<head>
<title>kvstore</title>
</head>
<body>
<h2><center>KVSTORE-HTTP SERVER IS WORKING !!!<center></h2>
<ul>
<li><h3>http://localhost:8080/insert/KEY=VALUE</h3>
<li><h3>http://localhost:8080/deletekv/KEY</h3>
<li><h3>http://localhost:8080/check/KEY</h3>
</ul>
</body>
<footer><sup>**</sup>Key and Value can be of string type only.And key-value store work properly for the above given commands.
</html>`,
)
}




func insert(res http.ResponseWriter, req *http.Request) {
key:=strings.Split(req.URL.Path[8:],"=")
var Key,Value string
if key[0]!=""{
Key=key[0]
}
if key[1]!=""{
Value=key[1]
}
if Key!="" && Value!=""{
mutex.Lock() 
hash[Key]=Value
fmt.Fprintf(res,"KEY = %s \t VALUE = %s",Key,Value)
mutex.Unlock()
} else{
fmt.Fprintf(res,"Give argumets properly in the form  KEY=VALUE")
}
}



func deletekv(res http.ResponseWriter, req *http.Request) {
key:=strings.Split(req.URL.Path[10:],"\n")
mutex.Lock()
if key[0]!="" && hash[key[0]]!=""{
fmt.Fprintf(res,"Pair deleted is KEY=%s \t VALUE=%s",key[0],hash[key[0]])
delete(hash,key[0])
}else{
fmt.Fprintf(res,"Given KEY=VALUE pair is not available in KVSTORE")
}
mutex.Unlock()
}



func check(res http.ResponseWriter, req *http.Request) {
key:=strings.Split(req.URL.Path[7:],"\n")
mutex.Lock()
if hash[key[0]]!=""{
fmt.Fprintf(res,"KEY=%s \t VALUE=%s",key[0],hash[key[0]])
}else{
fmt.Fprintf(res,"Given KEY=VALUE pair is not available in KVSTORE")
}
mutex.Unlock()
}
