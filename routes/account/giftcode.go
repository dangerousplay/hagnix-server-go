package account

import "github.com/kataras/iris"

type jGiftCodes struct {
	CharSlots   int   `json:"CharSlots"`
	VaultChests int   `json:"VaultChests"`
	Fame        int   `json:"Fame"`
	Gold        int   `json:"Gold"`
	Gifts       []int `json:"Gifts"`
}

func handleGiftCode(ctx iris.Context) {
	//TODO implement GiftCode
}
