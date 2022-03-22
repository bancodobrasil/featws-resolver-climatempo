package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"

	log "github.com/sirupsen/logrus"

	adapter "github.com/bancodobrasil/featws-resolver-adapter-go"
	"github.com/bancodobrasil/featws-resolver-adapter-go/types"
	"github.com/bancodobrasil/featws-resolver-climatempo/config"
)

var cfg = config.Config{}

func main() {

	err := config.LoadConfig(&cfg)
	if err != nil {
		log.Fatalf("Não foi possível carregar as configurações: %s\n", err)
	}
	adapter.Run(resolver, adapter.Config{
		Port: cfg.Port,
	})
}

func resolver(resolveInput types.ResolveInput, output *types.ResolveOutput) {
	sort.Strings(resolveInput.Load)

	if contains(resolveInput.Load, "weather") {
		resolveWeather(resolveInput, output)
	}
}

func resolveWeather(resolveInput types.ResolveInput, output *types.ResolveOutput) {
	locale, ok := resolveInput.Context["locale"]

	if !ok {
		output.Errors["weather"] = "The context 'locale' maybe be bounded for resolve 'weather'"
	} else {
		serviceLink := fmt.Sprintf("http://apiadvisor.climatempo.com.br/api/v1/weather/locale/%s/current?token=%s", locale, cfg.Token)

		resp, err := http.Get(serviceLink)
		if err != nil {
			log.Fatalln(err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		result := make(map[string]interface{})

		json.Unmarshal(body, &result)

		err2, ok := result["error"]

		if ok && err2.(bool) {

			output.Errors["weather"] = result
		} else {
			output.Context["weather"] = result
		}

	}
}

func contains(s []string, searchterm string) bool {
	i := sort.SearchStrings(s, searchterm)
	return i < len(s) && s[i] == searchterm
}
