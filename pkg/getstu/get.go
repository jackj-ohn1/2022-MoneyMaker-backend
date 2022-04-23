package getstu

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strings"
	"time"

	"golang.org/x/net/publicsuffix"
)

type Form struct {
	Execution string
	Lt        string
	EventId   string
}

func GetStu(id, password string) (string, error) {
	//urlstr := "http://spoc.ccnu.edu.cn/studentHomepage/getUserInfo"

	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})

	CommonErr(err)

	client := http.Client{
		Timeout: 10 * time.Second,
		Jar:     jar,
	}

	returnErr := GetCookie(&client, id, password)

	str := "http://kjyy.ccnu.edu.cn/clientweb/xcus/ic2/Default.aspx"

	req, err := http.NewRequest("GET", str, nil)
	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	re := regexp.MustCompile(`acc_info_name">(.*?)</span>`)

	name := re.FindAllStringSubmatch(string(body), -1)
	fmt.Println(name, len(name))

	if len(name) == 1 {
		if len(name[0]) != 2 {
			return "", errors.New("not exists")
		}
	} else {
		return "", errors.New("not exists")
	}

	return name[0][1], returnErr
}

func GetCookie(client *http.Client, username, password string) error {

	var user = url.Values{}

	form, err := GetbaseForm(client)
	if err != nil {
		log.Println(err)
		return err
	}

	user.Set("execution", form.Execution)
	user.Set("_eventId", form.EventId)
	user.Set("lt", form.Lt)
	user.Set("password", password)
	user.Set("username", username)

	req, err := http.NewRequest("POST", "https://account.ccnu.edu.cn/cas/login?service=http://kjyy.ccnu.edu.cn/loginall.aspx?page=", strings.NewReader(user.Encode()))

	if err != nil {
		log.Println(err)
		return err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Mobile Safari/537.36 Edg/99.0.1150.39")

	_, err = client.Do(req)

	if err != nil {
		log.Println(err)
		return err
	}

	fmt.Println("cookie获取成功!")
	return nil
}
func GetInfor(client *http.Client) {

	req1, _ := http.NewRequest("GET", "http://kjyy.ccnu.edu.cn/clientweb/xcus/ic2/Default.aspx", nil)

	resp1, _ := client.Do(req1)

	body1, _ := io.ReadAll(resp1.Body)

	log.Println(string(body1), "ok")

	time.Sleep(5 * time.Second)
}

func GetbaseForm(client *http.Client) (Form, error) {

	var form Form

	req, err := http.NewRequest("GET", "http://kjyy.ccnu.edu.cn/loginall.aspx", nil)

	if err != nil {
		log.Println(err)
		return form, err
	}

	resp, err := client.Do(req)

	if err != nil {
		log.Println(err)
		return form, err
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
		return form, err
	}
	// 获取 hidden 对应的表单值
	lt_reg := regexp.MustCompile(`name="lt" value="(.*?)"`)
	_eventId_reg := regexp.MustCompile(`name="_eventId" value="(.*?)"`)
	execution_reg := regexp.MustCompile(`name="execution" value="(.*?)"`)

	form.Lt = lt_reg.FindAllStringSubmatch(string(body), -1)[0][1]
	form.EventId = _eventId_reg.FindAllStringSubmatch(string(body), -1)[0][1]
	form.Execution = execution_reg.FindAllStringSubmatch(string(body), -1)[0][1]

	/*for _, v := range resp.Cookies() {
		jsessionid = v.Value
	}*/
	fmt.Println("表单获取成功!")
	return form, nil
}

func CommonErr(err error) {
	if err != nil {
		log.Println(err)
	}
}
