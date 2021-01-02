package filter

var (
	// BinPath is the location of the executable binary
	BinPath string

	filterFile = "/json/reactionroles.json"
)

// Filter phrase or word members will be muted for saying
type Filter struct {
	Value string `json:"value"`
}
