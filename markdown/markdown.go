package markdown

import (
	"fmt"
	"regexp"
	"strings"
)

// Markdown 转换配置
type MarkdownConfig struct {
	HeaderLevel int    // 标题级别 (1-6)
	ListSymbol  string // 列表符号 (-, *, +)
}

// 默认配置
var defaultConfig = MarkdownConfig{
	HeaderLevel: 2,
	ListSymbol:  "-",
}

// 转换HTML模板为Markdown
func ConvertToMarkdown(template string,keepMissing bool, data map[string]string, config ...MarkdownConfig) string {
	// 合并配置
	cfg := defaultConfig
	if len(config) > 0 {
		cfg = config[0]
	}

	// 处理占位符
    re := regexp.MustCompile(`{([^{}]+)}`)

	template = re.ReplaceAllStringFunc(template, func(m string) string {
		key := m[1 : len(m)-1]
        if val, exists := data[key]; exists {
            return val
        }
        if keepMissing {
            return m
        }
        return ""
	})

    fmt.Printf("template:%v\n",template)
	// 转换HTML标签
	template = strings.ReplaceAll(template, "<br/>", "\n")

	// 构建Markdown
	var builder strings.Builder
	lines := strings.Split(template, "\n")

	// 处理标题
	if len(lines) > 0 {
		builder.WriteString(fmt.Sprintf("%s %s\n\n",
			strings.Repeat("#", cfg.HeaderLevel),
			strings.TrimSpace(lines[0])))
	}

	// 处理内容行
	for _, line := range lines[1:] {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// 键值对处理
		if parts := strings.SplitN(line, ":", 2); len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			builder.WriteString(fmt.Sprintf("%s **%s**: %s\n",
				cfg.ListSymbol,
				key,
				value))
		} else {
			builder.WriteString(line + "\n")
		}
	}

	return strings.TrimSpace(builder.String())
}
/*
func main() {
	// 原始模板
	template := "{值班规则名称}<br/>时段: {值班时间段}<br/>今日值班人员: {值班人员}"

	// 填充数据
	data := map[string]string{
		"值班规则名称": "24小时应急响应值班提醒",
		"值班时间段":  "2023-07-20 18:00 至 2023-07-21 08:00",
		"值班人员":    "王五（主值班员） | 赵六（备勤）",
	}

	// 生成Markdown
	output := ConvertToMarkdown(template, true,data,MarkdownConfig{
		HeaderLevel: 1,
		//ListSymbol:  "•",
		ListSymbol:  "-",
	})

	fmt.Println(output)
}
*/
