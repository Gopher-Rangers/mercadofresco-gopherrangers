package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/warehouse"
	"github.com/goccy/go-json"
)

func Test_CreateWarehouse(t *testing.T) {
	t.Run("Deve retornar um status code 201, quando a entra de dados for bem-sucedida e retornará um warehouse.", func(t *testing.T) {

		// server :=

		data := warehouse.Warehouse{
			WarehouseCode:  "j753",
			Address:        "Rua das Margaridas",
			Telephone:      "4833334444",
			MinCapacity:    100,
			MinTemperature: 10,
		}

		dataJSON, _ := json.Marshal(data)

		bodyReader := strings.NewReader(string(dataJSON))

		req := httptest.NewRequest(http.MethodPost, "/api/v1/warehouse", bodyReader)

		res := httptest.NewRecorder()

	})
}
