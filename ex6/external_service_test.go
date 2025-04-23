package ex6

import (
	"testing"

	"github.com/golang/mock/gomock"
)

func TestGetData(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := NewMockServiceInterface(ctrl)
	mockService.EXPECT().GetData(123).Return("ok", nil)

	result, err := mockService.GetData(123)
	if err != nil {
		t.Errorf("Error getting data: %v", err)
	}
	if result != "ok" {
		t.Errorf("Expected 'test data', got '%s'", result)
	}

}
