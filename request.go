package geocoder

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/tonnerre/golang-pretty"
)

type requestBody struct {
	geocode string
	apikey  string
	sco     string
	kind    string
	format  string
	ll      string
	spn     string
	rspn    int
	results int
	skip    int
	lang    string
	key     string
}

func NewRequest(requestString string, maxResults int) *requestBody {
	r := requestBody{geocode: requestString, results: maxResults}
	r.lang = "en_US"
	systemLang := strings.Split(os.Getenv("LANG"), ".")
	if len(systemLang) > 0 && (systemLang[0] == "ru_RU" || systemLang[0] == "uk_UA" || systemLang[0] == "be_BY" || systemLang[0] == "en_US" || systemLang[0] == "tr_TR") {
		r.lang = systemLang[0]
	}
	return &r
}

func (r *requestBody) SetApiKey(apiKey string) {
	r.apikey = apiKey
}

func (r *requestBody) SetKey(key string) {
	r.key = key
}

func (r *requestBody) SetKind(kind string) {
	r.kind = kind
}

func (r *requestBody) SetScope(scope Scope) {
	r.ll = scope.Center().stringToScopeRequest()
	r.spn = scope.Size().stringToScopeRequest()
}

func (r *requestBody) SetScopeCenter(scopeCenter GeoPoint) {
	r.ll = scopeCenter.stringToScopeRequest()
}

func (r *requestBody) SetScopeSize(scopeSize ScopeSize) {
	r.spn = scopeSize.stringToScopeRequest()
}

func (r *requestBody) EnableSearchInScope(scopeSize ScopeSize) {
	r.rspn = 1
}

func (r *requestBody) Skip(skipResult int) {
	r.skip = skipResult
}

func (r *requestBody) SetLang(lang string) {
	if len(lang) > 0 && (lang == "ru_RU" || lang == "uk_UA" || lang == "be_BY" || lang == "en_US" || lang == "tr_TR") {
		r.lang = lang
	}
}

func (r *requestBody) sendRequest() Ymaps {

	url := fmt.Sprintf(
		"https://geocode-maps.yandex.ru/1.x/?geocode=%s&apikey=%s&sco=%s&kind=%s&format=%s&ll=%s&spn=%s&rspn=%d&results=%d&skip=%d&lang=%s&key=%s",
		r.geocode, r.apikey, r.sco, r.kind, r.format, r.ll, r.spn, r.rspn, r.results, r.skip, r.lang, r.key)
	fmt.Println(url)
	httpResp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer httpResp.Body.Close()
	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		panic(err)
	}
	var resp Ymaps
	xml.Unmarshal(body, &resp)
	return resp
}

func FindOne(requestString string) GeoObject {
	request := requestBody{geocode: requestString, results: 1}
	response := request.sendRequest()
	fmt.Printf("%# v", pretty.Formatter(response.GeoObjectCollection.FeatureMembers[0].GeoObject))
	return response.GeoObjectCollection.FeatureMembers[0].GeoObject
}

func Find(requestString string, maxResults int) []GeoObject {
	request := requestBody{geocode: requestString, results: maxResults}
	response := request.sendRequest()
	var result []GeoObject
	for _, item := range response.GeoObjectCollection.FeatureMembers {
		result = append(result, item.GeoObject)
	}
	fmt.Printf("%# v", pretty.Formatter(result))
	return result
}

func FindOneReverse(N, E float64, kind string) GeoObject {
	request := requestBody{geocode: fmt.Sprintf("%f, %f", N, E), kind: kind, results: 1}
	response := request.sendRequest()
	fmt.Printf("%# v", pretty.Formatter(response.GeoObjectCollection.FeatureMembers[0].GeoObject))
	return response.GeoObjectCollection.FeatureMembers[0].GeoObject
}

func FindOneFromScope(requestString string, scope Scope) GeoObject {
	request := requestBody{geocode: requestString, results: 1, ll: scope.Center().stringToScopeRequest(), spn: scope.Size().stringToScopeRequest(), rspn: 1}
	response := request.sendRequest()
	fmt.Printf("%# v", pretty.Formatter(response))
	return GeoObject{}
}
