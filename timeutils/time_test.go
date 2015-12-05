package timeutils

import (
	"fmt"
	"testing"
	"time"
)

func TestDateFormat(t *testing.T) {
	fmt.Println(DateFormat(time.Now(), "YYYY-MM-DD HH:mm:ss"))
	fmt.Println(DateFormat2(time.Now(), "YYYY-MM-DD HH:mm:ss"))
}
