
package main

import (

  // basics.
  "flag"

  // local packages.
  opener "github.com/kraasch/urlopener/pkg/opener"
)

var (
  // url to open.
  url = ""
  // flags.
  suppress = false
)

func main() {

  // parse flags.
  flag.StringVar(&url, "url", "", "URL to open.")
  flag.Parse()

  // open url.
  if url != "" {
    opener.OpenUrl(url)
  }

} // fin.

