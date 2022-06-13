package serializers

type RelationshipData struct {
	Data struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	}
}
