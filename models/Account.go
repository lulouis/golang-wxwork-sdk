package models

import (
	"golang-wxwork-sdk/utils"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//从Hr系统获取的信息
type EmployeeInfo struct {
	Id       bson.ObjectId `bson:"_id"`
	ID       string        `bson:"ID"`       //工号
	Name     string        `bson:"eName"`    //姓名
	Password string        `bson:"password"` //密码

}

func Login(ID string, password string) (usr *EmployeeInfo, err error) {

	fmt.Println("登陆开始")
	session, err := mgo.Dial(utils.MONGODB)
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	//查询人员表
	var list []*EmployeeInfo
	c := session.DB("HR").C("EmployeeInfo")
	fmt.Println(time.Now())
	err = c.Find(bson.M{"ID": ID, "eWorkState": bson.M{"$ne": "2"}}).All(&list)
	if len(list) > 0 {
		juheURL := "http://192.168.2.14/Marisfrolg.HRC/HRWorkFlow/GetEmpInfo?"
		//初始化参数
		param := url.Values{}
		//配置请求参数,方法内部已处理urlencode问题,中文参数可以直接传参  CodeList   model.ID
		param.Set("ID", ID)
		param.Set("pwd", password)
		//发送请求
		data, err := utils.Get(juheURL, param)
		v := string(data)
		//fmt.Println(v)
		if err != nil {
			fmt.Printf("请求失败,错误信息:\r\n%v", err)
		} else {
			var dem interface{}
			if err := json.Unmarshal([]byte(v), &dem); err == nil {
				fmt.Printf("返回信息数据:\r\n%v", dem)
				usr = list[0]
			} else {
				err = errors.New("密码错误，登陆失败")
			}
		}
	} else {
		err = errors.New("输入的工号不存在，登陆失败")
	}
	if usr == nil {
		err = errors.New("登陆失败")
	}
	return
}
