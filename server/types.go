package server

type Home struct {
	ID          string      `json:"id"`
	OwnerID     string      `json:"owner_id"`
	Accessories []Accessory `json:"accessories"`
}

type User struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	Email      string     `json:"email"`
	Location   string     `json:"location"`
	Scenes     []Scene    `json:"scenes"`
	Automation Automation `json:"automation"`
}

type AccessoryType string

const (
	TypeAirConditioningUnit AccessoryType = "AC_UNIT"
	TypeAirPurifier                       = "AIR_PURIFIER"
	TypeCamera                            = "CAMERA"
	TypeCoffeeMaker                       = "COFFEE_MAKER"
	TypeDishwasher                        = "DISHWASHER"
	TypeDryer                             = "DRYER"
	TypeFan                               = "FAN"
	TypeKettle                            = "KETTLE"
	TypeLight                             = "LIGHT"
	TypeOutlet                            = "OUTLET"
	TypeOven                              = "OVEN"
	TypeRefrigerator                      = "REFRIGERATOR"
	TypeSprinkler                         = "SPRINKLER"
	TypeSwitch                            = "SWITCH"
	TypeThermostat                        = "THERMOSTAT"
	TypeVacuum                            = "VACUUM"
	TypeWasher                            = "WASHER"
)

type Action string

const (
	ActionBrightness         Action = "BRIGHTNESS"
	ActionCameraStream              = "CAMERA_STREAM"
	ActionColorSpectrum             = "COLOR_SPECTRUM"
	ActionColorTemperature          = "COLOR_TEMPERATURE"
	ActionDock                      = "DOCK"
	ActionFanSpeed                  = "FAN_SPEED"
	ActionLocator                   = "LOCATOR"
	ActionModes                     = "MODES"
	ActionOnOff                     = "ON_OFF"
	ActionRunCycle                  = "RUN_CYCLE"
	ActionScene                     = "SCENE"
	ActionStartStop                 = "START_STOP"
	ActionTemperatureControl        = "TEMPERATURE_CONTROL"
	ActionTemperatureSetting        = "TEMPERATURE_SETTING"
	ActionToggles                   = "TOGGLES"
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

type Scene struct {
	ID           string           `json:"id"`
	Name         string           `json:"name"`
	TargetHomeID string           `json:"home_id"`
	States       map[string]State `json:"states"`
}

type Automation struct {
	ID           string           `json:"id"`
	Name         string           `json:"name"`
	TargetHomeID string           `json:"home_id"`
	Trigger      interface{}      `json:"trigger"`
	States       map[string]State `json:"states"`
}

type Room struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Accessories []Accessory `json:"accessories"`
}
