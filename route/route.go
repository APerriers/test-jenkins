package route

import (
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"os"
)

func Router() *gin.Engine {
	//设置gin模式
	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	//ping
	router.GET("/open/ping", func(c *gin.Context) {
		hostname, _ := os.Hostname()
		ips := make([]net.IP, 0)
		ifaces, _ := net.Interfaces()
		for _, i := range ifaces {
			addrs, _ := i.Addrs()
			// handle err
			for _, addr := range addrs {
				var ip net.IP
				switch v := addr.(type) {
				case *net.IPNet:
					ip = v.IP
				case *net.IPAddr:
					ip = v.IP
				}
				ips = append(ips, ip)
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"msg":      "pong",
			"code":     "200",
			"hostname": hostname,
			"IP":       ips,
		})
	})

	return router
}
