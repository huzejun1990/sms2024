// @Author huzejun 2024/1/13 18:31:00
package verification

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func GeneratorRandNo(digit int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var n int32
	for {
		n = r.Int31n(int32(math.Pow10(digit)))
		if n >= int32(math.Pow10(digit-1)) {
			break
		}
	}
	return fmt.Sprintf("%v", n)
}
