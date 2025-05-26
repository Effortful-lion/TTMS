package alipay

import (
	"TTMS/model/dto"
	"TTMS/pkg/common"
	"TTMS/pkg/resp"
	"TTMS/service"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/smartwalle/alipay/v3"
	"github.com/smartwalle/xid"
)

var client *alipay.Client

const (
	kAppId      = "2021000149606440" // 你的支付宝应用id
	kPrivateKey = "MIIEowIBAAKCAQEAiIfp7pyLgANBYU4xCPUPJxfudowHE71gOdIEJduRJl9iHsrxOrarK7ZR7touIv9qxdSNUtDAlKE6TA/IYXtR7OOEjVv2EbHjrUC+0OODc7oVu6A/FkTyqBSLsV/YbndlLa0d6bnLBJqlYd+/238KzAZNEMS4v1VdajiQCgOf43CLGIrzLJnrDRagoYHaY4ASR4YCOcLCI0Jl3f9NJFgodQBjM12rCQLcQqMNdc99g1Nkqo7lKEd2FE1OBXNLGdtVyS6liNwSrCYlddeTk77IkrlG2osZ7wr/tuhHT6uQXgT4Mk2EaA3KdiIMDgfyHt6vqkmOdyF96OM+boc+W+h0NQIDAQABAoIBACZ78ROiKqwtvrsX2DiHIvtJAp2FwOTuh3Yk3ahyKIrhbf3yBGchYERVAulJtKGGz2juDdvfHbmHbLBN7pSTtzcHtrBF0KmSM/P9AQxOZRo0xLV7aOz1CFUaXXnk41PX8bwbRQhBp5knoiBhemn7eGakDAsC+HRuMoOjZ+56sRHhmLxcuODRgvRKfAMVEDYF0Jgh+EGXH0+cBrjqGNFFyWFd6Y1GkjryBZPT/Uwng8zndQgHwDqLfDo0VgW5eWtW3KtPPtt33hILeI6sS73NlK6KfzTz3tU+MNv8D2QWAwznKBSRG46ofLy9RKwfiKYazOrA6uHGREQVUYCvTjg3hq0CgYEAxJD2it3GZkpP2W7B+V1ecu6vM6aE4MaodPHVu51tXAp2LBPrcTpTdjYmEfjlRW0DsjtouDoqJ7PhIHw32WQtp3LLNwNMnBOoXEh2BJ0FkeUc6ezLAsq6zH8kVv5bWg3j+/soXZ5kgY2K/yv9Mfa6Qw7YGN+uCEjdk6/DU1RSTmcCgYEAsc/3x8E4GoOJOvJgtixy0nhbvfAOyr/D1FWKpd2TvziefaTXmrk3LDUSFFx2QWTa8ZesUWaEbVqNhGlS9ZXpb9XQgZgpZ6+EFTpzhtBYlqTAxYl46xlvMvrVJS8iPjolp9ZmPbngVmAdolY2FoLrlAMI1gUw4oDHKLN5QI3qjwMCgYBycqz5obL4r9/aC4DpPhHV/V8JRbUsYr6GiKAXQsv/wR+Y/mghXwJlW5XC9RYY3CZAyhJo4YTV5Qf5MBocexWR32auuz8JyPBH+vCsMvihP7McCIETORKWKpjwBd0Oi/PoLQ4NpEGEpPL/K4+aXbuJe8UahnH/5N0ONlIV3IynlQKBgHpCqvZ4jqD8bHFRktnv+gtponsZArK55GatYqXlkt9aiO+PVogWbi9BVmsVoGIafMi9J0jZGpkuODwhR1+Bcn7pH3LcMywN7N3D0QeTXjZGlh29Bp1wYZtGCSSGIh9zT8z1iFtM4UnC0N18NbiI0H53D65I53H2ZfpAXVDbw7oNAoGBALy+U5zkMib06fn2uTJeMp5VdMWFBMUBi4d39rZ/gCLggDDYsO4Z8PzzR9DcCVBaYXwCmpYmXnIYSRqw5lrIL67Iy5CtA7GDcSwnQPxj6xW3oLOeqL033kBhfCHUXvDsQA4KpJCmNvKN96VD8uay8S5k5sgjmQ0UQJXofOOG5tm5"
	// TODO 设置回调地址域名
	kServerDomain = "http://frp-cup.com:43223" // 服务器地址/内网穿透地址，不定期更换
	// http://1.94.214.117:9999
	// http://frp-cup.com:43223
)

