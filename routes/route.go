package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/yekai1003/tokentask/dbs"

	"github.com/yekai1003/tokentask/utils"

	"github.com/yekai1003/tokentask/bcos"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

type RespMsg struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

//resp数据响应
func ResponseData(c *gin.Context, resp *RespMsg) {
	resp.Msg = utils.RecodeText(resp.Code)
	c.JSON(http.StatusOK, resp)
}

func Login(c *gin.Context) {
	resp := &RespMsg{
		"0",
		"OK",
		nil,
	}
	defer ResponseData(c, resp)
	user := dbs.User{}
	c.Bind(&user)
	fmt.Println("login:", user)
	ok, err := user.Query()
	if err != nil || !ok {
		resp.Code = utils.RECODE_DBERR
		return
	}
	resp.Data = "ok-ok"
	session := sessions.Default(c)
	session.Set("address", user.Address)
	session.Set("username", user.UserName)
	session.Set("password", user.Password)
	session.Save()
	return
}

func Register(c *gin.Context) {
	resp := &RespMsg{
		"0",
		"OK",
		nil,
	}
	defer ResponseData(c, resp)
	user := dbs.User{}
	c.Bind(&user)
	fmt.Println(user)
	address, err := bcos.NewAccount(user.Password)
	if err != nil {
		resp.Code = utils.RECODE_ETHERR
		return
	}
	user.Address = address
	err = user.Add()
	if err != nil {
		resp.Code = utils.RECODE_DBERR
		return
	}

	return
}

var token_begin uint = 2000

func Issue(c *gin.Context) {
	resp := &RespMsg{
		"0",
		"OK",
		nil,
	}
	defer ResponseData(c, resp)
	task := dbs.TaskInfo{}
	c.Bind(&task)
	token_begin++
	task.Task_id = token_begin
	session := sessions.Default(c)
	username := session.Get("username")
	task.Issuer = username.(string)

	err := task.Add()
	if err != nil {
		resp.Code = utils.RECODE_DBERR
		return
	}
	//缺少区块链调用部分

	return
}

func Modify(c *gin.Context) {
	resp := &RespMsg{
		"0",
		"OK",
		nil,
	}
	defer ResponseData(c, resp)

	// data, _ := ioutil.ReadAll(c.Request.Body)
	// fmt.Printf("ctx.Request.body: %v", string(data))
	taskmap := make(map[string]interface{})

	task := dbs.TaskInfo{}
	c.Bind(taskmap)
	fmt.Println(taskmap)
	err := dbs.TaskModify(task.Task_id, task.Status, task.Comment)
	if err != nil {
		fmt.Println("Modify err:", err)
		resp.Code = utils.RECODE_DBERR
		return
	}
	return
}

func TaskList(c *gin.Context) {
	resp := &RespMsg{
		"0",
		"OK",
		nil,
	}
	defer ResponseData(c, resp)
	pagestr := c.Query("page")
	page, _ := strconv.Atoi(pagestr)
	tasks := dbs.Task_query()
	//begin, end := 0, 10
	begin := (page - 1) * 10
	end := page * 10
	if end > len(tasks) {
		end = len(tasks)
	}
	fmt.Println("begin = ", begin, ", end = %d", end, ",page = ", page)
	fmt.Println(tasks)
	ts := struct {
		Total int         `json:"total"`
		Data  interface{} `json:"data"`
	}{
		Total: len(tasks),
		Data:  tasks[begin:end],
	}
	resp.Data = ts
	return
}
