package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

var client *http.Client
var ALPHA_VANTAGE_ENDPOINT = os.Getenv("ALPHA_VANTAGE_ENDPOINT")
var ALPHA_VANTAGE_API_KEY = os.Getenv("ALPHA_VANTAGE_API_KEY")

func init() {
	Logger()
	logger.Info().Msg("Crawling Orchestrator Starting ...")

	// Initiate MongoDB client
	_, err := MongoClient()
	if err != nil {
		logger.Info().Msg(err.Error())
	}

	// Initiate HTTP client
	client = &http.Client{}
}

func main() {

	// Fetch Stocks from database
	collection := mc.Db.Collection("companies")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		logger.Error().Msg(err.Error())
	}
	defer cursor.Close(ctx)

	var stocks []Stock
	if err = cursor.All(ctx, &stocks); err != nil {
		logger.Error().Msg(err.Error())
	}

	// For each symbol call the News & Sentiment API to get the latest news
	for _, stock := range stocks {
		logger.Info().Msg("Fetching news for: " + stock.Symbol + "....")
		feed, err := fetchLatestNews(stock.Symbol)
		if err != nil {
			logger.Error().Msg(err.Error())
		}

		// Create a slice of articles as should be represented in Mongo
		var articles []interface{}
		for _, news := range feed.Feed {
			articles = append(articles, Article{
				Title:                 news.Title,
				URL:                   news.URL,
				Source:                news.Source,
				TimePublished:         news.TimePublished,
				Authors:               news.Authors,
				Summary:               news.Summary,
				OverallSentimentScore: news.OverallSentimentScore,
			})
		}

		// Save articles for future reference
		collection := mc.Db.Collection("news")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		res, err := collection.InsertMany(ctx, articles)
		if err != nil {
			logger.Error().Caller().Msg(err.Error())
		} else {
			insertedCount := len(res.InsertedIDs)
			fmt.Println("Number of articles inserted:", insertedCount)
		}

	}

}

func fetchLatestNews(symbol string) (*Feed, error) {
	// Retrieve all articles URLs
	req, err := http.NewRequest("GET", ALPHA_VANTAGE_ENDPOINT, nil)
	if err != nil {
		return nil, err
	}

	// Add Request Header
	req.Header.Add("Accept", `application/json`)

	// Add Request Parameters
	q := req.URL.Query()
	q.Add("function", "NEWS_SENTIMENT")
	q.Add("tickers", symbol)
	q.Add("time_from", time.Now().AddDate(0, 0, -1).Format("20060102T1504"))
	q.Add("sort", "LATEST")
	q.Add("limit", "1000")
	q.Add("apikey", ALPHA_VANTAGE_API_KEY)
	req.URL.RawQuery = q.Encode()

	logger.Debug().Msg("Executing Request: " + req.URL.RawQuery)

	// Execute HTTP Request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	feed := Feed{}
	if err := json.Unmarshal(body, &feed); err != nil {
		return nil, err
	}

	return &feed, nil
}
