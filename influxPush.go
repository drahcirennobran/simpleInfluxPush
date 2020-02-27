package influxpush

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/influxdata/influxdb-client-go"
)

func Push(v uint) {
	influx, err := influxdb.New("http://192.168.0.67:8086", "") //, influxdb.WithHTTPClient(myHTTPClient))
	if err != nil {
		log.Fatal("influxdb.New")
		panic(err)
	}
	fmt.Printf("influx created")
	myMetrics := []influxdb.Metric{
		influxdb.NewRowMetric(
			map[string]interface{}{"conso": v},
			"system-metrics",
			map[string]string{},
			time.Date(2018, 3, 4, 5, 6, 7, 8, time.UTC)),
	}

	if _, err := influx.Write(context.Background(), "my-awesome-bucket", "my-very-awesome-org", myMetrics...); err != nil {
		log.Fatal(err)
	}
	influx.Close()
}
