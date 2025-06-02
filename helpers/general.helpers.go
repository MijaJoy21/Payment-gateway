package helpers

import (
	"fmt"
	"time"
)

func GenerateInvoiceId() string {
	now := time.Now()
	datePart := now.Format("20060102") // YYYYMMDD
	unixMillis := now.UnixNano() / int64(time.Millisecond)
	return fmt.Sprintf("INV-%s-%d", datePart, unixMillis)
}
