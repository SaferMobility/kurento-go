package kurento

type RTCIceCandidateAttributes struct {
	IpAddress        string
	PortNumber       int64
	Transport        string
	CandidateType    RTCStatsIceCandidateType
	Priority         int64
	AddressSourceUrl string
}