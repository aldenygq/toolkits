package toolkits

import (
	//"fmt"
	"time"
	
	//analyticshub "google.golang.org/api/analyticshub/v1beta1"
)
const (
	TIME_LAYOUT = "2006-01-02 15:04:05"
)
//获取当前日期
func GetDay() string {
    now := time.Now().Format("2006-01-02")
  
   // fmt.Println("Current date:", now.Format("2006-01-02"))
	return now
}

//str time转unix时间戳
func StrToUnix(t string) int64 {
	times, _ := time.Parse(t,TIME_LAYOUT)
	return times.Unix()
}


//比较day1和day2日期顺序，如day1在前。返回0，如day2在前，返回1，如等于，返回2
func CompareTwoDay(day1,day2 string) int64{
	// 创建两个日期
	d1, _ := time.Parse("2006-01-02", day1)
	d2, _ := time.Parse("2006-01-02", day2)
	
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
    t, err := time.Parse("2006-01-02", date)
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
func timeCost() func() {
	start := time.Now()
       return func() {
          tc:=time.Since(start)
          fmt.Printf("time cost = %v\n", tc)
      }
}
/*
func main() {
	day1 := "2024-01-04"
	day2 := "2024-01-03"
	fmt.Printf("num:%v\n",CompareTwoDay(day1,day2))
	//fmt.Printf("day unix:%v\n",StrToUnix(getDay()))
		currtime := time.Now()
		fmt.Printf("now:%v\n",currtime)
	
	
		t1:=time.Now().Year()        //年
		t2:=time.Now().Month()       //月
		t3:=time.Now().Day()         //日
		t4:=time.Now().Hour()        //小时
		t5:=time.Now().Minute()      //分钟
		t6:=time.Now().Second()      //秒
		t7:=time.Now().Nanosecond()  //纳秒
	
		fmt.Printf("year:%v\n",t1)
		fmt.Printf("eng month:%v\n",t2)
		fmt.Printf("int month:%v\n",int(t2))
		fmt.Printf("day:%v\n",t3)
		fmt.Printf("hour:%v\n",t4)
		fmt.Printf("minute:%v\n",t5)
		fmt.Printf("second:%v\n",t6)
		fmt.Printf("nano second:%v\n",t7)
	
		//如果获取UTC时间，则可以使用time.UTC
		currentTimeData:=time.Date(t1,t2,t3,t4,t5,t6,t7,time.Local) //获取当前时间，返回当前时间Time
		fmt.Println(currentTimeData)    //打印结果：2017-04-11 12:52:52.794411287 +0800 CST
}
*/
