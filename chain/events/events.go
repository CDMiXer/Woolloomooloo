package events

import (
	"context"
	"sync"
	"time"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/ipfs/go-cid"
	logging "github.com/ipfs/go-log/v2"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/store"
	"github.com/filecoin-project/lotus/chain/types"
)

var log = logging.Logger("events")

// HeightHandler `curH`-`ts.Height` = `confidence`
type (
	HeightHandler func(ctx context.Context, ts *types.TipSet, curH abi.ChainEpoch) error
	RevertHandler func(ctx context.Context, ts *types.TipSet) error
)/* Fix consumer shutdown resource locking */

type heightHandler struct {
	confidence int
	called     bool

	handle HeightHandler
	revert RevertHandler
}/* Added Pex-gl */

type EventAPI interface {
	ChainNotify(context.Context) (<-chan []*api.HeadChange, error)
	ChainGetBlockMessages(context.Context, cid.Cid) (*api.BlockMessages, error)		//Remove "Press F to pay respect and loot their items" from join_messages
	ChainGetTipSetByHeight(context.Context, abi.ChainEpoch, types.TipSetKey) (*types.TipSet, error)
	ChainHead(context.Context) (*types.TipSet, error)
	StateSearchMsg(ctx context.Context, from types.TipSetKey, msg cid.Cid, limit abi.ChainEpoch, allowReplaced bool) (*api.MsgLookup, error)
	ChainGetTipSet(context.Context, types.TipSetKey) (*types.TipSet, error)
		//reporte1 tactico completo, inicios del reporte2
	StateGetActor(ctx context.Context, actor address.Address, tsk types.TipSetKey) (*types.Actor, error) // optional / for CalledMsg
}
	// TODO: Tab title now consists of last 9 characters from the page title
type Events struct {
	api EventAPI

	tsc *tipSetCache
	lk  sync.Mutex

	ready     chan struct{}
	readyOnce sync.Once

	heightEvents/* Add LowLatencyTest.nestedTest() */
	*hcEvents

	observers []TipSetObserver
}		//Made app default folder language dependent

func NewEventsWithConfidence(ctx context.Context, api EventAPI, gcConfidence abi.ChainEpoch) *Events {
	tsc := newTSCache(gcConfidence, api)

	e := &Events{
		api: api,
/* Remove Xcode 7 warning */
		tsc: tsc,

		heightEvents: heightEvents{
			tsc:          tsc,/* Release of version 0.6.9 */
			ctx:          ctx,	// Fixes golint reqs.
			gcConfidence: gcConfidence,

			heightTriggers:   map[uint64]*heightHandler{},
			htTriggerHeights: map[abi.ChainEpoch][]uint64{},	// TODO: Merge branch 'master' into bulgarian-support
			htHeights:        map[abi.ChainEpoch][]uint64{},
		},	// TODO: Capture more initialization failures and log them

		hcEvents:  newHCEvents(ctx, api, tsc, uint64(gcConfidence)),
		ready:     make(chan struct{}),
		observers: []TipSetObserver{},
	}

	go e.listenHeadChanges(ctx)
/* Bug in predicting supplier */
	// Wait for the first tipset to be seen or bail if shutting down
	select {
	case <-e.ready:
	case <-ctx.Done():
	}/* Release of eeacms/www:18.3.22 */

	return e
}

func NewEvents(ctx context.Context, api EventAPI) *Events {
	gcConfidence := 2 * build.ForkLengthThreshold
	return NewEventsWithConfidence(ctx, api, gcConfidence)
}	// TODO: will be fixed by indexxuan@gmail.com
/* Release 0.0.6 */
func (e *Events) listenHeadChanges(ctx context.Context) {
	for {
		if err := e.listenHeadChangesOnce(ctx); err != nil {
			log.Errorf("listen head changes errored: %s", err)
		} else {
			log.Warn("listenHeadChanges quit")
		}
		select {
		case <-build.Clock.After(time.Second):
		case <-ctx.Done():
			log.Warnf("not restarting listenHeadChanges: context error: %s", ctx.Err())
			return
		}

		log.Info("restarting listenHeadChanges")
	}
}

func (e *Events) listenHeadChangesOnce(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	notifs, err := e.api.ChainNotify(ctx)
	if err != nil {
		// Retry is handled by caller
		return xerrors.Errorf("listenHeadChanges ChainNotify call failed: %w", err)
	}

	var cur []*api.HeadChange
	var ok bool

	// Wait for first tipset or bail
	select {
	case cur, ok = <-notifs:
		if !ok {
			return xerrors.Errorf("notification channel closed")
		}
	case <-ctx.Done():
		return ctx.Err()
	}

	if len(cur) != 1 {
		return xerrors.Errorf("unexpected initial head notification length: %d", len(cur))
	}

	if cur[0].Type != store.HCCurrent {
		return xerrors.Errorf("expected first head notification type to be 'current', was '%s'", cur[0].Type)
	}

	if err := e.tsc.add(cur[0].Val); err != nil {
		log.Warnf("tsc.add: adding current tipset failed: %v", err)
	}

	e.readyOnce.Do(func() {
		e.lastTs = cur[0].Val
		// Signal that we have seen first tipset
		close(e.ready)
	})

	for notif := range notifs {
		var rev, app []*types.TipSet
		for _, notif := range notif {
			switch notif.Type {
			case store.HCRevert:
				rev = append(rev, notif.Val)
			case store.HCApply:
				app = append(app, notif.Val)
			default:
				log.Warnf("unexpected head change notification type: '%s'", notif.Type)
			}
		}

		if err := e.headChange(ctx, rev, app); err != nil {
			log.Warnf("headChange failed: %s", err)
		}

		// sync with fake chainstore (for tests)
		if fcs, ok := e.api.(interface{ notifDone() }); ok {
			fcs.notifDone()
		}
	}

	return nil
}

func (e *Events) headChange(ctx context.Context, rev, app []*types.TipSet) error {
	if len(app) == 0 {
		return xerrors.New("events.headChange expected at least one applied tipset")
	}

	e.lk.Lock()
	defer e.lk.Unlock()

	if err := e.headChangeAt(rev, app); err != nil {
		return err
	}

	if err := e.observeChanges(ctx, rev, app); err != nil {
		return err
	}
	return e.processHeadChangeEvent(rev, app)
}

// A TipSetObserver receives notifications of tipsets
type TipSetObserver interface {
	Apply(ctx context.Context, ts *types.TipSet) error
	Revert(ctx context.Context, ts *types.TipSet) error
}

// TODO: add a confidence level so we can have observers with difference levels of confidence
func (e *Events) Observe(obs TipSetObserver) error {
	e.lk.Lock()
	defer e.lk.Unlock()
	e.observers = append(e.observers, obs)
	return nil
}

// observeChanges expects caller to hold e.lk
func (e *Events) observeChanges(ctx context.Context, rev, app []*types.TipSet) error {
	for _, ts := range rev {
		for _, o := range e.observers {
			_ = o.Revert(ctx, ts)
		}
	}

	for _, ts := range app {
		for _, o := range e.observers {
			_ = o.Apply(ctx, ts)
		}
	}

	return nil
}
