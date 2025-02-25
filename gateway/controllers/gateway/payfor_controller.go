/***************************************************
 ** @Desc : 代付、下发金额处理逻辑
 ** @Time : 2019/12/5 14:05
 ** @Author : yuebin
 ** @File : payfor_gateway
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/12/5 14:05
 ** @Software: GoLand
****************************************************/
package gateway

import (
	"fmt"
	"gateway/conf"
	"gateway/pay_for"
	"gateway/response"
	"github.com/beego/beego/v2/server/web"
	"strings"
)

type PayForGateway struct {
	web.Controller
}

/*
* 接受下游商户的代付请求
 */
func (c *PayForGateway) PayFor() {
	params := make(map[string]string)
	params["merchantKey"] = strings.TrimSpace(c.GetString("merchantKey"))
	params["realname"] = strings.TrimSpace(c.GetString("realname"))
	params["cardNo"] = strings.TrimSpace(c.GetString("cardNo"))
	//params["bankCode"] = strings.TrimSpace(c.GetString("bankCode"))
	params["accType"] = strings.TrimSpace(c.GetString("accType"))
	//params["province"] = strings.TrimSpace(c.GetString("province"))
	//params["city"] = strings.TrimSpace(c.GetString("city"))
	//params["bankAccountAddress"] = strings.TrimSpace(c.GetString("bankAccountAddress"))
	params["amount"] = strings.TrimSpace(c.GetString("amount"))
	//params["mobileNo"] = strings.TrimSpace(c.GetString("mobileNo"))
	params["merchantOrderId"] = strings.TrimSpace(c.GetString("merchantOrderId"))
	params["sign"] = strings.TrimSpace(c.GetString("sign"))

	payForResponse := new(response.PayForResponse)
	res, msg := checkParams(params)
	if !res {
		payForResponse.ResultCode = "01"
		payForResponse.ResultMsg = msg
	} else {

		payForResponse = pay_for.AutoPayFor(params, conf.SELF_API)
	}

	c.Data["json"] = payForResponse
	_ = c.ServeJSON()

}

/*
* 代付结果查询，
 */
func (c *PayForGateway) PayForQuery() {
	params := make(map[string]string)
	params["merchantKey"] = strings.TrimSpace(c.GetString("merchantKey"))
	params["timestamp"] = strings.TrimSpace(c.GetString("timestamp"))
	params["merchantOrderId"] = strings.TrimSpace(c.GetString("merchantOrderId"))
	params["sign"] = strings.TrimSpace(c.GetString("sign"))

	c.Data["json"] = pay_for.PayForResultQuery(params)
	_ = c.ServeJSON()
}

/*
* 商户查找余额
 */
func (c *PayForGateway) Balance() {
	params := make(map[string]string)
	params["merchantKey"] = strings.TrimSpace(c.GetString("merchantKey"))
	params["timestamp"] = strings.TrimSpace(c.GetString("timestamp"))
	params["sign"] = strings.TrimSpace(c.GetString("sign"))

	balanceResponse := new(response.BalanceResponse)
	res, msg := checkParams(params)
	if !res {
		balanceResponse.ResultCode = "-1"
		balanceResponse.ResultMsg = msg
		c.Data["json"] = balanceResponse
	} else {
		c.Data["json"] = pay_for.BalanceQuery(params)
	}
	_ = c.ServeJSON()
}

func checkParams(params map[string]string) (bool, string) {
	for k, v := range params {
		if v == "" || len(v) == 0 {
			return false, fmt.Sprintf("字段： %s 为必填！", k)
		}
	}
	return true, ""
}
