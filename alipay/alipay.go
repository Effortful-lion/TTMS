package alipay

import (
	"TTMS/dao/mysql"
	"TTMS/dao/redis"
	"TTMS/model/dto"
	"TTMS/pkg/common"
	"TTMS/pkg/resp"
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/smartwalle/alipay/v3"
	"github.com/smartwalle/xid"
)

var client *alipay.Client

const (
	kAppId      = "2021000149606440" // 你的支付宝应用id
	kPrivateKey = "MIIEowIBAAKCAQEAiIfp7pyLgANBYU4xCPUPJxfudowHE71gOdIEJduRJl9iHsrxOrarK7ZR7touIv9qxdSNUtDAlKE6TA/IYXtR7OOEjVv2EbHjrUC+0OODc7oVu6A/FkTyqBSLsV/YbndlLa0d6bnLBJqlYd+/238KzAZNEMS4v1VdajiQCgOf43CLGIrzLJnrDRagoYHaY4ASR4YCOcLCI0Jl3f9NJFgodQBjM12rCQLcQqMNdc99g1Nkqo7lKEd2FE1OBXNLGdtVyS6liNwSrCYlddeTk77IkrlG2osZ7wr/tuhHT6uQXgT4Mk2EaA3KdiIMDgfyHt6vqkmOdyF96OM+boc+W+h0NQIDAQABAoIBACZ78ROiKqwtvrsX2DiHIvtJAp2FwOTuh3Yk3ahyKIrhbf3yBGchYERVAulJtKGGz2juDdvfHbmHbLBN7pSTtzcHtrBF0KmSM/P9AQxOZRo0xLV7aOz1CFUaXXnk41PX8bwbRQhBp5knoiBhemn7eGakDAsC+HRuMoOjZ+56sRHhmLxcuODRgvRKfAMVEDYF0Jgh+EGXH0+cBrjqGNFFyWFd6Y1GkjryBZPT/Uwng8zndQgHwDqLfDo0VgW5eWtW3KtPPtt33hILeI6sS73NlK6KfzTz3tU+MNv8D2QWAwznKBSRG46ofLy9RKwfiKYazOrA6uHGREQVUYCvTjg3hq0CgYEAxJD2it3GZkpP2W7B+V1ecu6vM6aE4MaodPHVu51tXAp2LBPrcTpTdjYmEfjlRW0DsjtouDoqJ7PhIHw32WQtp3LLNwNMnBOoXEh2BJ0FkeUc6ezLAsq6zH8kVv5bWg3j+/soXZ5kgY2K/yv9Mfa6Qw7YGN+uCEjdk6/DU1RSTmcCgYEAsc/3x8E4GoOJOvJgtixy0nhbvfAOyr/D1FWKpd2TvziefaTXmrk3LDUSFFx2QWTa8ZesUWaEbVqNhGlS9ZXpb9XQgZgpZ6+EFTpzhtBYlqTAxYl46xlvMvrVJS8iPjolp9ZmPbngVmAdolY2FoLrlAMI1gUw4oDHKLN5QI3qjwMCgYBycqz5obL4r9/aC4DpPhHV/V8JRbUsYr6GiKAXQsv/wR+Y/mghXwJlW5XC9RYY3CZAyhJo4YTV5Qf5MBocexWR32auuz8JyPBH+vCsMvihP7McCIETORKWKpjwBd0Oi/PoLQ4NpEGEpPL/K4+aXbuJe8UahnH/5N0ONlIV3IynlQKBgHpCqvZ4jqD8bHFRktnv+gtponsZArK55GatYqXlkt9aiO+PVogWbi9BVmsVoGIafMi9J0jZGpkuODwhR1+Bcn7pH3LcMywN7N3D0QeTXjZGlh29Bp1wYZtGCSSGIh9zT8z1iFtM4UnC0N18NbiI0H53D65I53H2ZfpAXVDbw7oNAoGBALy+U5zkMib06fn2uTJeMp5VdMWFBMUBi4d39rZ/gCLggDDYsO4Z8PzzR9DcCVBaYXwCmpYmXnIYSRqw5lrIL67Iy5CtA7GDcSwnQPxj6xW3oLOeqL033kBhfCHUXvDsQA4KpJCmNvKN96VD8uay8S5k5sgjmQ0UQJXofOOG5tm5"
	// TODO 设置回调地址域名
	kServerDomain = "http://1.94.214.117:9999" // 服务器地址/内网穿透地址，不定期更换
	// http://1.94.214.117:9999
	// http://frp-cup.com:43223
	TicketExpiredTime = 5 * time.Minute
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
	var req dto.TicketPayReq
	if err := c.ShouldBind(&req); err != nil {
		resp.ResponseErrorWithMsg(c, resp.CodeInvalidParams, err.Error())
		return
	}
	money := req.Money
	customer_id := req.CustomerID
	plan_id := req.PlanID
	seat_id := req.SeatID
	
	moneyStr := fmt.Sprintf("%.2f", money)
	// 生成订单id
	var tradeNo = fmt.Sprintf("%d", xid.Next())
	// 生成 redis 的 key
	key := tradeNo

	// redis 存储 查订单的必要信息
	err := redis.InsertTicketOther(key, customer_id, plan_id, seat_id, TicketExpiredTime)
	if err != nil {
		resp.ResponseErrorWithMsg(c, resp.CodeError,err.Error())
		return
	}

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

	
	//c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("订单 %s 支付成功", outTradeNo)})
	c.Redirect(http.StatusMovedPermanently, "http://1.94.214.117:3000/manage")
}



