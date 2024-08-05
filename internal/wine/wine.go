package wine

import (
	"time"

	"github.com/guregu/null/v5/zero"
)

type Wine struct {
	ID          int       `json:"id"`
	TypeID      zero.Int  `json:"wine_type"`
	Name        string    `json:"name"`
	CountryID   zero.Int  `json:"country"`
	Score       zero.Int  `json:"score"`
	Producer    string    `json:"producer"`
	Year        time.Time `json:"year"`
	Alcohol     float64   `json:"alcohol"`
	AromaID     zero.Int  `json:"aroma"`
	IntensityID zero.Int  `json:"intensity"`
	FlavourID   zero.Int  `json:"flavour"`
	SweetnessID zero.Int  `json:"sweetness"`
	AcidityID   zero.Int  `json:"acidity"`
	TanninID    zero.Int  `json:"tannin"`
	BodyID      zero.Int  `json:"body"`
	ClarityID   zero.Int  `json:"clarity"`
	Finish      zero.Int  `json:"finish"`
	ColourID    zero.Int  `json:"colour"`
	DepthID     zero.Int  `json:"colour_depth"`
	BalanceID   zero.Int  `json:"balance"`
}
type WineType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Country struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Aroma struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Intensity struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Flavour struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Sweetness struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Acidity struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Tannin struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Body struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Clarity struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Finish struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Depth struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Balance struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
