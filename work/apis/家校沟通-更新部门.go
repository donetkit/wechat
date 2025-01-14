package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqUpdateDepartmentSchool 更新部门请求
// 文档：https://developer.work.weixin.qq.com/document/path/100159#更新部门
type ReqUpdateDepartmentSchool struct {
	Name             string                   `json:"name"`
	ParentId         int                      `json:"parentid"`
	Id               int                      `json:"id"`
	RegisterYear     int                      `json:"register_year"`
	StandardGrade    int                      `json:"standard_grade"`
	Order            int                      `json:"order"`
	NewId            int                      `json:"new_id,omitempty"`
	DepartmentAdmins []*UpdateDepartmentAdmin `json:"department_admins"`
}

type UpdateDepartmentAdmin struct {
	Op      int    `json:"op"`
	Userid  string `json:"userid"`
	Type    int    `json:"type"`
	Subject string `json:"subject"`
}

var _ bodyer = ReqUpdateDepartmentSchool{}

func (x ReqUpdateDepartmentSchool) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespUpdateDepartmentSchool 更新部门响应
// 文档：https://developer.work.weixin.qq.com/document/path/100159#更新部门

func (x RespUpdateDepartmentSchool) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecUpdateDepartmentSchool 更新部门
// 文档：https://developer.work.weixin.qq.com/document/path/100159#更新部门
func (c *ApiClient) ExecUpdateDepartmentSchool(req ReqUpdateDepartmentSchool) (RespUpdateDepartmentSchool, error) {
	var resp RespUpdateDepartmentSchool
	err := c.executeWXApiPost("/cgi-bin/school/department/update", req, &resp, true)
	if err != nil {
		return RespUpdateDepartmentSchool{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespUpdateDepartmentSchool{}, bizErr
	}
	return resp, nil
}

type RespUpdateDepartmentSchool struct {
	CommonResp
}

var _ bodyer = RespUpdateDepartmentSchool{}
