package kurento

// Represents the state of the RTCDataChannel
type RTCDataChannelState string

// Implement fmt.Stringer interface
func (t RTCDataChannelState) String() string {
	return string(t)
}

const (
	RTCDATACHANNELSTATE_connecting RTCDataChannelState = "connecting"
	RTCDATACHANNELSTATE_open       RTCDataChannelState = "open"
	RTCDATACHANNELSTATE_closing    RTCDataChannelState = "closing"
	RTCDATACHANNELSTATE_closed     RTCDataChannelState = "closed"
)