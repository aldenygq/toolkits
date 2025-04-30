package toolkits

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/crypto/bcrypt"
)
//int64类型转*int64
func Int64ToPointInt64(i int64) *int64 {
	ptrValue := new(int64)
    	*ptrValue = i
	return  ptrValue
}


//string字符串转int
func StrToInt(index string) (int,error) {
	result,err := strconv.Atoi(index)
	if err != nil {
		return 0,err
	}
	return result
}

//string字符串转int64
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
func MapToJson(input map[string]interface{}) (string,error) {
	data,err := json.Marshal(input)
	if err != nil {
		return "",err	
	}
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

// CompareHashAndPassword 安全地验证BCrypt哈希密码与明文是否匹配
// 参数：
//   hashedPassword: BCrypt哈希后的密码字符串
//   plainPassword: 用户输入的明文密码
// 返回值：
//   error: 验证成功返回nil，失败返回具体错误原因
func CompareHashAndPassword(hashedPassword string, plainPassword string) error {
   if hashedPassword == "" || plainPassword == "" {
       return errors.New("空密码参数")
   }    
   err := bcrypt.CompareHashAndPassword(
        []byte(hashedPassword),
        []byte(plainPassword),
    )
    
    if err != nil {
        // 对系统级错误进行封装，保留原始错误信息
        if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
            return fmt.Errorf("密码验证失败: %w", err)
        }
        return fmt.Errorf("密码系统错误: %w", err)
    }
    return nil
}

//slice根据指定分隔符生成指定格式字符串，如输入["123456","56789"]和符号@，输出"@123456@56789"
func StringBuild(numbers []string, symbol string) string {
	var builder strings.Builder
	for _, num := range numbers {
		builder.WriteString(symbol)
		builder.WriteString(num)
	}
	return builder.String()
}

//隐藏手机号部分字符串，通常是将中间几位替换为星号（*），以保护用户隐私。
//参数：
//phone：手机号
//startIndex：星号开始字符串索引位置
//endIndex：星号结束字符串索引位置
func HidePhoneWithFormat(phone string,startIndex,endIndex int) string {
    //提取数字
    re := regexp.MustCompile(`\d+`)
    digits := re.FindString(phone)
    //校验手机号是否合法
    if phone == "" || len(digits) < 11 {
        return phone
    }
    if startIndex < 0 || endIndex > 10 {
        return phone
    }


    // 隐藏中间部分
    hidden := digits[:startIndex] + "****" + digits[endIndex:]

    // 恢复原格式（将隐藏后的数字放回原位置）
    result := ""
    digitIndex := 0
    for _, char := range phone {
        if re.MatchString(string(char)) {
            if digitIndex < len(hidden) {
                result += string(hidden[digitIndex])
                digitIndex++
            } else {
                result += string(char)
            }
        } else {
            result += string(char)
        }
    }
    return result
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
