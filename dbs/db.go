package dbs

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var dbconn *sql.DB

type TaskInfo struct {
	Task_id   uint   `json:"task_id"`
	Issuer    string `json:"issuer"`
	Task_user string `json:"task_user"`
	Bonus     uint   `json:"bonus"`
	Status    uint   `json:"task_status"`
	TaskName  string `json:"task_name"`
	Comment   string `json:"comment"`
}

type User struct {
	Password string `json:"password"`
	UserName string `json:"username"`
	Address  string `json:"address"`
}

func init() {
	mysql_server := os.Getenv("MYSQL_SERVER")
	if mysql_server == "" {
		log.Panic("Failed to Getenv:MYSQL_SERVER")
	}
	db, err := sql.Open("mysql", "admin:123456@tcp("+mysql_server+")/tokentask?charset=utf8")
	if err != nil {
		log.Panic("failed to open mysql ", err)
	}
	if err = db.Ping(); err != nil {
		log.Panic("failed to ping mysql ", err)
	}
	dbconn = db

}

func (u User) Add() error {
	_, err := dbconn.Exec("insert into t_user(user_name,password,address) values(?,?,?)",
		u.UserName, u.Password, u.Address)
	if err != nil {
		fmt.Println("failed to register ", err, u)
		return err
	}
	return nil
}

func (u *User) Query() (bool, error) {
	rows, err := dbconn.Query("select address from t_user where user_name=? and password=?",
		u.UserName, u.Password)
	if err != nil {
		fmt.Println("failed to select t_user ", err)
		return false, err
	}
	//有结果集
	if rows.Next() {
		err = rows.Scan(&u.Address)
		if err != nil {
			fmt.Println("failed to scan select t_user ", err)
			return false, err
		}
		return true, nil
	}
	return false, err
}

func Task_query() []TaskInfo {
	var tasks []TaskInfo

	rows, err := dbconn.Query("select task_id, issue_user, ifnull(task_user,'NULL'), bonus, status, task_name, ifnull(comment,'NULL') from t_tasks order by task_id")
	if err != nil {
		fmt.Println("failed for query sql:", err)
		return tasks
	}

	for rows.Next() {
		var task TaskInfo
		err := rows.Scan(&task.Task_id, &task.Issuer, &task.Task_user, &task.Bonus, &task.Status, &task.TaskName, &task.Comment)
		if err != nil {
			fmt.Println("failed to Scan result set", err)
			break
		}
		tasks = append(tasks, task)
	}

	return tasks
}

func (t TaskInfo) Add() error {
	_, err := dbconn.Exec("insert into t_tasks(task_id,issue_user,bonus,task_name,status) values(?,?,?,?,0)",
		t.Task_id, t.Issuer, t.Bonus, t.TaskName)
	if err != nil {
		fmt.Println("Failed to add task:", err)
		return err
	}
	return nil
}

func TaskModify(taskID, status uint, remark string) error {
	_, err := dbconn.Exec("update t_tasks set status = ?, task_name=task_name+? where task_id = ?",
		status, remark, taskID)
	if err != nil {
		fmt.Println("Failed to modify task:", err)
		return err
	}
	return nil
}
