package feeds

import (
	"encoding/xml"
	"errors"
	"github.com/AnthonyPorthouse/rssCrawl/http"
	"io/ioutil"
)

func GetFeed(URL string) (*Feed, error) {
	feed, err := http.GetFeed(URL)
	if err != nil {
		return nil, err
	}
	defer feed.Body.Close()
	body, _ := ioutil.ReadAll(feed.Body)

	rss := new(RSSFeed)
	atom := new(AtomFeed)

	switch {
	case xml.Unmarshal(body, rss) == nil && rss.Version != "":
		feed, err := transformRSS(rss)
		if err != nil {
			return nil, err
		}

		return feed, nil
	case xml.Unmarshal(body, atom) == nil && atom.ID != "":
		feed, err := transformAtom(atom)
		if err != nil {
			return nil, err
		}

		return feed, nil

	default:
		return nil, errors.New("unable to unmarshal feed")
	}
}

func transformRSS(rss *RSSFeed) (*Feed, error) {
	feed := Feed{
		Title: rss.Channel.Title,
		URL:   rss.Channel.Link,
		Items: []FeedItem{},
	}

	for _, item := range rss.Channel.Items {
		fi := FeedItem{
			ID:      item.GUID,
			URL:     item.Link,
			Content: item.Description,
		}

		feed.Items = append(feed.Items, fi)
	}

	return &feed, nil
}

func transformAtom(atom *AtomFeed) (*Feed, error) {
	feed := Feed{
		Title: atom.Title,
		URL:   atom.Link.Href,
		Items: []FeedItem{},
	}

	for _, item := range atom.Entries {
		fi := FeedItem{
			ID:      item.ID,
			URL:     item.Link,
			Content: item.Content,
		}

		feed.Items = append(feed.Items, fi)
	}

	return &feed, nil
}

type Feed struct {
	Title string
	URL   string
	Items []FeedItem
}

type FeedItem struct {
	ID      string
	URL     string
	Content string
}
