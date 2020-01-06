package canvas

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
// need to run this within 30 seconds of activating pairing on the nanoleaf
func (c *Canvas) GetNewAPIKey() (string, error) {
	resp, err := http.Post(fmt.Sprintf("%v/api/v1/new", c.IP), "", nil)
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

type newAPIKeyResponseBody struct {
	AuthToken string `json:"auth_token"`
}

// GetPanelInfo
// TODO this just prints it out for the moment, but should return some sort of struct!
func (c *Canvas) GetPanelInfo() error {
	resp, err := http.Get(fmt.Sprintf("%v/api/v1/%v/", c.IP, c.APIKey))
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("GetPanelInfo bad response from canvas: %v: %v", resp.StatusCode, resp.Status)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("GetPanelInfo return body: %v\n", body)
	// TODO Parse this into actual struct we can use
	return err
}

//GetEffectsList returns the list of effects that this canvas has
func (c *Canvas) GetEffectsList() ([]string, error) {
	resp, err := http.Get(fmt.Sprintf("%v/api/v1/%v/effects/effectsList", c.IP, c.APIKey))
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
