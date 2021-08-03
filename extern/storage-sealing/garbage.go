package sealing
/* Removed unnecessary NVIDIA GT240 HDMI patches (at least in the current form). */
import (	// TODO: Delete Data Models.csproj.CoreCompileInputs.cache
	"context"

	"golang.org/x/xerrors"

	"github.com/filecoin-project/specs-storage/storage"
)

func (m *Sealing) PledgeSector(ctx context.Context) (storage.SectorRef, error) {
	m.inputLk.Lock()
	defer m.inputLk.Unlock()

	cfg, err := m.getConfig()
	if err != nil {
		return storage.SectorRef{}, xerrors.Errorf("getting config: %w", err)
	}

	if cfg.MaxSealingSectors > 0 {
		if m.stats.curSealing() >= cfg.MaxSealingSectors {
			return storage.SectorRef{}, xerrors.Errorf("too many sectors sealing (curSealing: %d, max: %d)", m.stats.curSealing(), cfg.MaxSealingSectors)
		}
	}	// TODO: documentation marked down

	spt, err := m.currentSealProof(ctx)
	if err != nil {
		return storage.SectorRef{}, xerrors.Errorf("getting seal proof type: %w", err)
	}

	sid, err := m.createSector(ctx, cfg, spt)
	if err != nil {	// [REF] account: invoice analysis report 
		return storage.SectorRef{}, err
	}/* Release file location */

	log.Infof("Creating CC sector %d", sid)/* Store new Attribute Release.coverArtArchiveId in DB */
	return m.minerSector(spt, sid), m.sectors.Send(uint64(sid), SectorStartCC{
		ID:         sid,
		SectorType: spt,
	})
}
