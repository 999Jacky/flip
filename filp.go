package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"os"
	"strconv"
	"strings"

	"test/filp/qtool"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	Ver := 1

	// 參數設定
	all := flag.Bool("a", false, "一次填完")
	setYear := flag.String("y", "0", "設定期間")
	DebugMode := flag.Bool("d", false, "除錯模式")
	flag.Parse()

	var qanser []string
	fmt.Println(*all)

	// Cookie & httpClient
	CookieJar, _ := cookiejar.New(nil)
	client := &http.Client{
		CheckRedirect: nil,
		Jar:           CookieJar,
	}

	// Future
	fmt.Print("取得資訊...")
	statE, err := http.Get("http://flip.999jacky.co.place/")
	if err != nil {
		fmt.Println("異常")
		fmt.Scanln()
		os.Exit(1)
	}
	statBody, _ := ioutil.ReadAll(statE.Body)
	getstat := qtool.StatJson{Enable: false}
	json.Unmarshal(statBody, &getstat)
	if len(getstat.Msg) > 0 {
		fmt.Println(getstat.Msg)
	}
	if getstat.Enable == false {
		fmt.Println("停用")
		fmt.Scanln()
		os.Exit(1)
	}
	if getstat.Ver > Ver {
		fmt.Println("發現更新")
		qtool.Open("http://999jacky.co.place/flip/file")
	}
	if getstat.Ver < Ver {
		fmt.Println("!Beta!")
	}

	// SetYear
	if *setYear != "0" {
		getstat.Year = *setYear
	}

	fmt.Println("完成")

	if *DebugMode {
		log.Printf("Respone：%s\n", statBody)
	}

	// Login
	account := ""
	password := ""
	fmt.Print("帳：")
	fmt.Scanln(&account)
	fmt.Print("密：")
	fmt.Scanln(&password)

	if !*DebugMode {
		qtool.CallClear()
	}

	fmt.Println(account + "：")
	fmt.Println("期間：" + getstat.Year)

	// 登入_取得 csrf-field
	req, _ := http.NewRequest("POST", "http://flip.stust.edu.tw/mobile/login?next=%2Fmobile", strings.NewReader(""))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("data-ajax", "false")
	resp, _ := client.Do(req)
	ldoc, _ := goquery.NewDocumentFromResponse(resp)
	cfeid, _ := ldoc.Find(`#fs-form > input[type="hidden"]:nth-child(1)`).Attr("value")
	resp.Body.Close()
	str := "csrf-field=" + cfeid + "&account=" + account + "&password=" + password + "&submit=fmSubmit"
	if *DebugMode {
		log.Printf("Csrf-field=%s\n", cfeid)
	}

	// 登入_Post
	lreq, _ := http.NewRequest("POST", "http://flip.stust.edu.tw/mobile/login?next=%2Fmobile", strings.NewReader(str))
	lreq.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	logRespone, _ := client.Do(lreq)
	lrbody, _ := ioutil.ReadAll(logRespone.Body)

	qnext := true

	if strings.Contains(string(lrbody), "帳號不存在") {
		fmt.Println("登入失敗")
		qnext = false
	}

	// 尋找問卷
	allq := 0
	ndoneq := 0
	var qLink []string
	for page := 1; qnext; page++ {

		if *DebugMode {
			log.Printf("第%d頁：\n", page)
		}

		teststr := "http://flip.stust.edu.tw/user/" + account + "/myLearningRecord?showType=poll&coursePeriodType=semester&period=" + getstat.Year + "&courseID=0&page=" + strconv.Itoa(page)
		req2, _ := http.NewRequest("POST", teststr, strings.NewReader(""))
		resp2, _ := client.Do(req2)

		doc, err := goquery.NewDocumentFromResponse(resp2)
		resp2.Body.Close()
		if err != nil {
			fmt.Println(err)
		}

		doc.Find("#pollListTbody").Each(func(i int, s *goquery.Selection) {
			for j := 0; j < 20; j++ {
				findstr := "#pollList_tr" + strconv.Itoa(j)
				ifs := s.Find(findstr + ">td:nth-child(7)").Text()
				band := s.Find(findstr).Find("td.td.major").Text()
				if band == "" {
					continue
				}
				dlink, _ := s.Find(findstr).Find("td.td.major > a").Attr("href")
				dlink = "http://flip.stust.edu.tw" + dlink
				allq++
				if strings.Contains(band, "學習動機問卷") && ifs != "已完成" {
					ndoneq++
					qLink = append(qLink, dlink)
				}

				if *DebugMode {
					log.Printf("%s,%s:%s\n", ifs, band, dlink)
				}

			}
		})

		// 檢查下一頁按鈕
		doc.Find("#xbox-inline > div.module.app-user.app-user-myLearningRecord > div > div > div.body > div > ul > li.fs-page-next.fs-page-content.fs-page-disable").Each(func(j int, s *goquery.Selection) {
			qnext = false
		})
		// 無內容
		if allq == 0 {
			qnext = false
		}

	}
	fmt.Printf("問卷總共:%d,未完成學習動機問卷:%d\n\n", allq, ndoneq)

	if *DebugMode {
		log.Println("以尋找到未完成學習動機問卷連結：")
		for _, v := range qLink {
			log.Println(v)
		}
	}

	// media頁
	for j := 0; j < len(qLink); j++ {

		req3, _ := http.NewRequest("POST", qLink[j], strings.NewReader(""))
		resp3, _ := client.Do(req3)

		qdoc, err := goquery.NewDocumentFromResponse(resp3)
		resp3.Body.Close()
		qTitle := qdoc.Find("head > title").Text()
		qTitle = strings.Split(qTitle, "_")[0]
		if err != nil {
			fmt.Println(err)
		}

		// poll頁
		Qlink, _ := qdoc.Find("#xbox-inline > div > div > div.body > div > iframe").Attr("src")
		Qlink = "http://flip.stust.edu.tw" + Qlink

		req4, _ := http.NewRequest("POST", Qlink, strings.NewReader(""))
		resp4, _ := client.Do(req4)

		q2doc, _ := goquery.NewDocumentFromResponse(resp4)
		resp4.Body.Close()
		Q2link, _ := q2doc.Find("#testing").Attr("data-url")
		qBodyLink := "http://flip.stust.edu.tw" + Q2link

		if *DebugMode {
			log.Printf("第%d個：\n", j+1)
			log.Printf("Media連結：%s\n", Qlink)
			log.Printf("Poll連結：%s\n", qBodyLink)
		}

		// 問卷內容
		req5, _ := http.NewRequest("POST", qBodyLink, strings.NewReader(""))
		resp5, _ := client.Do(req5)

		q3doc, _ := goquery.NewDocumentFromResponse(resp5)
		resp5.Body.Close()
		qScript := q3doc.Find("body > script:nth-child(18)").Text()
		csrfField, _ := q3doc.Find("#CForm_0_csrf-field").Attr("value")
		pageID, _ := q3doc.Find("#poll_fill_0_pageId").Attr("value")
		pollID, _ := q3doc.Find("#poll_fill_0_pollID").Attr("value")
		userID, _ := q3doc.Find("#poll_fill_0_userID").Attr("value")

		if *DebugMode {
			log.Printf("script大小=%d\n", len(qScript))
			log.Printf("csrf=%s\n", csrfField)
			log.Printf("pageID=%s\n", pageID)
			log.Printf("pollID=%s\n", pollID)
			log.Printf("userID=%s\n", userID)
		}

		// 處理Script
		postUrl := qtool.GetPostUrl(qScript)
		qJsonBody, count := qtool.Getbody(qScript)
		if qJsonBody == "fail" {
			fmt.Println("Script處理錯誤")
			continue
		}

		if *DebugMode {
			if qJsonBody == "fail" {
				log.Printf("Script_Fail!,%d\n", count)
			} else {
				log.Printf("Script以處理:%d\n", count)
			}
		}

		qParsed := qtool.JsonParse(qJsonBody)
		qAllID := qtool.GetID(qParsed)

		// 輸入答案
		checkInput := true
		if len(qanser) == 16 && *all {
			checkInput = false
		}
		for checkInput {
			if j < 10 {
				fmt.Println("0" + strconv.Itoa(j+1) + "：" + qTitle)
			} else {
				fmt.Println(strconv.Itoa(j+1) + "：" + qTitle)
			}
			fmt.Println("    輸入16組數字(以,分開)")
			fmt.Println("    1=7,2=6,....,7=1")
			fmt.Print("    Input：")
			input := ""
			fmt.Scanln(&input)
			qanser = strings.Split(input, ",")
			if len(qanser) == 16 {
				checkInput = !qtool.CheckInput(qanser)
			}
			if checkInput {
				fmt.Println("  請輸入16組1~7數字")
				fmt.Println()
			}
		}
		postData := qtool.ToPostURl(csrfField, pageID, pollID, qAllID, qanser, userID)

		if *DebugMode {
			log.Printf("輸入：%s\n", qanser)
			log.Printf("---PostData-Start\n%s", postData)
			log.Println("---PostData-End")
		}

		// 送出
		postUrl = "http://flip.stust.edu.tw" + postUrl + "&" + postData
		// fmt.Println(postUrl)

		req6, _ := http.NewRequest("POST", postUrl, strings.NewReader(""))
		req6.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		resp6, _ := client.Do(req6)
		body6, _ := ioutil.ReadAll(resp6.Body)

		// 確認回傳
		servRespone := qtool.Qrespone{}
		json.Unmarshal(body6, &servRespone)
		if servRespone.Ret.Status == "true" {
			fmt.Println("填寫成功")
		} else {
			fmt.Println("填寫失敗")
		}

		if *DebugMode {
			log.Println("filpRespone：" + servRespone.Ret.Status)
		}

	}
	fmt.Println()
	fmt.Println("全部完成")
	fmt.Scanln()
}