// 支付成功的通知接口
func Notify(c *gin.Context) {
	if err := c.Request.ParseForm(); err != nil {
		fmt.Println("解析表单数据失败", err)
		return
	}

	var notification, err = client.DecodeNotification(c.Request.Form)
	if err != nil {
		fmt.Println("解析异步通知发生错误", err)
		return
	}

	fmt.Println("解析异步通知成功:", notification.NotifyId)

	var p = alipay.NewPayload("alipay.trade.query")
	p.AddBizField("out_trade_no", notification.OutTradeNo)

	var rsp *alipay.TradeQueryRsp
	if err = client.Request(context.Background(), p, &rsp); err != nil {
		fmt.Printf("异步通知验证订单 %s 信息发生错误: %s \n", notification.OutTradeNo, err.Error())
		return
	}
	if rsp.IsFailure() {
		fmt.Printf("异步通知验证订单 %s 信息发生错误: %s-%s \n", notification.OutTradeNo, rsp.Msg, rsp.SubMsg)
		return
	}

	fmt.Printf("订单 %s 支付成功 \n", notification.OutTradeNo)

	// TODO 保存票信息
	// 验证订单信息成功，从订单中拿订单号
	pay_id := notification.OutTradeNo // 支付宝订单号


	fmt.Println("支付成功")
	// TODO 支付成功，插入票数据库
	// 从 pay_id 通过 redis 获取 customer_id plan_id seat_id
	mapStr, err := redis.GetTicketOther(pay_id)
	if err != nil { 
		resp.ResponseErrorWithMsg(c, resp.CodeError, "获取订单信息失败")
		return
	}
	customer_id, _ := strconv.ParseInt(mapStr["customerID"], 10, 64)
	plan_id, _ := strconv.ParseInt(mapStr["planID"], 10, 64)
	seat_id, _ := strconv.ParseInt(mapStr["seatID"], 10, 64)
	// 然后通过 customer_id plan_id seat_id 从 redis 中获取 ticket 对象
	mapStr2, err := redis.GetTicket(customer_id, plan_id, seat_id)
	if err != nil {
		resp.ResponseErrorWithMsg(c, resp.CodeError, err.Error())
		return
	}
	fmt.Println(mapStr2)
	// 从mapStr2中获取playID
	playIDStr, ok := mapStr2["playID"]
	if !ok {
		resp.ResponseErrorWithMsg(c, resp.CodeError, "缺少playID")
		return
	}
	play_id, err := strconv.ParseInt(playIDStr, 10, 64)
	if err != nil {
		resp.ResponseErrorWithMsg(c, resp.CodeError, "playID格式错误")
		return
	}
	// 从mapStr2中获取authID（注意：原代码中mapStr2["authID"]是字符串，需转换为int64再转int8）
	authIDStr, ok := mapStr2["authID"]
	if !ok {
		resp.ResponseErrorWithMsg(c, resp.CodeError, "缺少authID")
		return
	}
	authIDInt64, err := strconv.ParseInt(authIDStr, 10, 64)
	if err != nil {
		resp.ResponseErrorWithMsg(c, resp.CodeError, "authID格式错误")
		return
	}
	auth_id := int8(authIDInt64) // 转换为int8
	customerName := mapStr2["customerName"]
	priceStr := mapStr2["price"]
	priceFloat64, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		resp.ResponseErrorWithMsg(c, resp.CodeError, "price格式错误")
		return
	}
	expire := mapStr2["expire"]

	// 最后将 ticket 插入数据库中
	_, err = mysql.NewTicketDao().InsertTicket(customer_id, plan_id, seat_id, customerName, priceFloat64, common.ParseStringTime(expire), play_id, int8(auth_id))
	if err != nil {
		resp.ResponseErrorWithMsg(c, resp.CodeError, "订单插入失败")
		return
	}

	// 修改座位状态
	// 根据 plan_id 获取 hall_id
	plan, err := mysql.NewPlanDao().SelectPlanByID(plan_id)
	if err != nil {
		resp.ResponseErrorWithMsg(c, resp.CodeError, "查询计划失败导致无法锁定座位")
		return
	}
	seat, err := mysql.NewSeatDao().SelectSeatByID(seat_id)
	if err != nil {
		resp.ResponseErrorWithMsg(c, resp.CodeError, "查询座位失败导致无法锁定座位")
		return
	}
	hall_id := plan.HallID
	_, _ = mysql.NewSeatDao().SoldSeat(hall_id, seat.SeatRow, seat.SeatCol)

	client.ACKNotification(c.Writer)
}

