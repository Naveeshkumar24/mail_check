package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("domain, hasMx,hasSPF,spfRecord,hasDMARC,dmarcRecord\n")
	for scanner.Scan() {
		checkdomain(scanner.Text())
	}
	if err := scanner.Err(); err != nil {

		log.Fatal("Error :Could not read from the input:%v\n", err)
	}
}
func checkdomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error:%v", err)
	}
	if len(mxRecords) > 0 {
		hasMX = true
	}
	txtRecord, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error:%v", err)
	}
	for _, record := range txtRecord {

		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
		}
	}
	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error:%v", err)
	}
	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
		}

	}
	fmt.Println(domain, hasMX, hasSPF, hasDMARC, spfRecord, dmarcRecord)
}
