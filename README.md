# rfcsearch

[RFC](https://www.ietf.org/rfc.html) search results in JSON format.

This is an API server that scraps data from RFC Editor's
[search](https://www.rfc-editor.org/search/rfc_search.php) page and
returns the results in JSON format. It takes advantage of
[PuerkitoBio's](https://github.com/PuerkitoBio) excellent
[goquery](https://github.com/PuerkitoBio/goquery) package to parse the
HTML results page and gather data for each RFC.


```
$curl -s "https://rfcsearch-gorun.rhcloud.com/?q=coffee" | json_pp
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
