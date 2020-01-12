package canvas

type newAPIKeyResponseBody struct {
	AuthToken string `json:"auth_token"`
}

// PanelInfo is the full output of the PanelInfo from the GetPanelInfo query
type PanelInfo struct {
	Name            string `json:"name"`
	SerialNo        string `json:"serialNo"`
	Manufacturer    string `json:"manufacturer"`
	FirmwareVersion string `json:"firmwareVersion"`
	HardwareVersion string `json:"hardwareVersion"`
	Model           string `json:"model"`
	CloudHash       struct {
	} `json:"cloudHash"`
	Discovery struct {
	} `json:"discovery"`
	Effects struct {
		EffectsList []string `json:"effectsList"`
		Select      string   `json:"select"`
	} `json:"effects"`
	FirmwareUpgrade struct {
	} `json:"firmwareUpgrade"`
	PanelLayout struct {
		GlobalOrientation struct {
			Value int `json:"value"`
			Max   int `json:"max"`
			Min   int `json:"min"`
		} `json:"globalOrientation"`
		Layout struct {
			NumPanels    int `json:"numPanels"`
			SideLength   int `json:"sideLength"`
			PositionData []struct {
				PanelID   int `json:"panelId"`
				X         int `json:"x"`
				Y         int `json:"y"`
				O         int `json:"o"`
				ShapeType int `json:"shapeType"`
			} `json:"positionData"`
		} `json:"layout"`
	} `json:"panelLayout"`
	Schedules struct {
	} `json:"schedules"`
	State struct {
		Brightness struct {
			Value int `json:"value"`
			Max   int `json:"max"`
			Min   int `json:"min"`
		} `json:"brightness"`
		ColorMode string `json:"colorMode"`
		Ct        struct {
			Value int `json:"value"`
			Max   int `json:"max"`
			Min   int `json:"min"`
		} `json:"ct"`
		Hue struct {
			Value int `json:"value"`
			Max   int `json:"max"`
			Min   int `json:"min"`
		} `json:"hue"`
		On struct {
			Value bool `json:"value"`
		} `json:"on"`
		Sat struct {
			Value int `json:"value"`
			Max   int `json:"max"`
			Min   int `json:"min"`
		} `json:"sat"`
	} `json:"state"`
}
