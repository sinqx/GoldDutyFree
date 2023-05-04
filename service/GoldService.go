package service

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	pb "sg/proto"
)

func GetGoldPrice(gold *pb.Gold) (*pb.Price, error) {
	url := "https://www.goldapi.io/api/XAU/USD/"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Error creating HTTP request: %v", err)

	}

	req.Header.Set("x-access-token", "goldapi-h2mk9rlh7fi2if-io")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending HTTP request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading HTTP response: %v", err)
	}

	type GoldInfo struct {
		Price float32 `json:"price"`
	}
	var goldInfo GoldInfo
	err = json.Unmarshal(body, &goldInfo)
	if err != nil {
		log.Fatalf("Error deserializing JSON: %v", err)
	}

	var ounceToKg float32 = 31.1035
	var totalPrice float32
	var price = goldInfo.Price

	switch gold.Weight {
	case pb.GoldWeight_kg:
		totalPrice = price * ounceToKg * gold.Quantity
	case pb.GoldWeight_gram:
		totalPrice = price * ounceToKg * gold.Quantity / 1000
	case pb.GoldWeight_ounce:
		totalPrice = price * gold.Quantity
	default:
		log.Fatalf("Выбран неизвестный вес")
	}

	return &pb.Price{Price: totalPrice}, nil
}

//func findPrice(weight string, cost uint32) uint32 {  // для вычисления  каррат
//	switch weight {
//	case "kg":
//		findPrice()
//	case "ounce":
//		fmt.Println("Выбран вес в унциях")
//	case "gram24k":
//		fmt.Println("Выбран вес в 24-каратных граммах")
//	case "gram22k":
//		fmt.Println("Выбран вес в 22-каратных граммах")
//	case "gram21k":
//		fmt.Println("Выбран вес в 21-каратных граммах")
//	case "gram18k":
//		fmt.Println("Выбран вес в 18-каратных граммах")
//	default:
//		fmt.Println("Выбран неизвестный вес")
//	}
//	return weight * cost
//}

//Пример json ответа с goldApi
//
//{
//"timestamp": 1683191459,
//"metal": "XAU",
//"currency": "USD",
//"exchange": "FOREXCOM",
//"symbol": "FOREXCOM:XAUUSD",
//"prev_close_price": 2039.27,
//"open_price": 2039.27,
//"low_price": 2030.39,
//"high_price": 2081.46,
//"open_time": 1683158400,
//"price": 2034.22,
//"ch": -5.05,
//"chp": -0.25,
//"ask": 2034.62,
//"bid": 2033.94,
//"price_gram_24k": 65.4017,
//"price_gram_22k": 59.9516,
//"price_gram_21k": 57.2265,
//"price_gram_20k": 54.5014,
//"price_gram_18k": 49.0513
//}
