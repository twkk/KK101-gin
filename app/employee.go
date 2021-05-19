package main

import (
	"net/http"
	"strconv"
	"encoding/json"
    "fmt"
	"github.com/gin-gonic/gin"
)

var counter = 0

type User struct {
    ID int
    Username string
	Email string
	Mobile string
    Money float64
    Skills []string
    Position map[string]string
    Identification Identification
}

type Identification struct {
    Phone bool
    Email bool
}


func main() {
	
	router := gin.Default()
	router.GET("/hire/:input", hire)
	router.GET("/fire/:input", fire)
	router.GET("/inquire/", employeeget)

	router.Run(":80")
}




// employeeget 取得員工基本資料
func employeeget(context *gin.Context) {
	var Positions = map[string]string{
	"Title": "Manager",
	"Department": "RDDep2",
	}
	user := User {
        ID:     1,
        Username:   "Tony",
		Email:		"tony@stockinst.com",
		Mobile:		"0982080104",
        Skills: []string{"program", "rich", "play"},
        Position: Positions,
        Identification: Identification {
            Phone: true,
            Email: false,
        },
    }
	b, err := json.Marshal(user)
    if err != nil {
        fmt.Println("error:", err)
    }
    fmt.Println(string(b))
	
	context.JSON(http.StatusOK, gin.H{
		"Username": user.Username,
		"Email":  	user.Email,
		"Mobile": 	user.Mobile,
		"Position":	user.Position,
	})
	
}

	
    
	


// hire 新增人員
func hire(context *gin.Context) {
	var status string
	var msg string

	input := context.Param("input")
	amount, err := strconv.Atoi(input)

	if err == nil {
		if amount <= 0 {
			amount = 0
			status = "failed"
			msg = "操作失敗，存款金額需大於0元！"
		} else {
			counter += amount
			status = "ok"
			msg = "已成功存款" + strconv.Itoa(amount) + "元"
		}
	} else {
		amount = 0
		status = "failed"
		msg = "操作失敗，輸入有誤！"
	}
	context.JSON(http.StatusOK, gin.H{
		"amount":  amount,
		"status":  status,
		"message": msg,
	})
}

// fire 離職
func fire(context *gin.Context) {
	var status string
	var msg string

	input := context.Param("input")
	amount, err := strconv.Atoi(input)

	if err == nil {
		if amount <= 0 {
			amount = 0
			status = "failed"
			msg = "操作失敗，提款金額需大於0元！"
		} else {
			if counter-amount < 0 {
				amount = 0
				status = "failed"
				msg = "操作失敗，餘額不足！"
			} else {
				counter -= amount
				status = "ok"
				msg = "成功提款" + strconv.Itoa(amount) + "元"
			}
		}
	} else {
		amount = 0
		status = "failed"
		msg = "操作失敗，輸入有誤！"
	}
	context.JSON(http.StatusOK, gin.H{
		"amount":  amount,
		"status":  status,
		"message": msg,
	})
}