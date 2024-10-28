package store

import (
	"context"
	"fmt"
	"github.com/eampleev23/gophkeeper/internal/models"
	store_mocks "github.com/eampleev23/gophkeeper/internal/store/internal/mocks-reflect"
	"go.uber.org/mock/gomock"
	"reflect"
	"testing"
)

func TestDBStore_GetTextDataItemByID(t *testing.T) {
	cases := []struct {
		Name string

		Ctx     context.Context
		UserID  int
		InputID int

		StoreGetResponse *models.TextDataItem
		StoreGetError    error

		CalculateResponse *models.TextDataItem

		StoreSetError error

		ExpectedResponse *models.TextDataItem
		ExpectedError    error
	}{
		{
			Name:             "normal",
			Ctx:              context.Background(),
			UserID:           1,
			InputID:          1,
			StoreGetResponse: nil,
			StoreGetError:    nil,

			CalculateResponse: &models.TextDataItem{
				ID:               1,
				MetaValue:        "meta_value",
				TextContent:      "text_content",
				NonceTextContent: "nonce_text_content",
				OwnerID:          1,
			},

			StoreSetError: nil,

			ExpectedResponse: &models.TextDataItem{
				ID:               1,
				MetaValue:        "meta_value",
				TextContent:      "text_content",
				NonceTextContent: "nonce_text_content",
				OwnerID:          1,
			},
			ExpectedError: nil,
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("test case #%d: %s", i, tc.Name), func(t *testing.T) {
			// контроллер для отправки ошибок в тесте
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			//storeMock := mocks.NewMockStore(ctrl)
			storeMock := store_mocks.NewMockStore(ctrl)
			storeGet := storeMock.EXPECT().
				GetTextDataItemByID(tc.Ctx, tc.UserID, tc.InputID).
				Times(1).
				Return(tc.StoreGetResponse, tc.StoreGetError)

			foobarMock := mocks.NewMockFoobar(ctrl)
			if tc.StoreGetError == nil && tc.StoreGetResponse == nil {
				foobarCalculate := foobarMock.EXPECT().
					Calculate(tc.Req).
					After(storeGet).
					Times(1).
					Return(tc.FoobarCalculateResponse)

				_ = storeMock.EXPECT().
					SetFoobar(tc.Req, tc.FoobarCalculateResponse).
					After(foobarCalculate).
					Times(1).
					Return(tc.StoreSetError)
			}

			repo := NewRepo(storeMock)
			repo.foobar = foobarMock

			resp, err := repo.GetFoobar(tc.Req)
			if err := compareErrors(tc.ExpectedError, err); err != nil {
				t.Error(err)
				return
			}
			if !reflect.DeepEqual(resp, tc.ExpectedResponse) {
				t.Errorf("expected foobar response %v, got %v", tc.ExpectedResponse, resp)
				return
			}
		})
	}
}
