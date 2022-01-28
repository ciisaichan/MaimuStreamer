package bilibili

type PlayUrl struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		CurrentQuality     int      `json:"current_quality"`
		AcceptQuality      []string `json:"accept_quality"`
		CurrentQn          int      `json:"current_qn"`
		QualityDescription []struct {
			Qn   int    `json:"qn"`
			Desc string `json:"desc"`
		} `json:"quality_description"`
		Durl []struct {
			URL        string `json:"url"`
			Length     int    `json:"length"`
			Order      int    `json:"order"`
			StreamType int    `json:"stream_type"`
			P2PType    int    `json:"p2p_type"`
		} `json:"durl"`
	} `json:"data"`
}
