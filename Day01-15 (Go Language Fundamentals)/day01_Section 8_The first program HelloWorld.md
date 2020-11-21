# The first program: HelloWorld
> @author：Han Ru
>
> Copyright: Beijing Qianfeng Internet Technology Co., Ltd.
## One, go project engineering structure
After configuring the working directory, you can code and develop. Before that, let's look at the general project structure of go. The structure here is mainly the directory structure where source code and resource files are stored.
### 1.1 gopath directory
The gopath directory is the directory where we store the source code we write. There are usually 3 subdirectories under this directory: src, bin, pkg.
> src ---- Each subdirectory in it is a package. Inside the package is the Go source code file
>
> pkg ---- generated after compilation, the object file of the package
>
> bin ---- The generated executable file.
### 1.2 Writing the first program
The study of every programming language starts with a "Hello, World." program. This example first appeared in the C language bible "The C Programming Language" published in 1978. There is also a very good story about "Hello, World." that is that all programmers expect computers to have real intelligence one day, and then say "from the heart" to the people who created them, Hello, World .
1. In the HOME/go directory, (in the GOPATH directory), create a directory called src, then create a folder called hello in this directory, create a file called helloworld.go in this directory, and double-click to open , Enter the following:
```go
package main
import "fmt"
func main() {
fmt.Println("Hello, World!")
}
```
2. Execute the go program
There are several ways to execute go programs
Method 1: Use the go run command
​ step1: Open the terminal:
​ Use the shortcut key win+R under window, enter cmd to open the command line prompt
​ You can use shortcut keys under linux: ctrl+alt+T
​ Command+space under mac, enter terminal
​ step2: Enter the directory where helloworld.go is located
​ step3: Enter the go run helloworld.go command and observe the results.
Method 2: Use the go build command
​ step1: Open the terminal: in any file path, run:
​ go install hello
​ You can also enter the path of the project (application package), and then run:
​ go install
Note that when compiling and generating go programs, go will actually go to two places to find packages:
Under the src folder under GOROOT, and under the src folder under GOPATH.
In the program package, automatically find the main function of the main package as the program entry, and then compile it.
​ step2: Run go program
​ Under /home/go/bin/ (if there is no bin directory before, it will be created automatically), you will find a hello executable file, run it with the following command:
​ ./hello
​
### 1.3 Explanation of the first program
#### 3.2.1 package
-The files under the same package belong to the same project file and can be used directly without the `import` package
-The package name of all files under the same package are the same
-The name of the file `package` under the same package is recommended to be the name of the directory, but it may not be
#### 3.2.2 import
import "fmt" tells the Go compiler that this program needs to use the functions of the fmt package, which implements the formatting IO (input/output) function
It can be a relative path or an absolute path. It is recommended to use an absolute path (starting from the project root directory)
1. Point operation
We sometimes see the following way to import packages
```go
import(
. "fmt"
)
```
The meaning of this point operation is that after the package is imported, when you call the function of this package, you can omit the prefixed package name, that is, you adjust
The used `fmt.Println("hello world")` can be omitted and written as `Println("hello world")`
2. Alias ​​operations
Alias ​​operation, as the name implies, we can name the package another name that we can easily remember
```go
import(
f "fmt"
)
```
For alias operation, the prefix becomes our prefix when calling the package function, namely `f.Println("hello world")`
3. _Operation
This operation is often an operator that makes many people puzzled, please see the import below
```go
import (
"database/sql"
_ "GitHub.com/ZU Express/No MySQL/God RV"
)
```
_The operation is actually to introduce the package, instead of directly using the functions in the package, but to call the init function in the package
#### 3.3.3 main
main() is the entry point of the program.
### 1.4 Package description
We know that the source code is stored in the src directory of GOPATH, so how to distinguish between multiple projects? The answer is to use packages to organize our project directory structure. Anyone who has ever developed Java knows that if you use packages to organize your code, there will be no duplication of packages starting with the website domain name. For example, if Qianfeng’s website is `http://www.mobiletrain.org`, we can use `mobiletrain. Create a folder with the name of org`, and put my own Go projects in this folder, so that it will not conflict with other people's projects, and the package name is also unique.
If you have your own domain name, you can also use your own domain name. If you don't have a personal domain name, the current popular practice is to use your personal github name, because everyone's is unique, so there will be no duplication.
![package1](http://7xtcwd.com1.z0.glb.clouddn.com/package1.png)
As above, the src directory is followed by folders named by domain names. Take the github.com folder as an example again. It is a folder named github user to store the go source code written by this github user.
Qianfeng Go language learning group: 784190273
Author B station:
HTTPS://space.proportion.com/353694001
Corresponding video address:
Https://oooooo.proportion.com/video/av56018934
Https://oooooo.proportion.com/video/av47467197
Source code:
HTTPS://GitHub.com/Ruby with 1314/go_foundation
