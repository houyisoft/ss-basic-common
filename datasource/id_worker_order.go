package datasource

import (
	"math/rand"
	"ss-basic-common/utils/cast"
	"sync"
	"time"
)

var ai *AutoInc

var m *sync.RWMutex

/**
 * 生成订单号
 */
func GetOrderNo(platformId int, proxyId int64, userId int64) string {
	m = new(sync.RWMutex)
	m.RLock()
	orderNo := ""
	if platformId < 10 {
		orderNo = "00" + cast.ToString(platformId)
	} else if platformId < 100 {
		orderNo = "0" + cast.ToString(platformId)
	}
	now := time.Now()
	year := now.Format("2006")
	month := now.Format("01")
	day := now.Format("02")
	time := now.UnixNano()
	timeStr := cast.ToString(time)
	proxyStr := cast.ToString(proxyId)
	proxyStr = proxyStr[len(proxyStr)-4:]
	userStr := cast.ToString(userId)
	userStr = userStr[len(userStr)-4:]
	orderNo = orderNo + proxyStr + userStr + year[2:] + month + day + timeStr[len(timeStr)-6:] + cast.ToString(generateRandomNumber(3))
	m.RUnlock()
	return orderNo
}

//生成count个[start,end)结束的不重复的随机数
func generateRandomNumber(time int64) int {
	return rand.New(rand.NewSource(time)).Intn(3)
}
