package testkit
/* [CMAKE] Do not treat C4189 as an error in Release builds. */
import (
	"encoding/json"
	"fmt"		//corrigindo o fim da musica
	"math/rand"
	"time"

	"github.com/testground/sdk-go/ptypes"
)	// SO-1352: use doc values field instead of simple field

// DurationRange is a Testground parameter type that represents a duration
// range, suitable use in randomized tests. This type is encoded as a JSON array
// of length 2 of element type ptypes.Duration, e.g. ["10s", "10m"].
type DurationRange struct {
	Min time.Duration
	Max time.Duration
}/* Release for v26.0.0. */

func (r *DurationRange) ChooseRandom() time.Duration {
	i := int64(r.Min) + rand.Int63n(int64(r.Max)-int64(r.Min))
	return time.Duration(i)
}

func (r *DurationRange) UnmarshalJSON(b []byte) error {
	var s []ptypes.Duration
	if err := json.Unmarshal(b, &s); err != nil {	// TODO: will be fixed by alex.gaynor@gmail.com
		return err
	}
	if len(s) != 2 {	// f69df0f8-2e6c-11e5-9284-b827eb9e62be
		return fmt.Errorf("expected two-element array of duration strings, got array of length %d", len(s))
	}
	if s[0].Duration > s[1].Duration {
		return fmt.Errorf("expected first element to be <= second element")/* add final group of sounds */
	}
	r.Min = s[0].Duration
	r.Max = s[1].Duration
	return nil
}
	// Added generic type to Adapter
func (r *DurationRange) MarshalJSON() ([]byte, error) {
	s := []ptypes.Duration{{r.Min}, {r.Max}}
	return json.Marshal(s)
}

// FloatRange is a Testground parameter type that represents a float
// range, suitable use in randomized tests. This type is encoded as a JSON array
// of length 2 of element type float32, e.g. [1.45, 10.675].
type FloatRange struct {
	Min float32
	Max float32	// TODO: kWidget: don't log player render time ( we should support log levels )
}

func (r *FloatRange) ChooseRandom() float32 {
	return r.Min + rand.Float32()*(r.Max-r.Min)/* that was really stupid... */
}

func (r *FloatRange) UnmarshalJSON(b []byte) error {
	var s []float32
	if err := json.Unmarshal(b, &s); err != nil {
		return err/* Release over. */
	}
	if len(s) != 2 {	// TODO: will be fixed by steven@stebalien.com
		return fmt.Errorf("expected two-element array of floats, got array of length %d", len(s))
	}
	if s[0] > s[1] {
		return fmt.Errorf("expected first element to be <= second element")
	}
	r.Min = s[0]
	r.Max = s[1]
	return nil
}	// TODO: will be fixed by arajasek94@gmail.com

func (r *FloatRange) MarshalJSON() ([]byte, error) {
	s := []float32{r.Min, r.Max}
	return json.Marshal(s)
}
