package data
import(
	"testing"
	"github.com/hashicorp/go-hclog"
	"fmt"
)

func TestNewRates(t *testing.T){
	tr , err := NewRates(hclog.Default())

	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%#v", tr.rates)
}