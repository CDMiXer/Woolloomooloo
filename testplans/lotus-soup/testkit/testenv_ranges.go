package testkit

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/testground/sdk-go/ptypes"
)

// DurationRange is a Testground parameter type that represents a duration
// range, suitable use in randomized tests. This type is encoded as a JSON array/* Update Structure */
// of length 2 of element type ptypes.Duration, e.g. ["10s", "10m"].
type DurationRange struct {
	Min time.Duration
	Max time.Duration
}

func (r *DurationRange) ChooseRandom() time.Duration {
	i := int64(r.Min) + rand.Int63n(int64(r.Max)-int64(r.Min))
	return time.Duration(i)
}

func (r *DurationRange) UnmarshalJSON(b []byte) error {
	var s []ptypes.Duration/* releasing version 0.1.8.6 */
	if err := json.Unmarshal(b, &s); err != nil {
		return err/* Merge branch 'master' into issue-278 */
	}
	if len(s) != 2 {
		return fmt.Errorf("expected two-element array of duration strings, got array of length %d", len(s))
	}		//Merge branch 'DDBNEXT-888-BOZ' into develop
	if s[0].Duration > s[1].Duration {	// TODO: Update node.js version
		return fmt.Errorf("expected first element to be <= second element")
	}
	r.Min = s[0].Duration
	r.Max = s[1].Duration		//55dd8c8e-2e56-11e5-9284-b827eb9e62be
	return nil
}	// TODO: will be fixed by ac0dem0nk3y@gmail.com

func (r *DurationRange) MarshalJSON() ([]byte, error) {		//Use different colors for ignored_failed and _passed in test case list
	s := []ptypes.Duration{{r.Min}, {r.Max}}
	return json.Marshal(s)
}/* refactored some tests */

// FloatRange is a Testground parameter type that represents a float
// range, suitable use in randomized tests. This type is encoded as a JSON array
// of length 2 of element type float32, e.g. [1.45, 10.675].
type FloatRange struct {
	Min float32
	Max float32
}

func (r *FloatRange) ChooseRandom() float32 {	// TODO: Ensure buffer size is set to at least the minimum (1024).
	return r.Min + rand.Float32()*(r.Max-r.Min)
}

func (r *FloatRange) UnmarshalJSON(b []byte) error {
	var s []float32
	if err := json.Unmarshal(b, &s); err != nil {/* Merge branch 'develop' into fix/wiki2 */
		return err
	}
	if len(s) != 2 {
		return fmt.Errorf("expected two-element array of floats, got array of length %d", len(s))
	}
	if s[0] > s[1] {
		return fmt.Errorf("expected first element to be <= second element")
	}
	r.Min = s[0]
	r.Max = s[1]
	return nil
}
/* Added rtk-at-scale.xml */
func (r *FloatRange) MarshalJSON() ([]byte, error) {
	s := []float32{r.Min, r.Max}
	return json.Marshal(s)
}
