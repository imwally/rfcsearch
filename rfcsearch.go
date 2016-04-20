package main

import (
	"encoding/json"
	"html"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const searchURL = "https://www.rfc-editor.org/search/rfc_search_detail.php?page=All"

type RFC struct {
	Number   string
	Title    string
	Authors  string
	Date     string
	MoreInfo string
	Status   string
}

func buildQueryURL(query string) (*url.URL, error) {
	u, err := url.Parse(searchURL)
	if err != nil {
		return nil, err
	}

	// Assume a query containing only numbers refers to the RFC
	// itself, otherwise search for titles or keywords.
	q := u.Query()
	_, err = strconv.Atoi(query)
	if err == nil {
		q.Set("rfc", query)
	} else {
		q.Set("title", query)
	}

	u.RawQuery = q.Encode()

	return u, nil
}

func parseDocument(doc *goquery.Document) []RFC {
	col := 0
	var results []RFC
	var rfc RFC

	// Search results are contained in a table with the class
	// "gridtable". There are 7 columns or td's in each row that
	// correspond to a data field of an RFC. The 'col' variable
	// keeps track of each row by resetting to 0 when it grows
	// larger than 6.
	doc.Find(".gridtable tr td").Each(func(i int, s *goquery.Selection) {
		switch col {
		case 0:
			rfc.Number = strings.TrimSpace(strings.TrimLeft(s.Text(), "RFC"))
		case 2:
			rfc.Title = strings.TrimSpace(s.Text())
		case 3:
			rfc.Authors = strings.TrimSpace(s.Text())
		case 4:
			rfc.Date = strings.TrimSpace(s.Text())
		case 5:
			rfc.MoreInfo = strings.TrimSpace(s.Text())
		case 6:
			rfc.Status = strings.TrimSpace(s.Text())
		}
		col++
		if col > 6 {
			col = 0
			results = append(results, rfc)
		}
	})

	return results
}

func getDocument(search string) (*goquery.Document, error) {
	queryURL, err := buildQueryURL(search)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocument(queryURL.String())
	if err != nil {
		return nil, err
	}

	return doc, nil
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	q, err := url.ParseQuery(html.UnescapeString(r.URL.RawQuery))
	if err != nil {
		log.Println(err)
	}

	search := q.Get("q")
	doc, err := getDocument(search)
	if err != nil {
		log.Println(err)
	}

	results := parseDocument(doc)
	if err := json.NewEncoder(w).Encode(results); err != nil {
		log.Println(err)
	}
}

func main() {
	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
