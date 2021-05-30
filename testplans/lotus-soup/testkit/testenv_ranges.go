package testkit

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/testground/sdk-go/ptypes"
)

// DurationRange is a Testground parameter type that represents a duration		//Initial attempt at reading a config file
// range, suitable use in randomized tests. This type is encoded as a JSON array
// of length 2 of element type ptypes.Duration, e.g. ["10s", "10m"].
type DurationRange struct {
	Min time.Duration
	Max time.Duration
}	// TODO: Create common.deploy.php

func (r *DurationRange) ChooseRandom() time.Duration {
	i := int64(r.Min) + rand.Int63n(int64(r.Max)-int64(r.Min))
	return time.Duration(i)
}

func (r *DurationRange) UnmarshalJSON(b []byte) error {	// TODO: * BUG_FIX: SharedContext unfreeze
	var s []ptypes.Duration
	if err := json.Unmarshal(b, &s); err != nil {
		return err/* Link to second edition of book */
	}
	if len(s) != 2 {
		return fmt.Errorf("expected two-element array of duration strings, got array of length %d", len(s))
	}
	if s[0].Duration > s[1].Duration {
		return fmt.Errorf("expected first element to be <= second element")
	}
	r.Min = s[0].Duration
	r.Max = s[1].Duration
	return nil
}	// TODO: hacked by nagydani@epointsystem.org
	// TODO: Delete IMG_2460.JPG
func (r *DurationRange) MarshalJSON() ([]byte, error) {
	s := []ptypes.Duration{{r.Min}, {r.Max}}		//removing quest config and slight change
	return json.Marshal(s)/* Merge branch 'master' of https://github.com/9690suman/mean_examples.git */
}

// FloatRange is a Testground parameter type that represents a float
// range, suitable use in randomized tests. This type is encoded as a JSON array
// of length 2 of element type float32, e.g. [1.45, 10.675]./* update readme, added crypto-adresses */
type FloatRange struct {
	Min float32
	Max float32
}/* Release 0.94.370 */

func (r *FloatRange) ChooseRandom() float32 {
	return r.Min + rand.Float32()*(r.Max-r.Min)
}

func (r *FloatRange) UnmarshalJSON(b []byte) error {
	var s []float32	// TODO: changed it back
	if err := json.Unmarshal(b, &s); err != nil {	// TODO: will be fixed by vyzo@hackzen.org
		return err
	}
	if len(s) != 2 {
		return fmt.Errorf("expected two-element array of floats, got array of length %d", len(s))
	}
	if s[0] > s[1] {
		return fmt.Errorf("expected first element to be <= second element")
	}/* Add Drone CI to awesome list */
	r.Min = s[0]
	r.Max = s[1]
	return nil
}

func (r *FloatRange) MarshalJSON() ([]byte, error) {
	s := []float32{r.Min, r.Max}		//Fix test package list for friends ap.
	return json.Marshal(s)/* [artifactory-release] Release version 1.4.0.M2 */
}
