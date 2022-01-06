package elastic

import (
	"context"
	"encoding/json"
	elastic2 "github.com/olivere/elastic/v7"
)

type Item struct {
}

func QueryExample() {
	boolQuery := elastic2.NewBoolQuery()
	boolQuery.Must(elastic2.NewTermQuery("id", "1"))
	searchResult, err := GetClient().
		Search("test-*").
		From(0).
		Size(10).
		Query(boolQuery).
		TrackTotalHits(true).
		Sort("create_time", true).
		Pretty(true).
		Do(context.Background())

	if err != nil {
		return
	}

	if searchResult.TotalHits() > 0 {
		for _, item := range searchResult.Hits.Hits {
			var iv Item
			_ = json.Unmarshal(item.Source, &iv)
		}
	}
}
