package example

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main_open() {
	fmt.Println("======================================================================")
	fmt.Println("main")
	fmt.Println("======================================================================")
	gin.SetMode("debug")
	//routersInit := routers.InitRouter()
	readTimeout := 30 * time.Second
	writerTimeout := 30 * time.Second
	endPoint := fmt.Sprintf(":%d", 80)
	maxHeaderBytes := 1 << 20
	server := &http.Server{
		Addr: endPoint,
		//Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writerTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}
	log.Printf("[info] start http server listening %s", endPoint)
	server.ListenAndServe()

}
