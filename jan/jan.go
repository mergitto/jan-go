package jan

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dghubble/sling"
)

const (
	conditionNew = "new" // 新品
)

func Exec() {
	hostURL := "https://shopping.yahooapis.jp/ShoppingWebService/V3/itemSearch"
	appID := os.Getenv("YAHOO_API_APP_ID")
	janCode := "4902011830880"
	fmt.Printf("https://shopping.yahooapis.jp/ShoppingWebService/V3/itemSearch?appid=%s&jan_code=%s\n", appID, janCode)

	type query struct {
		AppID     string `url:"appid"`
		JanCode   string `url:"jan_code"`
		Condition string `url:"condition"`
	}

	type response struct {
		Hits []struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Price       int    `json:"price"`
			Image       struct {
				Medium string `json:"medium"`
			} `json:"image"`
		} `json:"hits"`
	}

	res := response{}
	var errRes interface{}
	resp, err := sling.New().Get(hostURL).QueryStruct(query{AppID: appID, JanCode: janCode, Condition: conditionNew}).Receive(&res, &errRes)
	if err != nil {
		log.Fatalf("err[%s]", err)
	}
	if resp.StatusCode == http.StatusOK {
		log.Println(resp.StatusCode)
	}

	for _, hit := range res.Hits {
		fmt.Printf("name: %s\n", truncateString(hit.Name, 20))
		fmt.Printf("description: %s\n", truncateString(hit.Description, 20))
		fmt.Printf("price: %d\n", hit.Price)
		fmt.Printf("image_url: %s\n", hit.Image.Medium)
		fmt.Println()
	}
}

func truncateString(str string, number int) string {
	return string([]rune(str)[:number])
}
