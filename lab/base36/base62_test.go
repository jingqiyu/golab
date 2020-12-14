package base36

import (
	"fmt"
	"testing"
)

func TestDeBase62Url(t *testing.T) {
	base62 := Base62(1916000086559)
	fmt.Println(base62)

	fmt.Println(Base62Url("http://www.baidu.com"))
	fmt.Println(DeBase62Url(Base62Url("http://www.baidu.com")))
}
