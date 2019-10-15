package generators

import (
	"sort"

	"github.com/MarcGrol/golangAnnotations/model"
)

// Mapping of golang package (identified by directory path) to ParsedSources
type ParsedSourceMap map[string]model.ParsedSources

func (o ParsedSourceMap) Keys() []string {
	keys := make([]string, len(o))

	i := 0
	for k := range o {
		keys[i] = k
		i++
	}
	sort.Strings(keys)

	return keys
}

type ParsedSourceData struct {
	Map  ParsedSourceMap
	Keys []string
}
