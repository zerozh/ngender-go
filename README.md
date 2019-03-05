# NGender Golang Version

A Golang version of [https://github.com/observerss/ngender](https://github.com/observerss/ngender)

## Usage

```go
package main

import (
	"fmt"
	
	"ngender-go"
)

func main(){
	gender, probability := ngender.Guess("赵本山")
	fmt.Println(gender, probability) // "male", 0.983622
	
	gender, probability = ngender.Guess("宋丹丹")
	fmt.Println(gender, probability) // "female", 0.975948
	
	gender, probability = ngender.Guess("James")
	fmt.Println(gender, probability) // "unknown", 0
}
```

## Performance

```
BenchmarkGuess-6   	 5000000	       265 ns/op
```
