package models

import (
	"encoding/json"
	"fmt"

	"golang-wxwork-sdk/utils"
	"net/url"
)

// 获取部门列表
// 请求方式：GET（HTTPS）
// 请求地址：https://qyapi.weixin.qq.com/cgi-bin/department/list?access_token=ACCESS_TOKEN&id=ID

func GetDepartmentList() (result *DepartmentGetResult, err error) {
	reqUrl := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/department/list")
	token := GetAccessToken()
	params := url.Values{}
	params.Set("access_token", token.Access_token)
	res, err := utils.HttpGet(reqUrl, params)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &result)
	if err != nil {
		return
	}

	return

}

// 创建部门
// 请求方式：POST（HTTPS）
// 请求地址：https://qyapi.weixin.qq.com/cgi-bin/department/create?access_token=ACCESS_TOKEN

func DepartmentAdd(department Department2) (result *DepartmentAddResult, err error) {
	reqUrl := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/department/create")
	token := GetAccessToken()
	params := url.Values{}
	params.Set("access_token", token.Access_token)
	jdata, err := json.Marshal(department)
	res, err := utils.HttpPostBody(reqUrl, params, string(jdata))
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &result)
	if err != nil {
		return
	}

	return

}

// 更新部门
// 请求方式：POST（HTTPS）
// 请求地址：https://qyapi.weixin.qq.com/cgi-bin/department/update?access_token=ACCESS_TOKEN

func DepartmentUpdate(department Department) (result *PublicResult, err error) {
	reqUrl := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/department/update")
	token := GetAccessToken()
	params := url.Values{}
	params.Set("access_token", token.Access_token)
	params.Set("id", string(department.Id))
	jdata, err := json.Marshal(department)
	res, err := utils.HttpPostBody(reqUrl, params, string(jdata))
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &result)
	if err != nil {
		return
	}

	return

}

// 删除部门
// 请求方式：GET（HTTPS）
// 请求地址：https://qyapi.weixin.qq.com/cgi-bin/department/delete?access_token=ACCESS_TOKEN&id=ID

func DepartmentDelete(id int) (result *PublicResult, err error) {
	reqUrl := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/department/delete")
	token := GetAccessToken()
	params := url.Values{}
	params.Set("access_token", token.Access_token)
	params.Set("id", fmt.Sprintf("%d", id))
	fmt.Println(id)
	res, err := utils.HttpGet(reqUrl, params)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &result)
	if err != nil {
		return
	}

	return

}
