package toolkits

import (
	"fmt"
	"regexp"
)

//校验电话号码是否合法
func CheckMobile(phone string) bool {
	// ^1第一位为1
	// [345789]{1} 后接一位345789 的数字
	// \\d \d的转义 表示数字 {9} 接9位
	// $ 结束符
	if len(phone) != 11 {
		return false
	}
	regRuler := "^1[3456789]{1}\\d{9}$"
	
	// 正则调用规则
	reg := regexp.MustCompile(regRuler)
	
	// 返回 MatchString 是否匹配
	return reg.MatchString(phone)
}

//邮箱是否合法
func CheckEmail(email string) bool {
	// 正则表达式匹配电子邮箱
	// 匹配标准：至少一个字母数字或点的字符串，后面跟@，然后是至少一个字母数字或-的字符串，后面跟点，然后是至少两个字母的顶级域名。
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9.!#$%&'*+/=?^_` + "`" + `{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`)
	return emailRegex.MatchString(email)
}

// 验证密码长度及复杂度，长度大于8
func ValidatePassword(password string) bool {
	// 至少8个字符，至少包含一个数字和一个特殊字符
	pattern := `^(?=.*[0-9])(?=.*[!@#$%^&*])(.{8,})$`
	matched, err := regexp.MatchString(pattern, password)
	if err != nil {
		fmt.Printf("密码不合法: %s\n", err)
		return false
	}
	return matched
}
