package qtool

import (
	"encoding/json"
)

type Qust struct {
	Num1 struct {
		ID       string `json:"id"`
		PollID   string `json:"pollID"`
		Sn       string `json:"sn"`
		Question string `json:"question"`
		Notes    string `json:"notes"`
		Type     string `json:"type"`
		JSON     string `json:"json"`
	} `json:"1"`
	Num2 struct {
		ID       string `json:"id"`
		PollID   string `json:"pollID"`
		Sn       string `json:"sn"`
		Question string `json:"question"`
		Notes    string `json:"notes"`
		Type     string `json:"type"`
		JSON     string `json:"json"`
	} `json:"2"`
	Num3 struct {
		ID       string `json:"id"`
		PollID   string `json:"pollID"`
		Sn       string `json:"sn"`
		Question string `json:"question"`
		Notes    string `json:"notes"`
		Type     string `json:"type"`
		JSON     string `json:"json"`
	} `json:"3"`
	Num4 struct {
		ID       string `json:"id"`
		PollID   string `json:"pollID"`
		Sn       string `json:"sn"`
		Question string `json:"question"`
		Notes    string `json:"notes"`
		Type     string `json:"type"`
		JSON     string `json:"json"`
	} `json:"4"`
	Num5 struct {
		ID       string `json:"id"`
		PollID   string `json:"pollID"`
		Sn       string `json:"sn"`
		Question string `json:"question"`
		Notes    string `json:"notes"`
		Type     string `json:"type"`
		JSON     string `json:"json"`
	} `json:"5"`
	Num6 struct {
		ID       string `json:"id"`
		PollID   string `json:"pollID"`
		Sn       string `json:"sn"`
		Question string `json:"question"`
		Notes    string `json:"notes"`
		Type     string `json:"type"`
		JSON     string `json:"json"`
	} `json:"6"`
	Num7 struct {
		ID       string `json:"id"`
		PollID   string `json:"pollID"`
		Sn       string `json:"sn"`
		Question string `json:"question"`
		Notes    string `json:"notes"`
		Type     string `json:"type"`
		JSON     string `json:"json"`
	} `json:"7"`
	Num8 struct {
		ID       string `json:"id"`
		PollID   string `json:"pollID"`
		Sn       string `json:"sn"`
		Question string `json:"question"`
		Notes    string `json:"notes"`
		Type     string `json:"type"`
		JSON     string `json:"json"`
	} `json:"8"`
	Num9 struct {
		ID       string `json:"id"`
		PollID   string `json:"pollID"`
		Sn       string `json:"sn"`
		Question string `json:"question"`
		Notes    string `json:"notes"`
		Type     string `json:"type"`
		JSON     string `json:"json"`
	} `json:"9"`
	Num10 struct {
		ID       string `json:"id"`
		PollID   string `json:"pollID"`
		Sn       string `json:"sn"`
		Question string `json:"question"`
		Notes    string `json:"notes"`
		Type     string `json:"type"`
		JSON     string `json:"json"`
	} `json:"10"`
	Num11 struct {
		ID       string `json:"id"`
		PollID   string `json:"pollID"`
		Sn       string `json:"sn"`
		Question string `json:"question"`
		Notes    string `json:"notes"`
		Type     string `json:"type"`
		JSON     string `json:"json"`
	} `json:"11"`
	Num12 struct {
		ID       string `json:"id"`
		PollID   string `json:"pollID"`
		Sn       string `json:"sn"`
		Question string `json:"question"`
		Notes    string `json:"notes"`
		Type     string `json:"type"`
		JSON     string `json:"json"`
	} `json:"12"`
	Num13 struct {
		ID       string `json:"id"`
		PollID   string `json:"pollID"`
		Sn       string `json:"sn"`
		Question string `json:"question"`
		Notes    string `json:"notes"`
		Type     string `json:"type"`
		JSON     string `json:"json"`
	} `json:"13"`
	Num14 struct {
		ID       string `json:"id"`
		PollID   string `json:"pollID"`
		Sn       string `json:"sn"`
		Question string `json:"question"`
		Notes    string `json:"notes"`
		Type     string `json:"type"`
		JSON     string `json:"json"`
	} `json:"14"`
	Num15 struct {
		ID       string `json:"id"`
		PollID   string `json:"pollID"`
		Sn       string `json:"sn"`
		Question string `json:"question"`
		Notes    string `json:"notes"`
		Type     string `json:"type"`
		JSON     string `json:"json"`
	} `json:"15"`
	Num16 struct {
		ID       string `json:"id"`
		PollID   string `json:"pollID"`
		Sn       string `json:"sn"`
		Question string `json:"question"`
		Notes    string `json:"notes"`
		Type     string `json:"type"`
		JSON     string `json:"json"`
	} `json:"16"`
}

func JsonParse(str string) Qust {
	var q Qust
	jsonstr := []byte(str)
	json.Unmarshal(jsonstr, &q)
	return q
}
