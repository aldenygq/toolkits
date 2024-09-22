package toolkits

import (
	"fmt"
	"time"
	"strings"
	
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


//获取当前周几,输出英文
func CurrentWeek() string {
	return time.Now().Weekday().String()
}
//Weekday英文转阿拉伯数字
// WeekToNumber 将星期的英文缩写转换为对应的数字（1 = 星期日，2 = 星期一，...）
func WeekToNumber(week string) (int, error) {
	switch week {
	case "Sun", "Sunday":
		return 1, nil
	case "Mon", "Monday":
		return 2, nil
	case "Tue", "Tuesday":
		return 3, nil
	case "Wed", "Wednesday":
		return 4, nil
	case "Thu", "Thursday":
		return 5, nil
	case "Fri", "Friday":
		return 6, nil
	case "Sat", "Saturday":
		return 7, nil
	default:
		return 0, fmt.Errorf("invalid week name: %s", week)
	}
}
// NumberToWeek 将数字转换为对应的星期英文缩写
func NumberToWeek(number int) (string, error) {
	switch number {
	case 1:
		return "Sunday", nil
	case 2:
		return "Monday", nil
	case 3:
		return "Tuesday", nil
	case 4:
		return "Wednesday", nil
	case 5:
		return "Thursday", nil
	case 6:
		return "Friday", nil
	case 7:
		return "Saturday", nil
	default:
		return "", fmt.Errorf("invalid week number: %d", number)
	}
}
//Weekday转英文string
func WeekdayToStr(week time.Weekday) string {
	return week.String()
}
//英文string转Weekday
func StringToWeekday(str string) (time.Weekday, error) {
	str = strings.TrimSpace(str)

	switch str {
	case "Sunday", "Sun":
		return time.Sunday, nil
	case "Monday", "Mon":
		return time.Monday, nil
	case "Tuesday", "Tue":
		return time.Tuesday, nil
	case "Wednesday", "Wed":
		return time.Wednesday, nil
	case "Thursday", "Thu":
		return time.Thursday, nil
	case "Friday", "Fri":
		return time.Friday, nil
	case "Saturday", "Sat":
		return time.Saturday, nil
	default:
		return 0, fmt.Errorf("invalid weekday name: %s", str)
	}
}

//WeekdayToChinese
func WeekdayToChinese(weekday time.Weekday) string {
	switch weekday {
	case time.Sunday:
		return "星期日"
	case time.Monday:
		return "星期一"
	case time.Tuesday:
		return "星期二"
	case time.Wednesday:
		return "星期三"
	case time.Thursday:
		return "星期四"
	case time.Friday:
		return "星期五"
	case time.Saturday:
		return "星期六"
	default:
		return "未知"
	}
}
//ChineseToNum
func ChineseToNumber(weekday string) int {
	switch weekday {
	case "星期日":
		return 0
	case "星期一":
		return 1
	case "星期二":
		return 2
	case "星期三":
		return 3
	case "星期四":
		return 4
	case "星期五":
		return 5
	case "星期六":
		return 6
	default:
		return 10
	}
}

