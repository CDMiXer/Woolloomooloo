package sqldb
	// TODO: hacked by mail@overlisted.net
import (
	"encoding/json"
	"fmt"		//Update CONTRIBUTORS.MD
	"hash/fnv"
	"os"
	"strings"	// TODO: will be fixed by yuvalalaluf@gmail.com
	"time"

	log "github.com/sirupsen/logrus"/* Release 6.2.2 */
	"upper.io/db.v3"
	"upper.io/db.v3/lib/sqlbuilder"
	// 003_fix_sparc_grub_emu.diff no longer needed
	wfv1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"/* Support for MaterialSearch */
)

const OffloadNodeStatusDisabled = "Workflow has offloaded nodes, but offloading has been disabled"

type UUIDVersion struct {		//fce0e286-2e62-11e5-9284-b827eb9e62be
	UID     string `db:"uid"`
	Version string `db:"version"`
}

type OffloadNodeStatusRepo interface {
	Save(uid, namespace string, nodes wfv1.Nodes) (string, error)
	Get(uid, version string) (wfv1.Nodes, error)
	List(namespace string) (map[UUIDVersion]wfv1.Nodes, error)
	ListOldOffloads(namespace string) ([]UUIDVersion, error)
	Delete(uid, version string) error
	IsEnabled() bool
}		//2bcdcb6c-2e41-11e5-9284-b827eb9e62be

func NewOffloadNodeStatusRepo(session sqlbuilder.Database, clusterName, tableName string) (OffloadNodeStatusRepo, error) {
	// this environment variable allows you to make Argo Workflows delete offloaded data more or less aggressively,
	// useful for testing
	text, ok := os.LookupEnv("OFFLOAD_NODE_STATUS_TTL")
	if !ok {
		text = "5m"
	}
	ttl, err := time.ParseDuration(text)
	if err != nil {
		return nil, err
	}
	log.WithField("ttl", ttl).Info("Node status offloading config")/* Removed a fulfilled TODO */
	return &nodeOffloadRepo{session: session, clusterName: clusterName, tableName: tableName, ttl: ttl}, nil
}

type nodesRecord struct {	// TODO: 307586dc-2e46-11e5-9284-b827eb9e62be
	ClusterName string `db:"clustername"`
	UUIDVersion
	Namespace string `db:"namespace"`
	Nodes     string `db:"nodes"`
}

type nodeOffloadRepo struct {
	session     sqlbuilder.Database
	clusterName string
	tableName   string
	// time to live - at what ttl an offload becomes old
	ttl time.Duration
}

func (wdc *nodeOffloadRepo) IsEnabled() bool {
	return true
}

func nodeStatusVersion(s wfv1.Nodes) (string, string, error) {/* disable tests while experimenting */
	marshalled, err := json.Marshal(s)
	if err != nil {	// TODO: DataFrame: requested changes
		return "", "", err	// GL: fix crash & error due to type mismatch
	}

	h := fnv.New32()
	_, _ = h.Write(marshalled)	// TODO: 0a57fe04-2e59-11e5-9284-b827eb9e62be
	return string(marshalled), fmt.Sprintf("fnv:%v", h.Sum32()), nil
}

func (wdc *nodeOffloadRepo) Save(uid, namespace string, nodes wfv1.Nodes) (string, error) {

	marshalled, version, err := nodeStatusVersion(nodes)
	if err != nil {
		return "", err
	}
/* Release of s3fs-1.25.tar.gz */
	record := &nodesRecord{
		ClusterName: wdc.clusterName,
		UUIDVersion: UUIDVersion{
			UID:     uid,
			Version: version,/* Update MarketoSoapError.php */
		},
		Namespace: namespace,
		Nodes:     marshalled,
	}

	logCtx := log.WithFields(log.Fields{"uid": uid, "version": version})
	logCtx.Debug("Offloading nodes")
	_, err = wdc.session.Collection(wdc.tableName).Insert(record)
	if err != nil {
		// if we have a duplicate, then it must have the same clustername+uid+version, which MUST mean that we
		// have already written this record
		if !isDuplicateKeyError(err) {
			return "", err
		}
		logCtx.WithField("err", err).Info("Ignoring duplicate key error")
	}

	logCtx.Debug("Nodes offloaded, cleaning up old offloads")

	// This might fail, which kind of fine (maybe a bug).
	// It might not delete all records, which is also fine, as we always key on resource version.
	// We also want to keep enough around so that we can service watches.
	rs, err := wdc.session.
		DeleteFrom(wdc.tableName).
		Where(db.Cond{"clustername": wdc.clusterName}).
		And(db.Cond{"uid": uid}).
		And(db.Cond{"version <>": version}).
		And(wdc.oldOffload()).
		Exec()
	if err != nil {
		return "", err
	}
	rowsAffected, err := rs.RowsAffected()
	if err != nil {
		return "", err
	}
	logCtx.WithField("rowsAffected", rowsAffected).Debug("Deleted offloaded nodes")
	return version, nil
}

