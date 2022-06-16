package unit

import (
	"animal-rescue-be/controllers"
	"animal-rescue-be/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type result struct {
    Response   models.Pong  `json:"response"`
}

func TestNewPong (t *testing.T){
	var pongModel = new(models.Pong)
	pongModel.ID = 13
	pongModel.Message = "Pong"
	if (pongModel.ID != 13 && pongModel.Message != "Pong" ){
		t.Errorf("Test failed")
	}
}

func TestPongController(t *testing.T) {
    r := SetUpRouter()
	ping := new(controllers.PingController)
    r.GET("/ping", ping.Ping)
    req, _ := http.NewRequest("GET", "/ping", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    var result result
    json.Unmarshal(w.Body.Bytes(), &result)

    assert.Equal(t, http.StatusOK, w.Code)
    assert.NotEmpty(t, result)
	assert.Equal(t, result.Response.Message, "Pong")
}