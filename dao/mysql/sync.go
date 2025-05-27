package mysql

import (
	"fmt"
	"time"
)

func SyncData() {
	go func() {
		for {
			time.Sleep(2 * time.Minute)
			if err := SyncPlan(); err != nil {
				fmt.Println("SyncData SyncPlan error:", err)
			}
			if err := SyncTicket(); err != nil {
				fmt.Println("SyncData SyncTicket error:", err)
			}
		}
	}()
}

func SyncPlan() error {
	return NewPlanDao().DeletePlanBeforeNow()
}

func SyncTicket() error {
	return NewTicketDao().DeleteTicketUsed()
}