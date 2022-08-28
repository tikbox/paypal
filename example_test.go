package paypal_test

import (
	"context"
	"fmt"
	"github.com/tikbox/paypal/v4"
	"testing"
)

func TestExample(t *testing.T) {
	Example()
}

func Example() {
	// Initialize client
	c, err := paypal.NewClient("AecAXJnfUqNBOT68RczPoY3AAHDPfTfCjtC6Q5xtt72FfX34gwMy7Qm4X_HyctUA1xVPgKnBFrdQvNsL", "EHT7Kh0nfloIA5MEdYEUslIEuE0Vu_umUi8NFxfYeX909En02N26lwOAtvyrlDbgNSZP12Ip9c0lpx6X", paypal.APIBaseSandBox)
	if err != nil {
		panic(err)
	}

	// Retrieve access token
	resp, err := c.GetAccessToken(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("resp.Token=", resp.Token)

	token, err := c.GetClientToken(context.Background(), resp.Token)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("token=", token)
}

func ExampleClient_CreatePayout_Venmo() {
	// Initialize client
	c, err := paypal.NewClient("clientID", "secretID", paypal.APIBaseSandBox)
	if err != nil {
		panic(err)
	}

	// Retrieve access token
	_, err = c.GetAccessToken(context.Background())
	if err != nil {
		panic(err)
	}

	// Set payout item with Venmo wallet
	payout := paypal.Payout{
		SenderBatchHeader: &paypal.SenderBatchHeader{
			SenderBatchID: "Payouts_2018_100007",
			EmailSubject:  "You have a payout!",
			EmailMessage:  "You have received a payout! Thanks for using our service!",
		},
		Items: []paypal.PayoutItem{
			{
				RecipientType:   "EMAIL",
				RecipientWallet: paypal.VenmoRecipientWallet,
				Receiver:        "receiver@example.com",
				Amount: &paypal.AmountPayout{
					Value:    "9.87",
					Currency: "USD",
				},
				Note:         "Thanks for your patronage!",
				SenderItemID: "201403140001",
			},
		},
	}

	c.CreatePayout(context.Background(), payout)
}
