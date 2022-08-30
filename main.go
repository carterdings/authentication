package main

import (
	"flag"
	stdlog "log"
	"os"

	"github.com/carterdings/authentication/logic"
	"github.com/carterdings/authentication/repo/log"
	"github.com/gin-gonic/gin"
)

var (
	addr    string
	logPath string
)

func main() {
	flag.StringVar(&addr, "addr", "127.0.0.1:8080", "server ip:port")
	flag.StringVar(&logPath, "log", "gin.log", "log file path")
	if !flag.Parsed() {
		flag.Parse()
	}

	initAndServe(addr, logPath)
}

func initAndServe(addr, logPath string) {
	privKey, pubKey, err := logic.GenRsaKey()
	if err != nil {
		stdlog.Fatal(err)
	}
	s := &Service{AuthService: logic.NewServie(privKey, pubKey)}

	gin.DisableConsoleColor()
	f, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		stdlog.Fatal(err)
	}
	log.InitLogger(f)

	router := gin.Default()

	router.POST("/user/create", s.CreateUser)
	router.POST("/user/delete", s.DeleteUser)
	router.POST("/user/add_role", s.AddRoleToUser)
	router.POST("/user/check_role", s.CheckRole)
	router.POST("/user/all_roles", s.AllRoles)
	router.POST("/role/create", s.CreateRole)
	router.POST("/role/delete", s.DeleteRole)
	router.POST("/auth/authenticate", s.Authenticate)
	router.POST("/auth/invalidate", s.Invalidate)

	router.Run(addr)
}
