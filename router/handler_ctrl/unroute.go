/*
	Copyright NetFoundry Inc.

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	https://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package handler_ctrl

import (
	"github.com/michaelquigley/pfxlog"
	"github.com/openziti/channel/v4"
	"github.com/openziti/ziti/common/pb/ctrl_pb"
	"github.com/openziti/ziti/router/forwarder"
	"google.golang.org/protobuf/proto"
)

type unrouteHandler struct {
	forwarder *forwarder.Forwarder
}

func newUnrouteHandler(forwarder *forwarder.Forwarder) *unrouteHandler {
	return &unrouteHandler{forwarder: forwarder}
}

func (h *unrouteHandler) ContentType() int32 {
	return int32(ctrl_pb.ContentType_UnrouteType)
}

func (h *unrouteHandler) HandleReceive(msg *channel.Message, ch channel.Channel) {
	removeRoute := &ctrl_pb.Unroute{}
	if err := proto.Unmarshal(msg.Body, removeRoute); err == nil {
		pfxlog.ContextLogger(ch.Label()).WithField("circuitId", removeRoute.CircuitId).Debug("received unroute")
		h.forwarder.Unroute(removeRoute.CircuitId, removeRoute.Now)
	} else {
		pfxlog.ContextLogger(ch.Label()).Errorf("unexpected error (%v)", err)
	}
}
