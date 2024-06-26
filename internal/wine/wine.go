package wine

type Wine struct {
	ID     int `json:"id"`
	TypeID int `json:"wine_type"`
	Type   WineType
	Name   string `json:"name"`
	/*Country  Country   `json:"country"`
	Producer string    `json:"producer"`
	Year     time.Time `json:"year"`
	Alcohol  float64   `json:"alcohol"`
	Nose     Nose      `json:"nose"`
	Palate   Palate    `json:"palate"`
	Colour   Colour    `json:"colour"`
	Score    int       `json:"score"`*/
}
type WineType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Country string

// Public Countries
const (
	France     Country = "France"
	Italy      Country = "Italy"
	Spain      Country = "Spain"
	Australia  Country = "Australia"
	NewZealand Country = "New_Zealand"
	Chile      Country = "Chile"
)

type Colour struct {
	Red     RedColour
	White   WhiteColour
	Rose    RoseColour
	Clarity Clarity
}
type RedColour string

// Public Red
const (
	Purple RedColour = "purple"
	Ruby   RedColour = "ruby"
	Garnet RedColour = "garnet"
	Tawny  RedColour = "tawny"
)

type WhiteColour string

// Public White
const (
	Straw  WhiteColour = "straw"
	Yellow WhiteColour = "yellow"
	Golden WhiteColour = "golden"
)

type RoseColour string

// Public Rose
const (
	Blush  RoseColour = "blush"
	Salmon RoseColour = "salmon"
	Pink   RoseColour = "pink"
)

type Clarity string

// Public Clarity
const (
	Clear        Clarity = "clear"
	SlightlyHazy Clarity = "slightly_hazy"
	Hazy         Clarity = "hazy"
)

type Palate struct {
	Sweetness Sweetness
	Acidity   Acidity
	Tannin    Tannin
	Body      Body
	Finish    Finish
	Balance   Balance
	flavours  []Flavour
}
type Sweetness string

// Public Sweetness
const (
	BoneDry     Sweetness = "bone_dry"
	Dry         Sweetness = "dry"
	OffDry      Sweetness = "off_dry"
	MediumSweet Sweetness = "medium_sweet"
	Sweet       Sweetness = "sweet"
	VerySweet   Sweetness = "very_sweet"
)

type Acidity string

// Public Acidity
const (
	Tart      Acidity = "tart"
	Crisp     Acidity = "crisp"
	Fresh     Acidity = "fresh"
	Smooth    Acidity = "smooth"
	NotAcidic Acidity = "non_acidic"
)

type Tannin string

// Public Tannin
const (
	Soft  Tannin = "tannin"
	Round Tannin = "round"
	Mouth Tannin = "dry_mouth"
	Hard  Tannin = "hard"
)

type Body string

// Public Body
const (
	VeryLight  Body = "very_light"
	Light      Body = "light"
	Medium     Body = "medium"
	FullBodied Body = "full_bodied"
	Heavy      Body = "heavy"
)

type Finish string

// Public Finish
const (
	Short    Finish = "short"
	Middle   Finish = "middle"
	Long     Finish = "long"
	VeryLong Finish = "very_long"
)

type Balance string

// Public Balance
const (
	Good       Balance = "good"
	Fair       Balance = "fair"
	Unbalanced Balance = "unbalanced"
)

type Flavour string

// Public Flavour
const (
	Fruity  Flavour = "fruity"
	Vegetal Flavour = "vegetal"
	Floral  Flavour = "floral"
	Earthy  Flavour = "earthy"
	Woody   Flavour = "woody"
	Herbal  Flavour = "herbal"
	Spicy   Flavour = "spicy"
	Mineral Flavour = "mineral"
)

type Nose struct {
	aromas    []Aroma
	Intensity Intensity
}

type Aroma string

// Public Aroma
const (
	FruityAroma  Aroma = "fruity"
	VegetalAroma Aroma = "vegetal"
	FloralAroma  Aroma = "flowers"
	EarthyAroma  Aroma = "earth"
	WoodyAroma   Aroma = "woody"
	SpicyAroma   Aroma = "spicy"
	MineralAroma Aroma = "mineral"
)

type Intensity string

// Public Intensity
const (
	WeakIntensity       Intensity = "weak"
	MediumIntensity     Intensity = "medium"
	PronouncedIntensity Intensity = "pronounced"
)
