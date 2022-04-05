package main

import "fmt"

func main() {
	names := []string{"Abdi", "Abdu", "Abde"}
	for i := 0; i < 3; i++ {
		fmt.Println(names[0], "!")
	}
	for _, value := range names {
		fmt.Println(value, ", hadir!")
	}

	names = append(names, "fuad")
	for i := 0; i < len(names); i++ {
		fmt.Println("nama ke-", i, ": ", names[i])
	}

	country := make(map[string]string)
	country["russia"] = "moskow"
	country["indonesia"] = "jakarta"
	for key, v := range country {
		fmt.Println(v, "adalah ibukota dari", key)
	}

	// break and continue
	for _, v := range names {
		if v == "Abdu" {
			fmt.Println("Ini dia si Abdu!")
			continue
		} else if v == "Abde" {
			fmt.Println("Sudah semua ya")
			break // fuad tidak dicetak
		} else {
			fmt.Println("Bukan Abdu")
		}
	}
}
