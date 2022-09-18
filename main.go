package main

import (
        //      "fmt"
        "context"
        "fmt"
        "httpserver/metrics"
        "io"
        "math/rand"
        "net/http"
        "os"
        "os/signal"
        "syscall"
        "time"

        //"strings"
        //      "work/logger"

        "github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
        //      logger.Init()
        //      logger.Lg.Info("begin the process")
        metrics.Register()
        mux := http.NewServeMux()
        mux.HandleFunc("/healthz", healthCheck)
        mux.HandleFunc("/GetComputeInfo", GetComputeInfo)
        mux.HandleFunc("/image", images)
        mux.HandleFunc("/hello", rootHandler)
        mux.Handle("/metrics", promhttp.Handler())

        server := &http.Server{
                Addr:    ":8080",
                Handler: mux,
        }
        go server.ListenAndServe()
        listenSignal(context.Background(), server)
        select {}
}

func images(w http.ResponseWriter, r *http.Request) {
        timer := metrics.NewTimer()
        defer timer.ObserveTotal()
        randInt := rand.Intn(2000)
        time.Sleep(time.Millisecond * time.Duration(randInt))
        w.Write([]byte(fmt.Sprintf("<h1>%d<h1>", randInt)))
}

func listenSignal(ctx context.Context, httpSrv *http.Server) {
        sigs := make(chan os.Signal, 1)
        signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

        select {
        case <-sigs:
                fmt.Println("notify sigs")
                httpSrv.Shutdown(ctx)
                fmt.Println("http shutdown gracefully")
        }
}
func randInt(min, max int) int {
        rand.Seed(time.Now().Unix())
        return rand.Intn(max-min) + min
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
        timer := metrics.NewTimer()
        defer timer.ObserveTotal()
        user := r.URL.Query().Get("user")
        delay := randInt(10, 2000)
        time.Sleep(time.Millisecond * time.Duration(delay))
        if user != "" {
                io.WriteString(w, fmt.Sprintf("hello [#{user}]\n"))
        } else {
                io.WriteString(w, "hello [stranger]\n")
        }
        io.WriteString(w, "==========================Details of the http request header;====================================\n")
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
        io.WriteString(w, "200 ok")

}

func GetComputeInfo(w http.ResponseWriter, r *http.Request) {
        os.Setenv("VERSION", "v1")
        version := os.Getenv("VERSION")
        for k, v := range r.Header {
                w.Header().Set(k, v[0])

        }
        w.Header().Set("version", version)

}

~

