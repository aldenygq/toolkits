package toolkits

import (
	//"fmt"
	"time"
	
	//analyticshub "google.golang.org/api/analyticshub/v1beta1"
)
const (
	TIME_LAYOUT = "2006-01-02 15:04:05"
	DAY_LAYOUT = "2006-01-02"
)
//获取当前日期
func GetDay() string {
    now := time.Now().Format(DAY_LAYOUT)
  
   // fmt.Println("Current date:", now.Format("2006-01-02"))
	return now
}

//str time转unix时间戳
func StringToUnix(t string) int64 {
	times, _ := time.Parse(t,TIME_LAYOUT)
	return times.Unix()
}
//unix时间戳转string
func UnixToString(num int64) string {
	return time.Unix(num, 0).Format(TIME_LAYOUT)
}

//比较day1和day2日期顺序，如day1在前。返回0，如day2在前，返回1，如等于，返回2
func CompareTwoDay(day1,day2 string) int64{
	// 创建两个日期
	d1, _ := time.Parse(DAY_LAYOUT, day1)
	d2, _ := time.Parse(DAY_LAYOUT, day2)
	
	// 比较日期
	if d1.Before(d2) {
		return 0
	} else if d1.After(d2) {
		return 1
	} else {
		return 2
	}
	return 3
}

//指定日期确认周几
func GetWeekDay(date string) (int,error) {
    
    // 解析日期字符串
    t, err := time.Parse(DAY_LAYOUT, date)
    if err != nil {
        return int(t.Weekday()),err 
    }
    return int(t.Weekday()),nil 
}
//统计函数耗时
//引用方式
//func Func(n int) int {
//      defer timeCost()()//注意，是对 timeCost()返回的函数进行调用，因此需要加两对小括>    号
//      ......
//}
func TimeCost() func() {
	start := time.Now()
        return func() {
          tc:=time.Since(start)
          fmt.Printf("time cost = %v\n", tc)
      }
}

//校验指定日期是否是周末
// IsWeekend 检查指定的日期是否是周末
func IsWeekend(dateStr string) (bool, error) {
	// 解析日期字符串，默认格式为 "2006-01-02"，这是Go中的特殊时间格式
	date, err := time.Parse(DAY_LAYOUT, dateStr)
	if err != nil {
		return false, err
	}

	// 检查日期是否是星期六(6)或星期日(0)
	dayOfWeek := date.Weekday()
	isWeekend := dayOfWeek == time.Saturday || dayOfWeek == time.Sunday

	return isWeekend, nil
}
