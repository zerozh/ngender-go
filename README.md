# NGender Golang Version

A Golang version of [https://github.com/observerss/ngender](https://github.com/observerss/ngender)

## Usage

`go get -u github.com/zerozh/ngender-go`

```go
package main

import (
	"fmt"
	
	"github.com/zerozh/ngender-go"
)

func main(){
	gender, probability := ngender.Guess("赵本山")
	fmt.Println(gender, probability) // "male", 0.9836229687547046
	
	gender, probability = ngender.Guess("宋丹丹")
	fmt.Println(gender, probability) // "female", 0.9759486128949907
	
	gender, probability = ngender.Guess("James")
	fmt.Println(gender, probability) // "unknown", 0
	
	gender, probability = ngender.GuessGivenName("俊杰")
	fmt.Println(gender, probability) // "male", 0.922575579531011
    	
	gender, probability = ngender.GuessGivenName("晓娟")
	fmt.Println(gender, probability) // "female", 0.9721899067171618
}
```

## Performance

```
BenchmarkGuess-6   	 5000000	       265 ns/op
BenchmarkGuessGivenName-6   	 5000000	       267 ns/op
```
