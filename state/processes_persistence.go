// Copyright 2015 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package state

import (
	"fmt"

	"github.com/juju/errors"
	"github.com/juju/names"
	jujutxn "github.com/juju/txn"
	"gopkg.in/juju/charm.v5"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2/txn"

	"github.com/juju/juju/process"
)

// TODO(ericsnow) Implement persistence using a TXN abstraction (used
// in the business logic) with ops factories available from the
// persistence layer.

type processesPersistenceBase interface {
	getCollection(name string) (stateCollection, func())
	run(transactions jujutxn.TransactionSource) error
}

type processesPersistence struct {
	st    processesPersistenceBase
	charm names.CharmTag
	unit  names.UnitTag
}

func (pp processesPersistence) ensureDefinitions(definitions ...charm.Process) error {
	// Add definition if not already added (or ensure matches).
	var ops []txn.Op
	for _, definition := range definitions {
		ops = append(ops, pp.newInsertDefinitionOp(definition))
	}
	buildTxn := func(attempt int) ([]txn.Op, error) {
		if attempt > 0 {
			// The last attempt aborted so clear out any ops that failed
			// the DocMissing assertion and try again.
			coll, closeColl := pp.st.getCollection(workloadProcessesC)
			defer closeColl()
			var okOps []txn.Op
			for _, op := range ops {
				var doc processDefinitionDoc
				err := coll.FindId(op.Id).One(&doc)
				if err == mgo.ErrNotFound {
					okOps = append(okOps, op)
				} else if err != nil {
					return nil, errors.Trace(err)
				}
				// Otherwise the op is dropped.
			}
			if len(okOps) == 0 {
				return nil, jujutxn.ErrNoOperations
			}
			ops = okOps
		}
		return ops, nil
	}
	if err := pp.st.run(buildTxn); err != nil {
		return errors.Trace(err)
	}

	return nil
}

func (pp processesPersistence) insert(info process.Info) error {
	var ops []txn.Op
	// TODO(ericsnow) Add unitPersistence.newEnsureAliveOp(pp.unit)?
	// TODO(ericsnow) Add pp.newEnsureDefinitionOp(info.Process)?
	ops = append(ops, pp.newInsertProcessOps(info)...)
	buildTxn := func(attempt int) ([]txn.Op, error) {
		return ops, nil
	}
	if err := pp.st.run(buildTxn); err != nil {
		return errors.Trace(err)
	}
	return nil
}

func (pp processesPersistence) setStatus(id string, status process.Status) error {
	var ops []txn.Op
	ops = append(ops, pp.newSetRawStatusOp(id, status))
	buildTxn := func(attempt int) ([]txn.Op, error) {
		if attempt > 0 {
			proc, err := pp.proc(id)
			if err != nil {
				return nil, errors.Trace(err)
			}
			if proc.Life != Alive {
				return nil, jujutxn.ErrNoOperations
			}
			// Otherwise try again...
		}
		return ops, nil
	}
	if err := pp.st.run(buildTxn); err != nil {
		return errors.Trace(err)
	}
	return nil
}

func (pp processesPersistence) list(ids ...string) ([]process.Info, error) {
	// TODO(ericsnow) finish!
	return nil, errors.Errorf("not finished")
}

func (pp processesPersistence) remove(id string) error {
	// TODO(ericsnow) finish!
	return errors.Errorf("not finished")
}

func (pp processesPersistence) proc(id string) (*processDoc, error) {
	coll, closeColl := pp.st.getCollection(workloadProcessesC)
	defer closeColl()
	id = pp.processID(id)
	var doc processDoc
	if err := coll.FindId(id).One(&doc); err != nil {
		return nil, errors.Trace(err)
	}
	return &doc, nil
}

func (pp processesPersistence) definitionID(name string) string {
	// The URL will always parse successfully.
	charmURL, _ := charm.ParseURL(pp.charm.Id())
	return fmt.Sprintf("%s#%s", charmGlobalKey(charmURL), name)
}

func (pp processesPersistence) processID(id string) string {
	return fmt.Sprintf("%s#%s", unitGlobalKey(pp.unit.Id()), id)
}

func (pp processesPersistence) launchID(id string) string {
	return pp.processID(id) + "#launch"
}

func (pp processesPersistence) newInsertDefinitionOp(definition charm.Process) txn.Op {
	doc := pp.newProcessDefinitionDoc(definition)
	return txn.Op{
		C:      workloadProcessesC,
		Id:     doc.DocID,
		Assert: txn.DocMissing,
		Insert: doc,
	}
}

