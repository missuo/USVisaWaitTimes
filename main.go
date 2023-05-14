/*
 * @Author: Vincent Young
 * @Date: 2023-05-12 23:21:03
 * @LastEditors: Vincent Young
 * @LastEditTime: 2023-05-14 23:25:24
 * @FilePath: /USVisaWaitTimes/main.go
 * @Telegram: https://t.me/missuo
 *
 * Copyright © 2023 by Vincent, All Rights Reserved.
 */
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func QueryVisaWaitTimes(cid string) ([]string, error) {
	url := fmt.Sprintf("https://travel.state.gov/content/travel/resources/database/database.getVisaWaitTimes.html?cid=%s&aid=VisaWaitTimesHomePage", cid)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %s", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body failed: %s", err)
	}
	result := strings.Split(strings.ReplaceAll(string(body), " ", ""), "|")
	for i := range result {
		result[i] = strings.TrimSpace(result[i])
	}
	return result, nil
}

func main() {
	l := flag.String("l", "", "Search for embassies by keyword")
	s := flag.String("s", "", "Get visa wait times by city name")
	flag.Parse()

	visaType := []struct {
		visaName string
	}{
		{"Interview Required Visitors (B1/B2)"},
		{"Interview Required Students/Exchange Visitors (F, M, J)"},
		{"Interview Required Petition-Based Temporary Workers (H, L, O, P, Q)"},
		{"Interview Required Crew and Transit (C, D, C1/D)"},
		{"Interview Waiver Students/Exchange Visitors (F, M, J)"},
		{"Interview Waiver Petition-Based Temporary Workers (H, L, O, P, Q)"},
		{"Interview Waiver Crew and Transit (C, D, C1/D)"},
		{"Interview Waiver Visitors (B1/B2)"},
	}
	embassyInfo := []struct {
		code  string
		value string
	}{
		{"P142", "N'Djamena"},
		{"abidjan", "Abidjan"},
		{"P2", "Abu Dhabi"},
		{"P3", "Abuja"},
		{"P4", "Accra"},
		{"adana", "Adana"},
		{"P5", "Addis Ababa"},
		{"P6", "Algiers"},
		{"P7", "Almaty"},
		{"P8", "Amman"},
		{"P9", "Amsterdam"},
		{"P10", "Ankara"},
		{"P11", "Antananarivo"},
		{"P225", "Apia"},
		{"P12", "Ashgabat"},
		{"P13", "Asmara"},
		{"astana", "Astana"},
		{"P15", "Asuncion"},
		{"athens", "Athens"},
		{"P17", "Auckland"},
		{"P226", "Baghdad"},
		{"P19", "Baku"},
		{"P20", "Bamako"},
		{"P21", "Bandar Seri Begawan"},
		{"P22", "Bangkok"},
		{"bangui", "Bangui"},
		{"banjul", "Banjul"},
		{"barcelona", "Barcelona"},
		{"P24", "Beijing"},
		{"P25", "Beirut"},
		{"P26", "Belfast"},
		{"P27", "Belgrade"},
		{"P28", "Belmopan"},
		{"P29", "Berlin"},
		{"P30", "Bern"},
		{"P31", "Bishkek"},
		{"P32", "Bogota"},
		{"P33", "Brasilia"},
		{"P34", "Bratislava"},
		{"P35", "Brazzaville"},
		{"P36", "Bridgetown"},
		{"P37", "Brussels"},
		{"P38", "Bucharest"},
		{"P39", "Budapest"},
		{"P40", "Buenos Aires"},
		{"P41", "Bujumbura"},
		{"P42", "Cairo"},
		{"P43", "Calgary"},
		{"canberra", "Canberra"},
		{"P44", "Cape Town"},
		{"P45", "Caracas"},
		{"P46", "Casablanca"},
		{"P47", "Chengdu"},
		{"P48", "Chennai ( Madras)"},
		{"P49", "Chiang Mai"},
		{"P50", "Chisinau"},
		{"P51", "Ciudad Juarez"},
		{"P52", "Colombo"},
		{"P53", "Conakry"},
		{"P54", "Copenhagen"},
		{"P55", "Cotonou"},
		{"P223", "Curacao"},
		{"P56", "Dakar"},
		{"P57", "Damascus"},
		{"P58", "Dar Es Salaam"},
		{"P59", "Dhahran"},
		{"P60", "Dhaka"},
		{"P227", "Dili"},
		{"P61", "Djibouti"},
		{"P62", "Doha"},
		{"P63", "Dubai"},
		{"P64", "Dublin"},
		{"P65", "Durban"},
		{"P66", "Dushanbe"},
		{"erbil", "Erbil"},
		{"P67", "Florence"},
		{"P68", "Frankfurt"},
		{"P69", "Freetown"},
		{"fukuoka", "Fukuoka"},
		{"P70", "Gaborone"},
		{"P71", "Georgetown"},
		{"P72", "Guadalajara"},
		{"P73", "Guangzhou"},
		{"P74", "Guatemala City"},
		{"P75", "Guayaquil"},
		{"P76", "Halifax"},
		{"P77", "Hamilton"},
		{"P78", "Hanoi"},
		{"P79", "Harare"},
		{"P80", "Havana"},
		{"P81", "Helsinki"},
		{"P82", "Hermosillo"},
		{"P83", "Ho Chi Minh City"},
		{"P84", "Hong Kong"},
		{"P85", "Hyderabad"},
		{"P86", "Islamabad"},
		{"P87", "Istanbul"},
		{"P88", "Jakarta"},
		{"P89", "Jeddah"},
		{"P90", "Jerusalem"},
		{"P91", "Johannesburg"},
		{"P228", "Juba"},
		{"P229", "Kabul"},
		{"P93", "Kampala"},
		{"kaohsiung", "Kaohsiung"},
		{"P94", "Karachi"},
		{"P95", "Kathmandu"},
		{"P96", "Khartoum"},
		{"P97", "Kigali"},
		{"P98", "Kingston"},
		{"P99", "Kinshasa"},
		{"P100", "Kolkata"},
		{"P101", "Kolonia"},
		{"P102", "Koror"},
		{"P103", "Krakow"},
		{"P104", "Kuala Lumpur"},
		{"P105", "Kuwait"},
		{"P106", "Kyiv"},
		{"P107", "La Paz"},
		{"P108", "Lagos"},
		{"lahore", "Lahore"},
		{"P109", "Libreville"},
		{"P110", "Lilongwe"},
		{"P111", "Lima"},
		{"P112", "Lisbon"},
		{"P113", "Ljubljana"},
		{"lome", "Lome"},
		{"P115", "London"},
		{"P116", "Luanda"},
		{"P117", "Lusaka"},
		{"P118", "Luxembourg"},
		{"P119", "Madrid"},
		{"P120", "Majuro"},
		{"P121", "Malabo"},
		{"managua", "Managua"},
		{"P123", "Manama"},
		{"P124", "Manila"},
		{"P125", "Maputo"},
		{"marseille", "Marseille"},
		{"P126", "Maseru"},
		{"P127", "Matamoros"},
		{"P128", "Mbabane"},
		{"P129", "Melbourne"},
		{"P130", "Merida"},
		{"mexicali_tpf", "Mexicali Tpf"},
		{"P131", "Mexico City"},
		{"P132", "Milan"},
		{"P133", "Minsk"},
		{"P134", "Monrovia"},
		{"P135", "Monterrey"},
		{"P136", "Montevideo"},
		{"P137", "Montreal"},
		{"P138", "Moscow"},
		{"P139", "Mumbai (Bombay)"},
		{"P140", "Munich"},
		{"P141", "Muscat"},
		{"P142", "N`Djamena"},
		{"P143", "Naha"},
		{"P144", "Nairobi"},
		{"P145", "Naples"},
		{"P146", "Nassau"},
		{"P147", "New Delhi"},
		{"P148", "Niamey"},
		{"P149", "Nicosia"},
		{"P150", "Nogales"},
		{"P151", "Nouakchott"},
		{"P152", "Nuevo Laredo"},
		{"P153", "Osaka-Kobe"},
		{"P154", "Oslo"},
		{"P155", "Ottawa"},
		{"P156", "Ouagadougou"},
		{"P157", "Panama City"},
		{"P158", "Paramaribo"},
		{"P159", "Paris"},
		{"P160", "Perth"},
		{"P161", "Phnom Penh"},
		{"P162", "Podgorica"},
		{"ponta_delgada", "Ponta Delgada"},
		{"P164", "Port Au Prince"},
		{"P165", "Port Louis"},
		{"P166", "Port Moresby"},
		{"P167", "Port Of Spain"},
		{"porto_alegre", "Porto Alegre"},
		{"P168", "Prague"},
		{"P169", "Praia"},
		{"pretoria", "Pretoria"},
		{"P231", "Pristina"},
		{"P170", "Quebec"},
		{"P171", "Quito"},
		{"P172", "Rangoon"},
		{"P173", "Recife"},
		{"P174", "Reykjavik"},
		{"P175", "Riga"},
		{"P230", "Rio De Janeiro"},
		{"P177", "Riyadh"},
		{"P178", "Rome"},
		{"P179", "San Jose"},
		{"P180", "San Salvador"},
		{"P181", "Sanaa"},
		{"P182", "Santiago"},
		{"P183", "Santo Domingo"},
		{"P184", "Sao Paulo"},
		{"P224", "Sapporo"},
		{"P185", "Sarajevo"},
		{"P186", "Seoul"},
		{"P187", "Shanghai"},
		{"P188", "Shenyang"},
		{"P189", "Singapore"},
		{"P190", "Skopje"},
		{"P191", "Sofia"},
		{"st_georges", "St Georges"},
		{"P192", "St Petersburg"},
		{"P193", "Stockholm"},
		{"P194", "Surabaya"},
		{"P195", "Suva"},
		{"P196", "Sydney"},
		{"P197", "Taipei"},
		{"P198", "Tallinn"},
		{"P199", "Tashkent"},
		{"P200", "Tbilisi"},
		{"P201", "Tegucigalpa"},
		{"P202", "Tel Aviv"},
		{"tijuana", "Tijuana"},
		{"P203", "Tijuana Tpf"},
		{"P204", "Tirana"},
		{"P205", "Tokyo"},
		{"P206", "Toronto"},
		{"P207", "Tripoli"},
		{"P208", "Tunis"},
		{"P209", "Ulaanbaatar"},
		{"usun-new_york", "Usun-New York"},
		{"P210", "Valletta"},
		{"P211", "Vancouver"},
		{"P212", "Vienna"},
		{"P213", "Vientiane"},
		{"P214", "Vilnius"},
		{"P215", "Vladivostok"},
		{"P216", "Warsaw"},
		{"P217", "Windhoek"},
		{"P218", "Yaounde"},
		{"yekaterinburg", "Yekaterinburg"},
		{"P220", "Yerevan"},
		{"P221", "Zagreb"},
	}
	if *l != "" && *s != "" {
		fmt.Println("Error: Cannot use both -l and -s flags at the same time.")
		return
	}
	if *l != "" {
		for _, item := range embassyInfo {
			if strings.Contains(strings.ToLower(item.value), strings.ToLower(*l)) {
				fmt.Printf("%s\n", item.value)
			}
		}
	} else if *s != "" {
		var cityCode string
		var embassyName string
		for _, item := range embassyInfo {
			if strings.Contains(strings.ToLower(item.value), strings.ToLower(*s)) {
				cityCode = item.code
				embassyName = item.value
				break
			}
		}
		if cityCode == "" {
			fmt.Printf("Error: City '%s' not found.\n", *s)
			return
		}
		result, err := QueryVisaWaitTimes(cityCode)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		} else {
			now := time.Now()
			fmt.Printf("Embassy: %s \nTime: %s\n\n", embassyName, now.Format("2006-01-02 15:04:05"))
			for i := 0; i < len(result); i++ {
				waitTime := strings.Replace(result[i], "Days", " Calendar Days", -1)
				waitTime = strings.Replace(waitTime, "SameDay", "Same Day", -1)
				fmt.Printf("%s: %s\n", visaType[i].visaName, waitTime)
			}
		}
		fmt.Printf("\nData From: https://travel.state.gov\nGitHub: https://github.com/missuo/USVisaWaitTimes\n\nMade with \033[31m♥\033[0m by Vincent\n")
	} else {
		fmt.Println("Error: Must specify either -l or -s flag.")
	}
}
