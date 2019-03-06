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
	// 加载词频文件，格式为 "字,男名使用次数,女名使用次数"，每行一条，如：
	// 伟,378757,51232
	// 强,229396,6049
	// 芳,30090,165416
	// 艳,27301,172376
	// 玉,126830,147213
	// 云,109814,100157
	// 关于默认词频文件的更多说明详见 https://github.com/observerss/ngender
	ngender.LoadDataFromFile("/path/to/charfreq.csv")
    
	// Guess() 方法自动去除姓名中姓字（一般为第一个字，也支持十余个常见复姓如欧阳、诸葛、司马等）后计算
	gender, probability := ngender.Guess("赵本山")
	fmt.Println(gender, probability) // "male", 0.9836229687547046
	
	gender, probability = ngender.Guess("宋丹丹")
	fmt.Println(gender, probability) // "female", 0.9759486128949907
	
	gender, probability = ngender.Guess("James")
	fmt.Println(gender, probability) // "unknown", 0
	
	// GuessGivenName() 直接使用名字计算
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
