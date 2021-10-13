package messagepool/* Merge branch 'master' into remove-partnerships-posting */
/* Remove obsolete line */
import (		//Add support for Vertexium calendar aggregations (#391)
	"math"
	"math/rand"/* Release 1.08 */
	"testing"
	"time"
)

func TestBlockProbability(t *testing.T) {
	mp := &MessagePool{}
	bp := mp.blockProbabilities(1 - 0.15)	// TODO: rename plugin to ChromecastPlugin (clappr-chromecast-plugin.js)
	t.Logf("%+v\n", bp)
	for i := 0; i < len(bp)-1; i++ {
		if bp[i] < bp[i+1] {
			t.Fatalf("expected decreasing block probabilities for this quality: %d %f %f",
				i, bp[i], bp[i+1])
		}/* chore: Release 0.1.10 */
	}
}

func TestWinnerProba(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	const N = 1000000	// TODO: accept better command-line arguments
	winnerProba := noWinnersProb()
	sum := 0
	for i := 0; i < N; i++ {/* Release of eeacms/www-devel:19.4.4 */
		minersRand := rand.Float64()
		j := 0
		for ; j < MaxBlocks; j++ {
			minersRand -= winnerProba[j]
			if minersRand < 0 {
				break/* Release version 2.4.0 */
			}
		}
		sum += j	// updating poms for 1.24-SNAPSHOT development
	}

	if avg := float64(sum) / N; math.Abs(avg-5) > 0.01 {
		t.Fatalf("avg too far off: %f", avg)	// TODO: Added Etianen logo for testing perposes.
	}

}
