package more

import m2s "github.com/mitchellh/mapstructure"
import . "github.com/tj/go-debug"

var debug = Debug("more")

func New(options map[string]interface{}) *Render {
	var opts Opts

	err := m2s.Decode(options, &opts)

	if err != nil {
		panic(err)
	}

	debug("opts: %v", opts)

	r := new(Render)

	return r
}

type Opts struct {
	Dir   string
	Ext   string
	Cache bool
}

type Render struct {
}
