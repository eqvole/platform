package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/spf13/viper"

	"github.com/eqvole/platform/pkg/common"
	"github.com/eqvole/platform/pkg/rest"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	conf := common.Opts{}
	err = viper.Unmarshal(&conf)
	if err != nil {
		panic(err)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		<-sigChan
		cancel()
	}()

	srv := http.Server{
		Handler: rest.NewRouter(),
		Addr:    fmt.Sprintf(":%d", conf.Port),
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			log.Print(err)
			cancel()
		}
	}()
	<-ctx.Done()
}
