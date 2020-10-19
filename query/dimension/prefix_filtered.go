package dimension

import (
	"encoding/json"

	"github.com/grafadruid/go-druid/query"
	"github.com/grafadruid/go-druid/query/types"
)

type PrefixFiltered struct {
	Base
	Delegate query.Dimension `json:"delegate"`
	Prefix   string          `json:"prefix"`
}

func NewPrefixFiltered() *PrefixFiltered {
	p := &PrefixFiltered{}
	p.SetType("prefixFiltered")
	return p
}

func (p *PrefixFiltered) SetDimension(dimension string) *PrefixFiltered {
	p.Base.SetDimension(dimension)
	return p
}

func (p *PrefixFiltered) SetOutputName(outputName string) *PrefixFiltered {
	p.Base.SetOutputName(outputName)
	return p
}

func (p *PrefixFiltered) SetOutputType(outputType types.OutputType) *PrefixFiltered {
	p.Base.SetOutputType(outputType)
	return p
}

func (p *PrefixFiltered) SetDelegate(delegate query.Dimension) *PrefixFiltered {
	p.Delegate = delegate
	return p
}

func (p *PrefixFiltered) SetPrefix(prefix string) *PrefixFiltered {
	p.Prefix = prefix
	return p
}

func (p *PrefixFiltered) UnmarshalJSON(data []byte) error {
	var tmp struct {
		Base
		Delegate json.RawMessage `json:"delegate"`
		Prefix   string          `json:"prefix"`
	}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	d, err := Load(tmp.Delegate)
	if err != nil {
		return err
	}
	p.Base = tmp.Base
	p.Delegate = d
	p.Prefix = tmp.Prefix
	return nil
}
