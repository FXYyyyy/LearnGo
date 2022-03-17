package word

import (
	"strings"
	"unicode"
)

// UnderscoreToUpperCamelCase
/* @Description: 下划线转大写驼峰
*  @param s
*  @return string
*/
func UnderscoreToUpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)
	return strings.Replace(s, " ", "", -1)
}

// UnderscoreToLowerCamelCase
/* @Description: 下划线转小写驼峰
*  @param s
*  @return string
*/
func UnderscoreToLowerCamelCase(s string) string {
	s = UnderscoreToUpperCamelCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

// CamelCaseToUnderscore
/* @Description: 驼峰转下划线
*  @param s
*  @return string
*/
func CamelCaseToUnderscore(s string) string {
	var output []rune
	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) {
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(r))
	}
	return string(output)
}

// ToUpper
/* @Description: 转大写
*  @param s
*  @return string
*/
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// ToLower
/* @Description: 转小写
*  @param s
*  @return string
*/
func ToLower(s string) string {
	return strings.ToLower(s)
}


