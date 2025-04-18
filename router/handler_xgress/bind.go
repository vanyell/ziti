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

package handler_xgress

import (
	"github.com/openziti/ziti/router/forwarder"
	"github.com/openziti/ziti/router/metrics"
	"github.com/openziti/ziti/router/xgress"
)

type bindHandler struct {
	dataPlaneAdapter   xgress.DataPlaneAdapter
	closeHandler       xgress.CloseHandler
	metricsPeekHandler xgress.PeekHandler
	forwarder          *forwarder.Forwarder
}

func NewBindHandler(dataPlaneAdapter xgress.DataPlaneAdapter, closeHandler xgress.CloseHandler, forwarder *forwarder.Forwarder) *bindHandler {
	return &bindHandler{
		dataPlaneAdapter:   dataPlaneAdapter,
		closeHandler:       closeHandler,
		metricsPeekHandler: metrics.NewXgressPeekHandler(forwarder.MetricsRegistry()),
		forwarder:          forwarder,
	}
}

func (bindHandler *bindHandler) HandleXgressBind(x *xgress.Xgress) {
	x.SetDataPlaneAdapter(bindHandler.dataPlaneAdapter)
	x.AddPeekHandler(bindHandler.metricsPeekHandler)

	x.AddCloseHandler(bindHandler.closeHandler)

	bindHandler.forwarder.RegisterDestination(x.CircuitId(), x.Address(), x)
}
