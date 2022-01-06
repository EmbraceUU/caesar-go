package elastic

import (
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
)

var (
	client *elastic.Client
)

func SetUp() {
	var err error
	client, err = elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false),
		elastic.SetBasicAuth("admin", "admin"),
		elastic.SetURL([]string{"127.0.0.1:9200"}...),
	)
	if err != nil {
		log.Fatal(fmt.Sprintf("init elastic failed, err is %s ", err.Error()))
		return
	}
}

func GetClient() *elastic.Client {
	return client
}
