package wine

import (
	"github.com/guregu/null/v5/zero"
)

type SafeInt struct {
	zero.Int
}

type SafeFloat struct {
	zero.Float
}

func (i *SafeInt) UnmarshalJSON(b []byte) error {
	if string(b) == `""` {
		i.Int = zero.IntFrom(0)
		return nil
	}
	return i.Int.UnmarshalJSON(b)
}

func (i *SafeFloat) UnmarshalJSON(b []byte) error {
	if string(b) == `""` {
		i.Float = zero.FloatFrom(0)
		return nil
	}
	return i.Float.UnmarshalJSON(b)
}

type Wine struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	TypeID      SafeInt   `json:"wine_type,omitempty"`
	CountryID   SafeInt   `json:"country,omitempty"`
	GrapeIDs    []SafeInt `json:"grapes,omitempty"`
	Score       SafeInt   `json:"score,omitempty"`
	Producer    string    `json:"producer,omitempty"`
	Year        SafeInt   `json:"year,omitempty"`
	Alcohol     SafeFloat `json:"alcohol,omitempty"`
	ColourID    SafeInt   `json:"colour,omitempty"`
	ClarityID   SafeInt   `json:"clarity,omitempty"`
	AromaID     SafeInt   `json:"aroma,omitempty"`
	IntensityID SafeInt   `json:"intensity,omitempty"`
	FlavourID   SafeInt   `json:"flavour,omitempty"`
	SweetnessID SafeInt   `json:"sweetness,omitempty"`
	AcidityID   SafeInt   `json:"acidity,omitempty"`
	TanninID    SafeInt   `json:"tannin,omitempty"`
	BodyID      SafeInt   `json:"body,omitempty"`
	FinishID    SafeInt   `json:"finish,omitempty"`
	BalanceID   SafeInt   `json:"balance,omitempty"`
	Notes       string    `json:"notes,omitempty"`
	ImageURL    string    `json:"image_url,omitempty"`
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

type Balance struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Colour struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Grapes struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
