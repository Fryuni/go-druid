package postaggregation

import (
	"encoding/json"

	"github.com/grafadruid/go-druid/query"
)

type Arithmetic struct {
	Base
	Fn       string                 `json:"fn"`
	Fields   []query.PostAggregator `json:"fields"`
	Ordering string                 `json:"ordering"`
}

func NewArithmetic() *Arithmetic {
	a := &Arithmetic{}
	a.SetType("arithmetic")
	return a
}

func (a *Arithmetic) SetName(name string) *Arithmetic {
	a.Base.SetName(name)
	return a
}

func (a *Arithmetic) SetFn(fn string) *Arithmetic {
	a.Fn = fn
	return a
}

func (a *Arithmetic) SetFields(fields []query.PostAggregator) *Arithmetic {
	a.Fields = fields
	return a
}

func (a *Arithmetic) SetOrdering(ordering string) *Arithmetic {
	a.Ordering = ordering
	return a
}

func (a *Arithmetic) UnmarshalJSON(data []byte) error {
	var tmp struct {
		Base
		Fn       string            `json:"fn"`
		Fields   []json.RawMessage `json:"fields"`
		Ordering string            `json:"ordering"`
	}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	var err error
	var p query.PostAggregator
	pp := make([]query.PostAggregator, len(tmp.Fields))
	for i := range tmp.Fields {
		if p, err = Load(tmp.Fields[i]); err != nil {
			return err
		}
		pp[i] = p
	}
	a.Base = tmp.Base
	a.Fn = tmp.Fn
	a.Fields = pp
	a.Ordering = tmp.Ordering
	return nil
}
