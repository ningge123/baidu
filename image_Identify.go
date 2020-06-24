package baidu

import "net/url"

const (
	animalIdentify 					 = "/rest/2.0/image-classify/v1/animal"
	plantIdentify  					 = "/rest/2.0/image-classify/v1/plant"
	advancedCeneralIdentify 	 	 = "/rest/2.0/image-classify/v2/advanced_general"
	ingredientIdentify 				 = "/rest/2.0/image-classify/v1/classify/ingredient"
	dishIdentify					 = "/rest/2.0/image-classify/v2/dish"
	redwineIdentify					 = "/rest/2.0/image-classify/v1/redwine"
)

type AnimalIdentifyResponse struct {
	CommonError
	LogId 				uint64 	`json:"log_id"`
	Result []struct{
		Score 			string 	`json:"score"`
		Name  			string  `json:"name"`
		BaikeInfo struct{
			BaikeUrl  	string 	`json:"baike_url"`
			ImageUrl    string 	`json:"image_url"`
			Description string 	`json:"description"`
		} `json:"baike_info"`
	} `json:"result"`
}

type PlantIdentifyResponse struct {
	CommonError
	LogId 				uint64 `json:"log_id"`
	Result []struct{
		Score 			float32 `json:"score"`
		Name  			string  `json:"name"`
		BaikeInfo struct{
			BaikeUrl  	string `json:"baike_url"`
			ImageUrl    string `json:"image_url"`
			Description string `json:"description"`
		} `json:"baike_info"`
	} `json:"result"`
}

type AdvancedCeneralIdentifyResponse struct {
	CommonError
	LogId 				uint64 `json:"log_id"`
	ResultNum			uint32 `json:"result_num"`
	Result []struct{
		Keyword			string `json:"keyword"`
		Score 			float32 `json:"score"`
		Name  			string  `json:"name"`
		Root			string 	`json:"root"`
		BaikeInfo struct{
			BaikeUrl  	string `json:"baike_url"`
			ImageUrl    string `json:"image_url"`
			Description string `json:"description"`
		} `json:"baike_info"`
	} `json:"result"`
}

type DishIdentifyResponse struct {
	CommonError
	LogId 				uint64 `json:"log_id"`
	ResultNum			uint32 `json:"result_num"`
	Result []struct{
		Calorie			float32 `json:"calorie"`
		Name  			string  `json:"name"`
		probability		float32 `json:"probability"`
		BaikeInfo struct{
			BaikeUrl  	string `json:"baike_url"`
			ImageUrl    string `json:"image_url"`
			Description string `json:"description"`
		} `json:"baike_info"`
	} `json:"result"`
}

type RedwineIdentifyResponse struct {
	CommonError
	LogId 						uint64 `json:"log_id"`
	Result struct{
		Hasdetail				uint `json:"hasdetail"`
		WineNameCn				string `json:"wineNameCn"`
		WineNameEn				string `json:"wineNameEn"`
		CountryCn				string `json:"countryCn"`
		CountryEn				string `json:"countryEn"`
		RegionCn				string `json:"regionCn"`
		RegionEn				string `json:"regionEn"`
		SubRegionCn     		string `json:"subRegionCn"`
		SubRegionEn     		string `json:"subRegionEn"`
		WineryCn	    		string `json:"wineryCn"`
		WineryEn 				string `json:"wineryEn"`
		ClassifyByColor			string `json:"classifyByColor"`
		ClassifyBySugar			string `json:"classifyBySugar"`
		Color					string `json:"color"`
		GrapeCn					string `json:"grapeCn"`
		GrapeEn					string	`json:"grapeEn"`
		TasteTemperature		string `json:"tasteTemperature"`
		Description 			string `json:"description"`
	} `json:"result"`
}

// 动物识别
func AnimalIdentify(token, image, topNum, baikeNum string) (*AnimalIdentifyResponse, error) {
	api := baseURL + animalIdentify

	urlToken, err := tokenApi(api, token)

	if err != nil {
		return nil, err
	}

	params := make(url.Values)
	params["image"] = []string{image}
	params["top_num"] = []string{topNum}
	params["baike_num"] = []string{baikeNum}

	res := new(AnimalIdentifyResponse)
	if err := postForm(urlToken, params, res); err != nil {
		return nil, err
	}

	return res, nil
}

// 植物识别
func PlantIdentify(token, image, baikeNum string) (*PlantIdentifyResponse, error) {
	api := baseURL + plantIdentify

	urlToken, err := tokenApi(api, token)

	if err != nil {
		return nil, err
	}

	params := make(url.Values)
	params["image"] = []string{image}
	params["baike_num"] = []string{baikeNum}

	res := new(PlantIdentifyResponse)
	if err := postForm(urlToken, params, res); err != nil {
		return nil, err
	}

	return res, nil
}

// 通用物体和场景识别高级版
func AdvancedCeneralIdentify(token, image, baikeNum string) (*AdvancedCeneralIdentifyResponse, error)  {
	api := baseURL + advancedCeneralIdentify

	urlToken, err := tokenApi(api, token)

	if err != nil {
		return nil, err
	}

	params := make(url.Values)
	params["image"] = []string{image}
	params["baike_num"] = []string{baikeNum}

	res := new(AdvancedCeneralIdentifyResponse)
	if err := postForm(urlToken, params, res); err != nil {
		return nil, err
	}

	return res, nil
}

// 果蔬识别
func IngredientIdentify(token, image, topNum string) (*AdvancedCeneralIdentifyResponse, error)  {
	api := baseURL + ingredientIdentify

	urlToken, err := tokenApi(api, token)

	if err != nil {
		return nil, err
	}

	params := make(url.Values)
	params["image"] = []string{image}
	params["top_num"] = []string{topNum}

	res := new(AdvancedCeneralIdentifyResponse)
	if err := postForm(urlToken, params, res); err != nil {
		return nil, err
	}

	return res, nil
}

// 菜品识别
func DishIdentify(token, image, filterThreshold, topNum, baikeNum string) (*DishIdentifyResponse, error)  {
	api := baseURL + dishIdentify

	urlToken, err := tokenApi(api, token)

	if err != nil {
		return nil, err
	}

	params := make(url.Values)
	params["image"] = []string{image}
	params["filter_threshold"] = []string{filterThreshold}
	params["top_num"] = []string{topNum}
	params["baike_num"] = []string{baikeNum}

	res := new(DishIdentifyResponse)
	if err := postForm(urlToken, params, res); err != nil {
		return nil, err
	}

	return res, nil
}

// 红酒识别
func RedwineIdentify(token, image string) (*RedwineIdentifyResponse, error) {
	api := baseURL + redwineIdentify

	urlToken, err := tokenApi(api, token)

	if err != nil {
		return nil, err
	}

	params := make(url.Values)
	params["image"] = []string{image}

	res := new(RedwineIdentifyResponse)
	if err := postForm(urlToken, params, res); err != nil {
		return nil, err
	}

	return res, nil
}