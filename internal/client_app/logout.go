package client_app

import "net/http"

func (clientApp *ClientApp) Logout(response *http.Response) {
	if response != nil {
		for _, v := range response.Cookies() {
			if v.Name == "token" {
				v.Name = "old"
				v.Value = "old"
				//main()
				continue
			}
		}
		return
	}
}
