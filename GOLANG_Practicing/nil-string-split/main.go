package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("\naaaa")
	fmt.Println(os.Getenv("PGSQL_DB_DSN_BRANTECT_API_APPROVE"))
	fmt.Println(os.Getenv("PGSQL_USER_BRANTECT_API_APPROVE"))
	fmt.Println(os.Getenv("PGSQL_PASSWORD_BRANTECT_API_APPROVE"))
	fmt.Println(os.Getenv("PGSQL_DBNAME_BRANTECT_API_APPROVE"))
	fmt.Println(os.Getenv("BRANTECT_API_JWT_URL"))
	fmt.Println(os.Getenv("API_USE_SECRET_BRANTECT_API_JWT"))
	fmt.Println("aaaa")

	domainNm := "sony.gmo"
	domainSplits := strings.Split(domainNm, ".")
	tld := domainSplits[len(domainSplits)-1]
	sld := "-"
	if len(domainSplits) > 2 {
		sld = domainSplits[len(domainSplits)-2]
	}

	fmt.Println("TLD: ", tld)
	fmt.Println("SLD: ", sld)

}
