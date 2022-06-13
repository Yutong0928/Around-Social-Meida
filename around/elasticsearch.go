package main

import (
    "context"

    "github.com/olivere/elastic/v7"
)

const (
        ES_URL = "http://34.125.129.194:9200"
)

func readFromES(query elastic.Query, index string) (*elastic.SearchResult, error) {
    client, err := elastic.NewClient(
        elastic.SetURL(ES_URL),
        elastic.SetBasicAuth("xiaoyuqi", "xiaoyuqi123"))
    if err != nil {
        return nil, err
    }

    searchResult, err := client.Search().
        Index(index).
        Query(query).
        Pretty(true).
        Do(context.Background())
    if err != nil {
        return nil, err
    }

    return searchResult, nil
}

func saveToES(i interface{}, index string, id string) error{
    client, err := elastic.NewClient(
        elastic.SetURL(ES_URL),
        elastic.SetBasicAuth("xiaoyuqi", "xiaoyuqi123"))
    if err != nil {
        return err
    }

    _, err = client.Index().
        Index(index).
        Id(id).
        BodyJson(i).
        Do(context.Background())
    return err
}

func deleteFromES(query elastic.Query, index string) error {
    client, err := elastic.NewClient(
        elastic.SetURL(ES_URL),
        elastic.SetBasicAuth("xiaoyuqi", "xiaoyuqi123"))
    if err != nil {
        return err
    }

    _, err = client.DeleteByQuery().
        Index(index).
        Query(query).
        Pretty(true).
        Do(context.Background())

    return err
}