package qtool

import "strconv"

type Qrespone struct {
	Ret struct {
		Status string `json:"status"`
		Action struct {
			Redir string `json:"redir"`
		} `json:"action"`
	} `json:"ret"`
}

type StatJson struct {
	Enable bool   `json:"Enable"`
	Ver    int    `json:"Ver"`
	Year   string `json:"Year"`
	Msg    string `json:"Msg"`
}

func CheckInput(qanser []string) bool {
	for i := 0; i < 16; i++ {
		tnum, err := strconv.Atoi(qanser[i])
		if err != nil {
			return false
		}
		if tnum < 1 || tnum > 7 {
			return false
		}
	}
	return true
}
