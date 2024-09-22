package holiday
import (
    "testing"
    "fmt"
)
//go test -v -test.run Test_NewHolidayInfo
func Test_NewHolidayInfo(t *testing.T) {
    path := "./holiday.json"
	err := NewHolidayInfo(path)
	if err != nil {
		fmt.Printf("new holiday info failed:%v\n",err)
		return 
	}
	//节假日列表信息
    fmt.Printf("holidays:%v\n",Holidays)
	//是否是周末
	fmt.Printf("is weekend:%v\n",IsWeekend("2024-09-22"))
	//获取年度节假日信息
	fmt.Printf("get year holiday list:%v\n",GetYearHolidays("2024"))
	//获取月度节假日信息
	fmt.Printf("get month holiday list:%v\n",GetMonthHolidays("2024","5"))
	//获取年度调休日信息
	fmt.Printf("get year make-up list:%v\n",GetYearMakeUpDay("2024"))
	//获取月度调休日信息
	fmt.Printf("get month make-up list:%v\n",GetMonthMakeUpDay("2024","9"))
	//校验某日是否为非工作日
	fmt.Printf("is holiday:%v\n",IsHoliday("2024-09-22"))
	//校验某日是否为工作日
	fmt.Printf("is office day:%v\n",IsOfficeDay("2024-09-22"))
}
