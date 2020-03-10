package main

import (
	"errors"
	"net/http"
	"time"

	"github.com/demos/api_server/config"
	"github.com/demos/api_server/model"
	"github.com/demos/api_server/router"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cfg = pflag.StringP("config", "c", "", "api server config file")
)

func main() {
	//init config
	pflag.Parse()
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	// init db
	model.DB.Init()
	defer model.DB.Close()
	// init router
	gin.SetMode(viper.GetString("runmode"))
	g := gin.New()
	middlewares := []gin.HandlerFunc{}
	router.Load(g,
		middlewares...,
	)
	g.Run(viper.GetString("port"))

	// Ping the server to make sure the router is working.
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Info("The router has been deployed successfully.")
	}()

	log.Infof("Start to listening the incoming requests on http address: %s", viper.GetString("port"))
	log.Info(http.ListenAndServe(viper.GetString("port"), g).Error())
}

// pingServer pings the http server to make sure the router is working.
func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get("http://127.0.0.1" + viper.GetString("port") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep for a second to continue the next ping.
		log.Info("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}

	return errors.New("Cannot connect to the router.")
}
