package models

// 参数	必须	说明
// access_token	是	调用接口凭证
// name	是	部门名称。长度限制为1~32个字符，字符不能包括\:?”<>｜
// parentid	是	父部门id，32位整型
// order	否	在父部门中的次序值。order值大的排序靠前。有效的值范围是[0, 2^32)
// id	否	部门id，32位整型，指定时必须大于1。若不填该参数，将自动生成id

type Department struct {
	Name     string `json:"name"`
	Parentid int    `json:"parentid"`
	Order    int    `json:"order"`
	Id       int    `json:"id"`
}

type Department2 struct {
	Name     string `json:"name"`
	Parentid int    `json:"parentid"`
	Order    int    `json:"order"`
}

type DepartmentAddResult struct {
	PublicResult
	Id int
}

type DepartmentGetResult struct {
	PublicResult
	Department []*Department
}
