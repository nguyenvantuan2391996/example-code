package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	pprof.Register(r)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "pong"})
	})

	log.Println("Server running on :8080")
	r.Run(":8080")
}

//| Route                       | Description               |
//| --------------------------- | ------------------------- |
//| `/debug/pprof/`             | Index pprof               |
//| `/debug/pprof/cmdline`      | Command line args         |
//| `/debug/pprof/profile`      | CPU profile (30s default) |
//| `/debug/pprof/symbol`       | Symbol table              |
//| `/debug/pprof/trace`        | Execution trace           |
//| `/debug/pprof/heap`         | Heap profile              |
//| `/debug/pprof/goroutine`    | Goroutines stack          |
//| `/debug/pprof/block`        | Block profile             |
//| `/debug/pprof/threadcreate` | Thread create profile     |

// go tool pprof http://localhost:8080/debug/pprof/profile?seconds=30
// go tool pprof pprof.___go_build_test_bus.samples.cpu.001.pb.gz
// (pprof) top
// (pprof) list Foo
// (pprof) web
// (pprof) png
// (pprof) pdf
// (pprof) svg
