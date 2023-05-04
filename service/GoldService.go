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

	var price float32
	err = json.Unmarshal(body, &price)
	if err != nil {
		log.Fatalf("Error deserializing JSON: %v", err)
	}

	var ounceToKg float32 = 31.1035
	var totalPrice float32

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
