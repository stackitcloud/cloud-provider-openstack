/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package openstack

import (
	"github.com/gophercloud/gophercloud/v2/openstack/blockstorage/v3/backups"
	"github.com/gophercloud/gophercloud/v2/openstack/blockstorage/v3/snapshots"
	"github.com/gophercloud/gophercloud/v2/openstack/blockstorage/v3/volumes"
	"github.com/gophercloud/gophercloud/v2/openstack/compute/v2/servers"
	"github.com/stretchr/testify/mock"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/cloud-provider-openstack/pkg/util/metadata"
)

var fakeVol1 = volumes.Volume{
	ID:               "261a8b81-3660-43e5-bab8-6470b65ee4e9",
	Name:             "fake-duplicate",
	Status:           "available",
	AvailabilityZone: "nova",
	Size:             1,
}

var fakeSnapshot = snapshots.Snapshot{
	ID:       "261a8b81-3660-43e5-bab8-6470b65ee4e8",
	Name:     "fake-snapshot",
	Status:   "available",
	Size:     1,
	VolumeID: "CSIVolumeID",
	Metadata: make(map[string]string),
}

var fakemap = make(map[string]string)

var fakeBackup = backups.Backup{
	ID:         "eb5e4e9a-a4e5-4728-a748-04f9e2868573",
	Name:       "fake-snapshot",
	Status:     "available",
	Size:       1,
	VolumeID:   "CSIVolumeID",
	SnapshotID: "261a8b81-3660-43e5-bab8-6470b65ee4e8",
	Metadata:   &fakemap,
}

// revive:disable:exported
// OpenStackMock is an autogenerated mock type for the IOpenStack type
// ORIGINALLY GENERATED BY mockery with hand edits
type OpenStackMock struct {
	mock.Mock
}

// revive:enable:exported

