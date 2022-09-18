package gbfs

type Metadata struct {
	LastUpdated int    `json:"last_updated"`
	TTL         int    `json:"ttl"`
	Version     string `json:"version"`
}
