package redpacket

import (
	"fmt"
	"testing"
)

func TestDoubleAverage(t *testing.T) {
	fmt.Println(GetRedPacket(10, 100))
	fmt.Println(GetRedPacket(4, 100))
	fmt.Println(GetRedPacket(4, 12))
	fmt.Println(GetRedPacket(4, 12))
	fmt.Println(GetRedPacket(4, 12))
	fmt.Println(GetRedPacket(9, 100))
}