// 支付宝支付初始化
func InitAliPay() {
	var err error

	if client, err = alipay.New(kAppId, kPrivateKey, false); err != nil {
		fmt.Println("初始化支付宝失败", err)
		return
	}

	// 加载证书
	if err = client.LoadAppCertPublicKeyFromFile("./cert/appPublicCert.crt"); err != nil {
		fmt.Println("加载证书发生错误", err)
		return
	}
	if err = client.LoadAliPayRootCertFromFile("./cert/alipayRootCert.crt"); err != nil {
		fmt.Println("加载证书发生错误", err)
		return
	}
	if err = client.LoadAlipayCertPublicKeyFromFile("./cert/alipayPublicCert.crt"); err != nil {
		fmt.Println("加载证书发生错误", err)
		return
	}

	// 加载内容加密密钥
	if err = client.SetEncryptKey("vHFPzZZS5FaTnt9zyxE/Bg=="); err != nil {
		fmt.Println("加载内容加密方式发生错误", err)
		return
	}
}

// 调用支付宝支付接口
// TODO 当退出支付或者支付失败时，删除原订单记录
func Pay(c *gin.Context) {
	var req *dto.TicketPayReq
	if err := c.ShouldBindJSON(&req); err!= nil {
		resp.ResponseError(c, resp.CodeInvalidParams)
		return
	}
	money := req.Money
	ticket_id := req.TicketID
	common.SetContext(c, "ticket_id", ticket_id)
	moneyStr := fmt.Sprintf("%.2f", money)
	// 生成订单id
	var tradeNo = fmt.Sprintf("%d", xid.Next())

	var p = alipay.TradePagePay{}
	p.NotifyURL = kServerDomain + "/alipay/notify"
	p.ReturnURL = kServerDomain + "/alipay/callback"
	p.Subject = "本次支付订单号:" + tradeNo
	p.OutTradeNo = tradeNo                   //订单号，一个订单号只能支付一次  time.Now().String()
	p.TotalAmount = moneyStr                  // 金额
	p.ProductCode = "FAST_INSTANT_TRADE_PAY" //销售产品码，与支付宝签约的产品码名称,目前仅支持FAST_INSTANT_TRADE_PAY

	url, _ := client.TradePagePay(p)
	//c.Redirect(http.StatusTemporaryRedirect, url.String())
	//return url.String() 
	resp.ResponseSuccess(c, url.String())
}

// 支付宝支付成功的回调接口
func Callback(c *gin.Context) {
	fmt.Println("支付宝回调")
	if err := c.Request.ParseForm(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "解析表单数据失败"})
		return
	}

	if err := client.VerifySign(c.Request.Form); err != nil {
		fmt.Println("回调验证签名发生错误", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "回调验证签名发生错误"})
		return
	}

	fmt.Println("回调验证签名通过")

	var outTradeNo = c.Request.Form.Get("out_trade_no")
	var p = alipay.TradeQuery{}
	p.OutTradeNo = outTradeNo

	rsp, err := client.TradeQuery(context.Background(), p)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("验证订单 %s 信息发生错误: %s", outTradeNo, err.Error())})
		return
	}

	if rsp.IsFailure() {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("验证订单 %s 信息发生错误: %s-%s", outTradeNo, rsp.Msg, rsp.SubMsg)})
		return
	}
	fmt.Println("支付成功")
	// TODO 支付成功，插入票数据库
	
	
	//c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("订单 %s 支付成功", outTradeNo)})
	c.Redirect(http.StatusMovedPermanently, "http://baidu.com")

}

