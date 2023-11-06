package redpacket

import (
	"encoding/xml"
	"fmt"
	"github.com/donetkit/wechat/pay/config"
	"github.com/donetkit/wechat/util"
	"strconv"
)

// RedPacketGateWay 发放红包接口
// https://pay.weixin.qq.com/wiki/doc/api/tools/cash_coupon.php?chapter=13_4&index=3
var RedPacketGateWay = "https://api.mch.weixin.qq.com/mmpaymkttransfers/sendredpack"

// RedPacket struct extends context
type RedPacket struct {
	*config.Config
}

// NewRedPacket return an instance of RedPacket package
func NewRedPacket(cfg *config.Config) *RedPacket {
	return &RedPacket{cfg}
}

// Params 调用参数
type Params struct {
	MchBillno   string // 商户订单号
	SendName    string // 商户名称
	ReOpenID    string
	TotalAmount int
	TotalNum    int
	Wishing     string
	ClientIP    string
	ActName     string
	Remark      string

	RootCa string // ca证书
}

// request 接口请求参数
type request struct {
	NonceStr    string `xml:"nonce_str"`
	Sign        string `xml:"sign"`
	MchID       string `xml:"mch_id"`
	MchBillno   string `xml:"mch_billno"`
	Wxappid     string `xml:"wxappid"`
	SendName    string `xml:"send_name"`
	ReOpenID    string `xml:"re_openid"`
	TotalAmount int    `xml:"total_amount"`
	TotalNum    int    `xml:"total_num"`
	Wishing     string `xml:"wishing"`
	ClientIP    string `xml:"client_ip"`
	ActName     string `xml:"act_name"`
	Remark      string `xml:"remark"`
}

// Response 接口返回
type Response struct {
	ReturnCode  string `xml:"return_code"`
	ReturnMsg   string `xml:"return_msg"`
	ResultCode  string `xml:"result_code,omitempty"`
	ErrCode     string `xml:"err_code,omitempty"`
	ErrCodeDes  string `xml:"err_code_des,omitempty"`
	MchBillno   string `xml:"mch_billno,omitempty"`
	MchID       string `xml:"mch_id,omitempty"`
	Wxappid     string `xml:"wxappid"`
	ReOpenID    string `xml:"re_openid"`
	TotalAmount int    `xml:"total_amount"`
	SendListid  string `xml:"send_listid"`
}

// SendRedPacket 发放红包
func (RedPacket *RedPacket) SendRedPacket(p *Params) (rsp *Response, err error) {
	nonceStr := util.RandomStr(32)
	param := make(map[string]string)

	param["nonce_str"] = nonceStr
	param["mch_id"] = RedPacket.MchID
	param["wxappid"] = RedPacket.AppID
	param["mch_billno"] = p.MchBillno
	param["send_name"] = p.SendName
	param["re_openid"] = p.ReOpenID
	param["total_amount"] = strconv.Itoa(p.TotalAmount)
	param["total_num"] = strconv.Itoa(p.TotalNum)
	param["wishing"] = p.Wishing
	param["client_ip"] = p.ClientIP
	param["act_name"] = p.ActName
	param["remark"] = p.Remark
	//param["scene_id"] = "PRODUCT_2"

	sign, err := util.ParamSign(param, RedPacket.Key)
	if err != nil {
		return
	}

	req := request{
		NonceStr:    nonceStr,
		Sign:        sign,
		MchID:       RedPacket.MchID,
		Wxappid:     RedPacket.AppID,
		MchBillno:   p.MchBillno,
		SendName:    p.SendName,
		ReOpenID:    p.ReOpenID,
		TotalAmount: p.TotalAmount,
		TotalNum:    p.TotalNum,
		Wishing:     p.Wishing,
		ClientIP:    p.ClientIP,
		ActName:     p.ActName,
		Remark:      p.Remark,
	}

	rawRet, err := util.PostXMLWithTLS(RedPacketGateWay, req, p.RootCa, RedPacket.MchID)
	if err != nil {
		return
	}
	err = xml.Unmarshal(rawRet, &rsp)
	if err != nil {
		return
	}
	if rsp.ReturnCode == "SUCCESS" {
		if rsp.ResultCode == "SUCCESS" {
			err = nil
			return
		}
		err = fmt.Errorf("send RedPacket error, errcode=%s,errmsg=%s", rsp.ErrCode, rsp.ErrCodeDes)
		return
	}
	err = fmt.Errorf("[msg : xmlUnmarshalError] [rawReturn : %s] [sign : %s]", string(rawRet), sign)
	return
}
