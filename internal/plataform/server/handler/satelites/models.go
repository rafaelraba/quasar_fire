package satellites

type SatelliteRequest struct {
    Name string `json:"name"`
    Distance float64 `json:"distance"`
    Message []string `json:"message"`
}
type Position struct {
    X float32 `json:"x"`
    Y float32 `json:"y"`
}
type TopSecretRequest struct {
    Satellites []SatelliteRequest `json:"satellites"`
}
type TopSecretResponse struct {
    Position Position `json:"position"`
    Message string `json:"message"`
}

type TopSecretSplitRequest struct {
    Distance float64 `json:"distance"`
    Message []string `json:"message"`

}

type ResponseInfo struct {
    Status      string `json:"status"`
    Description string `json:"description"`
}


