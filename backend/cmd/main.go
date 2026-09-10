package main

import (
	"log"

	"fnMusicServe/interal/repository"
	"fnMusicServe/interal/router"
)

func main() {
	db := repository.InitDB()

	r := router.InitRouter(db)
	log.Println("服务启动,端口:8061")
	if err := r.Run(":8061"); err != nil {
		log.Fatalf("服务启动失败:%v", err)
	}
}
