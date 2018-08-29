package server

type Home struct {
	ID          string      `json:"id"`
	Owner       string      `json:"owner"`
	Accessories []Accessory `json:"accessories"`
}

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Location string `json:"location"`
}

type AccessoryType string

const (
	AirConditioningUnit AccessoryType = "AC_UNIT"
	AirPurifier                       = "AIR_PURIFIER"
	Camera                            = "CAMERA"
	CoffeeMaker                       = "COFFEE_MAKER"
	Dishwasher                        = "DISHWASHER"
	Dryer                             = "DRYER"
	Fan                               = "FAN"
	Kettle                            = "KETTLE"
	Light                             = "LIGHT"
	Outlet                            = "OUTLET"
	Oven                              = "OVEN"
	Refrigerator                      = "REFRIGERATOR"
	Sprinkler                         = "SPRINKLER"
	Switch                            = "SWITCH"
	Thermostat                        = "THERMOSTAT"
	Vacuum                            = "VACUUM"
	Washer                            = "WASHER"
)

type Action string

const (
	Brightness         Action = "BRIGHTNESS"
	CameraStream              = "CAMERA_STREAM"
	ColorSpectrum             = "COLOR_SPECTRUM"
	ColorTemperature          = "COLOR_TEMPERATURE"
	Dock                      = "DOCK"
	FanSpeed                  = "FAN_SPEED"
	Locator                   = "LOCATOR"
	Modes                     = "MODES"
	OnOff                     = "ON_OFF"
	RunCycle                  = "RUN_CYCLE"
	StartStop                 = "START_STOP"
	TemperatureControl        = "TEMPERATURE_CONTROL"
	TemperatureSetting        = "TEMPERATURE_SETTING"
	Toggles                   = "TOGGLES"
)

type Accessory struct {
	ID    string        `json:"id"`
	Type  AccessoryType `json:"type"`
	State State         `json:"state"`
}

type State struct {
	Properties []Property `json:"properties"`
}

type Property struct {
	Action Action      `json:"action"`
	Value  interface{} `json:"value"`
}
