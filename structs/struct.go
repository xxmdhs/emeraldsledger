package structs

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

var Shanhai, _ = time.LoadLocation("Asia/Shanghai")

type McbbsAd struct {
	Uid      string
	Username string
	Count    int
	Time     int64
	Cause    string
	Type     string
	Link     string
}

func (m *McbbsAd) Hash() string {
	var s = m.Uid + m.Username + strconv.Itoa(m.Count) + strconv.FormatInt(m.Time, 10) + m.Cause + m.Type + m.Link
	h := sha256.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
