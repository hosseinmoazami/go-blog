package routing

import (
	"blog/pkg/config"
	"fmt"
	"log"
)

func Serve() {
	r := GetRouter()

	configs := config.Get()

	r.MaxMultipartMemory = configs.Server.MemoryLimitForMultipartForm << 20 // 8 MiB

	err := r.Run(fmt.Sprintf("%s:%s", configs.Server.Host, configs.Server.Port))
	if err != nil {
		log.Fatal("Error in routing")
	}
}
