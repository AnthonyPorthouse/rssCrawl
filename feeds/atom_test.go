package feeds

import (
	"encoding/xml"
	"testing"
	"time"
)

type test struct {
	feed []byte

	// Required
	id      string
	title   string
	updated string

	// Recommended
	authors int
	link    string

	// Optional
	categories   int
	contributors int
	generator    string
	icon         string
	logo         string
	rights       string
	subtitle     string

	entries int
}

var tests = []test{
	{
		feed: []byte(`<?xml version="1.0" encoding="utf-8"?>
<feed xmlns="http://www.w3.org/2005/Atom">

  <id>urn:uuid:60a76c80-d399-11d9-b93C-0003939e0af6</id>
  <title>Example Feed</title>
  <link href="http://example.org/"/>
  <updated>2003-12-13T18:30:02Z</updated>

  <author>
  </author>

  <entry>
  </entry>

</feed>`),
		id:      "urn:uuid:60a76c80-d399-11d9-b93C-0003939e0af6",
		title:   "Example Feed",
		link:    "http://example.org/",
		updated: "2003-12-13T18:30:02Z",
		authors: 1,
		entries: 1,
	},
}

func TestAtomFeed(t *testing.T) {
	for _, test := range tests {
		atom := new(AtomFeed)
		err := xml.Unmarshal(test.feed, atom)

		// Check unmarshalling issues
		if err != nil {
			t.Error(
				"For", string(test.feed),
				"Expected", nil,
				"Got", err,
			)
		}

		// Check Required Fields
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
		if atom.Updated.Format(time.RFC3339) != test.updated {
			t.Error(
				"For", string(test.feed),
				"Expected", test.updated,
				"Got", atom.Updated.Format(time.RFC3339),
			)
		}

		// Check Recommended Fields
		if len(atom.Authors) != test.authors {
			t.Error(
				"For", string(test.feed),
				"Expected", test.authors,
				"Got", len(atom.Authors),
			)
		}
		if atom.Link.Href != test.link {
			t.Error(
				"For", string(test.feed),
				"Expected", test.link,
				"Got", atom.Link.Href,
			)
		}

		// Check Optional Fields
		if len(atom.Categories) != test.categories {
			t.Error(
				"For", string(test.feed),
				"Expected", test.categories,
				"Got", len(atom.Categories),
			)
		}
		if len(atom.Contributors) != test.contributors {
			t.Error(
				"For", string(test.feed),
				"Expected", test.contributors,
				"Got", len(atom.Contributors),
			)
		}
		if atom.Generator != test.generator {
			t.Error(
				"For", string(test.feed),
				"Expected", test.generator,
				"Got", atom.Generator,
			)
		}
		if atom.Icon != test.icon {
			t.Error(
				"For", string(test.feed),
				"Expected", test.icon,
				"Got", atom.Icon,
			)
		}
		if atom.Logo != test.logo {
			t.Error(
				"For", string(test.feed),
				"Expected", test.logo,
				"Got", atom.Logo,
			)
		}
		if atom.Rights != test.rights {
			t.Error(
				"For", string(test.feed),
				"Expected", test.rights,
				"Got", atom.Rights,
			)
		}
		if atom.Subtitle != test.subtitle {
			t.Error(
				"For", string(test.feed),
				"Expected", test.subtitle,
				"Got", atom.Subtitle,
			)
		}

		if len(atom.Entries) != test.entries {
			t.Error(
				"For", string(test.feed),
				"Expected", test.entries,
				"Got", len(atom.Entries),
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
