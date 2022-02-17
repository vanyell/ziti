package persistence

import (
	"github.com/google/uuid"
	"github.com/openziti/fabric/controller/db"
	"github.com/openziti/foundation/storage/boltz"
	"go.etcd.io/bbolt"
	"testing"
)

func Test_EdgeRouterEvents(t *testing.T) {
	ctx := NewTestContext(t)
	defer ctx.Cleanup()
	ctx.Init()

	eventChecker := boltz.NewTestEventChecker(&ctx.Assertions)
	eventChecker.AddHandlers(ctx.stores.Router)
	eventChecker.AddHandlers(ctx.stores.EdgeRouter)

	fp := uuid.NewString()
	edgeRouter := &EdgeRouter{
		Router: db.Router{
			BaseExtEntity: boltz.BaseExtEntity{
				Id: uuid.NewString(),
			},
			Name:        uuid.NewString(),
			Fingerprint: &fp,
		},
	}

	err := ctx.db.Update(func(tx *bbolt.Tx) error {
		return ctx.stores.EdgeRouter.Create(boltz.NewMutateContext(tx), edgeRouter)
	})
	ctx.NoError(err)
	eventChecker.RequireEvent(boltz.TestEntityTypeParent, edgeRouter, boltz.EventCreate)
	eventChecker.RequireEvent(boltz.TestEntityTypeChild, edgeRouter, boltz.EventCreate)
	eventChecker.RequireNoEvent()

	// check events generated by updating from the edge-router store
	newFp := uuid.NewString()
	err = ctx.db.Update(func(tx *bbolt.Tx) error {
		edgeRouter.Name = uuid.NewString()
		edgeRouter.Fingerprint = &newFp
		return ctx.stores.EdgeRouter.Update(boltz.NewMutateContext(tx), edgeRouter, nil)
	})
	ctx.NoError(err)

	entity := eventChecker.RequireEvent(boltz.TestEntityTypeParent, edgeRouter, boltz.EventUpdate)
	r, ok := entity.(*db.Router)
	ctx.True(ok)
	ctx.Equal(edgeRouter.Name, r.Name)
	ctx.NotNil(r.Fingerprint)
	ctx.Equal(*edgeRouter.Fingerprint, *r.Fingerprint)

	entity = eventChecker.RequireEvent(boltz.TestEntityTypeChild, edgeRouter, boltz.EventUpdate)
	er, ok := entity.(*EdgeRouter)
	ctx.True(ok)
	ctx.Equal(edgeRouter.Name, er.Name)
	ctx.NotNil(er.Fingerprint)
	ctx.Equal(*edgeRouter.Fingerprint, *er.Fingerprint)

	eventChecker.RequireNoEvent()

	// check events generated by updating from the router store
	newFp = uuid.NewString()
	err = ctx.db.Update(func(tx *bbolt.Tx) error {
		edgeRouter.Name = uuid.NewString()
		edgeRouter.Fingerprint = &newFp
		return ctx.stores.Router.Update(boltz.NewMutateContext(tx), &edgeRouter.Router, nil)
	})
	ctx.NoError(err)

	entity = eventChecker.RequireEvent(boltz.TestEntityTypeParent, edgeRouter, boltz.EventUpdate)
	r, ok = entity.(*db.Router)
	ctx.True(ok)
	ctx.Equal(edgeRouter.Name, r.Name)
	ctx.NotNil(r.Fingerprint)
	ctx.Equal(*edgeRouter.Fingerprint, *r.Fingerprint)

	entity = eventChecker.RequireEvent(boltz.TestEntityTypeChild, edgeRouter, boltz.EventUpdate)
	er, ok = entity.(*EdgeRouter)
	ctx.True(ok)
	ctx.Equal(edgeRouter.Name, er.Name)
	ctx.NotNil(er.Fingerprint)
	ctx.Equal(*edgeRouter.Fingerprint, *er.Fingerprint)

	// ensure setting fingerprint to nil works
	err = ctx.db.Update(func(tx *bbolt.Tx) error {
		edgeRouter.Fingerprint = nil
		return ctx.stores.EdgeRouter.Update(boltz.NewMutateContext(tx), edgeRouter, nil)
	})
	ctx.NoError(err)

	entity = eventChecker.RequireEvent(boltz.TestEntityTypeParent, edgeRouter, boltz.EventUpdate)
	r, ok = entity.(*db.Router)
	ctx.True(ok)
	ctx.Equal(edgeRouter.Name, r.Name)
	ctx.Nil(edgeRouter.Fingerprint)
	ctx.Nil(r.Fingerprint)

	entity = eventChecker.RequireEvent(boltz.TestEntityTypeChild, edgeRouter, boltz.EventUpdate)
	er, ok = entity.(*EdgeRouter)
	ctx.True(ok)
	ctx.Equal(edgeRouter.Name, er.Name)
	ctx.Nil(er.Fingerprint)

	// check delete events
	err = ctx.db.Update(func(tx *bbolt.Tx) error {
		return ctx.stores.EdgeRouter.DeleteById(boltz.NewMutateContext(tx), edgeRouter.Id)
	})
	ctx.NoError(err)
	eventChecker.RequireEvent(boltz.TestEntityTypeParent, edgeRouter, boltz.EventDelete)
	eventChecker.RequireEvent(boltz.TestEntityTypeChild, edgeRouter, boltz.EventDelete)
	eventChecker.RequireNoEvent()

	//// check delete again, this time invoked from the child store
	//mgr = &Manager{
	//	Employee: Employee{
	//		Id:   uuid.NewString(),
	//		Name: "mgr",
	//	},
	//}
	//
	//err = test.db.Update(func(tx *bbolt.Tx) error {
	//	return test.mgrStore.Create(NewMutateContext(tx), mgr)
	//})
	//test.NoError(err)
	//eventChecker.requireEvent(entityTypeEmployee, mgr.Id, EventCreate)
	//eventChecker.requireEvent(entityTypeManager, mgr.Id, EventCreate)
	//eventChecker.requireNoEvent()
	//
	//err = test.db.Update(func(tx *bbolt.Tx) error {
	//	return test.empStore.DeleteById(NewMutateContext(tx), mgr.Id)
	//})
	//test.NoError(err)
	//eventChecker.requireEvent(entityTypeEmployee, mgr.Id, EventDelete)
	//eventChecker.requireEvent(entityTypeManager, mgr.Id, EventDelete)
	//eventChecker.requireNoEvent()
}
