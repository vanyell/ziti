package events

import (
	"github.com/openziti/fabric/controller/network"
	"github.com/openziti/fabric/trace"
	"github.com/openziti/foundation/util/cowslice"
	"github.com/pkg/errors"
	"reflect"
)

func init() {
	RegisterEventType("metrics", registerMetricsEventHandler)
	RegisterEventType("services", registerServiceEventHandler)
	RegisterEventType("fabric.usage", registerUsageEventHandler)
	RegisterEventType("fabric.sessions", registerSessionEventHandler)
	RegisterEventType("fabric.terminators", registerTerminatorEventHandler)
	RegisterEventType("fabric.traces", func(val interface{}, _ map[interface{}]interface{}) error {
		handler, ok := val.(trace.EventHandler)
		if !ok {
			return errors.Errorf("type %v doesn't implement github.com/openziti/fabric/trace/EventHandler interface.", reflect.TypeOf(val))
		}
		AddTraceEventHandler(handler)
		return nil
	})
}

func AddSessionEventHandler(handler network.SessionEventHandler) {
	cowslice.Append(network.SessionEventHandlerRegistry, handler)
}

func RemoveSessionEventHandler(handler network.SessionEventHandler) {
	cowslice.Delete(network.SessionEventHandlerRegistry, handler)
}

func AddTraceEventHandler(handler trace.EventHandler) {
	cowslice.Append(trace.EventHandlerRegistry, handler)
}

func RemoveTraceEventHandler(handler trace.EventHandler) {
	cowslice.Delete(trace.EventHandlerRegistry, handler)
}

func AddTerminatorEventHandler(handler TerminatorEventHandler) {
	cowslice.Append(terminatorEventHandlerRegistry, handler)
}

func RemoveTerminatorEventHandler(handler TerminatorEventHandler) {
	cowslice.Delete(terminatorEventHandlerRegistry, handler)
}

func AddServiceEventHandler(handler ServiceEventHandler) {
	cowslice.Append(serviceEventHandlerRegistry, handler)
}

func RemoveServiceEventHandler(handler ServiceEventHandler) {
	cowslice.Delete(serviceEventHandlerRegistry, handler)
}
