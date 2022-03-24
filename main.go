package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type data struct {
	Data string `json:"data"`
}

var (
	SearchDomain     = flag.String("d", "", "exalple [domainhouse.exe -d baidu.com]")
	SearchDomainFile = flag.String("f", "", "exalple [domainhouse.exe -f target.txt]")
)

func main() {
	flag.Parse()
	if (*SearchDomain == "" && *SearchDomainFile == "") || (*SearchDomain != "" && *SearchDomainFile != "") {
		flag.Usage()
	}
	if *SearchDomain != "" {
		QueryDomain()
	}
	if *SearchDomainFile != "" {
		SearchReadfile()
	}
}

func QueryDomain() {
	res, err := http.Get("https://domainhouse.buzz/query.php?token=9378409013576a0cb7c6fac863d5dfe03f0d288a&domain=" + *SearchDomain)
	if err != nil {
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	doamin := data{}
	err = json.Unmarshal(body, &doamin)
	if err != nil {
		return
	}
	fmt.Println(doamin.Data)
}

func SearchReadfile() {
	fi, err := os.Open(*SearchDomainFile)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()
	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		*SearchDomain = string(a)
		QueryDomain()
	}

}
