package main

import (
	"github.com/gocolly/colly"
	"regexp"
)

func main() {
	url := "http://91db42a7-fb12-4d5e-85e9-ba29fcaae091.challenge.ctf.show/"
	var left_string string
	c := colly.NewCollector()
	for i := 32; i < 128; i++ {
		str := string(i)
		data := map[string]string{
			"ctf_show": str,
		}
		err := c.Post(url, data)
		if err != nil {
			return
		}
		c.OnHTML("html", func(e *colly.HTMLElement) {
			print(e.Text)
			compile, _ := regexp.Compile(`\?\?\?\?\?`)
			if !compile.MatchString(e.Text) {

				left_string += string(i)
			}
		})

		c.OnRequest(func(r *colly.Request) {
			println(r.Method)
			r.Headers.Set("Content-Type", "application/x-www-form-urlencoded")
			//r.ProxyURL = "http://127.0.0.1:8080"
		})
		c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"
		//cookies := c.Cookies("url")
		//c.SetCookies("", cookies)
		c.OnError(func(_ *colly.Response, err error) {
			panic(err)
		})
	}
	println(left_string)

}
