package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	///////////////////////OS////////////////////////
	args := os.Args
	fmt.Println(args)
	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println(hostname)
	} else {
		fmt.Println("error: ", err)
	}

	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	fmt.Println(username, " : ", password)

	/////////////////////FLAG////////////////////////
	// flag.String(key, defaulValue, usageDescription)
	var name *string = flag.String("username", "root", "Put your username")
	var pass *string = flag.String("pass", "root", "Put your pass")
	// wajib parse jika menggunakan flag string
	flag.Parse()

	fmt.Println(*name)
	fmt.Println(*pass)

	//////////////////////STRINGS//////////////////////
	fmt.Println(strings.Contains("Rafly Rigan Nagachi", "Naga"))
	fmt.Println(strings.Split("Rafly Rigan Nagachi", " "))
	fmt.Println(strings.ToLower("Rafly Rigan Nagachi"))
	fmt.Println(strings.ToUpper("Rafly Rigan Nagachi"))
	fmt.Println(strings.Title("Rafly Rigan Nagachi"))
	fmt.Println(strings.Trim("   Rafly    Rigan Nagachi   ", " "))
	fmt.Println(strings.TrimSpace("   Rafly Rigan Nagachi    *    "))
	fmt.Println(strings.Replace("Rafly Rigan Nagachi", "R", "C", -1))
	fmt.Println(strings.ReplaceAll("Rafly Rigan Nagachi", "R", "D"))
	fmt.Println(strings.Count("Rafly Rigan Nagachi", "R"))

	/////////////////////STRCONV////////////////////////
	strBool, err := strconv.ParseBool("true")
	strBoolInv := strconv.FormatBool(true)
	strInt, err := strconv.ParseInt("1234", 10, 32)
	fmt.Println(strBool)
	fmt.Println(strBoolInv)
	fmt.Println()
	if err == nil {
		fmt.Println(strInt)
	} else {
		fmt.Println(err.Error())
	}

	strIntInv := strconv.FormatInt(123, 16)
	fmt.Println(strIntInv)

	///////////////////////MATH/////////////////////////
	fmt.Println("Matematika")
	fmt.Println(math.Round(1.23))
	fmt.Println(math.Round(1.76))
	fmt.Println(math.Floor(1.23))
	fmt.Println(math.Ceil(1.23))
	fmt.Println(math.Max(14, 14.5))
	fmt.Println(math.Min(21, 22))
	fmt.Println(math.Pi)
}
