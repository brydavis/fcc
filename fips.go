// https://www.fcc.gov/developers/census-block-conversions-api
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	lat, long := 47.6097, -122.3331 // Coordinates of Interest

	fmt.Println(Fips(lat, long))
}

func Fips(lat, long float64) map[string]interface{} {
	results := make(map[string]interface{})

	res, err := http.Get(fmt.Sprintf("http://data.fcc.gov/api/block/find?format=json&latitude=%0.5f&longitude=%0.5f&showall=true", lat, long))
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	defer res.Body.Close()

	contents, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	var v interface{}
	json.Unmarshal(contents, &v)

	results["block"] = v.(map[string]interface{})["Block"].(map[string]interface{})["FIPS"]
	results["county"] = v.(map[string]interface{})["County"].(map[string]interface{})["FIPS"]
	results["state"] = v.(map[string]interface{})["State"].(map[string]interface{})["FIPS"]

	return results

}
