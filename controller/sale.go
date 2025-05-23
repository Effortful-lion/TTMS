package controller

import "github.com/gin-gonic/gin"

type SaleController struct {
}

func NewSaleController() *SaleController {
	return &SaleController{}
}

func (s *SaleController) SellHandler(c *gin.Context) {
	panic("implement me")
}