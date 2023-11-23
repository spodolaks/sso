package sso

import (
	"fmt"

	"github.com/spodolaks/sso/internal/config"
)


func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)

	//TODO init logger

	//TODO init app
}