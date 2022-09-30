package kurento

// Types of candidates
type RTCStatsIceCandidateType string

// Implement fmt.Stringer interface
func (t RTCStatsIceCandidateType) String() string {
	return string(t)
}

const (
	RTCSTATSICECANDIDATETYPE_host            RTCStatsIceCandidateType = "host"
	RTCSTATSICECANDIDATETYPE_serverreflexive RTCStatsIceCandidateType = "serverreflexive"
	RTCSTATSICECANDIDATETYPE_peerreflexive   RTCStatsIceCandidateType = "peerreflexive"
	RTCSTATSICECANDIDATETYPE_relayed         RTCStatsIceCandidateType = "relayed"
)