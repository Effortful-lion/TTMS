package service

import (
	"TTMS/dao/mysql"
	"TTMS/model/do"
	"TTMS/model/dto"
	"TTMS/pkg/common"
	"errors"
	"time"
)

type TicketService struct {
}

func NewTicketService() *TicketService {
	return &TicketService{}
}

func (t *TicketService) CountOnceSeat(plan_id int64) (float64, error) {
	// 通过 plan_id 查 所有票的数量
	total, err := mysql.NewTicketDao().CountTicketByPlanID(plan_id)
	if err!= nil {
		return 0, err
	}
	// 通过 plan_id 和 ticket_status 查 所有已核销票的数量
	part, err := mysql.NewTicketDao().CountUsedTicketByPlanID(plan_id, do.TicketStatusUsed)
	if err!= nil {
		return 0, err
	}
	// 计算 占比
	return common.CalculatePercentageFloat(float64(part), float64(total)), nil
}

func (t *TicketService) CountSeat(plan_id int64) (float64, error) {
	// 通过 play_id 查 所有票的数量
	total, err := mysql.NewTicketDao().CountTicketByPlayID(plan_id)
	if err!= nil {
		return 0, err
	}
	// 通过 play_id 和 ticket_status 查 所有已核销票的数量
	part, err := mysql.NewTicketDao().CountUsedTicketByPlayID(plan_id, do.TicketStatusUsed)
	if err!= nil {
		return 0, err
	}
	// 计算 占比
	return common.CalculatePercentageFloat(float64(part), float64(total)), nil
}

func (t *TicketService) CountOnceTicketPercentageByID(plan_id int64) (float64, error) {
	// 通过 plan_id 查 play_id
	plan, err := mysql.NewPlanDao().SelectPlanByID(plan_id)
	if err!= nil {
		return 0, err
	}
	play_id := plan.PlayID
	// 分别 计算 单场的 总金额 和 总金额
	// 单场：
	onceMoney, err := mysql.NewTicketDao().CountOnceTicketEvery(play_id, plan_id)
	if err!= nil {
		return 0, err
	}
	// 总金额：
	totalMoney, err := mysql.NewTicketDao().CountTicketEvery(play_id)
	if err!= nil {
		return 0, err
	}
	// 计算 占比
	percentage := common.CalculatePercentageFloat(onceMoney, totalMoney)
	return percentage, nil
}

func (t *TicketService) CountOnceTicketByID(plan_id int64) (*dto.TicketCountResp, error) {
	// 通过 plan_id 查 play_id
	plan, err := mysql.NewPlanDao().SelectPlanByID(plan_id)
	if err!= nil {
		return nil, err
	}
	play_id := plan.PlayID
	name, err := mysql.NewPlayDao().SelectNameByID(play_id)
	if err != nil {
		return nil, err
	}
	money, err := mysql.NewTicketDao().CountOnceTicketEvery(play_id, plan_id)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	return &dto.TicketCountResp{
		PlayName:   name,
		TotalMoney: money,
		CountTime:  now,
	}, nil
}


func (t *TicketService) CountTicketByID(play_id int64) (*dto.TicketCountResp, error) {
	name, err := mysql.NewPlayDao().SelectNameByID(play_id)
	if err != nil {
		return nil, err
	}
	money, err := mysql.NewTicketDao().CountTicketEvery(play_id)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	return &dto.TicketCountResp{
		PlayName:   name,
		TotalMoney: money,
		CountTime:  now,
	}, nil
}

func (t *TicketService) CountTicket() (*dto.TicketCountListResp, error) {
	// 统计所有 play 的 play_id
	plays, err := mysql.NewPlayDao().SelectAllPlay()
	if err != nil {
		return nil, err
	}
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
	if err != nil {
		return nil, err
	}
	// 赋值 map
	res := &dto.TicketCountListResp{
		TicketCountList: make([]*dto.TicketCountResp, len(play_ids)),
	}
	now := time.Now()
	for i := range data {
		money := data[i]
		name := names[i]
		res.TicketCountList = append(res.TicketCountList, &dto.TicketCountResp{
			PlayName:   name,
			TotalMoney: money,
			CountTime:  now,
		})
	}
	return res, nil
}

