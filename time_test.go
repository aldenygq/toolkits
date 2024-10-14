package toolkits
import (
    "fmt"
    "testing"
    "time"
)
//go test -v -test.run Test_GetDay
func Test_GetDay(t *testing.T) {
    fmt.Printf("current day:%v\n",GetDay())
}

//go test -v -test.run Test_StringToUnix
func Test_StringToUnix(t *testing.T) {
    fmt.Printf("unix time:%v\n",StringToUnix("2024-09-24 13:24:56"))
}

//go test -v -test.run Test_UnixToString
func Test_UnixToString(t *testing.T) {
    fmt.Printf("string time:%v\n",UnixToString(1727184296))
}

///go test -v -test.run Test_CompareTwoDay
func Test_CompareTwoDay(t *testing.T) {
    d1 := "2024-09-24"
    d2 := "2024-08-24"
    fmt.Printf("result:%v\n",CompareTwoDay(d1,d2))
}

//go test -v -test.run Test_GetWeekDay
func Test_GetWeekDay(t *testing.T) {
    date := "2024-09-24"
    num,err := GetWeekDay(date)
    if err != nil {
        fmt.Printf("get week day failed:%v\n",err)
        return
    }
    fmt.Printf("weekday:%v\n",num)
}

// go test -v -test.run Test_CurrentWeek
func Test_CurrentWeek(t *testing.T) {
    fmt.Printf("current day:%v\n",CurrentWeek())
}

// go test -v -test.run Test_WeekToNumber
func Test_WeekToNumber(t *testing.T) {
    week := "Sunday"
    num,err := WeekToNumber(week)
    if err != nil {
        fmt.Printf("string week to num failed:%v\n",err)
        return
    }

    fmt.Printf("num:%v\n",num)
}
//go test -v -test.run Test_NumberToWeek
func Test_NumberToWeek(t *testing.T) {
    num := 1
    week,err := NumberToWeek(num)
    if err != nil {
        fmt.Printf("get string week failed:%v\n",err)
        return
    }
    fmt.Printf("week:%v\n",week)
}

//go test -v -test.run Test_WeekdayToStr
func Test_WeekdayToStr(t *testing.T) {
    var tim time.Weekday = 2
    fmt.Printf("week:%v\n",WeekdayToStr(tim))
}

//go test -v -test.run Test_StringToWeekday
func Test_StringToWeekday(t *testing.T) {
    week := "Thursday"
    num,err := StringToWeekday(week)
    if err != nil {
        fmt.Printf("string week to Weekday failed:%v\n",err)
        return
    }

    fmt.Printf("num:%v\n",int(num))
}

//go test -v -test.run Test_WeekdayToChinese
func Test_WeekdayToChinese(t *testing.T) {
    var tim time.Weekday = 2
    fmt.Printf("week:%v\n",WeekdayToChinese(tim))
}

//go test -v -test.run Test_ChineseToNumber
func Test_ChineseToNumber(t *testing.T) {
    week := "星期一"
    fmt.Printf("week:%v\n",ChineseToNumber(week))
}

//go test -v -test.run Test_DayNumInMonth
func Test_DayNumInMonth(t *testing.T) {
    var year int = 2024
    var month int = 5
    fmt.Printf("days:%v\n",DayNumInMonth(year,month))
}

//go test -v -test.run Test_DayListInMonth
func Test_DayListInMonth(t *testing.T) {
     var year int = 2024
     var month int = 5
     fmt.Printf("days list:%v\n",DayListInMonth(year,month))
}

//go test -v -test.run Test_TimeDuration
func Test_TimeDuration(t *testing.T) {
    start := 1728781932
    end := 1728872614

    info,err := TimeDuration(start,end)
    if err != nil {
        fmt.Printf("err:%v\n",err)
        return
    }
    fmt.Printf("duration:%v\n",info)
}
