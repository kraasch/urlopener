
# urlopener

## usage

In code as Go package:

```go
import (
  opener "github.com/kraasch/urlopener/pkg/opener"
)

func main() {
  url := "https://amazon.de/"
  opener.OpenUrl(url)
}
```

On the command-line:

```bash
urlopener -url 'https://amazon.de/'
```

## misc

Based on:

 - https://stackoverflow.com/questions/39320371/how-start-web-server-to-open-page-in-browser-in-golang
 - https://gist.github.com/sevkin/9798d67b2cb9d07cb05f89f14ba682f8