func (t *TicketService) VerifyTicket(ticketID int64) error {
	// 查票
	ticket, err := mysql.NewTicketDao().GetTicketByID(ticketID)
	if err != nil {
		return err
	}
	// 检查票是否已被核销
	if ticket.TicketStatus == 1 {
		return errors.New("票已被核销，不可重复核销")
	}
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
		if err != nil {
			return nil, err
		}
		play, err := mysql.NewPlayDao().SelectPlayByID(plan.PlayID)
		if err != nil {
			return nil, err
		}
		play_name := play.PlayName
		hall, err := mysql.NewHallDao().SelectHall(plan.HallID)
		if err != nil {
			return nil, err
		}
		hall_name := hall.HallName
		seat, err := mysql.NewSeatDao().SelectSeatByID(do.SeatID)
		if err != nil {
			return nil, err
		}
		seat_row := seat.SeatRow
		seat_col := seat.SeatCol
		ticketInfoList = append(ticketInfoList, &dto.TicketInfoResp{
			TicketID:         do.TicketID,
			CustomerID:       do.CustomerID,
			PlanID:           do.PlanID,
			SeatID:           do.SeatID,
			PlayID:           do.PlayID,
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
	if err != nil {
		return err
	}
	// 检查票是否已被核销
	if ticket.TicketStatus == 1 {
		return errors.New("票已被核销, 无法退票")
	}
	// 检查票是否已过期
	if ticket.TicketExpireTime.Before(time.Now()) {
		return errors.New("票已过期, 无法退票")
	}
	// 删除票
	ticketID := req.TicketID
	err = mysql.NewTicketDao().CancelTicket(ticketID)
	if err != nil {
		return errors.New("删除票失败")
	}
	// 修改座位状态
	plan_id := ticket.PlanID
	seat_id := ticket.SeatID
	plan, err := mysql.NewPlanDao().SelectPlanByID(plan_id)
	if err != nil {
		return errors.New("查询计划失败导致无法退票")
	}
	hall_id := plan.HallID
	seat, err := mysql.NewSeatDao().SelectSeatByID(seat_id)
	if err != nil {
		return errors.New("查询座位失败导致无法退票")
	}
	seat_row := seat.SeatRow
	seat_col := seat.SeatCol
	err = mysql.NewSeatDao().CancelSeat(hall_id, seat_row, seat_col)
	if err != nil {
		return errors.New("修改座位状态失败导致无法退票")
	}
	return nil
}

func (t *TicketService) BuyTicket(customerID int64, auth string, req *dto.TicketBuyReq) error {
	plan_id := req.PlanID
	seat_row := req.SeatRow
	seat_col := req.SeatCol
	// 查询 plan ，得到 play_id 和 hall_id，然后查询 play_name 和 hall_name
	planDao := mysql.NewPlanDao()
	plan, err := planDao.SelectPlanByID(plan_id)
	if err != nil {
		return err
	}
	ticket_price := plan.PlanPrice
	plan_start_time := plan.PlanStartTime
	ticket_expire_time := plan_start_time.Add(do.TicketExpiredTime)
	var customer_name string
	switch auth {
	case common.AuthAdmin:
		customer, err := mysql.NewEmployDao().SelectEmployByID(customerID)
		if err != nil {
			return err
		}
		customer_name = customer.EmployName
	case common.AuthUser:
		customerDao := mysql.NewCustomerDao()
		customer, err := customerDao.SelectCustomerByID(customerID)
		if err != nil {
			return err
		}
		customer_name = customer.CustomerName
	}
	// 执行 座位 的增加操作并 返回 座位id
	seatDao := mysql.NewSeatDao()
	seat_id, err := seatDao.SoldSeat(plan.HallID, seat_row, seat_col)
	if err != nil {
		return errors.New("选座失败")
	}
	// 执行票的增加操作
	auth_id := common.GetRoleID(auth)
	ticketDao := mysql.NewTicketDao()
	err = ticketDao.InsertTicket(customerID, plan_id, seat_id, customer_name, ticket_price, ticket_expire_time, plan.PlayID, auth_id)
	if err != nil {
		return err
	}
	return nil
}
