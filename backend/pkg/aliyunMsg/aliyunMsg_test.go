package aliyunMsg

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestMsg(t *testing.T) {
	randCode := fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))
	fmt.Println(randCode)
}
