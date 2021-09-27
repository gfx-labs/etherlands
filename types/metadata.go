package types

type ClusterMetadata struct {
	OriginX int64 `json:"origin_x"`
	OriginZ int64 `json:"origin_z"`
	LengthX int64 `json:"length_x"`
	LengthZ int64 `json:"length_z"`

	Offsets [][2]int64 `json:"offsets"`
	PlotIds []uint64   `json:"plot_ids"`
}

type DistrictMetadata struct {
	Owner       string            `json:"owner"`
	Name        string            `json:"name"`
	Contains    []uint64          `json:"contains"`
	Clusters    []ClusterMetadata `json:"clusters"`
	Description string            `json:"description"`
	Image       string            `json:"image"`
	ExternalURL string            `json:"external_url"`
	Attributes  []Attribute       `json:"attributes"`
}

type Attribute struct {
	DisplayType string `json:"display_type,omitempty"`
	TraitType   string `json:"trait_type"`
	Value       int64  `json:"value"`
}
