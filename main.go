package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	// 指定要访问的网页 URL

	fmt.Println("anticher V 1.1 仙区自由组织开发")

	fmt.Println("请输入作业精灵url")
	var url string
	fmt.Scanln(&url)
	//url := "http://www.1010jiajiao.com/daan/book/08725100400000000014.html"

	// 发起 HTTP 请求获取网页内容
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		return
	}
	defer resp.Body.Close()

	// 读取网页内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// 定义正则表达式用于匹配 URL
	pattern := `(http[s]?://[^"']+)["']`

	// 编译正则表达式
	re := regexp.MustCompile(pattern)

	// 在响应的 HTML 内容中查找匹配的 URL
	matches := re.FindAllStringSubmatch(string(body), -1)

	// 打印匹配的 URL
	fmt.Println("正则匹配答案页面中")
	for _, match := range matches {

		durl := string(match[1])

		if strings.Contains(durl, "chapter_") {
			fmt.Println(durl)
			//干翻 教育局

			// 发起 HTTP 请求获取网页内容
			resp, err := http.Get(durl)
			if err != nil {
				fmt.Println("Error fetching URL:", err)
				return
			}
			defer resp.Body.Close()

			// 读取网页内容
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("Error reading response body:", err)
				return
			}

			// 定义正则表达式用于匹配 URL
			pattern := `(http[s]?://[^"']+)["']`

			// 编译正则表达式
			re := regexp.MustCompile(pattern)

			// 在响应的 HTML 内容中查找匹配的 URL
			matches := re.FindAllStringSubmatch(string(body), -1)

			// 打印匹配的 URL
			fmt.Println("Found URLs:")

			for _, match := range matches {

				eurl := string(match[1])

				// 检测字符串中是否包含 .jpg?x-oss-process=
				if strings.Contains(eurl, ".jpg?x-oss-process=") {
					fmt.Println("seig heil")

					// 发起 HTTP 请求获取图片内容
					resp, err := http.Get(eurl)
					if err != nil {
						fmt.Println("Error fetching URL:", err)
						return
					}
					defer resp.Body.Close()

					// 创建保存图片的文件
					fileName := filepath.Base(eurl + ".jpg")
					out, err := os.Create(fileName)
					if err != nil {
						fmt.Println("Error creating file:", err)
						return
					}
					defer out.Close()

					// 将 HTTP 响应的内容写入到文件中
					_, err = io.Copy(out, resp.Body)
					if err != nil {
						fmt.Println("Error copying content to file:", err)
						return
					}

					fmt.Println("答案下载完成 仙区自由组织万岁，打倒腐败教育局和学校，以及邪恶培训机构，农民思想的家长！！", fileName)
				}
			}
		}
	}
}
