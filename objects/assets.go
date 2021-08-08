package objects

type AssetStatus string
type AssetType string

const (
	Holding  AssetStatus = "holding"
	Watching AssetStatus = "watching"

	Stock  AssetType = "stock"
	Crypto AssetType = "crypto"
)

type Asset struct {
	ID string `gorm:"primary_key" json:"id,omitempty"`

	Symbol string `json:"symbol,omitempty"`
	Name   string `json:"name,omitempty"`

	Type   AssetType   `json:"type,omitempty"`
	Status AssetStatus `json:"status,omitempty"`

	Price float64 `json:"price,omitempty"`
}
