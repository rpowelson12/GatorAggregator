package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
)

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	rssFeed := &RSSFeed{}
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return rssFeed, fmt.Errorf("error getting feed from URL: %v", err)
	}

	req.Header.Set("User-Agent", "gator")
	client := new(http.Client)

	res, err := client.Do(req)
	if err != nil {
		return rssFeed, fmt.Errorf("error requesting info: %v", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return rssFeed, fmt.Errorf("error reading response: %v", err)
	}

	err = xml.Unmarshal(body, rssFeed)
	if err != nil {
		return rssFeed, fmt.Errorf("error parsing xml: %v", err)
	}

	rssFeed.Channel.Title = html.UnescapeString(rssFeed.Channel.Title)
	rssFeed.Channel.Description = html.UnescapeString(rssFeed.Channel.Description)

	for i := range rssFeed.Channel.Item {
		rssFeed.Channel.Item[i].Title = html.UnescapeString(rssFeed.Channel.Item[i].Title)
		rssFeed.Channel.Item[i].Description = html.UnescapeString(rssFeed.Channel.Item[i].Description)
	}

	return rssFeed, nil
}
