package feeds

import (
	"encoding/xml"
	"testing"
	"time"
)

type test struct {
	feed    []byte
	id      string
	title   string
	link    string
	updated string
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
		id:      "urn:uuid:60a76c80-d399-11d9-b93C-0003939e0af6",
		title:   "Example Feed",
		link:    "http://example.org/",
		updated: "2003-12-13T18:30:02Z",
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

		if atom.ID != test.id {
			t.Error(
				"For", string(test.feed),
				"Expected", test.id,
				"Got", atom.ID,
			)
		}

		if atom.Title != test.title {
			t.Error(
				"For", string(test.feed),
				"Expected", test.title,
				"Got", atom.Title,
			)
		}

		if atom.Link.Href != test.link {
			t.Error(
				"For", string(test.feed),
				"Expected", test.link,
				"Got", atom.Link.Href,
			)
		}

		if atom.Updated.Format(time.RFC3339) != test.updated {
			t.Error(
				"For", string(test.feed),
				"Expected", test.updated,
				"Got", atom.Updated.Format(time.RFC3339),
			)
		}
	}
}

func TestAtomDate_UnmarshalXML(t *testing.T) {
	feed := []byte(`<?xml version="1.0" encoding="utf-8"?>
<feed xmlns="http://www.w3.org/2005/Atom">

  <title>Example Feed</title>
  <link href="http://example.org/"/>
  <updated>2003-12-13</updated>
  <author>
    <name>John Doe</name>
  </author>
  <id>urn:uuid:60a76c80-d399-11d9-b93C-0003939e0af6</id>
</feed>`)

	atom := new(AtomFeed)
	err := xml.Unmarshal(feed, atom)

	if err == nil {
		t.Error("For", string(feed),
			"Expected", "err",
			"Got", err)
	}
}
