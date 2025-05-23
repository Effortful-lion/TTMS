package service

import (
	"TTMS/dao/mysql"
	"TTMS/model/do"
	"TTMS/model/dto"
	"time"
)

type TicketService struct {
}

func NewTicketService() *TicketService {
	return &TicketService{}
}

func (t *TicketService) CountTicketByID(play_id int64) (*dto.TicketCountResp, error) {
	name, err := mysql.NewPlayDao().SelectNameByID(play_id)
	if err != nil {
		return nil, err
	} 
	money, err := mysql.NewTicketDao().CountTicketEvery(play_id)
	if err != nil{
		return nil, err
	}
	now := time.Now()
	return &dto.TicketCountResp{
		PlayName: name,
		TotalMoney: money,
		CountTime: now,
	}, nil
}

func (t *TicketService) CountTicket() (*dto.TicketCountListResp, error){
	// 统计所有 play 的 play_id
	plays, err := mysql.NewPlayDao().SelectAllPlay()
	if err != nil { return nil, err}
	play_ids := make([]int64, len(plays))
	for i := range plays {
		play_ids[i] = plays[i].PlayID
	}
	// 通过 play_id 的数组，返回一个对应的名字数组
	names, err := mysql.NewPlayDao().SelectNamesByIDs(play_ids)
	if err != nil {
		return nil, err
	}
	// 通过 plan_id 的数组，返回一个对应的金额数组
	data, err := mysql.NewTicketDao().CountTicket(play_ids)
	if err != nil { return nil, err}
	// 赋值 map
	res := &dto.TicketCountListResp{
		TicketCountList: make([]*dto.TicketCountResp, len(play_ids)),
	}
	now := time.Now()
	for i := range data {
		money := data[i]
		name := names[i]
		res.TicketCountList = append(res.TicketCountList, &dto.TicketCountResp{
			PlayName: name,
			TotalMoney: money,
			CountTime: now,
		})
	}
	return res, nil
}

func (t *TicketService) VerifyTicket(ticketID int64) error {
	return mysql.NewTicketDao().VerifyTicket(ticketID)
}

func (t *TicketService) GetTicketList(customerID int64) (*dto.TicketInfoListResp, error) {
	dolist, err := mysql.NewTicketDao().GetTicketList(customerID)
	if err != nil {
		return nil, err
	}
	// 转换为 TicketInfoResp
	var ticketInfoList []*dto.TicketInfoResp
	for _, do := range dolist {
		plan_id := do.PlanID
		plan, err := mysql.NewPlanDao().SelectPlanByID(plan_id)
		if err != nil {return nil, err}
		play, err := mysql.NewPlayDao().SelectPlayByID(plan.PlayID)
		if err != nil {return nil, err}
		play_name := play.PlayName
		hall, err := mysql.NewHallDao().SelectHall(plan.HallID)
		if err != nil {return nil, err}
		hall_name := hall.HallName
		seat, err := mysql.NewSeatDao().SelectSeatByID(do.SeatID)
		if err != nil {return nil, err}
		seat_row := seat.SeatRow
		seat_col := seat.SeatCol
		ticketInfoList = append(ticketInfoList, &dto.TicketInfoResp{
			TicketID: 	      do.TicketID,	
			CustomerID:       do.CustomerID,
			PlanID:           do.PlanID,
			SeatID:           do.SeatID,
			PlayID:  		  do.PlayID,		
			CustomerName:     do.CustomerName,
			HallName:         hall_name,
			PlayName:         play_name,
			SeatRow:          seat_row,
			SeatCol:          seat_col,
			TicketPrice:      do.TicketPrice,
			TicketStatus:     do.TicketStatus,
			TicketExpireTime: do.TicketExpireTime,
		})
	}
	// 转换为 TicketInfoListResp
	ticketInfoListResp := &dto.TicketInfoListResp{
		Tickets: make([]*dto.TicketInfoResp, len(ticketInfoList)),
	}
	// 最高效（分配好长度）
	copy(ticketInfoListResp.Tickets, ticketInfoList)
	// 次高效
	// for i := range ticketInfoList {
	// 	ticketInfoListResp.Tickets[i] = ticketInfoList[i]
	// }
	// 最低效的是依次赋值后再追加 append
	return ticketInfoListResp, nil
}

func (t *TicketService) CancelTicket(req *dto.TicketCancelReq) error {
	// 查票
	ticket, err := mysql.NewTicketDao().GetTicketByID(req.TicketID)
	if err != nil{
		return err
	}
	// 删除票
	ticketID := req.TicketID
	err = mysql.NewTicketDao().CancelTicket(ticketID)
	if err != nil {return err }
	// 修改座位状态
	plan_id := ticket.PlanID
	seat_id := ticket.SeatID
	plan, err := mysql.NewPlanDao().SelectPlanByID(plan_id)
	if err != nil {return err }
	hall_id := plan.HallID
	seat, err := mysql.NewSeatDao().SelectSeatByID(seat_id)
	if err != nil {return err }
	seat_row := seat.SeatRow
	seat_col := seat.SeatCol
	err = mysql.NewSeatDao().CancelSeat(hall_id, seat_row, seat_col)
	if err != nil {return err }
	return nil
}

func (t *TicketService) BuyTicket(customerID int64, req *dto.TicketBuyReq) (error) {
	plan_id := req.PlanID
	seat_row := req.SeatRow
	seat_col := req.SeatCol
	// 查询 plan ，得到 play_id 和 hall_id，然后查询 play_name 和 hall_name
	planDao := mysql.NewPlanDao()
	plan, err := planDao.SelectPlanByID(plan_id)
	if err != nil {return err }
	ticket_price := plan.PlanPrice
	plan_start_time := plan.PlanStartTime
	ticket_expire_time := plan_start_time.Add(do.TicketExpiredTime)
	// 根据 customerID 获得 customer_name
	customerDao := mysql.NewCustomerDao()
	customer, err := customerDao.SelectCustomerByID(customerID)
	if err != nil {return err }
	customer_name := customer.CustomerName
	// 执行 座位 的增加操作并 返回 座位id
	seatDao := mysql.NewSeatDao()
	seat_id, err := seatDao.SoldSeat(plan.HallID, seat_row, seat_col)
	if err != nil {return err }
	// 执行票的增加操作
	ticketDao := mysql.NewTicketDao()
	err = ticketDao.InsertTicket(customerID, plan_id, seat_id, customer_name, ticket_price, ticket_expire_time, plan.PlayID)
	if err != nil {return err }
	return nil
}