# tokentask

## 1. 注册
curl -c cookie -H "Content-type: application/json" -X POST -d '{"username":"yekai","password":"123"}' "http://localhost:8080/register"

## 2. 登陆
curl -c cookie -H "Content-type: application/json" -X POST -d '{"username":"yekai","password":"123"}' "http://localhost:8080/login"


## 3. 发布
curl -c cookie -H "Content-type: application/json" -X POST -d '{"task_name":"kill fuhongxue","task_bonus":100}' "http://localhost:8080/issue"


## 4. 列表
curl  "http://localhost:8080/tasklist?page=1"

## 5. 更改状态
curl  -H "Content-type: application/json" -X POST -d '{"task_id":1,"task_status":1, "comment":"ok"}' "http://localhost:8080/update"


curl  -H "Content-type: application/json" -X POST -d '{"task_id":3,"task_status":1, "comment":"干的漂亮"}' "http://118.89.103.58:8080/update"