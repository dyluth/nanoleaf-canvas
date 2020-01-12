package canvas

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// Canvas represents a connection to a nanoleaf canvas
type Canvas struct {
	IP     string
	APIKey string
	client *http.Client
}

// New creates a new connection to a nanoleaf canvas
func New(ip, apiKey string) *Canvas {
	c := Canvas{IP: ip, APIKey: apiKey}
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	c.client = &http.Client{Transport: tr, Timeout: 10 * time.Second}
	return &c
}

// GetNewAPIKey - to get a new API key fron teh nanoleaf.
// A user is authorized to access the OpenAPI if they can demonstrate physical access of Panels.
// This is achieved by:
// Holding the on-off button down for 5-7 seconds until the LED starts flashing in a pattern
// need to run this within 30 seconds of the above.
func (c *Canvas) GetNewAPIKey() (string, error) {
	resp, err := http.Post(fmt.Sprintf("http://%v:16021/api/v1/new/", c.IP), "text/plain", nil)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("GetNewAPIKey bad response from canvas: %v: %v", resp.StatusCode, resp.Status)
	}
	//body looks like: {"auth_token" : <auth_token>}
	apiKey := newAPIKeyResponseBody{}
	err = json.NewDecoder(resp.Body).Decode(&apiKey)
	defer resp.Body.Close()
	c.APIKey = apiKey.AuthToken
	return apiKey.AuthToken, err
}

// GetPanelInfo
// TODO this just prints it out for the moment, but should return some sort of struct!
func (c *Canvas) GetPanelInfo() (PanelInfo, error) {
	resp, err := http.Get(fmt.Sprintf("http://%v:16021/api/v1/%v/", c.IP, c.APIKey))
	if err != nil {
		return PanelInfo{}, err
	}
	if resp.StatusCode != 200 {
		return PanelInfo{}, fmt.Errorf("GetPanelInfo bad response from canvas: %v: %v", resp.StatusCode, resp.Status)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	// TODO Parse this into actual struct we can use
	info := PanelInfo{}
	err = json.Unmarshal(body, &info)
	if err != nil {
		return PanelInfo{}, err
	}
	/* // print out the json structure returned
	infoPretty, err := json.MarshalIndent(info, "", "  ")
	if err != nil {
		return PanelInfo{}, err
	}
	fmt.Printf("Panel Info:\n%v\n", string(infoPretty))
	*/
	return info, err
}

//GetEffectsList returns the list of effects that this canvas has
func (c *Canvas) GetEffectsList() ([]string, error) {
	resp, err := http.Get(fmt.Sprintf("http://%v:16021/api/v1/%v/effects/effectsList", c.IP, c.APIKey))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("GetPanelInfo bad response from canvas: %v: %v", resp.StatusCode, resp.Status)
	}
	effects := []string{}
	err = json.NewDecoder(resp.Body).Decode(&effects)
	defer resp.Body.Close()
	return effects, err
}

// SetEffect sets the canvas to play the effect of the chosen name
func (c *Canvas) SetEffect(effect string) error {
	payload := strings.NewReader(fmt.Sprintf("{\"select\" : \"%v\"}", effect))

	client := &http.Client{}
	req, err := http.NewRequest("PUT", fmt.Sprintf("http://%v:16021/api/v1/%v/effects", c.IP, c.APIKey), payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "text/plain")

	res, err := client.Do(req)
	defer res.Body.Close()
	_, err = ioutil.ReadAll(res.Body)
	return err
}
