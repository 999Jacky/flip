package qtool

import (
	"strings"
)

func GetID(Qust Qust) []string {
	var ID []string
	ID = append(ID, Qust.Num1.ID)
	ID = append(ID, Qust.Num2.ID)
	ID = append(ID, Qust.Num3.ID)
	ID = append(ID, Qust.Num4.ID)
	ID = append(ID, Qust.Num5.ID)
	ID = append(ID, Qust.Num6.ID)
	ID = append(ID, Qust.Num7.ID)
	ID = append(ID, Qust.Num8.ID)
	ID = append(ID, Qust.Num9.ID)
	ID = append(ID, Qust.Num10.ID)
	ID = append(ID, Qust.Num11.ID)
	ID = append(ID, Qust.Num12.ID)
	ID = append(ID, Qust.Num13.ID)
	ID = append(ID, Qust.Num14.ID)
	ID = append(ID, Qust.Num15.ID)
	ID = append(ID, Qust.Num16.ID)
	return ID
}
func GetPostUrl(str string) string {
	strIndex := strings.Index(str, "_url_submit_")
	if strIndex > 0 {
		postStr := ""
		t := false
		for i := strIndex + 12; i < strIndex+250; i++ {
			if string(str[i]) == "'" {
				if t {
					return postStr
				} else {
					t = true
				}
			} else {
				if t {
					postStr += string(str[i])
				}
			}
		}
	}
	return "fail"
}
