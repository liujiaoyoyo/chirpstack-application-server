package das

import (
	"github.com/brocaar/chirpstack-application-server/internal/integration/loracloud/client/helpers"
)

// UplinkRequest implements the LoRa Cloud uplink/send request.
type UplinkRequest map[helpers.EUI64]interface{} // UplinkMsg || UplinkMsgModem || UplinkMsgJoining

// UplinkResponse implements the LoRa Cloud uplink/send response.
type UplinkResponse struct {
	Result UplinkDeviceMapResponse `json:"result"`
}

// UplinkDeviceMapResponse implements the LoRa Cloud uplink/send respone per DevEUI.
type UplinkDeviceMapResponse map[helpers.EUI64]UplinkResponseItem

// UplinkMsg implements the LoRa Cloud UplinkMsg object.
// The purpose of this message is to create a downlink opportunity for LoRa Cloud.
type UplinkMsg struct {
	MsgType   string  `json:"msgtype"` // Must be set to "updf"
	FCnt      uint32  `json:"fcnt"`
	Port      uint8   `json:"port"`
	DR        uint8   `json:"dr"`
	Freq      uint32  `json:"freq"`
	Timestamp float64 `json:"timestamp"` // Seconds since UTC
	Payload   string  `json:"payload"`   // Leave this blank
}

// UplinkMsgModem implements the LoRa Cloud UplinkMsg object containing a modem payload.
type UplinkMsgModem struct {
	MsgType   string           `json:"msgtype"` // Must be set to "modem"
	Payload   helpers.HEXBytes `json:"payload"`
	FCnt      uint32           `json:"fcnt"`
	Timestamp float64          `json:"timestamp"` // Seconds since UTC
	DR        uint8            `json:"dr"`
	Freq      uint32           `json:"freq"`
	// TODO GNSS fields!
}

// UplinkMsgJoining implements the LoRa Cloud UplinkMsg object indicating a session reset.
type UplinkMsgJoining struct {
	MsgType   string  `json:"msgtype"`   // Must be set to "joining"
	Timestamp float64 `json:"timestamp"` // Seconds since UTC
	DR        uint8   `json:"dr"`
	Freq      uint32  `json:"freq"`
}

// UplinkResponseItem holds the response for a single DevEUI.
type UplinkResponseItem struct {
	Result UplinkResponseResult `json:"result"`
	Error  string               `json:"error"`
}

// UplinkResponseResult holds the response result.
type UplinkResponseResult struct {
	File              interface{}   `json:"file"`
	StreamRecords     interface{}   `json:"stream_records"`
	FulfilledRequests interface{}   `json:"fulfilled_requests"`
	FPorts            interface{}   `json:"fports"`
	InfoFields        interface{}   `json:"info_fields"`
	PendingRequests   interface{}   `json:"pending_requests"`
	LogMessages       interface{}   `json:"log_messages"`
	Downlink          *LoRaDownlink `json:"dnlink"`
}

// LoRaDownlink implements the LoRa Cloud LoRaDownlink object.
type LoRaDownlink struct {
	Port    uint8            `json:"port"`
	Payload helpers.HEXBytes `json:"payload"`
}