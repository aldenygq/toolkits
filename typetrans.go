package toolkits

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	//"strings"

	"golang.org/x/crypto/bcrypt"
)
func Int64ToPointInt64(i int64) *int64 {
	ptrValue := new(int64)
    	*ptrValue = i
	return  ptrValue
}
//string to int
func StrToInt(index string) int {
	result,_:= strconv.Atoi(index)
	return result
}
//str转int64
func StrToInt64(str string) (int64, error) {
	var (
		num int64
		err error
	)
	num, err = strconv.ParseInt(str, 10, 64)
	if err != nil {
		return num, err
	}

	return num, nil
}
//float64 To string
func Float64ToString(e float64) string {
	return strconv.FormatFloat(e, 'E', -1, 64)
}
//int To string
func IntToString(e int) string {
	return strconv.Itoa(e)
}
//int64 To string
func Int64ToString(e int64) string {
	return strconv.FormatInt(e, 10)
}

//map to json
func MapToJson(input map[string]interface{}) string {
	data,_ := json.Marshal(input)
	return string(data)
}

//[]rune to string
func RuneToStr(r []rune) string {
	return string(r)
}
//string to []rune
func StrToRune(s string) []rune {
	return []rune(s)
}
//bool to string
func BoolToStr(b bool) string {
	//todo :bool to string
	sBool := strconv.FormatBool(true) //方法1
	return sBool
}

//string to bool
func StrToBool(str string) bool {
	//todo :string to bool
	//接受 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False 等字符串；
	//其他形式的字符串会返回错误
	b, _ := strconv.ParseBool("1")
	return b
}
//复杂结构转为string，可转类型包含map、struct
func StructToString(data interface{}) (string, error) {
	d, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	return string(d), nil
}
//取百分比,n表示取几位小数
func FloatRound(f float64, n int) float64 {
	format := "%." + strconv.Itoa(n) + "f"
	res, _ := strconv.ParseFloat(fmt.Sprintf(format, f), 64)
	return res
}

//检测字符串是否包含阿拉伯数字,包含为true，不包含为false
func CheckInt(str string) bool {
	reg := regexp.MustCompile(`[0-9]+`)
	result := reg.FindAllString(str, -1)
	if len(result) != 0 {
		return true
	}
	return false
}

//检测字符串是否包含汉字,包含为true，不包含为false
func CheckString(str string) bool {
	reg := regexp.MustCompile(`[\p{Han}]+`)
	result := reg.FindAllString(str, -1)
	if len(result) >= 1 {
		return true
	}
	return false
}
//校验hash密码
func CompareHashAndPassword(e string, p string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(e), []byte(p))
	if err != nil {
		return false, err
	}
	return true, nil
}

//阿拉伯数值转中文数值
func ToCN(num int) string {
	//1、数字为0
	if num==0{
		return "零"
	}
	var ans string
	//数字
	szdw:=[]string{"十","百","千","万","十万", "百万", "千万","亿"}
	//数字单位
	sz:=[]string{"零","一","二","三","四","五","六","七","八","九"}
	res:=make([]string,0)

	//数字单位角标
	idx:=-1
	for;num>0;{
		//当前位数的值
		x:=num%10
		//2、数字大于等于10
		// 插入数字单位，只有当数字单位角标在范围内，且当前数字不为0 时才有效
		if idx >= 0 && idx < len(szdw) && x != 0{
			res = append([]string{szdw[idx]},res... )
		}
		//3、数字中间有多个0
		// 当前数字为0，且后一位也为0 时，为避免重复删除一个零文字
		if x == 0 && len(res) != 0 && res[0] == "零"{
			res=res[1:]
		}
		// 插入数字文字
		res = append([]string{sz[x]},res... )
		num/=10
		idx++
	}
	//4、个位数为0
	if len(res)>1 && res[len(res)-1]=="零"{
		res=res[:len(res)-1]
	}
	//合并字符串
	for i := 0; i < len(res); i++ {
		ans=ans+res[i]
	}
	return ans
}
