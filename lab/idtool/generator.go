package idtool

import (
	"time"

	"github.com/sony/sonyflake"
)

func NextId() (uint64, error) {
	sf := sonyflake.NewSonyflake(sonyflake.Settings{
		StartTime: time.Unix(1514736000, 0),
	})

	return sf.NextID()
}
