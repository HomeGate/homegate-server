package server

import (
	"github.com/labstack/echo"
	"net/http"
)

type GoogleRequest struct {
	RequestId string         `json:"requestId"`
	Inputs    []GoogleIntent `json:"inputs"`
}

type GoogleIntent struct {
	Intent string `json:"intent"`
}

type GoogleDevice struct {
	ID              string           `json:"id"`
	Name            GoogleDeviceName `json:"name"`
	Type            string           `json:"type"`
	Traits          []GoogleTraits   `json:"traits"`
	WillReportState bool             `json:"willReportState"`
	RoomHint        string           `json:"roomHint"`
}

type GoogleDeviceName struct {
	Name      string   `json:"name"`
	Nicknames []string `json:"nicknames"`
}

type GoogleTraits string

type GoogleSyncResponse struct {
	RequestId string                    `json:"requestId"`
	Payload   GoogleSyncResponsePayload `json:"payload"`
}

type GoogleSyncResponsePayload struct {
	Devices []GoogleDevice `json:"devices"`
}

func handleGoogleRequests(c echo.Context) error {

	req := new(GoogleRequest)
	c.Bind(req)

	for _, input := range req.Inputs {
		println(input.Intent)
	}

	return c.JSON(http.StatusOK, GoogleSyncResponse{
		RequestId: req.RequestId,
		Payload: GoogleSyncResponsePayload{
			Devices: []GoogleDevice{
				{
					"000-00-000",
					GoogleDeviceName{
						"Lampe",
						[]string{
							"Deckenlampe",
						},
					},
					"action.devices.types.LIGHT",
					[]GoogleTraits{
						"action.devices.traits.OnOff",
					},
					false,
					"Wohnzimmer",
				},
			},
		},
	})
}
