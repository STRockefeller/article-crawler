# article-crawler

## Description

This package is a simple web crawler written in Go, that extracts text content from a given URL by recursively traversing the HTML document tree and selecting certain HTML tags. The tags selected for extraction include `p`, `h1`, `h2`, `h3`, `h4`, `h5`, `h6`, `ul`, `ol`, `pre`, and `blockquote`.

## Installation

To use this package, you will need to have Go installed on your system. Once you have Go installed, you can add the package to your project using the following command:

```go get github.com/STRockefeller/article-crawler```

## Usage

To use the crawler, simply call the `Crawl` function with the URL you want to crawl as its argument. The function will return a string containing the extracted text content.

```go
package main

import (
	"fmt"
	"github.com/STRockefeller/article-crawler"
)

func main() {
	url := "https://example.com"
	text := crawler.Crawl(url)
	fmt.Println(text)
}
```

## Contributing

Contributions to this package are welcome. If you find a bug or have a feature request, please open an issue or submit a pull request.

## License

This package is licensed under the MIT license. See the LICENSE file for more information.
