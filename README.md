# tokentask

## 1. 注册
curl -c cookie -H "Content-type: application/json" -X POST -d '{"username":"yekai","password":"123"}' "http://localhost:8080/register"

## 2. 登陆
curl -c cookie -H "Content-type: application/json" -X POST -d '{"username":"yekai","password":"123"}' "http://localhost:8080/login"


## 2. 登陆
curl -c cookie -H "Content-type: application/json" -X POST -d '{"task_name":"kill fuhongxue","task_bonus":100}' "http://localhost:8080/issue"


curl  "http://localhost:8080/tasklist?page=1"