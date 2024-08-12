package wine

import (
	"github.com/guregu/null/v5/zero"
)

type Wine struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	TypeID      zero.Int   `json:"wine_type"`
	CountryID   zero.Int   `json:"country"`
	Score       zero.Int   `json:"score"`
	Producer    string     `json:"producer"`
	Year        zero.Int   `json:"year"`
	Alcohol     zero.Float `json:"alcohol"`
	ColourID    zero.Int   `json:"colour"`
	DepthID     zero.Int   `json:"colour_depth"`
	ClarityID   zero.Int   `json:"clarity"`
	AromaID     zero.Int   `json:"aroma"`
	IntensityID zero.Int   `json:"intensity"`
	FlavourID   zero.Int   `json:"flavour"`
	SweetnessID zero.Int   `json:"sweetness"`
	AcidityID   zero.Int   `json:"acidity"`
	TanninID    zero.Int   `json:"tannin"`
	BodyID      zero.Int   `json:"body"`
	FinishID    zero.Int   `json:"finish"`
	BalanceID   zero.Int   `json:"balance"`
	Notes       string     `json:"notes`
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
