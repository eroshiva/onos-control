// Copyright 2019-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package manager is is the main coordinator for the ONOS control subsystem.
package manager

import (
	log "k8s.io/klog"
)

var mgr Manager

// Manager single point of entry for the config system.
type Manager struct {
}

// NewManager initializes the network control manager subsystem.
func NewManager() (*Manager, error) {
	log.Info("Creating Manager")
	mgr = Manager{}

	return &mgr, nil
}

// LoadManager creates a configuration subsystem manager primed with stores loaded from the specified files.
func LoadManager() (*Manager, error) {
	return NewManager()
}

// Run starts a synchronizer based on the devices and the northbound services.
func (m *Manager) Run() {
	log.Info("Starting Manager")
	// Start the main dispatcher system
}

//Close kills the channels and manager related objects
func (m *Manager) Close() {
	log.Info("Closing Manager")
}

// GetManager returns the initialized and running instance of manager.
// Should be called only after NewManager and Run are done.
func GetManager() *Manager {
	return &mgr
}
