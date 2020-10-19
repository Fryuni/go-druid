package filter

import (
	"encoding/json"

	"github.com/grafadruid/go-druid/query"
	"github.com/grafadruid/go-druid/query/extractionfn"
)

type Extraction struct {
	Base
	Dimension    string             `json:"dimension"`
	Value        string             `json:"value"`
	ExtractionFn query.ExtractionFn `json:"extractionFn"`
}

func NewExtraction() *Extraction {
	e := &Extraction{}
	e.SetType("extraction")
	return e
}

func (e *Extraction) SetDimension(dimension string) *Extraction {
	e.Dimension = dimension
	return e
}

func (e *Extraction) SetValue(value string) *Extraction {
	e.Value = value
	return e
}

func (e *Extraction) SetExtractionFn(extractionFn query.ExtractionFn) *Extraction {
	e.ExtractionFn = extractionFn
	return e
}

func (e *Extraction) UnmarshalJSON(data []byte) error {
	var tmp struct {
		Base
		Dimension    string          `json:"dimension"`
		Value        string          `json:"value"`
		ExtractionFn json.RawMessage `json:"extractionFn"`
	}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	ex, err := extractionfn.Load(tmp.ExtractionFn)
	if err != nil {
		return err
	}
	e.Base = tmp.Base
	e.Dimension = tmp.Dimension
	e.Value = tmp.Value
	e.ExtractionFn = ex
	return nil
}