// func Notify(c *gin.Context) {
// 	if err := c.Request.ParseForm(); err != nil {
// 		resp.ResponseErrorWithMsg(c, resp.CodeError, err.Error())
// 		// 删除票的信息
// 		ticket_id, _ := common.GetContext(c, "ticket_id")
// 		ticket_id_int, ok := ticket_id.(int64)
// 		if !ok {
// 			resp.ResponseErrorWithMsg(c, resp.CodeError, "ticket_id 类型转换失败")
// 			return	
// 		}
// 		if err = service.NewTicketService().CancelTicket(ticket_id_int); err!= nil {
// 			resp.ResponseErrorWithMsg(c, resp.CodeError, err.Error())
// 			return
// 		}
// 		return
// 	}
// 	var notification, err = client.DecodeNotification(c.Request.Form)
// 	if err != nil {
// 		resp.ResponseErrorWithMsg(c, resp.CodeError, fmt.Sprintf("解析异步通知发生错误: %s \n", err.Error()))
// 		// 删除票的信息
// 		ticket_id, _ := common.GetContext(c, "ticket_id")
// 		ticket_id_int, ok := ticket_id.(int64)
// 		if !ok {
// 			resp.ResponseErrorWithMsg(c, resp.CodeError, "ticket_id 类型转换失败")
// 			return	
// 		}
// 		if err = service.NewTicketService().CancelTicket(ticket_id_int); err!= nil {
// 			resp.ResponseErrorWithMsg(c, resp.CodeError, err.Error())
// 			return
// 		}
// 		return
// 	}

// 	fmt.Println("解析异步通知成功:", notification.NotifyId)

// 	var p = alipay.NewPayload("alipay.trade.query")
// 	p.AddBizField("out_trade_no", notification.OutTradeNo)

// 	var rsp *alipay.TradeQueryRsp
// 	if err = client.Request(context.Background(), p, &rsp); err != nil {
// 		resp.ResponseErrorWithMsg(c, resp.CodeError, fmt.Sprintf("异步通知验证订单 %s 信息发生错误: %s \n", notification.OutTradeNo, err.Error()))
// 		// 删除票的信息
// 		ticket_id, _ := common.GetContext(c, "ticket_id")
// 		ticket_id_int, ok := ticket_id.(int64)
// 		if !ok {
// 			resp.ResponseErrorWithMsg(c, resp.CodeError, "ticket_id 类型转换失败")
// 			return	
// 		}
// 		if err = service.NewTicketService().CancelTicket(ticket_id_int); err!= nil {
// 			resp.ResponseErrorWithMsg(c, resp.CodeError, err.Error())
// 			return
// 		}
// 		return
// 	}
// 	if rsp.IsFailure() {
// 		resp.ResponseErrorWithMsg(c, resp.CodeError, fmt.Sprintf("异步通知验证订单 %s 信息发生错误: %s-%s \n", notification.OutTradeNo, rsp.Msg, rsp.SubMsg))
// 		// 删除票的信息
// 		ticket_id, _ := common.GetContext(c, "ticket_id")
// 		ticket_id_int, ok := ticket_id.(int64)
// 		if !ok {
// 			resp.ResponseErrorWithMsg(c, resp.CodeError, "ticket_id 类型转换失败")
// 			return	
// 		}
// 		if err = service.NewTicketService().CancelTicket(ticket_id_int); err!= nil {
// 			resp.ResponseErrorWithMsg(c, resp.CodeError, err.Error())
// 			return
// 		}
// 		return
// 	}
	
// 	fmt.Printf("订单 %s 支付成功 \n", notification.OutTradeNo)

// 	// TODO 保存票信息

// 	client.ACKNotification(c.Writer)
// }