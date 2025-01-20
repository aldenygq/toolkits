package toolkits

import (
	"fmt"
	"regexp"
	"unicode"
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

// 验证密码长度及复杂度，长度可以自定义
func ValidatePassword(password string,lenth int) error {
	// 长度大于8
	if len(password) <= lenth {
		return fmt.Errorf("密码长度必须大于%v",lenth)
	}

	// 包含至少一个小写字母
	hasLowercase := regexp.MustCompile(`[a-z]`)
	// 包含至少一个大写字母
	hasUppercase := regexp.MustCompile(`[A-Z]`)
	// 包含至少一个数字
	hasDigit := regexp.MustCompile(`[0-9]`)
	// 包含至少一个特殊符号
	hasSpecial := regexp.MustCompile(`[^!@#$%^&*/\>]`)

	if !hasLowercase.MatchString(password) {
		return fmt.Errorf("密码必须包含至少一个小写字母")
	}
	if !hasUppercase.MatchString(password) {
		return fmt.Errorf("密码必须包含至少一个大写字母")
	}
	if !hasDigit.MatchString(password) {
		return fmt.Errorf("密码必须包含至少一个数字")
	}
	if !hasSpecial.MatchString(password) {
		return fmt.Errorf("密码必须包含至少一个特殊符号")
	}

	// 如果所有检查都通过
	return nil
}


// containsChinese 判断字符串是否包含中文字符
func ContainsChinese(s string) bool {
    for _, c := range s {
        if unicode.Is(unicode.Scripts["Han"], c) {
            return true
        }
    }
    return false
}
// containsChinese 判断字符串是否包含英文字符
func ContainsEnglishLetters(s string) bool {
    for _, c := range s {
        if unicode.IsLetter(c) { // 检查字符是否为字母
            return true
        }
    }
    return false
}
// containsUppercase 判断字符串中是否包含至少一个大写字母
func ContainsUppercase(s string) bool {
    for _, r := range s {
        if unicode.IsUpper(r) {
            return true
        }
    }
    return false
}
// containsLowercase 判断字符串中是否包含至少一个大写字母
func ContainsLowercase(s string) bool {
    for _, r := range s {
        if unicode.IsLower(r) {
            return true
        }
    }
    return false
}

