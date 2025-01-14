package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqCreateDepartmentSchool 创建部门请求
// 文档：https://developer.work.weixin.qq.com/document/path/100158#创建部门

type ReqCreateDepartmentSchool struct {
	Name             string             `json:"name"`
	ParentId         int                `json:"parentid"`
	Id               int                `json:"id"`
	Type             int                `json:"type"`
	RegisterYear     int                `json:"register_year"`
	StandardGrade    int                `json:"standard_grade"`
	Order            int                `json:"order"`
	DepartmentAdmins []*DepartmentAdmin `json:"department_admins"`
}

type DepartmentAdmin struct {
	Userid  string `json:"userid"`
	Type    int    `json:"type"`
	Subject string `json:"subject"`
}

var _ bodyer = ReqCreateDepartmentSchool{}

func (x ReqCreateDepartmentSchool) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespCreateDepartmentSchool 创建部门响应
// 文档：https://developer.work.weixin.qq.com/document/path/100158#创建部门

type RespCreateDepartmentSchool struct {
	CommonResp
	Id int `json:"id"`
}

var _ bodyer = RespCreateDepartmentSchool{}

func (x RespCreateDepartmentSchool) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecCreateDepartmentSchool 创建部门
// 文档：https://developer.work.weixin.qq.com/document/path/100158#创建部门
func (c *ApiClient) ExecCreateDepartmentSchool(req ReqCreateDepartmentSchool) (RespCreateDepartmentSchool, error) {
	var resp RespCreateDepartmentSchool
	err := c.executeWXApiPost("/cgi-bin/school/department/create", req, &resp, true)
	if err != nil {
		return RespCreateDepartmentSchool{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespCreateDepartmentSchool{}, bizErr
	}
	return resp, nil
}
