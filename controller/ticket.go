package controller

import "github.com/gin-gonic/gin"

type TicketController struct {
}

func NewTicketController() *TicketController {
	return &TicketController{}
}

// 买票
func (t *TicketController) BuyHandler(c *gin.Context) {
	panic("implement me")	
}