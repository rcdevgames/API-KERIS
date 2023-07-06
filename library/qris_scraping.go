package library

import (
	"QAPI/logger"
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var jarz *cookiejar.Jar

type QrisParams struct {
	email  string
	pass   string
	cookie string
}

type MerchantData struct {
	Name string
	MNID string
}

func InitQris(email string, pass string) *QrisParams {
	// store data
	cookie := fmt.Sprintf("%x_cookie.txt", md5.Sum([]byte(email)))

	// Membuat cookie jar
	jar, err := cookiejar.New(nil)
	if err != nil {
		logger.Log.Err(err).Msg("Gagal Membuat Coockie Jar")
	}

	jarz = jar

	return &QrisParams{
		email:  email,
		pass:   pass,
		cookie: cookie,
	}
}

func (q *QrisParams) Merchant() (MerchantData, error) {
	today := time.Now().Format("2006-01-02")
	fromDate := today
	toDate := time.Now().AddDate(0, 0, 31).Format("2006-01-02")

	label := MerchantData{}
	try := 0
relogin:
	url := "https://merchant.qris.id/m/kontenr.php?idir=pages/historytrx.php"
	data := q.filterData(fromDate, toDate, "")

	res, err := q.request(url, "POST", data)
	if err != nil {
		logger.Log.Err(err).Msg("Gagal Get Merchant")
		return label, err
	}

	// fmt.Println(res)

	if !strings.Contains(res, "logout") {
		try++
		if try > 3 {
			return label, fmt.Errorf("Gagal login setelah 3x percobaan")
		}

		err = q.login()
		if err != nil {
			logger.Log.Err(err).Msg("Gagal Login di Get Merchant")
			return label, err
		}
		goto relogin
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(res))
	if err != nil {
		logger.Log.Err(err).Msg("Gagal Membaca DOM HTML pada Get Merchant")
		return label, err
	}

	labelNode := doc.Find(".callout-info > h5")
	if labelNode.Length() > 0 {
		labelStr := labelNode.Text()
		str := strings.Split(labelStr, " / ")
		label = MerchantData{
			Name: strings.TrimPrefix(str[0], "Merchant : "),
			MNID: strings.TrimPrefix(str[1], "mID : "),
		}
	}

	return label, nil
}

func (q *QrisParams) Mutasi(fromDate string, toDate string, amount int) ([]map[string]interface{}, map[string]string, error) {
	today := time.Now().Format("2006-01-02")
	var amountStr string

	// check data
	if (fromDate != "" && !regexp.MustCompile(`\d{4,}-\d{2,}-\d{2,}`).MatchString(fromDate)) ||
		(toDate != "" && !regexp.MustCompile(`\d{4,}-\d{2,}-\d{2,}`).MatchString(toDate)) {
		return nil, nil, fmt.Errorf("Harap input tanggal dengan format yang benar.\neg: %s", today)
	}
	if amount < 0 {
		amountStr = ""
		return nil, nil, fmt.Errorf("Harap input Nominal dengan bilangan bulat & minimal 0, eg: 100000")
	}

	fromDate = today
	if fromDate != "" {
		fromDate = toDate
	}
	if toDate == "" {
		toDate = time.Now().AddDate(0, 0, 31).Format("2006-01-02")
	}

	try := 0
relogin:
	url := "https://merchant.qris.id/m/kontenr.php?idir=pages/historytrx.php"
	data := q.filterData(fromDate, toDate, amountStr)

	res, err := q.request(url, "POST", data)
	if err != nil {
		logger.Log.Err(err).Msg("Gagal POST Mutasi")
		return nil, nil, err
	}

	if !strings.Contains(res, "logout") {
		try++
		if try > 3 {
			return nil, nil, fmt.Errorf("Gagal login setelah 3x percobaan")
		}

		err = q.login()
		if err != nil {
			logger.Log.Err(err).Msg("Gagal Login di POST Mutasi")
			return nil, nil, err
		}
		goto relogin
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(res))
	if err != nil {
		logger.Log.Err(err).Msg("Gagal Membaca DOM HTML pada POST Mutasi")
		return nil, nil, err
	}

	history := []map[string]interface{}{}
	doc.Find("#history > tbody > tr").Each(func(i int, row *goquery.Selection) {
		record := map[string]interface{}{}
		row.Find("td").Each(func(j int, col *goquery.Selection) {
			record[strings.ToLower(col.Parent().Find("th").Text())] = col.Text()
		})
		history = append(history, record)
	})

	label := map[string]string{}
	labelNode := doc.Find(".callout-info > h5")
	if labelNode.Length() > 0 {
		labelStr := labelNode.Text()
		str := strings.Split(labelStr, " / ")
		label = map[string]string{
			"nama_merchant": strings.TrimPrefix(str[0], "Nama Merchant: "),
			"id_merchant":   strings.TrimPrefix(str[1], "ID Merchant: "),
		}
	}

	return history, label, nil
}

func (q *QrisParams) request(url, method string, data url.Values) (string, error) {
	// Membuat HTTP client dengan cookie jar
	client := &http.Client{
		Jar: jarz,
	}
	fmt.Println(url)
	req, err := http.NewRequest(method, url, strings.NewReader(data.Encode()))
	if err != nil {
		logger.Log.Err(err).Msg("Gagal Membuat Request HTTP")
		return "", err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// req.Header.Set("Cookie", q.cookie)

	resp, err := client.Do(req)
	if err != nil {
		logger.Log.Err(err).Msg("Gagal Eksekusi Request")
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Log.Err(err).Msg("Gagal membaca response body")
		return "", err
	}

	return string(body), nil
}

func (q *QrisParams) login() error {
	url := "https://merchant.qris.id/m/login.php?pgv=go"
	data := q.loginData()

	res, err := q.request(url, "POST", data)
	if err != nil {
		logger.Log.Err(err).Msg("Gagal Request Login")
		return err
	}

	if !strings.Contains(res, "historytrx") {
		q.cookie = ""
		return fmt.Errorf("Tidak dapat login, Harap cek kembali email & password anda")
	}

	return nil
}

func (q *QrisParams) loginData() url.Values {
	data := url.Values{}
	data.Set("username", q.email)
	data.Set("password", q.pass)
	data.Set("submitBtn", "")

	return data
}

func (q *QrisParams) filterData(fromDate string, toDate string, amount string) url.Values {
	data := url.Values{}
	data.Set("datexbegin", fromDate)
	data.Set("datexend", toDate)
	data.Set("limitasidata", fmt.Sprintf("%d", 10))
	data.Set("searchtxt", amount)
	data.Set("Filter", "Filter")

	return data
}
