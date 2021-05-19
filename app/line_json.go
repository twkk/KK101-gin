package main

import (
	"github.com/gin-gonic/gin"
//	"html/template"
	"strconv"
	"net/http"
//	"log"
	"fmt"
)

var counter = 0
var rd1 		= []position{{Title: "manager",Department: "RD2"}}
var allEmployee1 []employee

func main() {
	allEmployee1 = append(allEmployee1, employee{ID: "201", Email: "xyz@gmail.com", Mobile: "0982080104", Position: rd1[0]})
	r := gin.Default()
	r.LoadHTMLGlob("views/*")
	r.GET("/users", func(c *gin.Context) {
		c.JSON(200, user{ID: 123, Name: "張三", Age: 20})
	})
	r.GET("/employee", func(c *gin.Context) {
		c.IndentedJSON(200, allEmployee1)
	})
	r.GET("/employee/1", func(c *gin.Context) {
		c.IndentedJSON(200, rd1)
	})
	r.GET("/Indented/users", func(c *gin.Context) {
		c.IndentedJSON(200, allEmployee1)
	})
	r.GET("/pureJson", func(c *gin.Context) {
		c.PureJSON(200, gin.H{
			"message": "<b>Hello, world!</b>",
			})
	})
	r.GET("/asciiJSON", func(c *gin.Context) {
		c.AsciiJSON(200, gin.H{"message": "hello 戲子無情"})
	})
	
	r.GET("/fire/:input", fire)
	r.GET("/hire/:input", hire)
	//http.HandleFunc("/modify", modify)         //設定存取的路由
	r.GET( "/modify", get_modify)
	r.Any( "/modify2", modify2)
	r.POST("/modify", post_modify)
	r.Run(":8080")
}


type user struct {
	ID   int
	Name string
	Age  int
}

type position struct{
	Title 		string
	Department	string
}
type employee struct {
	ID   	string
	Email 	string
	Mobile  string
	Position position
}

// fire 離職
func fire(context *gin.Context) {
	var status string
	var msg string
	var findout bool
	input := context.Param("input")
	amount, err := strconv.Atoi(input)
	for i,j := range allEmployee1{
		 if (j.ID==input){
			((&j).ID) =  "fire"
			allEmployee1[i].ID ="fired"
			fmt.Println("員工: ", j)
			findout = true
			msg ="成功資遣" + input
			status = "Success"
			amount = 0
			}
	}
	if (!findout){
		msg ="操作失敗，輸入錯誤無此員工"
		status = "failed"
		if err == nil {
			amount = 0
		}

	}		

	context.JSON(http.StatusOK, gin.H{
		"amount":  amount,
		"status":  status,
		"message": msg,
	})
}


// hire 新增人員
func hire(context *gin.Context) {
	var status string
	var msg string
	var amount int = 0
	input := context.Param("input")
	//amount, err := strconv.Atoi(input)
	var findout bool
	for _,j := range allEmployee1{
		 if (j.ID==input){
		 findout = true
		 }
	}	 
	if (findout) {
		msg = "操作失敗，username 重複，增加辨識字元"
		status = "failed"
		amount = 0
	}else{
		//TODO 用預設的要改成html 格子內輸入 
		status = "ok"
		allEmployee1 = append(allEmployee1, employee{ID: input, Email: "kk@gmail.com", Mobile: "0982080104", Position: rd1[0]})
		msg = "新增人員" + input 
	}
	
	context.JSON(http.StatusOK, gin.H{
		"amount":  amount,                                                                                                                                                                                                                                                                                                                                                                                         
		"status":  status,
		"message": msg,
	})
}

// edit 人員資料
func get_modify(c *gin.Context) {


		c.HTML(200, "modify.html", nil)
		}

type MyForm struct {
	Username string `form:"ename"`
	UserItem string `form:"eitem"`
	Userdata string `form:"edata"`
}

func modify2(c *gin.Context) {
    var myform MyForm
	var findout bool
	var noteresult string
    if err := c.BindQuery(&myform); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }
    
	noteresult = "can't find username"
	for i,j := range allEmployee1{
		if (j.ID==myform.Username){
			noteresult = "can't get pair item"
			if(myform.UserItem=="Mobile"){
				findout = true
				noteresult = "success"
				allEmployee1[i].Mobile =myform.Userdata
			}
		}
		
	}
	if(findout){
		fmt.Println("員工: ", "dd")
	}
		c.JSON(http.StatusOK, gin.H{
		"Result":noteresult,
		"Name"	:myform.Username,
        "Item"	:myform.UserItem,
        "Data"	:myform.Userdata,
		})
}

func post_modify(context *gin.Context) { 
		//請求的是登入資料，那麼執行登入的邏輯判斷
		var fakeForm MyForm
		//username :=	context.DefaultPostForm("Userdata","somebody")
		//fakeForm.Userdata=string([]byte(username))
		context.Bind(&fakeForm)
		//fakeForm.Userdata = context.DefaultPostForm("Userdata","somebody")
		context.JSON(200, gin.H{"項目": fakeForm.UserItem,"內容": fakeForm.Userdata})
		
        //fmt.Println("item:", context.Form["useritem"])
        //fmt.Println("new-content:", context.Form["userdata"])
		
		}