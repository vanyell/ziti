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

package xctrl

import (
	"github.com/openziti/channel/v4"
	"github.com/openziti/ziti/common/config"
	"github.com/openziti/storage/boltz"
)

// An Xctrl allows adding handlers to the router <-> controller connection
// on the controller side. This means you can support additional message
// types/flows to extend the basic fabric functionality.
//
// There is a corresponding Xrctrl interface for extending communication on
// the router side
type Xctrl interface {
	config.Subconfig
	channel.BindHandler
	Enabled() bool
	Run(ctrl channel.Channel, db boltz.Db, done chan struct{}) error
	NotifyOfReconnect(ch channel.Channel)
	GetTraceDecoders() []channel.TraceMessageDecoder
}
