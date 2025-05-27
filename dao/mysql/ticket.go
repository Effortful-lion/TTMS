package mysql

import (
	"TTMS/model/do"
	"errors"
	"fmt"
	"time"
)

type TicketDao struct {
}

func NewTicketDao() *TicketDao {
	return &TicketDao{}
}

func (td *TicketDao) DeleteTicketUsed()error {
	err := DB.Where("ticket_status = ?", 2).Delete(&do.Ticket{}).Error
	if err != nil {
		return fmt.Errorf("删除已售票失败: %v", err)
	}
	return nil
}

// 根据 plan_id 查票数量
func (td *TicketDao) CountTicketByPlanID(plan_id int64) (int64, error) {
	var res int64
	err := DB.Model(&do.Ticket{}).Where("plan_id =?", plan_id).Count(&res).Error
	if err != nil {
		return 0, errors.New("根据 plan_id 查票数量有误")
	}
	return res, nil
}

// 根据 play_id 查票数量
func (td *TicketDao) CountTicketByPlayID(play_id int64) (int64, error) {
	var res int64
	err := DB.Model(&do.Ticket{}).Where("play_id =?", play_id).Count(&res).Error
	if err != nil {
		return 0, errors.New("根据 play_id 查票数量有误")
	}
	return res, nil
}

// 根据 plan_id 和 ticket_status 查票
func (td *TicketDao) CountUsedTicketByPlanID(plan_id int64, ticket_status int8) (int64, error) {
	var res int64
	err := DB.Model(&do.Ticket{}).Where("plan_id = ? and ticket_status = ?", plan_id, ticket_status).Count(&res).Error
	if err != nil {
		return 0, errors.New("根据 plan_id 和 ticket_status 查票有误")
	}
	return res, nil
}

// 根据 play_id 和 ticket_status 查票
func (td *TicketDao) CountUsedTicketByPlayID(play_id int64, ticket_status int8) (int64, error) {
	var res int64
	err := DB.Model(&do.Ticket{}).Where("play_id = ? and ticket_status = ?", play_id, ticket_status).Count(&res).Error
	if err != nil {
		return 0, errors.New("根据 plan_id 和 ticket_status 查票有误")
	}
	return res, nil
}

// 票房统计：统计总金额(元)
func (td *TicketDao) CountTicket(play_ids []int64) ([]float64, error) {
	var err error
	var res = make([]float64, len(play_ids))
	for i := range play_ids {
		play_id := play_ids[i]
		// 根据 play_id 查询票总金额
		res[i], err = td.CountTicketEvery(play_id)
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (td *TicketDao) CountOnceTicketEvery(play_id, plan_id int64) (float64, error) {
	var res float64
	// 根据 play_id 和 plan_id 查询 单场的 票总金额
	var temp_ticket []do.Ticket
	err := DB.Model(&do.Ticket{}).Where("play_id = ? and plan_id = ?", play_id, plan_id).Find(&temp_ticket).Error
	if err != nil {
		return 0, errors.New("根据 play_id 和 plan_id 查询单场次票总金额有误")
	}
	for j := range temp_ticket {
		res += temp_ticket[j].TicketPrice
	}
	return res, nil
}

func (td *TicketDao) CountTicketEvery(play_id int64) (float64, error) {
	var res float64
	// 根据 play_id 查询票总金额
	var temp_ticket []do.Ticket
	err := DB.Model(&do.Ticket{}).Where("play_id = ?", play_id).Find(&temp_ticket).Error
	if err != nil {
		return 0, errors.New("根据 play_id 查询票总金额有误")
	}
	for j := range temp_ticket {
		res += temp_ticket[j].TicketPrice
	}
	return res, nil
}

// 判断票是否过期
func (td *TicketDao) CheckTicketExpire(ticketID int64) (bool, error) {
	ticket, err := td.GetTicketByID(ticketID)
	if err != nil {
		return false, err
	}
	// 过期返回 true
	return ticket.TicketExpireTime.Before(time.Now()), nil
}

// 核销票
func (td *TicketDao) VerifyTicket(ticketID int64) error {
	// 检查票的过期时间
	ok, err := td.CheckTicketExpire(ticketID)
	if err != nil {
		return err
	}
	if ok {
		err := td.ChangeTicketStatus(ticketID, do.TicketStatusCanceled)
		if err != nil {
			return err
		}
		return errors.New("票已过期")
	}
	// 修改票的状态
	return td.ChangeTicketStatus(ticketID, do.TicketStatusUsed)
}

func (td *TicketDao) ChangeTicketStatus(ticketID int64, status int8) error {
	return DB.Model(&do.Ticket{}).Where("ticket_id = ?", ticketID).Update("ticket_status", status).Error
}

// 查票id
func (td *TicketDao) GetTicketByID(ticketID int64) (*do.Ticket, error) {
	var ticket do.Ticket
	err := DB.Where("ticket_id = ?", ticketID).First(&ticket).Error
	if err != nil {
		if err.Error() == "record not found" {
			return nil, err
		}
		return nil, err
	}
	return &ticket, nil
}

// 查票列表
func (td *TicketDao) GetTicketList(customerID int64) ([]*do.Ticket, error) {
	var tickets []do.Ticket
	// 查询客户未过期且状态小于 2 的票
	err := DB.Where("customer_id = ? AND ticket_status < 2 AND ticket_expire_time > NOW()", customerID).
		Find(&tickets).Error
	if err != nil {
		return nil, err
	}

	// 转换为指针 slice
	ticketsPtr := make([]*do.Ticket, len(tickets))
	for i := range tickets {
		ticketsPtr[i] = &tickets[i]
	}
	return ticketsPtr, nil
}

func (td *TicketDao) CancelTicket(ticketID int64) error {
	return DB.Delete(&do.Ticket{}, ticketID).Error
}

func (td *TicketDao) InsertTicket(customerID int64, planID int64, seatID int64, customerName string, ticketPrice float64, ticketExpireTime time.Time, playID int64, role int8) (id int64, err error) {
	fmt.Println("票过期时间dao：", ticketExpireTime)
	ticket := &do.Ticket{
		CustomerID:       customerID,
		CustomerName:     customerName,
		PlayID:           playID,
		PlanID:           planID,
		SeatID:           seatID,
		TicketPrice:      ticketPrice,
		TicketExpireTime: ticketExpireTime,
		TicketStatus:     do.TicketStatusUnUsed,
		Role:             role,
	}
	err = DB.Create(ticket).Pluck("ticket_id", &ticket.TicketID).Error
	if err != nil {
		return 0, err
	}
	return ticket.TicketID, nil
}
