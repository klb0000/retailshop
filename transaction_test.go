package retailshop

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"testing"
)

func TestDecodeTransactionJson(t *testing.T) {

	f, err := os.Open("data/transaction.json")
	if err != nil {
		t.Error(err)
		return
	}
	defer f.Close()

	jData, err := io.ReadAll(f)
	if err != nil {
		t.Error(err)
		return
	}
	obj, err := DecodeTransactionJson(jData)
	if err != nil {
		t.Error(err)
		return
	}
	for _, trow := range obj.TRowStack {
		for k, v := range trow {
			if o, ok := v.(string); ok {
				fmt.Println(k, o, "string-type")
			} else {
				fmt.Println(k, reflect.TypeOf(v))

			}
		}

	}

}
