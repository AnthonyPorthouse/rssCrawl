package feeds

import (
	"encoding/xml"
	"testing"
	"time"
)

type test struct {
	feed  []byte
	title string
	link  string
	updated time.Time
}

var tests = []test{
	{
		feed: []byte(`<?xml version="1.0" encoding="utf-8"?>
<feed xmlns="http://www.w3.org/2005/Atom">

  <title>Example Feed</title>
  <link href="http://example.org/"/>
  <updated>2003-12-13T18:30:02Z</updated>
  <author>
    <name>John Doe</name>
  </author>
  <id>urn:uuid:60a76c80-d399-11d9-b93C-0003939e0af6</id>

  <entry>
    <title>Atom-Powered Robots Run Amok</title>
    <link href="http://example.org/2003/12/13/atom03"/>
    <id>urn:uuid:1225c695-cfb8-4ebb-aaaa-80da344efa6a</id>
    <updated>2003-12-13T18:30:02Z</updated>
    <summary>Some text.</summary>
  </entry>

</feed>`),
		title: "Example Feed",
		link:  "http://example.org/",
	},
}

func TestAtomFeed(t *testing.T) {
	for _, test := range tests {
		atom := new(AtomFeed)
		err := xml.Unmarshal(test.feed, atom)
		if err != nil {
			t.Error(
				"For", string(test.feed),
				"Expected", nil,
				"Got", err,
			)
		}
	}
}

func TestTitle(t *testing.T) {
	for _, test := range tests {
		atom := new(AtomFeed)
		xml.Unmarshal(test.feed, atom)
		if atom.Title != test.title {
			t.Error(
				"For", string(test.feed),
				"Expected", test.title,
				"Got", atom.Title,
			)
		}
	}
}

func TestLink(t *testing.T) {
	for _, test := range tests {
		atom := new(AtomFeed)
		xml.Unmarshal(test.feed, atom)
		if atom.Link.Href != test.link {
			t.Error(
				"For", string(test.feed),
				"Expected", test.link,
				"Got", atom.Link.Href,
			)
		}
	}
}
