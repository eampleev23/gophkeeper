package client

import (
	"encoding/json"
	"fmt"
	"github.com/eampleev23/gophkeeper/internal/models"
	"net/http"
	"strings"
)

var textDataItem models.TextDataItem

func (clientApp *ClientApp) ShowText(response *http.Response, inputID string) error {

	var textDataItemRequestStr = `{"id":`
	textDataItemRequestStr += inputID
	textDataItemRequestStr += `}`

	var textDataItemRequest = []byte(textDataItemRequestStr)

	responseData, _, err := clientApp.RequestToApi(textDataItemRequest, "api/user/get-text-data", http.MethodPost)
	if err != nil {
		fmt.Println("Ошибка при получении ответа от сервера, попробуйте обновить клиент")
		return err
	}

	err = json.Unmarshal(responseData, &textDataItem)

	unPackedTextContent := UnpackData(textDataItem.TextContent, textDataItem.NonceTextContent)
	unPackedTextContent = strings.TrimSuffix(unPackedTextContent, ",")
	fmt.Printf("\nЗапрашиваемые текстовые данные: %s\n\n", unPackedTextContent)

	clientApp.ShowDataItems(nil)
	return err
}
