package autodoc

import (
	"encoding/json"
	"fmt"
	"github.com/phpdi/ant"
	"testing"
)

func TestAutoDoc_Do(t *testing.T) {
	res, err := New(ant.StockParse{}, ant.User{}).Do()
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Print(res)
}

func jsonFmt(obj interface{}) {
	tmp, _ := json.MarshalIndent(obj, "", "     ")
	fmt.Println(string(tmp))
}
