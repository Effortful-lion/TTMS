package dto



type HallIsertReq struct {
	HallName string `json:"hall_name" binding:"required"` // 演出厅名称，唯一
	HallRow  int    `json:"hall_row" binding:"required"`  // 行数
	HallCol  int    `json:"hall_col" binding:"required"`  // 列数
}

type HallUpdateReq struct {
	HallID   int64 		 `json:"hall_id" binding:"required"`
	HallName string 	 `json:"hall_name" binding:"required"`
	HallRow  int     	 `json:"hall_row" binding:"required"`
	HallCol  int     	 `json:"hall_col" binding:"required"`
}

type HallInfoResp struct {
	HallID   int64 		 `json:"hall_id"`
	HallName string 	 `json:"hall_name"`
	HallRow  int     	 `json:"hall_row"`
	HallCol  int     	 `json:"hall_col"`
	HallTotal int    	 `json:"hall_total"`
}

type HallInfoListResp struct {
	Halls []*HallInfoResp `json:"halls"`
}