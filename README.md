# frequency-analysis


Produces frequency analysis and returns the 10 most common words in text.


## Installation

Run the following command from you terminal:


 ```bash
 go get github.com/koind/frequency-analysis
 ```

## Usage

Package usage example.

```go
package main

import (
	"github.com/koind/frequency-analysis"
)

func main() {
	text := "lorem and lorem and lorem lorem lorem lorem lorem lorem lorem lorem lorem lorem"
	
	result := frequency_analysis.Analyze(text)
	println(result) // map["lorem"]12
}
```

## Available Methods

The following methods are available:

##### koind/frequency-analysis

```go
Analyze(text string) map[string]int
```

## Tests

Run the following command from you terminal:

```
go test -v .
```