// AttachVolume provides a mock function with given fields: instanceID, volumeID
func (_m *OpenStackMock) AttachVolume(instanceID string, volumeID string) (string, error) {
	ret := _m.Called(instanceID, volumeID)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(instanceID, volumeID)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(instanceID, volumeID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateVolume provides a mock function with given fields: name, size, vtype, availability, tags
func (_m *OpenStackMock) CreateVolume(name string, size int, vtype string, availability string, snapshotID string, sourceVolID string, sourceBackupID string, tags map[string]string) (*volumes.Volume, error) {
	ret := _m.Called(name, size, vtype, availability, snapshotID, sourceVolID, sourceBackupID, tags)

	var r0 *volumes.Volume
	if rf, ok := ret.Get(0).(func(string, int, string, string, string, string, string, map[string]string) *volumes.Volume); ok {
		r0 = rf(name, size, vtype, availability, snapshotID, sourceVolID, sourceBackupID, tags)
	} else {
		r0 = ret.Get(0).(*volumes.Volume)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int, string, string, string, string, string, map[string]string) error); ok {
		r1 = rf(name, size, vtype, availability, snapshotID, sourceVolID, sourceBackupID, tags)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteVolume provides a mock function with given fields: volumeID
func (_m *OpenStackMock) DeleteVolume(volumeID string) error {
	ret := _m.Called(volumeID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(volumeID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetVolume provides a mock function with given fields: volumeID
func (_m *OpenStackMock) GetVolume(volumeID string) (*volumes.Volume, error) {
	return &fakeVol1, nil
}

// DetachVolume provides a mock function with given fields: instanceID, volumeID
func (_m *OpenStackMock) DetachVolume(instanceID string, volumeID string) error {
	ret := _m.Called(instanceID, volumeID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(instanceID, volumeID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAttachmentDiskPath provides a mock function with given fields: instanceID, volumeID
func (_m *OpenStackMock) GetAttachmentDiskPath(instanceID string, volumeID string) (string, error) {
	ret := _m.Called(instanceID, volumeID)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(instanceID, volumeID)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(instanceID, volumeID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WaitDiskAttached provides a mock function with given fields: instanceID, volumeID
func (_m *OpenStackMock) WaitDiskAttached(instanceID string, volumeID string) error {
	ret := _m.Called(instanceID, volumeID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(instanceID, volumeID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitVolumeTargetStatus provides a mock function with given fields: volumeID, tStatus
func (_m *OpenStackMock) WaitVolumeTargetStatus(volumeID string, tStatus []string) error {
	ret := _m.Called(volumeID, tStatus)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []string) error); ok {
		r0 = rf(volumeID, tStatus)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitVolumeTargetStatusWithCustomBackoff provides a mock function with given fields: volumeID, tStatus, backoff
func (_m *OpenStackMock) WaitVolumeTargetStatusWithCustomBackoff(volumeID string, tStatus []string, backoff *wait.Backoff) error {
	ret := _m.Called(volumeID, tStatus, backoff)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []string, *wait.Backoff) error); ok {
		r0 = rf(volumeID, tStatus, backoff)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitDiskDetached provides a mock function with given fields: instanceID, volumeID
func (_m *OpenStackMock) WaitDiskDetached(instanceID string, volumeID string) error {
	ret := _m.Called(instanceID, volumeID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(instanceID, volumeID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetVolumesByName provides a mock function with given fields: name
func (_m *OpenStackMock) GetVolumesByName(name string) ([]volumes.Volume, error) {

	ret := _m.Called(name)

	var r0 []volumes.Volume
	if rf, ok := ret.Get(0).(func(string) []volumes.Volume); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]volumes.Volume)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListSnapshots provides a mock function with given fields: limit, offset, filters
func (_m *OpenStackMock) ListSnapshots(filters map[string]string) ([]snapshots.Snapshot, string, error) {
	ret := _m.Called(filters)

	var r0 []snapshots.Snapshot
	if rf, ok := ret.Get(0).(func(map[string]string) []snapshots.Snapshot); ok {
		r0 = rf(filters)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]snapshots.Snapshot)
		}
	}
	var r1 string

	if rf, ok := ret.Get(1).(func(map[string]string) string); ok {
		r1 = rf(filters)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(string)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(map[string]string) error); ok {
		r2 = rf(filters)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// CreateSnapshot provides a mock function with given fields: name, volID, tags
func (_m *OpenStackMock) CreateSnapshot(name string, volID string, tags map[string]string) (*snapshots.Snapshot, error) {
	ret := _m.Called(name, volID, tags)

	var r0 *snapshots.Snapshot
	if rf, ok := ret.Get(0).(func(string, string, map[string]string) *snapshots.Snapshot); ok {
		r0 = rf(name, volID, tags)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*snapshots.Snapshot)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, map[string]string) error); ok {
		r1 = rf(name, volID, tags)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteSnapshot provides a mock function with given fields: snapID
func (_m *OpenStackMock) DeleteSnapshot(snapID string) error {
	ret := _m.Called(snapID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(snapID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *OpenStackMock) ListBackups(filters map[string]string) ([]backups.Backup, error) {
	ret := _m.Called(filters)

	var r0 []backups.Backup
	if rf, ok := ret.Get(0).(func(map[string]string) []backups.Backup); ok {
		r0 = rf(filters)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]backups.Backup)
		}
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(map[string]string) error); ok {
		r1 = rf(filters)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *OpenStackMock) CreateBackup(name, volID, snapshotID, availabilityZone string, tags map[string]string) (*backups.Backup, error) {
	ret := _m.Called(name, volID, snapshotID, availabilityZone, tags)

	var r0 *backups.Backup
	if rf, ok := ret.Get(0).(func(string, string, string, string, map[string]string) *backups.Backup); ok {
		r0 = rf(name, volID, snapshotID, availabilityZone, tags)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*backups.Backup)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string, string, map[string]string) error); ok {
		r1 = rf(name, volID, snapshotID, availabilityZone, tags)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *OpenStackMock) DeleteBackup(backupID string) error {
	ret := _m.Called(backupID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(backupID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListVolumes provides a mock function without param
func (_m *OpenStackMock) ListVolumes(limit int, marker string) ([]volumes.Volume, string, error) {
	ret := _m.Called(limit, marker)

	var r0 []volumes.Volume
	if rf, ok := ret.Get(0).(func(int, string) []volumes.Volume); ok {
		r0 = rf(limit, marker)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]volumes.Volume)
		}
	}

	var r1 string

	if rf, ok := ret.Get(1).(func(int, string) string); ok {
		r1 = rf(limit, marker)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(string)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(int, string) error); ok {
		r2 = rf(limit, marker)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

func (_m *OpenStackMock) GetAvailabilityZone() (string, error) {
	ret := _m.Called()
	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *OpenStackMock) GetInstanceID() (string, error) {
	return "", nil
}

func (_m *OpenStackMock) GetSnapshotByID(snapshotID string) (*snapshots.Snapshot, error) {

	return &fakeSnapshot, nil
}

func (_m *OpenStackMock) WaitSnapshotReady(snapshotID string) (string, error) {
	ret := _m.Called(snapshotID)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(snapshotID)
	} else {
		r0 = ret.String(0)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(snapshotID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *OpenStackMock) GetBackupByID(backupID string) (*backups.Backup, error) {

	return &fakeBackup, nil
}

func (_m *OpenStackMock) WaitBackupReady(backupID string, snapshotSize int, backupMaxDurationSecondsPerGB int) (string, error) {
	ret := _m.Called(backupID)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(backupID)
	} else {
		r0 = ret.String(0)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(backupID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *OpenStackMock) GetMaxVolLimit() int64 {
	return 256
}

func (_m *OpenStackMock) BackupsAreEnabled() (bool, error) {
	return true, nil
}

func (_m *OpenStackMock) GetInstanceByID(instanceID string) (*servers.Server, error) {
	return nil, nil
}

// ExpandVolume provides a mock function with given fields: instanceID, volumeID
func (_m *OpenStackMock) ExpandVolume(volumeID string, status string, size int) error {
	ret := _m.Called(volumeID, status, size)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, int) error); ok {
		r0 = rf(volumeID, status, size)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *OpenStackMock) GetMetadataOpts() metadata.Opts {
	var m metadata.Opts
	m.SearchOrder = "configDrive"
	return m
}

// GetBlockStorageOpts provides a mock function to return BlockStorageOpts
func (_m *OpenStackMock) GetBlockStorageOpts() BlockStorageOpts {
	return BlockStorageOpts{}
}
