package buyer

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Buyer struct {
	id           int    `json:"id"`
	cardNumberId string `json:"card_number_id"`
	firstName    string `json:"first_name"`
	lastName     string `json:"last_name"`
}

type repository struct{}

type Repository interface {
	GetAll() ([]Buyer, error)
	Save(id int, cardNumberId string, firstName string, lastName string) (Buyer, error)
}

func NewRepository() Repository {
	return &repository{}
}

func (repository) GetAll() ([]Buyer, error) {
	return getAllBuyersFromJson()
}

func (repository) Save(id int, cardNumberId string, firstName string, lastName string) (Buyer, error) {
	newBuyer := Buyer{id, cardNumberId, firstName, lastName}

	savedBuyer, err := saveBuyerJson(newBuyer)

	if err != nil {
		return newBuyer, err
	}

	return savedBuyer, nil
}

func saveBuyerJson(buyer Buyer) (Buyer, error) {
	text := fmt.Sprintf("%d;%s;%s;%s\n", buyer.id, buyer.cardNumberId, buyer.firstName, buyer.lastName)
	f, err := os.OpenFile("../../pkg/store/buyers.json",
		os.O_APPEND|os.O_WRONLY|os.O_CREATE,
		0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, err = f.WriteString(text); err != nil {
		panic(err)
	}
	return buyer, nil
}

func getAllBuyersFromJson() ([]Buyer, error) {
	fi := "../../pkg/store/buyers.json"
	f, err := os.Open(fi)
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		os.Exit(1)
		return nil, err
	}

	var buyers []Buyer

	r := bufio.NewReader(f)
	s, e := r.ReadString('\n')
	for e == nil {
		splitedData := strings.Split(s, ";")
		var buyer Buyer
		buyer.id, _ = strconv.Atoi(splitedData[0])
		buyer.cardNumberId = splitedData[1]
		buyer.firstName = splitedData[2]
		buyer.lastName = splitedData[3]
		buyers = append(buyers, buyer)
		s, e = r.ReadString('\n')
	}
	return buyers, nil
}
