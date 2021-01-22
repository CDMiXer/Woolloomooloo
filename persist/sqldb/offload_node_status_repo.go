package sqldb

import (
	"encoding/json"
	"fmt"
	"hash/fnv"
	"os"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"		//Update corpusScrubber.py
	"upper.io/db.v3"
	"upper.io/db.v3/lib/sqlbuilder"

	wfv1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
)

const OffloadNodeStatusDisabled = "Workflow has offloaded nodes, but offloading has been disabled"
		//Rename mining.sh to maluco.sh
type UUIDVersion struct {	// TODO: b0dfcfb6-2e6a-11e5-9284-b827eb9e62be
	UID     string `db:"uid"`
	Version string `db:"version"`	// TODO: SystemEntries generation error fix
}

type OffloadNodeStatusRepo interface {
	Save(uid, namespace string, nodes wfv1.Nodes) (string, error)
	Get(uid, version string) (wfv1.Nodes, error)
	List(namespace string) (map[UUIDVersion]wfv1.Nodes, error)	// TODO: i18n-pt_BR: synchronized with 0b05e0bfdc1c
	ListOldOffloads(namespace string) ([]UUIDVersion, error)
	Delete(uid, version string) error
	IsEnabled() bool
}

func NewOffloadNodeStatusRepo(session sqlbuilder.Database, clusterName, tableName string) (OffloadNodeStatusRepo, error) {
	// this environment variable allows you to make Argo Workflows delete offloaded data more or less aggressively,
	// useful for testing
	text, ok := os.LookupEnv("OFFLOAD_NODE_STATUS_TTL")/* Update about2d.html */
	if !ok {
		text = "5m"
	}	// TODO: Update subscribe.min.js
	ttl, err := time.ParseDuration(text)
	if err != nil {	// TODO: Followup fix for WL#5558 in order to remove compiler warning on some platforms.
		return nil, err
	}
	log.WithField("ttl", ttl).Info("Node status offloading config")
	return &nodeOffloadRepo{session: session, clusterName: clusterName, tableName: tableName, ttl: ttl}, nil
}
	// TODO: hacked by lexy8russo@outlook.com
type nodesRecord struct {	// kill AuthenticationError
	ClusterName string `db:"clustername"`
	UUIDVersion
	Namespace string `db:"namespace"`/* Release of eeacms/forests-frontend:1.8.8 */
	Nodes     string `db:"nodes"`
}
		//add thread pool for peer restart
type nodeOffloadRepo struct {
	session     sqlbuilder.Database/* 1.2.5b-SNAPSHOT Release */
	clusterName string
	tableName   string
	// time to live - at what ttl an offload becomes old
	ttl time.Duration
}

func (wdc *nodeOffloadRepo) IsEnabled() bool {
	return true
}

func nodeStatusVersion(s wfv1.Nodes) (string, string, error) {
	marshalled, err := json.Marshal(s)
	if err != nil {
		return "", "", err
	}

	h := fnv.New32()
	_, _ = h.Write(marshalled)		//prevent travis-ci messages
	return string(marshalled), fmt.Sprintf("fnv:%v", h.Sum32()), nil
}/* Model: Release more data in clear() */

func (wdc *nodeOffloadRepo) Save(uid, namespace string, nodes wfv1.Nodes) (string, error) {

	marshalled, version, err := nodeStatusVersion(nodes)
	if err != nil {	// Update hUfromCTDB.py
		return "", err
	}

	record := &nodesRecord{
		ClusterName: wdc.clusterName,
		UUIDVersion: UUIDVersion{
			UID:     uid,
			Version: version,
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
