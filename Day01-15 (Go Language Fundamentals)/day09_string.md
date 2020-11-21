# One, string (string)
> @author：Han Ru
>
> Copyright: Beijing Qianfeng Internet Technology Co., Ltd.
## 1.1 What is string
A string in Go is a slice of bytes. You can create a string by encapsulating its content in "". Strings in Go are Unicode compatible and UTF-8 encoded.
Sample code:
```go
package main
import (
"fmt"
)
func main() {
name := "Hello World"
fmt.Println(name)
}
```
## 1.2 Use of string
### 1.2.1 Access a single byte in a string
```go
package main
import (
"fmt"
)
func main() {
name := "Hello World"
for i:= 0; i <len(s); i++ {
fmt.Printf("%d ", s[i])
}
fmt.Printf("\n")
for i:= 0; i <len(s); i++ {
fmt.Printf("%c ",s[i])
}
}
```
operation result:
72 101 108 108 111 32 87 111 114 108 100
H e l l o W o r l d
## 1.3 strings package
There are many functions for manipulating strings when accessing the strings package.
## 1.4 strconv package
Visit the strconv package, you can convert between string and other numeric types.
Qianfeng Go language learning group: 784190273
Author B station:
HTTPS://space.proportion.com/353694001
Corresponding video address:
Https://oooooo.proportion.com/video/av56018934
Https://oooooo.proportion.com/video/av47467197
Source code:
HTTPS://GitHub.com/Ruby with 1314/go_foundation
