## Golang-100 days from novice to master
> Author: Han Ru, Davie, Steven
>
> Recently, many friends are looking for complete learning materials for the Go language, but it takes a lot of time to record videos and tutorials. I usually prepare for the Go language subject, so the time is relatively tight. Teacher Davie and I are responsible for part of it. The output of Golang content. From technical articles, to videos, to project code. They are also published on major platforms, but Zhihu is only convenient for reading technical articles, and B station is only convenient for watching videos. So we uploaded all of our learning materials on github, from the most basic entry to project design, hoping to help more partners who want to understand and learn the Go language, so that everyone can communicate and learn. We have established a learning discussion group (Go Language Learning Camp: 784190273) to join the group for learning and discussion.
>
> Because it is continuous creation, it will continue to be updated. Some chapter lists have no content yet, so stay tuned. .
>
> Creation is not easy, thank you for your support. If you find something after reading it, you can give the author a cup of coffee. If you have any questions, you can join the group discussion.
>
> Finally, thank you **Qianfeng Education Go Language Teaching Department** for your support.

![WechatIMG723_meitu_1](img/WechatIMG724_meitu_2.jpg)

### Go language application field and employment analysis
Go language is the second open source programming language (system development language) released by Google in 2009. 
It is a programming language based on compilation, garbage collection and concurrency.
Go language is optimized for the programming of multi-processor system applications. Programs compiled with Go can be comparable to the speed of C/C++ code, and are safer and support parallel processes. As a language that appeared in the 21st century, its near-C execution performance, near-analytic language development efficiency, and near-perfect compilation speed have become popular all over the world. Especially in cloud projects, most of them are developed using Golang. I have to say that Golang has long been deeply rooted in people's hearts. For a new project without historical burden, Golang may be the best choice.
**Golang's philosophy: "Less is more or less is less". **
-Easy learning curve
-Efficiency: fast compilation time, high development efficiency and operating efficiency
-"Famous background, pure blood"
-Freedom and efficiency: combined thinking, non-intrusive interface
-Powerful standard library
-Easy deployment: Binary files, Copy deployment
-Parallel and asynchronous programming is almost painless
**Currently in several popular fields, Go is useful. **
-Cloud computing infrastructure field
Representative projects: docker, kubernetes, etcd, consul, cloudflare CDN, Qiniu cloud storage, etc.
-Basic software
Representative projects [tidb, influxdb, cockroachdb, etc.
-Microservice
Representative projects: go-kit, micro, monzo bank's typhon, bilibili, etc.
-Internet infrastructure
Representative projects: Ethereum, hyperledger, etc.
-Distributed systems, database agents, middleware, etc., such as Etcd
-DevOps-Go / Python / Shell / Ruby
**As a Go language developer, the main areas of employment include:**
-Golang Development Engineer / Golang Development Engineer
-Golang server background development / game server development
-Cloud computing platform (golang) development engineer
-Blockchain development (golang) engineer
-Golang architect
**A few suggestions for beginners:**
-Make English as your working language.
-Practice makes perfect.
-All experience comes from mistakes.
-Don't be one of the leeches.
-Either stand out or kicked out.
### Day01~15-[Go Language Fundamentals](./Day01-15(Go Language Fundamentals))
#### Day01-[First Understanding of Go Language](./Day01-15(Go Language Fundamentals)/day01_Section 8_The first program HelloWorld.md)
-Introduction to Go Language-History of Go Language / Core Features of Go Language / Logo Version of Go Language / Application Fields of Go
-Go language environment construction-Windows system / Linux system / MacOS system
-HelloWorld program-Go language file structure format / fmt package / Print function
-The execution principle of Go-Go commands
-Install IDE-Goland tools / other IDE
-Comments-The role of comments / single-line comments / multi-line comments
- Coding Standards
#### Day02-[Basic Grammar](./Day01-15(Go language foundation)/day02_Basic Grammar.md)
-Base and conversion-Computer principle / Binary / Decimal / Octal / Hexadecimal / Base conversion
-Variables-Naming of variables / Use of variables / Analysis of variables / Notes on variables
-Constants-constant naming / constant use / constant analysis / constant attention / iota keyword
#### Day03-[Data Type & Operator](./Day01-15 (Go Language Foundation)/day03_Data Type & Operator.md)
-Data Type-Integer / Floating Point / Complex Number / String / Character Encoding
-Data type conversion-forced conversion / automatic conversion
-Operators-Arithmetic Operators / Assignment Operators / Comparison Operators / Logical Operators / Bitwise Operators / Shift Operators / Priority of Operators
-Expression-expression value / expression type
-Keyboard input and print output-Scanln() / Scanf() / Print() / Printf() / Println()
-Format placeholders-%v / %T / %t / %s / %f /% d / %p / %c. . .

#### Day04-[Branch Statement](./Day01-15(Go Language Fundamentals)/day04_Branch Statement.md)
-Application scenarios of branch structure-Condition / structure / code block / flow chart
-if statement-simple if / if-else structure / if-elif-else structure / nested if / other ways of writing
-switch statement-switch structure / case statement / break statement / fallthrough statement / other ways of writing switch

#### Day05-[Loop Statement](./Day01-15(Go Language Fundamentals)/day05_Loop Statement.md)
-Application scenarios of cyclic structure-Condition / structure / code block / flow chart
-for loop-basic structure / branch structure in loop / nested loop / other wording of for
-Loop control statement-break / continue
-goto statement-goto structure / goto points
-Random number generation-Random number generation
-Application case-1~100 summation / judging prime numbers / guessing the number game / printing the ninety-nine table / printing the triangle pattern / the number of daffodils / hundred money and hundred chickens

#### Day06-[Array](./Day01-15(Go Language Fundamentals)/day06_Array.md)
-Array-Array Concept / Use of Array / Syntax of Array / Length of Array / Points to Note for Array
-Array traversal-Array subscript / normal for traversal array / for...range traversal
-Sorting of arrays-Bubble sorting/selection sorting/insertion sorting. . .
-Multi-dimensional array-Two-dimensional array / use of two-dimensional array / traversal of two-dimensional array
-Array data type-Array is value type data
#### Day07-[Slice](./Day01-15(Go Language Fundamentals)/day07_Slice Use.md)
-Slices-The concept of slices / The use of slices / The syntax of slices / The length and capacity of slices / Points to note about slices
-The principle of slicing-the underlying array of the slice / create a slice on an existing array / intercept a slice
-Slice traversal-Slice subscript / normal for traversal slice / for...range traversal
-Slice related functions-make() / append() / copy() / len() / cap()
-Slice copy-Deep copy / Shallow copy
-Slice data type-Slice is reference type data
#### Day08-[Map](./Day01-15(Go Language Fundamentals)/day08_Map usage.md)
-Map-Concept of Map / Use of Map / Syntax of Map / Length and Capacity of Map / Points to Note for Map
-Map storage characteristics-key-value / Map key type / Key-value pairs in Map are out of order
-Creation of Map-Empty Map /
-Map operation-add data / modify data / get data / delete data
-Map data type-Map is reference type data
#### Day09-[string](./Day01-15(Go Language Fundamentals)/day09_string.md)
-Use of strings-Calculating length / subscript operation / slicing / common methods
-strings package
-strconv package
#### Day10-[Function](./Day01-15(Go Language Fundamentals)/day10_function.md)
-Function-Concept of function / Role of function
-Function syntax-Define function / call function
-Function parameters-Use of parameters / variable parameters / parameter transfer
-Return value of function-Return value / return statement / No return value / Return single value / Return multiple values
-Variable scope-local variable / global variable
-Recursive function-Recursive algorithm / recursive function implementation
-Advanced functions-The essence of functions / anonymous functions / higher-order functions / callback functions / closure structure
-Defer function-Defer function / delay parameter / delay of stack / defer note
#### Day11-[Package Management](./Day01-15(Go Language Fundamentals)/day11_包管理.md)
-Package management-package concept / package definition / package import / main package
-Package keyword-package / import
-Package import logic-init() function / multiple init() functions in the same package / multiple init() functions in different packages
-Manage external packages
#### Day12-[Pointer](./Day01-15(Go Language Fundamentals)/day12_Pointer.md)
-Pointer-The concept of pointer / Get the address of a variable / Operate the pointer to change the variable
-Pointer syntax-define pointer / * / get pointer value / & / pointer pointer
-Points to note about pointers-Pointer type / pointer address / null pointer
-Pointer application-Pointer as parameter / Pointer as function return value / Array pointer and pointer array / Pointer function and function pointer
#### Day13-[Structure](./Day01-15(Go Language Fundamentals)/day13-Structure.md)
-Structure-Structure concept / structure definition / structure initialization / structure access / make and new
-Anonymous fields of the structure
-Structure Nesting-Structure Nesting / Anonymous Nesting / Promoting Fields
-Use of structure-Structure pointer / structure as function parameter / structure as return value of function
#### Day14-[Methods and Interfaces](./Day01-15(Go Language Fundamentals)/day14_第1节_method.md)
-Method-Concept of method / Use of method
-Method syntax-method receiver / method and function
-Methods in structure nesting-"Inheritance" of methods / "Rewriting" of methods
-Interface-Concept of interface / Use of interface
-Interface syntax-interface / interface type
-Interface application-empty interface / type assertion / type keyword
#### Day15-[Error Handling](./Day01-15(Go Language Fundamentals)/day15_error handling.md)
-Wrong-Wrong concept / Wrong usage /
-Error type-error interface / error type indication / custom error
-Error handling-Return error function / error handling
-Related knowledge-panic() function / recover() function / defer() function
### Day16~20-[Advanced Go Language Fundamentals](./Day16-20(Advanced Go Language Fundamentals))
#### Day16-[I/O Operation](goon_ch.md)
-I/O-What is the use of I/O/os package
-File operation-Get file information / file creation
-I/O operation
-Related packages-bufio package / ioutil package
-I / O application-copy files / resume uploading / traverse folders
#### Day17-[Concurrent Programming Goroutine](goon_ch.md)
#### Day18-[Channel](goon_ch.md)
#### Day19-[Reflection Mechanism](goon_ch.md)
#### Day20-[Comprehensive Exercises](goon_ch.md)
### Day21~22-[Network Programming](./Day21-22(Network Programming))
### Day23~24-[MySQL DB Fundamentals](./Day23-24(MySQL DB Fundamentals))
### Day25-[Go Language Connect MySQL](./Day25(Go Connect MySQL))
### Day26~31-[Web front end](./Day26-31(Web front end))
#### Day26-[HTML](goon_ch.md)
#### Day27-[CSS](goon_ch.md)
#### Day28~30-[JavaScript](goon_ch.md)
#### Day31-[jQuery](goon_ch.md)
### Day32~35-[Go Web Development](./Day32-35(Go Web Development))
#### Day32-[Web First Meet](goon_ch.md)
#### Day33-[httpPackage Detailed](goon_ch.md)
#### Day34-[session and cookie](goon_ch.md)
#### Day35-[Text Processing](goon_ch.md)
### Day36~37-[beego frame](./Day36-37(beego frame))
#### Day36-[beego framework introduction and process analysis]()
-beego framework-beego introduction / beego installation / beego features
-bee tools-bee introduction / bee installation
-Usage of bee-bee command
-beego program flow analysis-beego program entry / go language execution sequence
-Beego framework function-request interception / routing distribution
-beego controller-processing logic
-beego.Run method-Parsing configuration / routing distribution / monitoring service
#### Day37-[beego framework summary and database connection configuration]()
-conf configuration-project data configuration / configuration data reading
-controllers-Controller introduction / Controller function / Controller definition
-models- Data layer function / model definition
-routers-routing layer function / routing classification
-Static resources-Static resource directory function / Static resource path setting
-Database installation and configuration-mysql database installation / mysql database basic commands / visualization tools
-Database driver-Database driver classification / mysql driver installation / connection configuration / database connection
### Day38~41-[Project Actual Combat One](./Day38-41(beego framework development blog system)/day38_Project construction, login registration and Session function development.md)
#### Day38-[Project construction, login registration and Session function development](./Day38-41(beego framework development blog system)/day38_Project construction, login registration and Session function development.md)
-bee tool usage-project creation / project operation
-mysql database operation-database connection configuration / read database configuration / connect to database
-models- database table design / implementation of database operation method package
-User registration-view form data submission / server receiving post data / operating database table to add data
-User login-login function controller / routing registration / server receiving Post data / database table condition query
-Session handling-session function enable configuration / add session data / get session data / delete session data
-BaseController-BaseController role / controller method execution order / Parepare method role
#### Day39-[Develop article, project home page and view article details function development](./Day38-41 (Beego framework development blog system)/day39_ write article, project home page and view article details function development.md)
-model layer-database table design / database addition operation
-Controller-Write article function controller definition / get method to parse html page / Post method to receive form data
-Routing layer-Registration routing analysis
-View layer-html function page / js logic judgment / js form submission
-Paging function-Paging design / database limit operation to achieve paging query
-Homepage content display-model conversion / pagination view function
-Markdown syntax-Common third-party libraries / markdown format programming syntax / markdown and html conversion
#### Day40-[Development of the function of writing articles, project homepage and viewing article details](./Day38-41(Develop blog system with beego framework)/day40_Modify articles, delete articles and article tags function development.md)
-Modify article function-Register and modify function routing / modify function controller logic implementation / get method to render page / post method to receive form data / database table data modification operation
-Delete function-Delete database table data according to conditions / view redirection after deletion
-Label function-Database query
#### Day41-[Development of article writing, project homepage and viewing article details](./Day38-41(beego framework development blog system)/day41_Homepage function extension, picture uploading and function development.md)
-Function extension-tags query / page query / multi-condition logic judgment / multi-condition query & use
-File upload-Data sheet design / js realizes file submission / server receives file data / file type judgment / file size judgment / file name / save file
-Project Summary-Beego framework composition / Beego debugging tool / Beego program execution process / Database operation / Beego project architecture / session processing / Template file syntax
### Day42~43-[Gin frame](./Day42-43(Gin frame))
### Day44-[MySQL database advanced](./Day44(MySQL database advanced))
### Day45-[Git](./Day45(Git))
### Day46~50-[Project Actual Combat Two](./Day46-50(Project Actual Combat Two))
### Day51-[Node.js](./Day51(Node.js))
### Day52-[Vue](./Day52(Vue))
### Day53-[Redis Database](./Day53(Redis Database))
### Day54~55-[Iris framework](./Day54-55(iris framework)/day54_web development introduction, iris framework installation, HTTP request and return, iris routing processing.md)
#### Day54-[web development introduction, iris framework installation, HTTP request and return, iris routing processing](./Day54-55(iris framework)/day54_web development introduction, iris framework installation, HTTP request and return, iris routing Processing.md)
-web development-project structure / development process / actual project introduction / project technology stack
-iris framework-iris introduction / iris features / iris frame installation / iris reference materials
-http request and processing-data request and classification / http1.0 and http1.1 / iris standard request processing / custom request processing / request processing data format package
-Routing processing-Context concept / regular expression routing
#### Day55-[Frame setting, mvc package, session usage, project construction and resource import](./Day54-55(iris framework)/day55_frame setting, mvc package, session usage, project construction and resource import. md)
-Routing group-Party implementation routing group / context.Next() method / taml configuration file / yaml configuration file setting / custom configuration file / custom configuration setting
-mvc package-mvc.Application role / mvc characteristics / life cycle / mvc.Configure configuration
-Session processing and use-The difference between session and cookie / session support data type / session creation / session use
-Actual combat project-Create enterprise management platform project / Catalog description / Project resource integration
### Day56~60-[Project Actual Combat Three](./Day56-60(Project Actual Combat Three))
### Day61-[Linux](./Day61(Linux))
### Day62~64-[Container](./Day62-64(Container))
#### Day62-[Virtualization VS Containerization](goon_ch.md)
#### Day63-[Docker](goon_ch.md)
#### Day64-[Kubernetes(k8s)](goon_ch.md)
### Day65~75-[Distributed](./Day65-75(Distributed))
#### Day65-[Distributed Theory](goon_ch.md)
#### Day66~67-[Distributed File System FastDFS](goon_ch.md)
#### Day68-[Nginx and Echo Proxy Deployment](goon_ch.md)
#### Day69~70-[Go development to achieve high availability etcd system](goon_ch.md)
#### Day60~75-[Project Actual Combat 4: Distributed Project](goon_ch.md)
### Day76~95-[Microservice](./Day76(Microservice features)/day76_Introduction to microservices and features.md)
#### Day76-[Introduction to Microservices, Monolithic Applications, Microservice Solutions, Microservice Practice Issues](./Day76(Microservice Features)/day76_Microservice Introduction and Features Introduction.md)
-Introduction to microservices-Building a single application / Internet company structure
-Single application dilemma-Application development trends / difficulties and dilemmas
-Solutions to complex problems-Microservice architecture / system scalability / advantages of microservices / lack of microservices
-Microservice Practice-Definition of Microservice / Problems to be Solved by Microservice Practice / Common Microservice Architecture
#### Day77-[Protobuf introduction, programming implementation, Protobuf syntax](./Day77(protobuf)/day77_Protobuf introduction.md)
-Introduction to Protobuf-Introduction to Protobuf / Protobuf application scenarios / Protobuf advantages / Protobuf disadvantages
-Programming to achieve Protobuf transmission-Environment preparation / Protobuf protocol syntax / Steps to use Protobuf
-Protobuf syntax-Protobuf protocol syntax / Protobuf serialization principle
#### Day78~79-[Microservice Management](./Day78-79(Microservice Management)/day78_微服务管理(上).md)
##### Day78-[Principle of Service Discovery, Consul Construction and Configuration, Service Management, Cluster Construction](./Day78-79(Microservice Management)/day78_微服务管理(上).md)
-Service Discovery / Reasons for Using Service Discovery / Service Discovery Scheme / Introduction to Consul / Distributed and Single Point of Failure / Consul Environment Configuration / Consul Internal Principles / Consul Startup / Consul Node Discovery
##### Day79-[Define microservices, service query, cluster construction](./Day78-79(microservice management)/day79_微服务管理(下).md)
-Service management / service definition / service registration and discovery / service query / registration of multiple services / multi-node service cluster / Docker environment construction
#### Day80-[RPC introduction, RPC programming and implementation, RPC and Protobuf combination](./Day80(RPC remote call mechanism)/day80_RPC remote procedure call.md)
-RPC introduction-RPC introduction and principle introduction / local procedure call / RPC technology implementation
-RPC programming and implementation-RPC official library / net/rpc library to implement RPC call programming
-RPC combined with Protobuf-Transmission data format definition / Protobuf format data combined with RPC
#### Day81~82-[gRPC remote procedure call](./Day81-82(gRPC remote call mechanism)/day81_gRPC remote call mechanism introduction.md)
##### Day81-[gRPC introduction and installation, use of gRPC framework](./Day81-82(gRPC remote call mechanism)/day81_gRPC remote call mechanism introduction.md)
-gRPC introduction and installation-What is gRPC / grpc-go introduction
-gRPC framework use-define service / compile .proto file / gRPC implement RPC programming
##### Day82-[gRPC programming use, TLS verification and Token authentication, interceptor and custom interceptor use](./Day81-82(gRPC remote call mechanism)/day82_gRPC use.md)
-gRPC programming use-server stream RPC / compile .protoc files / server code / client code generation / service registration and monitoring processing / client data reception / client stream mode / two-way stream mode
-TLS verification and Token authentication-Authorized authentication / SSL/TLS working principle / Token authentication process / Custom Token authentication
-Use of Interceptor-Introduction to Interceptor / Custom UnaryServerInterceptor / Interceptor Registration
#### Day83~85-[go-micro microservice framework](./Day83-85(go-micro microservice framework)/day83_go-micro framework introduction.md)
##### Day83-[micro framework introduction and use, micro creation of microservices](./Day83-85(go-micro microservice framework)/day83_go-micro framework introduction.md)
-Introduction to the micro framework-Background / Introduction to the micro / Micro components / Micro tool components: API, Web, Sidecar, Bot / Go-Micro features / Consul installation and environment preparation
-Create microservice-Definition of microservice / Initialize service instance / Options optional configuration / Define service interface, realize service business logic / Microservice call / Customize specified service discovery component
##### Day84-[Heartbeat mechanism and optional configuration, event-driven mechanism](./Day83-85(go-micro microservice framework)/day84_go-micro use(一).md)
-Heartbeat mechanism and optional configuration-The origin of heartbeat mechanism / Consul heartbeat configuration / TTL and interval time
-Event-driven mechanism-Publish/subscribe mechanism / Broker component design / go-plugins installation / MQTT introduction and environment construction / Programming to achieve message subscription and publishing
##### Day85-[micro framework Selector mechanism](./Day83-85(go-micro microservice framework)/day85_go-micro use (2).md)
-Selector mechanism-Load balancing algorithm / Mico's Selector / Selector definition / DefaultSelector / registrySelector
#### Day86-[RESTful Design and Use](./Day86(RESTful Design)/day86_RESTful Standard Design.md)
-RESTful design specification and use-Go-Micro API gateway / Micro tool installation / Micro API working principle / Reverse proxy API service startup / REST mapping
#### Day87-[Microservice Project Design](goon_ch.md)
#### Day88-[RPC remote call mechanism](./goon.md5)
#### Day89~95-[Project Actual Combat Five: Microservice Project](goon_ch.md)
### Day96~100-[Perfect Ending](./Day96-100(Perfect Ending))
#### Day96~97-[Project Deployment and Performance Tuning](goon_ch.md)
#### Day98-[Project Summary](goon_ch.md)
#### Day99-[Interview Guidance](goon_ch.md)
#### Day100-[English Interview](goon_ch.md)
> Thanks:
>
>​ Thanks to Qianfeng Education and colleagues in the Go language group of Qianfeng Education: Mr. Steven, Mr. Davie, etc. for their technical knowledge and help. # Golang-100-Days
