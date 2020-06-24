package baidu

import (
	"fmt"
	"testing"
)

func TestAnimalIdentify(t *testing.T) {
	sourceBufferString, _ := Base64EncodeImage("file path")

	tokenResponse, _ := GetAccessToken("appid", "secret")

	response, err := AnimalIdentify(tokenResponse.AccessToken, sourceBufferString, "6", "10")

	if err != nil {
		t.Error(err)
	}

	fmt.Println(response)
}

func TestPlantIdentify(t *testing.T) {
	sourceBufferString, _ := Base64EncodeImage("file path")

	tokenResponse, _ := GetAccessToken("appid", "secret")

	response, err := PlantIdentify(tokenResponse.AccessToken, sourceBufferString, "10")

	if err != nil {
		t.Error(err)
	}

	fmt.Println(response)
}

func TestAdvancedCeneralIdentify(t *testing.T) {
	sourceBufferString, _ := Base64EncodeImage("file path")

	tokenResponse, _ := GetAccessToken("appid", "secret")

	response, err := AdvancedCeneralIdentify(tokenResponse.AccessToken, sourceBufferString, "10")

	if err != nil {
		t.Error(err)
	}

	fmt.Println(response)
}

func TestIngredientIdentify(t *testing.T)  {
	sourceBufferString, _ := Base64EncodeImage("file path")

	tokenResponse, _ := GetAccessToken("appid", "secret")

	response, err := IngredientIdentify(tokenResponse.AccessToken, sourceBufferString, "10")

	if err != nil {
		t.Error(err)
	}

	fmt.Println(response)
}

func TestRedwineIdentify(t *testing.T)  {
	sourceBufferString, _ := Base64EncodeImage("file path")

	tokenResponse, _ := GetAccessToken("appid", "secret")

	response, err := RedwineIdentify(tokenResponse.AccessToken, sourceBufferString)

	if err != nil {
		t.Error(err)
	}

	fmt.Println(response)
}