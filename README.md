# rfcsearch

[RFC](https://www.ietf.org/rfc.html) search results in JSON format.

This is an API server written in Go that passes search parameters to
RFC Editor's
[search](https://www.rfc-editor.org/search/rfc_search.php) page and
returns the results in JSON format. It takes advantage of
[PuerkitoBio's](https://github.com/PuerkitoBio) excellent
[goquery](https://github.com/PuerkitoBio/goquery) package to parse the
HTML results page and gather data for each RFC.

A live version is currently running on
[Heroku](https://www.heroku.com) at
https://rfcsearch.herokuapp.com/

Queries are made on either keywords or RFC numbers. A query containing
only numbers is assumed to be the RFC itself and returns only that
RFC. If any characters appear in the query then keywords and titles
will be searched.

### Title / Keyword search

```
$curl -s "https://rfcsearch.herokuapp.com/?q=coffee" | json_pp
[
   {
      "title" : "The Hyper Text Coffee Pot Control Protocol for Tea Efflux Appliances (HTCPCP-TEA)",
      "number" : "7168",
      "link" : "http://tools.ietf.org/html/rfc7168",
      "authors" : "I. Nazar",
      "date" : "1 April 2014",
      "status" : "Informational",
      "moreinfo" : "Updates RFC 2324"
   },

...

```

### RFC search

```
curl -s "https://rfcsearch.herokuapp.com/?q=4492" | json_pp
[
   {
      "number" : "4492",
      "status" : "Informational",
      "date" : "May 2006",
      "authors" : "S. Blake-Wilson, N. Bolyard, V. Gupta, C. Hawk, B. Moeller",
      "title" : "Elliptic Curve Cryptography (ECC) Cipher Suites for Transport Layer Security (TLS)",
      "link" : "http://tools.ietf.org/html/rfc4492",
      "moreinfo" : "Updated by RFC 5246, RFC 7027, Errata"
   }
]

```
