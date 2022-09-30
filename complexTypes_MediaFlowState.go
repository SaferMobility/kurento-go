package kurento

// Flowing state of the media.
type MediaFlowState string

// Implement fmt.Stringer interface
func (t MediaFlowState) String() string {
	return string(t)
}

const (
	MEDIAFLOWSTATE_FLOWING     MediaFlowState = "FLOWING"
	MEDIAFLOWSTATE_NOT_FLOWING MediaFlowState = "NOT_FLOWING"
)