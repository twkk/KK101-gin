# KK101_Sample_Golang_Gin

KK101_Sample_Golang_Gin is my first sample code for learning Golang with Gin.
for https://github.com/heyu-exercise/Site-Reliability-Engineering

- exercise topic: Employee Data Management
- [implement exercise below]
  - [Write an simple RESTful API Golang gin]
  - [Have CRUD function-查詢,新增,刪除,修改]
  - [Build your application to container images]

# Contents
- [file]
- main function file app/line_json.go
- Dockerfile
- views/form.html

```bash
run_env 	 base: win10, go1.16.4
docker image base: golang:1.15.3-alpine

- import
github.com/gin-gonic/gin
```

## Usage

```bash
輸出員工json 格式化 localhost:8080//employee/1
輸出員工json 緊緻版 localhost:8080//employee
//查詢,新增,刪除,修改
curl -X GET "localhost:8080/modify2?ename=203&eitem=Mobile&edata=0921357888"//修改
curl -X GET "localhost:8080/modify2?ename=201&eitem=yunfen&edata=xyz"		//修改無對應
curl -X GET "localhost:8080/fire/201"										//刪除存在的
curl -X GET "localhost:8080/fire/test"										//刪除不存在
curl -X GET "localhost:8080/hire/new_member"								//新增
curl -X GET "localhost:8080//employee/1"									//查詢

```

```bash
Docker images
 bankworld	//RESTful sample1 bank		,bank1.go  		<< old version
 employee	//RESTful sample2 employee	,line_json.go	<< new version

git clone XXXX

cd XXXX
docker build -t gin-123 .
docker run --rm -p 8080:8080 -d gin-123
```

## TODO
-1.表單式新增未完成
-2.gin, net/http 混用?
-3.修正 Golang 需要大寫,json小寫 "'json/...'"
-4.刪除用比較簡便的修改名稱為離職替代,應該在輸出的地方加上篩選
		進入 db/csv 的時候只存有正常名稱的,
		要研究一下 delete an element from a Slice in Golang 
-5.manage your service with docker-compose
-6.Authentication and validation 
-7.leaning note
	
## License
[MIT](https://choosealicense.com/licenses/mit/)