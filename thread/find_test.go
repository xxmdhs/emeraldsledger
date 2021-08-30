package thread

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestFindPage(t *testing.T) {
	ads, err := FindPage(623778, 1, 10, 500)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(ads)
	fmt.Println(string(b))
}
