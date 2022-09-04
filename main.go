package main

import (
//	"fmt"
	"io"
	"net/http"
	"os"
	//"strings"
//	"work/logger"
)

func main() {
//	logger.Init()
//	logger.Lg.Info("begin the process")
	http.HandleFunc("/healthz", healthCheck)
	http.HandleFunc("/GetComputeInfo", GetComputeInfo)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
//		logger.Lg.Error("server error 500")
	}
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "200 ok")

}

func GetComputeInfo(w http.ResponseWriter, r *http.Request) {
	os.Setenv("VERSION","v1")
	version := os.Getenv("VERSION")
//	addr := r.RemoteAddr
//	ar := strings.Split(addr, ":")
//	ip := ar[0]
//	port := ar[1]
//	logger.Lg.Info(fmt.Sprintf("client ip is %s port %s", ip, port))
	for k, v := range r.Header {
		w.Header().Set(k, v[0])

	}
	w.Header().Set("version", version)

}

