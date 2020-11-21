# Use of packages in Go
> @author：Han Ru
>
> Copyright: Beijing Qianfeng Internet Technology Co., Ltd.
Go language uses package as a syntax element to organize source code. All syntax visibility is defined at the package level. Compared with Java, python and other languages, this is not an innovation, but compared with C traditional include , It appears to be much more advanced.
```
myblog
├── conf
│ └── app.conf
├── controllers
│ ├── aboutme_controller.go
│ ├── add_article_controller.go
│ ├── album_controller.go
│ ├── base_controller.go
│ ├── default.go
│ ├── delete_article_controller.go
│ ├── exit_controller.go
│ ├── home_controller.go
│ ├── login_controller.go
│ ├── register_controller.go
│ ├── show_article_controller.go
│ ├── tags_controller.go
│ ├── update_article_controller.go
│ └── upload_controller.go
├── main.go
├── models
│ ├── album_model.go
│ ├── article_model.go
│ ├── home_model.go
│ ├── tags_model.go
│ └── user_model.go
├── myblogweb
├── routers
│ └── router.go
├── static
│ ├── css
│ │ ├── blogsheet.css
│ │ └── lib
│ │ ├── highlight.css
│ │ └── login.css
│ ├── img
│ ├── js
│ │ ├── blog.js
│ │ ├── lib
│ │ │ │ ├── jquery-3.3.1.min.js
│ │ │ └── jquery.url.js
│ │ └── reload.min.js
│ └── upload
│ └── img
│ └── 2018
│ └── 12
│ └── 11
│ ├── 1544511378-bee2.png
├── tests
│ └── default_test.go
├── utils
│ ├── myUtils.go
│ └── mysqlUtils.go
└── views
├── aboultme.html
├── album.html
├── block
│ ├── home_block.html
│ └── nav.html
├── home.html
├── index.tpl
├── login.html
├── register.html
├── show_article.html
├── tags.html
└── write_article.html
```
The source code reuse of the Go language is built on the basis of packages. The package is completed through package, import, GOPATH operations.
## 1. Main package
The entry of the Go language, the package where the main() function is located is called main. If the main package wants to reference other codes, it needs to be imported!
## 2, package
The src directory organizes and saves Go source files in the form of code packages. Each code package has a one-to-one correspondence with the folder under the src directory. Each subdirectory is a code package.
> Code package package name and file directory name are not required to be consistent. For example, the file directory is called hello, but the package name of the code package can be declared as "main", but the package declared on the first line of the source file in the same directory must be the same!
Add a package definition to the first line of all .go files in the same directory to mark the package to which the file belongs, and demonstrate the syntax:
```go
package package name
```
The package needs to meet:
-Files of the same level in a directory belong to one package. In other words, the package name of all files under the same package are the same.
-The name of the file `package` under the same package is recommended to be the name of the directory, but it may not be. In other words, the package name can be different from the directory name.
-The package named main is the entry package of the application, and other packages cannot be used.
> The files under the same package belong to the same project file and can be used directly without the `import` package
Packages can be defined nested, corresponding to nested directories, but the package name should be consistent with the directory where it is located, for example:
```go
// File: qf/ruby/tool.go
package ruby
// Functions that can be exported
func FuncPublic() {
}
// Functions that cannot be exported
func funcPrivate() {
}
```
In the package, the first letter of the identifier is capitalized to determine whether it can be exported. Only capitals of the first letter can be exported and regarded as public resources.
## 3. import
To reference other packages, you can use the import keyword, which can be imported individually or in batches. Syntax demonstration:
A: Usually imported
```go
// single import
import "package"
// Batch Import
import (
"package1"
"package2"
)
```
B: Point operation
We sometimes see the following way to import packages
```go
import(
. "fmt"
)
```
The meaning of this point operation is that after the package is imported, when you call the function of this package, you can omit the prefixed package name, that is, you adjust
The used `fmt.Println("hello world")` can be omitted and written as `Println("hello world")`
C: Get an alias
The alias operation is just what the name implies, we can name the package another name that we can easily remember. When importing, you can define aliases for the package, and the syntax is demonstrated:
```go
import (
p1 "package1"
p2 "package2"
)
// When in use: alias operation, the prefix becomes our prefix when calling the package function
p1.Method()
```
D: _Operation
If you only need to perform initialization operations when importing the package, you do not need to use other functions, constants and other resources in the package. Then you can import the package anonymously.
This operation is often an operator that is confusing for many people. Please see the following import:
```go
import (
"database/sql"
_ "GitHub.com/ZU Express/No MySQL/God RV"
)
```
_The operation is actually to import the package, instead of directly using the functions in the package, but to call the init function in the package. In other words, using the underscore as the alias of the package will only execute init().
> The path name of the imported package can be a relative path or an absolute path. It is recommended to use an absolute path (starting from the project root directory).
## 4. GOPATH environment variable
When importing, it will retrieve src/package from the GO installation directory (that is, the directory set by the GOROOT environment variable) and the directory set by the GOPATH environment variable to import the package. If it does not exist, the import fails.
GOROOT is the location of the GO built-in package.
GOPATH is the location of the package defined by ourselves.
Usually when we are developing a Go project, debugging or compiling, we need to set GOPATH to point to our project directory, and the packages in the src directory in the directory can be imported.
## 5, init() package initialization
Below we introduce these two functions in detail: init() and main() are reserved functions in the go language. We can define the init() function in the source code. This function will be executed when the package is imported. For example, if the package is imported in main and there is init() in the package, then the code in init() will be executed before the main() function is executed to initialize the package. Specific information. E.g:
Package source code:
```go
src/userPackage/tool.go
package userPackage
import "fmt"
func init() {
fmt.Println("tool init")
}
```
Main function source code:
```go
src/main.go
package main
import (
"usePackage"
)
func main() {
fmt.Println("main run")
// use userPackage
usePackage.SomeFunc()
}
```
When executed, it will output "tool init" first, and then "main run".
Let's introduce the two functions of init() and main() in detail below. The differences in the go language are as follows:
Same point:
The two functions cannot have any parameters and return values ​​when they are defined.
This function can only be automatically called by the go program and cannot be referenced.
difference:
init can be applied to any package, and multiple definitions can be repeated.
The main function can only be used in the main package, and only one can be defined.
The order of execution of the two functions:
The go file in the main package will always be executed by default.
The order of init() calls to the same go file is from top to bottom.
For different files in the same package, sort the file names from small to large by character strings, and then call the init() function in each file in sequence.
For different packages, if they do not depend on each other, call the init() function in the package in the order of import in the main package.
If the package has dependencies, the calling order is the last dependent and the first to be initialized, for example: import order main -> A -> B -> C, then the initialization order is C -> B -> A -> main, one execution corresponds to The init method. The main package is always initialized by the last one, because it always depends on other packages
![20170831112523944](img/20170831112523944.png)
Picture quoted from the Internet
Avoid cyclic import, for example: A -> B -> C -> A.
A package is imported by multiple other packages, but it can only be initialized once
## 6. Manage external packages
go allows importing code from different code bases. For the external package to be imported by import, you can use the go get command to remove it and place it in the directory corresponding to GOPATH.
For example, if we want to connect to the mysql database through the go language, we need to download the mysql data package first, open the terminal and enter the following command:
```shell
localhost:~ ruby$ go get github.com/go-sql-driver/mysql
```
After installation, you can see the corresponding file package directory under the src of the gopath directory:
![Installation package 1](IMG/installation package 1.PNG)
>In other words, for the go language, it doesn't really care whether your code is internal or external. In short, it is in GOPATH. The path of any import package starts from GOPATH; the only difference is that it is internally dependent. The package is written by the developer, and externally dependent packages are obtained from go get.
## Extension
We can compile the package file through go install.
We know that a non-main package will generate a .a file after compilation (generated in a temporary directory, unless you use go install to install under `$GOROOT` or ​`$GOPATH`, otherwise you will not see .a), use Link to use in subsequent executable programs.
For example, the path of the source code part of the package in the Go standard library is: `$GOROOT/src`, and the path of the compiled .a file in the standard library is under `$GOROOT/pkg/darwin_amd64`.
Qianfeng Go language learning group: 784190273
Author B station:
HTTPS://space.proportion.com/353694001
Corresponding video address:
Https://oooooo.proportion.com/video/av56018934
Https://oooooo.proportion.com/video/av47467197
Source code:
Https://GitHub.com/Ruby with 1314/go_advanced
