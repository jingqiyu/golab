package huisu

import (
	"fmt"
	"testing"
	"time"
)

func TestPailie(t *testing.T) {
	Pailie([]int{1, 1, 3})

	i := time.Now().Unix() - 3381*86400

	fmt.Println(time.Unix(i, 0))

}
