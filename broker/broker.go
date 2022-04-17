package broker

import (
	"log"
	"os"
	"runtime"

	"github.com/fhmq/hmq/broker"
)

func Start() error {
	runtime.GOMAXPROCS(runtime.NumCPU())
	config, err := broker.ConfigureConfig(os.Args[1:])
	if err != nil {
		log.Fatal("configure broker config error: ", err)
	}

	b, err := broker.NewBroker(config)
	if err != nil {
		log.Fatal("New Broker error: ", err)
	}
	b.Start()

	return nil
}
