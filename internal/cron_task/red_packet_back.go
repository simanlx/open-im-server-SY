package cronTask

// 1. 定时查询红包的状态
// 2. 如果红包超时时间已经过了24小时，那么调用红包返回脚本
// 3. 如果红包没超过24小时，继续等待
// 4. 红包返回脚本实际上是一个封装好的请求
/*func mian() {
	for {
		rpktList := getAllRedPacket()
		for _, rpkt := range rpktList {
			if rpkt.createTime+24*60*60 < time.Now().Unix() {
				// 调用红包返回脚本
				returnRedPacket(rpkt)
			} else {
				// 继续等待
				continue
			}
		}
		time.Sleep(60 * time.Second)
	}
}*/
