package apis

import (
	"encoding/json"
)

// 自动生成的文件, 生成方式: make api doc=微信文档地址url
// 可自行修改生成的文件,以满足开发需求

// ReqGetApprovalDetailOa 获取审批申请详情请求
// 文档：https://developer.work.weixin.qq.com/document/path/92634#获取审批申请详情
type ReqGetApprovalDetailOa struct {
	// SpNo 审批单编号。，必填
	SpNo int `json:"sp_no"`
}

var _ bodyer = ReqGetApprovalDetailOa{}

func (x ReqGetApprovalDetailOa) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespGetApprovalDetailOa 获取审批申请详情响应
// 文档：https://developer.work.weixin.qq.com/document/path/92634#获取审批申请详情
type RespGetApprovalDetailOa struct {
	CommonResp
	Info struct {
		SpNo       string `json:"sp_no"`
		SpName     string `json:"sp_name"`
		SpStatus   int    `json:"sp_status"`
		TemplateId string `json:"template_id"`
		ApplyTime  int    `json:"apply_time"`
		Applyer    struct {
			Userid  string `json:"userid"`
			Partyid string `json:"partyid"`
		} `json:"applyer"`
		SpRecord []struct {
			SpStatus     int `json:"sp_status"`
			Approverattr int `json:"approverattr"`
			Details      []struct {
				Approver struct {
					Userid string `json:"userid"`
				} `json:"approver"`
				Speech   string        `json:"speech"`
				SpStatus int           `json:"sp_status"`
				Sptime   int           `json:"sptime"`
				MediaId  []interface{} `json:"media_id"`
			} `json:"details"`
		} `json:"sp_record"`
		Notifyer []struct {
			Userid string `json:"userid"`
		} `json:"notifyer"`
		ApplyData struct {
			Contents []struct {
				Control string `json:"control"`
				Id      string `json:"id"`
				Title   []struct {
					Text string `json:"text"`
					Lang string `json:"lang"`
				} `json:"title"`
				Value struct {
					Text        string        `json:"text"`
					Tips        []interface{} `json:"tips"`
					Members     []interface{} `json:"members"`
					Departments []interface{} `json:"departments"`
					Files       []interface{} `json:"files"`
					Children    []interface{} `json:"children"`
					StatField   []interface{} `json:"stat_field"`
				} `json:"value"`
			} `json:"contents"`
		} `json:"apply_data"`
		Comments []struct {
			CommentUserInfo struct {
				Userid string `json:"userid"`
			} `json:"commentUserInfo"`
			Commenttime    int      `json:"commenttime"`
			Commentcontent string   `json:"commentcontent"`
			Commentid      string   `json:"commentid"`
			MediaId        []string `json:"media_id"`
		} `json:"comments"`
		ProcessList struct {
			NodeList []struct {
				NodeType    int `json:"node_type"`
				SpStatus    int `json:"sp_status"`
				ApvRel      int `json:"apv_rel"`
				SubNodeList []struct {
					Userid   string   `json:"userid"`
					Speech   string   `json:"speech"`
					SpYj     int      `json:"sp_yj"`
					Sptime   int      `json:"sptime"`
					MediaIds []string `json:"media_ids"`
				} `json:"sub_node_list"`
			} `json:"node_list"`
		} `json:"process_list"`
	} `json:"info"`
}

var _ bodyer = RespGetApprovalDetailOa{}

func (x RespGetApprovalDetailOa) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExecGetApprovalDetailOa 获取审批申请详情
// 文档：https://developer.work.weixin.qq.com/document/path/92634#获取审批申请详情
func (c *ApiClient) ExecGetApprovalDetailOa(req ReqGetApprovalDetailOa) (RespGetApprovalDetailOa, error) {
	var resp RespGetApprovalDetailOa
	err := c.executeWXApiPost("/cgi-bin/oa/getapprovaldetail", req, &resp, true)
	if err != nil {
		return RespGetApprovalDetailOa{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespGetApprovalDetailOa{}, bizErr
	}
	return resp, nil
}
