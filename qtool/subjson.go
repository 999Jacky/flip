package qtool

import "strings"

func Getbody(str string) (string, int) {
	a := 0 // {
	count := 0
	jsonstr := ""
	str = strings.Trim(str, " ")
	strend := strings.Index(str, "result")

	if len(str) < 16 || strend < 0 || string(str[16]) != "{" {
		return "fail", strend
	}
	for i := 16; i < strend; i++ {
		count++
		switch string(str[i]) {
		case "{":
			a++
			jsonstr += "{"
		case "}":
			a--
			jsonstr += "}"
		case ",":
			if a == 0 {
				return jsonstr, count
			} else {
				jsonstr += ","
			}
		default:
			jsonstr += string(str[i])
		}
	}

	return "fail", strend
}
