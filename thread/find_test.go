package thread

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/xxmdhs/emeraldsledger/http"
)

func TestFindPage(t *testing.T) {
	l := http.NewLimitGet(10, 5000, 10)
	ads, err := FindPage(623778, 1, l)
	if err != nil {
		t.Error(err)
	}
	b, _ := json.Marshal(ads)

	fmt.Println(string(b))
}
