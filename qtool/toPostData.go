package qtool

func ToPostURl(csrfField string, pageID string, pollID string, qAllID []string, anser []string, userID string) string {
	/*
		Data:=make(map[string]string)
		Data["CForm_0_csrf-field"]=csrfField
		astr:="poll_fill_0_"
		Data[astr+"pagID"]=pageID
		Data[astr+"pollID"]=pollID
		bstr:="ques_result"
		cstr:="ques_type"
		for i:=0;i<16;i++{
			Data[astr+bstr+qAllID[i]]=anser[i]
			Data[astr+cstr+qAllID[i]]="single"
		}
		Data[astr+"userID"]=userID
		Data["_fmSubmit"]="yes"
	*/
	astr := "poll_fill_0_"
	bstr := "ques_result"
	cstr := "ques_type"
	dataUrl := ""
	for i := 0; i < 16; i++ {
		dataUrl += astr + cstr + qAllID[i] + "=single&"
		dataUrl += astr + bstr + qAllID[i] + "=" + anser[i] + "&"
	}
	dataUrl += astr + "pagID=" + pageID + "&"
	dataUrl += astr + "pollID=" + pollID + "&"
	dataUrl += astr + "userID=" + userID + "&"
	dataUrl += "_fmSubmit=yes&"
	dataUrl += "CForm_0_csrf-field=" + csrfField
	return dataUrl
}
