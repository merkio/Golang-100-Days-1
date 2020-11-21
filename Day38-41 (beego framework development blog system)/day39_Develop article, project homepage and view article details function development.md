# Write article, project homepage and view article details function development
**@author: Davie**
**Copyright: Beijing Qianfeng Internet Technology Co., Ltd.**
Click to write an article, we enter the page of writing an article, when the user clicks the button to submit, the article data should be stored in the database.
## One writing article function development
### 1.1 Database table design
First, we first design the database, the articles submitted by users, including the title, tags, introduction, content, creation time, etc.
In the mysqlUtils.go file, add the operation of the article table:
```go
//Create article table
func CreateTableWithArticle(){
sql:=`create table if not exists article(
id int(4) primary key auto_increment not null,
title varchar(30),
author varchar(20),
tags varchar(30),
short varchar(255),
content longtext,
createtime int(10)
);`
ModifyDB(sql)
}
```
### 1.2 Model layer implementation
#### 1.2.1 Article structure definition
Create a go file in the model directory: article_model.go
```go
package models
import "myblog/utils"
type Article struct {
Id int
Title string
Tags string
Short string
Content string
Author string
Createtime int64
//Status int //Status=0 means normal, 1 means delete, 2 means frozen
}
```
#### 1.2.2 Add articles and database operations
```go
//---------data processing-----------
func AddArticle(article Article) (int64, error) {
i, err := insertArticle(article)
return i, err
}
//-----------Database operation---------------
//Insert an article
func insertArticle(article Article) (int64, error) {
return utils.ModifyDB("insert into article(title,tags,short,content,author,createtime) values(?,?,?,?,?,?)",
article.Title, article.Tags, article.Short, article.Content, article.Author, article.Createtime)
}
```
Currently we are just writing articles, so what is needed is to add data.
### 1.3 Controller layer
#### 1.3.1 Add article controller definition
First create a controller file, add_article_controller.go.
```go
package controllers
import (
"fmt"
"myblog/models"
"time"
)
type AddArticleController struct {
BaseController
}
/*
When accessing the /add path, the Get method of AddArticleController is triggered
The responding page is through TpName
*/
func (this *AddArticleController) Get() {
this.TplName = "write_article.html"
}
```
#### 1.3.2 Controller realizes adding article logic processing
```go
//Return the json string through this.ServerJSON() method
func (this *AddArticleController) Post() {
//Get the data transmitted by the browser, and get the value through the name attribute of the form
title := this.GetString("title")
tags := this.GetString("tags")
short := this.GetString("short")
content := this.GetString("content")
fmt.Printf("title:%s,tags:%s\n", title, tags)
//Instantiate the model, put it in and out of the database
art := models.Article{0, title, tags, short, content, "Qianfeng Education", time.Now().Unix()}
_, err := models.AddArticle(art)
//Return data to the browser
var response map[string]interface{}
if err == nil {
//No mistake
response = map[string]interface{}{"code": 1, "message": "ok"}
} else {
response = map[string]interface{}{"code": 0, "message": "error"}
}
this.Data["json"] = response
this.ServeJSON()
}
```
If the user requests to write the article path, the write_article.html page will be displayed. After adding the information, click the submit button to submit the data.
### 1.4 Register to add article routing
Then register a new route:
```go
//write an essay
beego.Router("/article/add", &controllers.AddArticleController{})
```
### 1.5 View layer development
#### 1.5.1 Create a new write_article.html file
We create an html file (write_article.html) in the views directory for writing articles.
```html
<!DOCTYPE html>
<HTML浪="恩">
<head>
<meta charset="UTF-8">
<title>Write an article</title>
<link href="../static/css/blogsheet.css" rel="stylesheet">
<script src="../static/js/lib/jquery-3.3.1.min.js"></script>
<script src="../static/js/lib/jquery.url.js"></script>
<script src="../static/js/blog.js"></script>
</head>
<body>
{{template "block/nav.html" .}}
<div id="main">
<form id="write-art-form" method="post">
<div>Title</div>
<input type="text" placeholder="Please enter a title" name="title">
<div>tag</div>
<input type="text" placeholder="Please enter tags" name="tags">
<div>Introduction</div>
<textarea placeholder="Please enter a brief introduction" name="short"></textarea>
<div>Content</div>
<textarea id="content" placeholder="Please enter content" name="content"></textarea>
<input id="write-article-id" hidden name="id">
<br>
<button type="button" onclick="history.back()">Back</button>
<button type="submit" id="write-art-submit">Submit</button>
</form>
</div>
</body>
</html>
```
#### 1.5.2 Writing js files to realize the function of adding articles
Next, write the js script file and open the blog.js file in the static/js directory.
```js
//Add article form
$("#write-art-form").validate({
rules: {
title: "required",
tags: "required",
short: {
required: true,
minlength: 2
},
content: {
required: true,
minlength: 2
}
},
messages: {
title: "Please enter a title",
tags: "Please enter tags",
short: {
required: "Please enter a profile",
minlength: "Introduction content must be at least two characters"
},
content: {
required: "Please enter the content of the article",
minlength: "The article content is at least two characters"
}
},
submitHandler: function (form) {
var urlStr = "/article/add";
alert("urlStr:" + urlStr);
$(form).ajaxSubmit({
url: urlStr,
type: "post",
dataType: "json",
success: function (data, status) {
alert(":data:" + data.message);
setTimeout(function () {
window.location.href = "/"
}, 1000)
},
error: function (data, status) {
alert("err:" + data.message + ":" + status)
}
});
}
})
```
### 1.6 Project Operation
After running the project, open the browser, log in and enter the homepage, click to write an article:
![Entrance of Blog Writing Function](./img/WX20190521-145856@2x.png)
Then enter the article writing page:
![Write article details](./img/WX20190521-150025@2x.png)
Click the button to submit, and then query the database, the data has been inserted into it.
## Two project home page function realization
In the last lesson, we learned to realize the function of writing articles. In this lesson, we will implement the development of the home page function. The home page is the page to be displayed after the user logs in. The function that I want to achieve in the end is: click on the homepage, it will automatically query the database and display the articles; if there are more articles, we can realize paging.
### 2.1 Query article function
#### 2.1.1 Query article controller
We first modify the home_controller.go file, in the Get() method, first query all articles and display them on the page. Because there may be many articles, in order to have a better user experience, we need to query by page. The first page is queried by default.
```go
func (this *HomeController) Get() {
page, _ := this.GetInt("page")
if page <= 0 {
page = 1
}
var artList []models.Article
artList, _ = models.FindArticleWithPage(page)
this.Data["PageCode"] = 1
this.Data["HasFooter"] = true
fmt.Println("IsLogin:", this.IsLogin, this.Loginuser)
this.Data["Content"] = models.MakeHomeBlocks(artList, this.IsLogin)
this.TplName = "home.html"
}
```
#### 2.1.2 Model layer processing
We first add the query of the article in the article_model.go file, and we need to query by page:
```go
//-----------Query article---------
//Search article according to page number
func FindArticleWithPage(page int) ([]Article, error) {
//Get the number of articles per page from the configuration file
num, _ := beego.AppConfig.Int("articleListPageNum")
page--
fmt.Println("---------->page", page)
return QueryArticleWithPage(page, num)
}
/**
Paging query database
limit paging query statement,
Syntax: limit m, n
m represents how many digits to start from, regardless of the id value
n represents how many pieces of data are obtained
Note that there is where before limit
*/
func QueryArticleWithPage(page, num int) ([]Article, error) {
sql := fmt.Sprintf("limit %d,%d", page*num, num)
return QueryArticlesWithCon(sql)
}
func QueryArticlesWithCon(sql string) ([]Article, error) {
sql = "select id,title,tags,short,content,author,createtime from article "+ sql
rows, err := utils.QueryDB(sql)
if err != nil {
return nil, err
}
var artList []Article
for rows.Next() {
id := 0
title := ""
tags := ""
short := ""
content := ""
author := ""
var createtime int64
createtime = 0
rows.Scan(&id, &title, &tags, &short, &content, &author, &createtime)
art: = Article {id, title, tags, short, content, author, createtime}
artList = append(artList, art)
}
return artList, nil
}
```
#### 2.1.3 Home page display content structure definition
Create a go file in the models directory to control the content displayed on the home page:
```go
type HomeBlockParam struct {
Id int
Title string
Tags [] TagLink
Short string
Content string
Author string
CreateTime string
//View the address of the article
Link string
//Modify the address of the article
UpdateLink string
DeleteLink string
//Record whether to log in
IsLogin bool
}
//Tag link
type TagLink struct {
TagName string
TagUrl string
}
```
We need to convert the data queried from the database into the corresponding structure object, so first design the structure, here we need to consider if the user is logged in, then we can modify or delete an article. Of course, if you are not logged in, you can only view it. So in the design of the structure, we have created a direct link to modify and delete fields.
#### 2.1.4 Homepage content display function
Next, we add a method to display the content of the article on the page:
```go
//----------Homepage display content---------
func MakeHomeBlocks(articles []Article, isLogin bool) template.HTML {
htmlHome := ""
for _, art := range articles {
//Convert the database model to the model required for the home page template
homeParam := HomeBlockParam{}
home PA RAM.ID = art.ID
homeParam.Title = art.Title
homeParam.Tags = createTagsLinks(art.Tags)
fmt.Println("tag-->", art.Tags)
homeParam.Short = art.Short
homeParam.Content = art.Content
home PA RAM.author = art.author
homeParam.CreateTime = utils.SwitchTimeStampToData(art.Createtime)
homeParam.Link = "/article/" + strconv.Itoa(art.Id)
homeParam.UpdateLink = "/article/update?id=" + strconv.Itoa(art.Id)
homeParam.DeleteLink = "/article/delete?id=" + strconv.Itoa(art.Id)
homeParam.IsLogin = isLogin
//Processing variables
//ParseFile parses the file for inserting variables
t, _ := template.ParseFiles("views/block/home_block.html")
buffer := bytes.Buffer{}
//It is to replace the two in the html file with the data passed in
t.Execute(&buffer, homeParam)
htmlHome += buffer.String()
}
fmt.Println("htmlHome-->",htmlHome)
return template.HTML(htmlHome)
}
```
An additional method is needed:
```go
//Convert the tags string into the data structure required by the homepage template
func createTagsLinks(tags string) []TagLink {
var tagLink [] TagLink
tagsPamar := strings.Split(tags, "&")
for _, tag := range tagsPamar {
tagLink = append(tagLink, TagLink{tag, "/?tag=" + tag})
}
return tagLink
}
```
#### 2.1.4 View layer development
Next, we design the page. Just now in the model's MakeHomeBlocks() method, we need to use the template to fill and format the html page content, so we create another html page under views/block: home_block.html, with the following content:
```html
<div id="home-block-item">
<h2><a href="{{.Link}}">{{.Title}}</a></h2>
<div>
<span>{{.CreateTime}}</span>
<span>
{{range .Tags}}
<a href="{{.TagUrl}}">&nbsp{{.TagName}}</a>
{{end}}
</span>
</div>
<p><a href={{.Link}}>{{.Short}}</a></p>
{{if .IsLogin}}
<div class="home-block-item-update">
<a href='javascript:if(confirm("Are you sure to delete?")){location="{{.DeleteLink}}"}'>Delete</a>
<a href={{.UpdateLink}}>Edit</a>
</div>
{{end}}
</div>
```
We display the data of the article queried from the data. If the user is logged in, then we delete and modify it, because the user has these two permissions, otherwise it will not be displayed.
#### 2.1.5 Project Operation
##### 2.1.5.1 Prepare test data
We insert 10 pieces of data in the database:
![Data from the database](./img/WX20190522-113706@2x.png)
##### 2.1.5.1 Page display data number configuration
Next, we set up the configuration file, each page shows 6 (also 8 or 10...),
Modify the app.conf file in the conf directory:
```
appname = myblog
httpport = 8080
runmode = dev
#mysqlConfiguration
driverName = mysql
mysqluser = root
mysqlpwd = yu271400
host = 127.0.0.1
port = 3306
dbname = myblog
#Session
sessionon = true
sessionprovider = "file"
session then = "Qianfeng Education does not have a blog"
sessiongcmaxlifetime = 1800
sessionproviderconfig = "./tmp"
sessioncookielifetime = 1800
articleListPageNum = 6
```
Then start the project, open the browser and enter the URL: [http://127.0.0.1:8080/](http://127.0.0.1:8080/)
![Not logged in homepage effect](./img/WX20190522-141320@2x.png)
Although the user is not logged in, it can be viewed. Next, we click the login button to log in:
![Homepage effect after login](./img/WX20190522-141446@2x.png)
After logging in, the user can delete and modify the function.
### 2.2 Article pagination display function
By the end, we have been able to display the content of the first page, and then we will add the functions of the previous page and the next page.
#### 2.2.1 Paging structure definition
First add a paging structure object in home_model.go:
```go
type HomeFooterPageCode struct {
HasPre bool
HasNext bool
ShowPage string
PreLink string
NextLink string
}
```
#### 2.2.2 Querying Articles and Modifying Methods
Next add the method:
```go
//-----------Page turning-----------
//page is the current page number
func ConfigHomeFooterPageCode(page int) HomeFooterPageCode {
pageCode := HomeFooterPageCode{}
//Query the total number
num := GetArticleRowsNum()
//Read the number of items displayed on each page from the configuration file
pageRow, _ := beego.AppConfig.Int("articleListPageNum")
//Calculate the total number of pages
allPageNum := (num-1)/pageRow + 1
pageCode.ShowPage = fmt.Sprintf("%d/%d", page, allPageNum)
//The current page number is less than or equal to 1, then the button on the previous page cannot be clicked
if page <= 1 {
pageCode.HasPre = false
} else {
pageCode.HasPre = true
}
//The current page number is greater than or equal to the total number of pages, then the button of the next page cannot be clicked
if page >= allPageNum {
pageCode.HasNext = false
} else {
pageCode.HasNext = true
}
pageCode.PreLink = "/?page=" + strconv.Itoa(page-1)
pageCode.NextLink = "/?page=" + strconv.Itoa(page+1)
return pageCode
}
```
This code needs to query the total amount of all articles in the database, so we must first add the method of querying the total amount of data in the article_model.go file:
```go
//------Turn page------
//The number of rows in the storage table can only be changed by yourself. This value needs to be updated when an article is added or deleted
var artcileRowsNum = 0
//Only when the number of rows is first obtained, the number of rows in the statistical table is taken
func GetArticleRowsNum() int {
if artcileRowsNum == 0 {
artcileRowsNum = QueryArticleRowNum()
}
return artcileRowsNum
}
//Query the total number of articles
func QueryArticleRowNum() int {
row := utils.QueryRowDB("select count(id) from article")
num := 0
row.Scan(&num)
return num
}
```
We also have to consider a problem, that is, when adding or deleting articles, the total amount of data will change, so we must modify the method of adding articles:
First add a method to set the total number of pages:
```go
//Set the number of pages
func SetArticleRowsNum(){
artcileRowsNum = QueryArticleRowNum()
}
```
Then modify the method of adding articles:
```go
//---------Add article-----------
func AddArticle(article Article) (int64, error) {
i, err := insertArticle(article)
SetArticleRowsNum()
return i, err
}
```
#### 2.2.3 Home page controller adds paging processing logic
Modify the Get() method of home_controller.go:
```go
func (this *HomeController) Get() {
...
artList, _ = models.FindArticleWithPage(page)
this.Data["PageCode"] = models.ConfigHomeFooterPageCode(page)
this.Data["HasFooter"] = true
...
}
```
#### 2.2.4 Modify the home page view and add paging controls
Modify the home.html page:
```html
<!DOCTYPE html>
<HTML浪="恩">
<head>
<meta charset="UTF-8">
<title>Home</title>
<link href="../static/css/blogsheet.css" rel="stylesheet">
</head>
<body>
{{template "block/nav.html" .}}
<div id="main">
{{.Content}}
{{if .HasFooter}}
<div id="home-footer">
<a {{if .PageCode.HasPre}}href="{{.PageCode.PreLink}}" {{else}} class="disable" {{end}}>Previous page</a>
<span>{{.PageCode.ShowPage}}Page</span>
<a {{if .PageCode.HasNext}}href="{{.PageCode.NextLink}}" {{else}} class="disable" {{end}}>Next page</a>
</div>
{{end}}
</div>
</body>
</html>
```
Add links to previous and next pages.
#### 2.2.5 Project Operation
First, we insert 5 pieces of data into the database, and then modify the configuration file to display 5 pieces per page.
![Paging display](./img/WX20190522-141839@2x.png)
Next, we add a new article, click to write a blog:
![Write new article](./img/WX20190522-142111@2x.png)
The last page is displayed as an article we just added. Up to now, we can display the page number perfectly.
## Three article details function development
In the last lesson, we implemented the article list function and paging function of the project homepage. This lesson will continue with the development and implementation. When clicking on the article, the detailed content of the article should be displayed.
### 3.1 View article details
#### 3.1.1 Add routing settings
First set up routing:
```go
func init() {
...
//write an essay
beego.Router("/article/add", &controllers.AddArticleController{})
//Display article content
beego.Router("/article/:id", &controllers.ShowArticleController{})
}
```
#### 3.1.2 Add controller and logic implementation
Then create a go file in the controllers directory, show_article_controller.go:
```go
type ShowArticleController struct {
//Bee go.controller
BaseController
}
func (this *ShowArticleController) Get() {
idStr := this.Ctx.Input.Param(":id")
id, _ := strconv.Atoi(idStr)
fmt.Println("id:", id)
//Get the article information corresponding to id
art := models.QueryArticleWithId(id)
this.Data["Title"] = art.Title
this.Data["Content"] = art.Content
//this.Data["Content"] = utils.SwitchMarkdownToHtml(art.Content)
this.TplName="show_article.html"
}
```
#### 3.1.3 Modify model and add query method
Next, in the article_model.go file, add a method to query articles based on id:
```go
//----------Query article-------------
func QueryArticleWithId(id int) Article {
row := utils.QueryRowDB("select id,title,tags,short,content,author,createtime from article where id=" + strconv.Itoa(id))
title := ""
tags := ""
short := ""
content := ""
author := ""
var createtime int64
createtime = 0
row.Scan(&id, &title, &tags, &short, &content, &author, &createtime)
art := Article{id, title, tags, short, content, author, createtime}
return art
}
```
#### 3.1.4 Create article detail view
Next, we create the view, in the views directory, create a new html page file, show_article.html:
```html
<!DOCTYPE html>
<HTML浪="恩">
<head>
<meta charset="UTF-8">
<title>{{.Title}}</title>
<link href="../static/css/blogsheet.css" rel="stylesheet">
</head>
<body>
{{template "block/nav.html" .}}
<div id="main">
<h1>{{.Title}}</h1>
<div>{{.Content}}</div>
</div>
</body>
</html>
```
#### 3.1.5 Project Operation
Next we restart the project, refresh the page, and click on an article:
![Article Details](./img/WX20190522-144626@2x.png)
### 3.2 Development supports Markdown format
Although the page can display the content of the article, it looks very uncomfortable, like a pot of gruel, we display it in the markdown syntax format.
#### 3.2.1 Markdown syntax function related libraries
Let's first understand and familiarize ourselves with the libraries we need to use before proceeding with project development. include:
-Convert Markdown syntax: russross/blackfriday
-Find the content of Document: PuerkitoBio/goquery
-Syntax highlighting: sourcegraph/syntaxhighlight
-Insert module: html/template
-Execute external commands: os/exec
-File operation: path/filepath
-Create a web server: SimpleHTTPServer
-Parse the .yml configuration file: gopkg.in/yaml.v2
First, you need to install the markdown installation package:
Open the terminal and enter the following command:
```shell
go get github.com/russross/blackfriday
go get github.com/PuerkitoBio/goquery
go get github.com/sourcegraph/syntaxhighlight
```
#### 3.2.2 Tool library source code
After installation, you can also go to the src directory to view:
![View Tool Library Download](./img/WX20190522-144232@2x.png)
### 3.3 Syntax Introduction
#### 3.3.1 russross/blackfriday package
The third-party library russross/blackfriday is used to use markdown syntax in golang.
**markdown**: It is a markdown language that can be written by a common text editor. Through a simple markup syntax, it can make common text content have a certain format.
Markdown has a series of derivative versions that are used to extend the functions of Markdown (such as tables, footnotes, embedded HTML, etc.). These functions are not available in the original Markdown. They can convert Markdown into more formats, such as LaTeX, Docbook . The more famous Markdown enhanced versions include Markdown Extra, MultiMarkdown, Maruku, etc. These derivative versions are either based on tools, such as Pandoc; or based on websites, such as GitHub and Wikipedia, which are basically compatible in syntax, but have some changes in syntax and rendering effects.
test. Face:
```markdown
## One, russross/blackfriday package
```
Sample code:
```go
func main() {
fileread, _ := ioutil.ReadFile("extra/blackfriday conversion markdown/test.md")
//Convert Markdown syntax, such as converting "#" to "<h1></h1>"
subHtml := blackfriday.MarkdownCommon(fileread)
subHtmlStr := string(subHtml)
fmt.Println(subHtmlStr)
}
```
Output:
```
<H2> a, russross / blackfriday bag </ h2>
```
#### 3.3.2 PuerkitoBio/goquery package
Anyone who has done web development should have used or heard of jQuery, which provides a convenient API for manipulating the DOM. Using Go language for server-side development, sometimes you need to parse HTML files, such as crawling website content, writing a crawler, etc. If there is a jQuery-like library available at this time, it will be very convenient to manipulate the DOM, and it will be quick to get started. `PuerkitoBio/goquery` This library implements jQuery-like functions, allowing you to conveniently use Go language to manipulate HTML documents.
The library provides few types, but many methods, and we cannot explain each method one by one. Here, the use of the library is explained by simulating several usage scenarios.
##### 3.3.2.1 Document
Document represents an HTML document to be manipulated, but, unlike jQuery, it loads part of the DOM document.
```go
type Document struct {
*Selection
Url *url.URL
rootNode *html.Node // The root node of the document
}
```
Because a Selection type is embedded in the Document, the Document can directly use the Selection type method. There are five ways to obtain a Document instance.
##### 3.3.2.1.2 Selection
Selection represents a collection of nodes that meet specific conditions.
```go
type Selection struct {
Nodes []*html.Node
document *Document
prevSel *Selection
}
```
Generally, after obtaining a Document instance, obtain a Selection instance through the Dcoument.Find method, and then use chain syntax and methods to manipulate it like jQuery.
The methods provided by the Selection type can be divided into the following categories
-Position operations similar to functions
-Expand Selection collection (add selected nodes)
-Filtering method to reduce node collection
-Loop through selected nodes
-Modify documents
-Detect or obtain node attribute values
-Query or display the identity of a node
-Jump back and forth between document trees (commonly used method of finding nodes)
##### 3.3.2.1.3. Basic usage:
1. Create a document
```go
d,e := goquery.NewDocumentFromReader(reader io.Reader)
d,e := goquery.NewDocument(url string)
```
2. Find content
```go
ele.Find("#title") //Find according to id
ele.Find(".title") //Find according to class
ele.Find("h2").Find("a") //Chain call
```
3. Get content
```go
ele.Html()
ele.Text()
```
4. Get attributes
```go
ele.Attr("href")
ele.AttrOr(“href”, “”)
```
5. Traverse
```go
ele.Find(“.item”).Each(func(index int, ele *goquery.Selection){
})
```
Example:
```go
func main() {
doc, err := goquery.NewDocument("http://studygolang.com/topics")
if err != nil {
log.Fatal(err)
}
doc.Find(".topic").Each(func(i int, contentSelection *goquery.Selection) {
title := contentSelection.Find(".title a").Text()
//Find(".title a") is the same as Find(".title").Find("a")
fmt.Println("th", i+1, "the title of the post:", title)
//ret,_ := contentSelection.Html()
//fmt.Printf("\n\n\n%v", ret)
//SCUD.printed N(content selection.text())
})
//The final output is an html document:
//new, err := doc.Html()
}
```
The input string in Find is CSS selector, and its grammatical style is referred to http://www.w3school.com.cn/cssref/css_selectors.asp. Such as:
| Grammar | Expression |
| --------------- | --------------------------------- ---------------- |
| #firstname | Select all elements with id="firstname". |
| * | Select all elements. |
| p | Select all <p> elements. |
| div,p | Select all <div> elements and all <p> elements. |
| div p | Select all <p> elements inside the <div> element. |
| div>p | Select all <p> elements whose parent element is the <div> element. |
| div+p | Select all <p> elements immediately after the <div> element. |
| [target] | Select all elements with target attribute. |
| [target=_blank] | Select all elements with target="_blank". |
| a[src*=”abc”] | Select each <a> element whose src attribute contains the “abc” substring. |
| a[src$=”.pdf”] | Select all <a> elements whose src attribute ends with “.pdf”. |
#### 3.3.3 sourcegraph/syntaxhighlight package
The syntaxhighlight package provides syntax highlighting of code. It currently uses a language-independent lexical analyzer and performs well on JavaScript, Java, Ruby, Python, Go and C. The main AsHTML(src []byte) ([]byte, error) function, the output is the same CSS class of HTML and google-code-prettify, so any style sheet should also be applied to this package.
```go
func main() {
src := []byte(`
/* hello, world! */
var a = 3;
// b is a cool function
function b() {
return 7;
}`)
highlighted, err := syntaxhighlight.AsHTML(src)
if err != nil {
fmt.Println(err)
os.Exit(1)
}
fmt.Println(string(highlighted))
}
```
Output
```
<span class="com">/* hello, world! */</span>
<span class="kwd">var</span> <span class="pln">a</span> <span class="pun">=</span> <span class="dec">3</span> span><span class="pun">;</span>
<span class="com">// b is a cool function</span>
<span class="kwd">function</span> <span class="pln">b</span><span class="pun">(</span><span class="pun">)</span> span> <span class="pun">{</span>
<span class="kwd">return</span> <span class="dec">7</span><span class="pun">;</span>
<span class="pun">}</span>
```
Conversion by the following rules
```go
var DefaultHTMLConfig = HTMLConfig{
String: "str",
Keyword: "kwd",
Comment: "com",
Type: "typ",
Literal: "lit",
Punctuation: "pun",
Plaintext: "pln",
Tag: "tag",
HTMLTag: "htm",
HTMLAttrName: "atn",
HTMLAttrValue: "atv",
Decimal: "dec",
Whitespace: "",
}
```
### 3.4 Modify project code
#### 3.4.1 Import style
First import the style package on the show_article.html page:
```html
<!DOCTYPE html>
<HTML浪="恩">
<head>
...
<link href="../static/css/lib/highlight.css" rel="stylesheet">
</head>
```
#### 3.4.2 Add Markdown syntax conversion method
Next, add a method in the myUtils.go file in the utils directory:
```go
func SwitchMarkdownToHtml(content string) template.HTML {
markdown := blackfriday.MarkdownCommon([]byte(content))
//Get the html document
doc, _ := goquery.NewDocumentFromReader(bytes.NewReader(markdown))
/**
For document process query, the selector has the same syntax as css
The first parameter: i is the number of elements found
The second parameter: selection is the element to be queried
*/
doc.Find("code").Each(func(i int, selection *goquery.Selection) {
light, _ := syntaxhighlight.AsHTML([]byte(selection.Text()))
selection.SetHtml(string(light))
fmt.Println(selection.Html())
fmt.Println("light:", string(light))
fmt.Println("\n\n\n")
})
htmlString, _ := doc.Html()
return template.HTML(htmlString)
}
```
#### 3.4.3 Modify the controller program
Finally, modify the Get() method in the controller:
```go
func (this *ShowArticleController) Get() {
...
this.Data["Title"] = art.Title
//this.Data["Content"] = art.Content
this.Data["Content"] = utils.SwitchMarkdownToHtml(art.Content)
this.TplName="show_article.html"
}
```
### 3.5 Project operation effect
After restarting the project, then refresh the page:
![Support Markdown syntax effect](./img/WX20190522-143848@2x.png)
As you can see, the effect is better, which achieves our beautification effect on the page.
