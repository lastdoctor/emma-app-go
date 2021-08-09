package main

import (
	"github.com/lastdoctor/emma-app-go/internal/logger"
	"github.com/lastdoctor/emma-app-go/internal/repository/pg"
)

func main() {
	//ctx := context.Background()

	//Init repository repository
	//repository, err := repository.New(ctx)
	//if err != nil {
	//	logger.Logger().Error(err)
	//}
	dial, err := pg.Dial()
	if err != nil {
		return
	}
	logger.Logger().Info(dial.String())
}
