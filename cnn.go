//Package goCNNTop10 implements a basic client for the cnn API news
package goCNNTop10

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

//Item describes a Reddit item
type Item struct {
	Title       string
	URL         string
	Description string
}

type response struct {
	Articles []Item
}

//Get fetches the most recent Items posted to the specified subreddit
func Get() ([]Item, error) {
	url := "https://newsapi.org/v1/articles?source=cnn&sortBy=top&apiKey=bb652081e166447bb4857b339291e276"

	//url := fmt.Sprintf("http://reddit.com/r/%s.json", reddit)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	r := new(response)
	err = json.NewDecoder(resp.Body).Decode(r)
	if err != nil {
		return nil, err
	}
	items := make([]Item, len(r.Articles))
	for i, child := range r.Articles {
		items[i] = child
	}
	return items, nil
}

func (i Item) String() string {
	return fmt.Sprintf("%s\n%s\n%s\n", i.Title, i.Description, i.URL)
}
