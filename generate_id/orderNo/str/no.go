package str

import (
	"math"
	"strconv"
	"strings"
	"time"

	"gotest/generate_id/orderNo/datetime"
)

//生成单号
//06123xxxxx
//sum 最少10
func MakeYearDaysRand(sum int) string {
	//年
	strs := time.Now().Format("06")
	//一年中的第几天
	days := string(datetime.GetDaysInYearByThisYear())
	count := len(days)
	if count < 3 {
		//重复字符0
		days = strings.Repeat("0", 3-count) + days
	}
	//组合
	strs += days
	//剩余随机数
	sum = sum - 5
	if sum < 1 {
		sum = 5
	}
	//0~9999999的随机数
	ran := GetRand()

	pow := math.Pow(10, float64(sum)) - 1
	//fmt.Println("sum=>", sum)
	//fmt.Println("pow=>", pow)
	result := strconv.Itoa(ran.Intn(int(pow)))

	count = len(result)
	if count < sum {
		//重复字符0
		result = strings.Repeat("0", sum-count) + result
	}
	//组合
	strs += result
	return strs
}
