package application

import (
	"context"
	"fmt"
	"github.com/jaimera/poc-services/broker"
	"github.com/jaimera/poc-services/domain/contract"
	"github.com/jaimera/poc-services/repository"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type App struct {
	Logger      *logrus.Entry
	DataManager contract.DataManager
	HTTPClient  *http.Client
	services    *AppService
	KafkaServer contract.PortBroker // broker.KafkaConnection
	HttpPort    string
}

// BuildApp initilialize everything the application needs: server, db, logger...
func BuildApp() App {

	logger := logrus.WithFields(logrus.Fields{
		"AppName": "pocservices",
		"Version": "1.0",
	})

	logger.Infof("Initializing APP!")

	httpClient := new(http.Client)

	// init DB, database params should be stored in a safe place
	db, err := repository.Connect(
		"root",
		"root",
		"db_poc",
		"db",
		3306,
	)
	endAsErr(err, "Could not connect to database.")
	atInterruption(func() { db.Close() })

	app := App{
		Logger:      logger,
		DataManager: db,
		HTTPClient:  httpClient,
		HttpPort:    ":8080",
	}

	// init services
	app.services = NewAppService(&app)

	// init kafka configuration
	app.KafkaServer = broker.NewKafkaConnection(
		logger,
		"kafka:29092",
		"poc-services",
		0,
		"tcp",
	)

	go app.KafkaServer.Consume(context.Background(), app.GetServices().Port())

	return app
}

func (app App) GetServices() *AppService {
	return app.services
}

func atInterruption(fn func()) {
	go func() {
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, os.Interrupt)
		<-sc

		fn()
		os.Exit(0)
	}()
}

func endAsErr(err error, message string) {
	if err != nil {
		fmt.Println(message, err)
		time.Sleep(time.Millisecond * 50)
		os.Exit(1)
	}
}
