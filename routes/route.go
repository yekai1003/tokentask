package routes

import (
	"fmt"
	"net/http"

	"github.com/yekai1003/tokentask/dbs"

	"github.com/yekai1003/tokentask/utils"

	"github.com/yekai1003/tokentask/bcos"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

type RespMsg struct {
	Code string      `json:"errno"`
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

func Issue(c *gin.Context) {
	resp := &RespMsg{
		"0",
		"OK",
		nil,
	}
	defer ResponseData(c, resp)
	task := dbs.TaskInfo{}
	c.Bind(&task)
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
	//
	task := dbs.TaskInfo{}
	c.Bind(&task)
	fmt.Println(task)
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
	tasks := dbs.Task_query()
	resp.Data = tasks
	return
}
