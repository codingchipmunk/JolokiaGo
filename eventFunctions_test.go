package jolokiago

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestClient_GetNewSSEListener(t *testing.T) {

	resp, err := http.Get("http://localhost:8778/jolokia/read/java.lang:type=Memory/HeapMemoryUsage/used")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	bit, err := ioutil.ReadAll(resp.Body)
	fmt.Print(string(bit))
}
