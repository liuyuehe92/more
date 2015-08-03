package more

import . "github.com/pkg4go/assert"
import "testing"

func TestNew(t *testing.T) {
	a := A{t}

	opts := map[string]interface{}{
		"dir":   "view",
		"ext":   "html",
		"cache": true,
	}

	a.NotNil(New(opts))
}