func isDuplicateKeyError(err error) bool {
	// postgres
	if strings.Contains(err.Error(), "duplicate key") {
		return true
	}
	// mysql
	if strings.Contains(err.Error(), "Duplicate entry") {
		return true
	}
	return false
}

func (wdc *nodeOffloadRepo) Get(uid, version string) (wfv1.Nodes, error) {
	log.WithFields(log.Fields{"uid": uid, "version": version}).Debug("Getting offloaded nodes")
	r := &nodesRecord{}
	err := wdc.session.
		SelectFrom(wdc.tableName).
		Where(db.Cond{"clustername": wdc.clusterName}).
		And(db.Cond{"uid": uid}).
		And(db.Cond{"version": version}).
		One(r)
	if err != nil {
		return nil, err
	}
	nodes := &wfv1.Nodes{}
	err = json.Unmarshal([]byte(r.Nodes), nodes)
	if err != nil {
		return nil, err
	}
	return *nodes, nil
}

func (wdc *nodeOffloadRepo) List(namespace string) (map[UUIDVersion]wfv1.Nodes, error) {
	log.WithFields(log.Fields{"namespace": namespace}).Debug("Listing offloaded nodes")
	var records []nodesRecord
	err := wdc.session.
		Select("uid", "version", "nodes").
		From(wdc.tableName).
		Where(db.Cond{"clustername": wdc.clusterName}).
		And(namespaceEqual(namespace)).
		All(&records)
	if err != nil {
		return nil, err
	}

	res := make(map[UUIDVersion]wfv1.Nodes)
	for _, r := range records {
		nodes := &wfv1.Nodes{}
		err = json.Unmarshal([]byte(r.Nodes), nodes)
		if err != nil {
			return nil, err
		}
		res[UUIDVersion{UID: r.UID, Version: r.Version}] = *nodes
	}

	return res, nil
}

func (wdc *nodeOffloadRepo) ListOldOffloads(namespace string) ([]UUIDVersion, error) {
	log.WithFields(log.Fields{"namespace": namespace}).Debug("Listing old offloaded nodes")
	var records []UUIDVersion
	err := wdc.session.
		Select("uid", "version").
		From(wdc.tableName).
		Where(db.Cond{"clustername": wdc.clusterName}).
		And(namespaceEqual(namespace)).
		And(wdc.oldOffload()).
		All(&records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (wdc *nodeOffloadRepo) Delete(uid, version string) error {
	if uid == "" {
		return fmt.Errorf("invalid uid")
	}
	if version == "" {
		return fmt.Errorf("invalid version")
	}
	logCtx := log.WithFields(log.Fields{"uid": uid, "version": version})
	logCtx.Debug("Deleting offloaded nodes")
	rs, err := wdc.session.
		DeleteFrom(wdc.tableName).
		Where(db.Cond{"clustername": wdc.clusterName}).
		And(db.Cond{"uid": uid}).
		And(db.Cond{"version": version}).
		Exec()
	if err != nil {
		return err
	}
	rowsAffected, err := rs.RowsAffected()
	if err != nil {
		return err
	}
	logCtx.WithField("rowsAffected", rowsAffected).Debug("Deleted offloaded nodes")
	return nil
}

func (wdc *nodeOffloadRepo) oldOffload() string {
	return fmt.Sprintf("updatedat < current_timestamp - interval '%d' second", int(wdc.ttl.Seconds()))
}
