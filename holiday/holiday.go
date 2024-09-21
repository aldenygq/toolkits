package holiday

import (
    "encoding/json"
    "os"
    "fmt"
    "io/ioutil"
    "time"
    "strings"
)
const (
    DAY_LAYOUT = "2006-01-02"
)
var Holidays []Holiday = make([]Holiday,0)
type Holiday struct {
    Date string `json:"date"`
    Name string `json:"name"`
    IsOfficeDay bool `json:"is_office_day"`
}
func NewHolidayInfo(datefile string) error {
   // 打开文件
   file, err := os.Open(datefile)
   if err != nil {
       fmt.Printf("Error opening file: %s\n", err)
       return err
   }
   defer file.Close()

   // 读取文件内容
   jsonData, err := ioutil.ReadAll(file)
   if err != nil {
       fmt.Printf("Error reading file: %s\n", err)
       return err
   }

   // 定义一个用于存放解析后数据的变量
   var holidays []Holiday = make([]Holiday,0)

   // 解析JSON数据到结构体
   err = json.Unmarshal(jsonData, &holidays)
   if err != nil {
       fmt.Printf("Error parsing JSON: %s\n", err)
       return err
   }
   Holidays = holidays
   return nil
}
//校验指定日期是否是周末
//IsWeekend 检查指定的日期是否是周末
func IsWeekend(dateStr string) bool {
	// 解析日期字符串，默认格式为 "2006-01-02"，这是Go中的特殊时间格式
	date, err := time.Parse(DAY_LAYOUT, dateStr)
	if err != nil {
		return false
	}

	// 检查日期是否是星期六(6)或星期日(0)
	dayOfWeek := date.Weekday()
	isWeekend := dayOfWeek == time.Saturday || dayOfWeek == time.Sunday

	return isWeekend
}
func GetYearHolidays(year string) []Holiday {
    var hs []Holiday = make([]Holiday,0)
    for _,v := range Holidays {
        ds := strings.Split(v.Date,"-")
        if ds[0] == year && !v.IsOfficeDay{
            hs = append(hs,v)
        }
    }
    return hs
}
func GetMonthHolidays(year,month string) []Holiday {
    var hs []Holiday = make([]Holiday,0)
    for _,v := range Holidays {
        ds := strings.Split(v.Date,"-")
        if len(month) == 1 {
            month = "0" + month
        }
        if ds[0] == year && ds[1] == month && !v.IsOfficeDay{
            hs = append(hs,v)
        }
    }
    return hs
}
func GetYearMakeUpDay(year string) []Holiday {
    var hs []Holiday = make([]Holiday,0)
    for _,v := range Holidays {
        ds := strings.Split(v.Date,"-")
       if ds[0] == year && v.IsOfficeDay{
            hs = append(hs,v)
       }
    }

    return hs
}
func GetMonthMakeUpDay(year,month string) []Holiday {
    var hs []Holiday = make([]Holiday,0)
    for _,v := range Holidays {
        ds := strings.Split(v.Date,"-")
        if len(month) == 1 {
            month = "0" + month
        }
        if ds[0] == year && ds[1] == month && v.IsOfficeDay{
            hs = append(hs,v)
        }
    }
    return hs
}
func IsHoliday(dateStr string) bool {
    for _,v := range Holidays {
        if dateStr == v.Date {
            if v.IsOfficeDay {
                return false
            }
        }
    }
    return true
}
//是否是工作日
func IsOfficeDay(dateStr string) bool {
    //是否是周末
    if IsWeekend(dateStr) {
        //周末是否调休
        if IsHoliday(dateStr) {
            return false
        }
    }

    return true
}