func Notify(c *gin.Context) {
	if err := c.Request.ParseForm(); err != nil {
		resp.ResponseErrorWithMsg(c, resp.CodeError, err.Error())
		// 删除票的信息
		ticket_id, _ := common.GetContext(c, "ticket_id")
		ticket_id_int, ok := ticket_id.(int64)
		if !ok {
			resp.ResponseErrorWithMsg(c, resp.CodeError, "ticket_id 类型转换失败")
			return	
		}
		if err = service.NewTicketService().CancelTicket(ticket_id_int); err!= nil {
			resp.ResponseErrorWithMsg(c, resp.CodeError, err.Error())
			return
		}
		return
	}
	var notification, err = client.DecodeNotification(c.Request.Form)
	if err != nil {
		resp.ResponseErrorWithMsg(c, resp.CodeError, fmt.Sprintf("解析异步通知发生错误: %s \n", err.Error()))
		// 删除票的信息
		ticket_id, _ := common.GetContext(c, "ticket_id")
		ticket_id_int, ok := ticket_id.(int64)
		if !ok {
			resp.ResponseErrorWithMsg(c, resp.CodeError, "ticket_id 类型转换失败")
			return	
		}
		if err = service.NewTicketService().CancelTicket(ticket_id_int); err!= nil {
			resp.ResponseErrorWithMsg(c, resp.CodeError, err.Error())
			return
		}
		return
	}

	fmt.Println("解析异步通知成功:", notification.NotifyId)

	var p = alipay.NewPayload("alipay.trade.query")
	p.AddBizField("out_trade_no", notification.OutTradeNo)

	var rsp *alipay.TradeQueryRsp
	if err = client.Request(context.Background(), p, &rsp); err != nil {
		resp.ResponseErrorWithMsg(c, resp.CodeError, fmt.Sprintf("异步通知验证订单 %s 信息发生错误: %s \n", notification.OutTradeNo, err.Error()))
		// 删除票的信息
		ticket_id, _ := common.GetContext(c, "ticket_id")
		ticket_id_int, ok := ticket_id.(int64)
		if !ok {
			resp.ResponseErrorWithMsg(c, resp.CodeError, "ticket_id 类型转换失败")
			return	
		}
		if err = service.NewTicketService().CancelTicket(ticket_id_int); err!= nil {
			resp.ResponseErrorWithMsg(c, resp.CodeError, err.Error())
			return
		}
		return
	}
	if rsp.IsFailure() {
		resp.ResponseErrorWithMsg(c, resp.CodeError, fmt.Sprintf("异步通知验证订单 %s 信息发生错误: %s-%s \n", notification.OutTradeNo, rsp.Msg, rsp.SubMsg))
		// 删除票的信息
		ticket_id, _ := common.GetContext(c, "ticket_id")
		ticket_id_int, ok := ticket_id.(int64)
		if !ok {
			resp.ResponseErrorWithMsg(c, resp.CodeError, "ticket_id 类型转换失败")
			return	
		}
		if err = service.NewTicketService().CancelTicket(ticket_id_int); err!= nil {
			resp.ResponseErrorWithMsg(c, resp.CodeError, err.Error())
			return
		}
		return
	}
	
	fmt.Printf("订单 %s 支付成功 \n", notification.OutTradeNo)
	client.ACKNotification(c.Writer)
}

// // 支付成功的通知接口
// func Notify2(c *gin.Context) {
// 	if err := c.Request.ParseForm(); err != nil {
// 		fmt.Println("解析表单数据失败", err)
// 		return
// 	}

// 	var notification, err = client.DecodeNotification(c.Request.Form)
// 	if err != nil {
// 		fmt.Println("解析异步通知发生错误", err)
// 		return
// 	}

// 	fmt.Println("解析异步通知成功:", notification.NotifyId)

// 	var p = alipay.NewPayload("alipay.trade.query")
// 	p.AddBizField("out_trade_no", notification.OutTradeNo)

// 	var rsp *alipay.TradeQueryRsp
// 	if err = client.Request(context.Background(), p, &rsp); err != nil {
// 		fmt.Printf("异步通知验证订单 %s 信息发生错误: %s \n", notification.OutTradeNo, err.Error())
// 		return
// 	}
// 	if rsp.IsFailure() {
// 		fmt.Printf("异步通知验证订单 %s 信息发生错误: %s-%s \n", notification.OutTradeNo, rsp.Msg, rsp.SubMsg)
// 		return
// 	}

// 	fmt.Printf("订单 %s 支付成功 \n", notification.OutTradeNo)

// 	client.ACKNotification(c.Writer)
// }