func (pp processesPersistence) newInsertProcessOps(info process.Info) []txn.Op {
	var ops []txn.Op
	ops = append(ops, pp.newInsertLaunchOp(info))
	ops = append(ops, pp.newInsertProcOp(info))
	return ops
}

func (pp processesPersistence) newInsertLaunchOp(info process.Info) txn.Op {
	doc := pp.newLaunchDoc(info)
	return txn.Op{
		C:      workloadProcessesC,
		Id:     doc.DocID,
		Assert: txn.DocMissing,
		Insert: doc,
	}
}

func (pp processesPersistence) newInsertProcOp(info process.Info) txn.Op {
	doc := pp.newProcessDoc(info)
	return txn.Op{
		C:      workloadProcessesC,
		Id:     doc.DocID,
		Assert: txn.DocMissing,
		Insert: doc,
	}
}

func (pp processesPersistence) newSetRawStatusOp(id string, status process.RawStatus) txn.Op {
	id = pp.processID(id)
	return txn.Op{
		C:      workloadProcessesC,
		Id:     id,
		Assert: isAliveDoc,
		Update: bson.D{{"$set", bson.D{{"pluginstatus", status.Value}}}},
	}
}

type processDefinitionDoc struct {
	DocID   string `bson:"_id"`
	EnvUUID string `bson:"env-uuid"`

	UnitID string `bson:"unitid"`

	Name        string            `bson:"name"`
	Description string            `bson:"description"`
	Type        string            `bson:"type"`
	TypeOptions map[string]string `bson:"typeoptions"`
	Command     string            `bson:"command"`
	Image       string            `bson:"image"`
	Ports       []string          `bson:"ports"`
	Volumes     []string          `bson:"volumes"`
	EnvVars     map[string]string `bson:"envvars"`
}

func (pp processesPersistence) newProcessDefinitionDoc(definition charm.Process) *processDefinitionDoc {
	id := pp.definitionID(definition.Name)

	var ports []string
	for _, p := range definition.Ports {
		// TODO(ericsnow) Ensure p.Endpoint is in state?
		ports = append(ports, fmt.Sprintf("%d:%d:%s", p.External, p.Internal, p.Endpoint))
	}

	var volumes []string
	for _, v := range definition.Volumes {
		// TODO(ericsnow) Ensure v.Name is in state?
		volumes = append(volumes, fmt.Sprintf("%s:%s:%s:%s", v.ExternalMount, v.InternalMount, v.Mode, v.Name))
	}

	return &processDefinitionDoc{
		DocID:  id,
		UnitID: pp.unit.Id(),

		Name:        definition.Name,
		Description: definition.Description,
		Type:        definition.Type,
		TypeOptions: definition.TypeOptions,
		Command:     definition.Command,
		Image:       definition.Image,
		Ports:       ports,
		Volumes:     volumes,
		EnvVars:     definition.EnvVars,
	}
}

type processLaunchDoc struct {
	DocID   string `bson:"_id"`
	EnvUUID string `bson:"env-uuid"`

	PluginID  string `bson:"pluginid"`
	RawStatus string `bson:"rawstatus"`
}

func (pp processesPersistence) newLaunchDoc(info process.Info) *processLaunchDoc {
	id := pp.launchID(info.ID())
	return &processLaunchDoc{
		DocID: id,

		PluginID:  info.Details.ID,
		RawStatus: info.Details.Status.Value,
	}
}

type processDoc struct {
	DocID   string `bson:"_id"`
	EnvUUID string `bson:"env-uuid"`

	Life         Life   `bson:"life"`
	Status       string `bson:"status"`
	PluginStatus string `bson:"pluginstatus"`
}

func (pp processesPersistence) newProcessDoc(info process.Info) *processDoc {
	id := pp.processID(info.ID())

	var status string
	switch info.Status {
	case process.StatusPending:
		status = "pending"
	case process.StatusActive:
		status = "active"
	case process.StatusFailed:
		status = "failed"
	case process.StatusStopped:
		status = "stopped"
	default:
		// TODO(ericsnow) disallow? don't worry (shouldn't happen)?
		status = "unknown"
	}

	return &processDoc{
		DocID: id,

		Life:         Alive,
		Status:       status,
		PluginStatus: info.Details.Status.Value,
	}
}
