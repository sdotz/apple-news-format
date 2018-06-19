package styles

type ListItemType string

const (
	ListItemTypeBullet            ListItemType = "bullet"
	ListItemTypeDecimal           ListItemType = "decimal"
	ListItemTypeLowerAlphabetical ListItemType = "lower_alphabetical"
	ListItemTypeUpperAlphabetical ListItemType = "upper_alphabetical"
	ListItemTypeLowerRoman        ListItemType = "lower_roman"
	ListItemTypeUpperRoman        ListItemType = "upper_roman"
	ListItemTypeCharacter         ListItemType = "character"
	ListItemTypeNone              ListItemType = "none"
)

type ListItemStyle struct {
	Character string       `json:"character"`
	Type      ListItemType `json:"type"`
}
