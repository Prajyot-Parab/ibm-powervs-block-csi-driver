/*
Copyright 2021 The Kubernetes Authors.

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

package cloud

import (
	"github.com/IBM-Cloud/power-go-client/power/models"
)

type Cloud interface {
	CreateDisk(volumeName string, diskOptions *DiskOptions) (disk *Disk, err error)
	DeleteDisk(volumeID string) (success bool, err error)
	AttachDisk(volumeID string, nodeID string) (err error)
	DetachDisk(volumeID string, nodeID string) (err error)
	ResizeDisk(volumeID string, reqSize int64) (newSize int64, err error)
	WaitForVolumeState(volumeID, state string) error
	GetDiskByName(name string) (disk *Disk, err error)
	GetDiskByID(volumeID string) (disk *Disk, err error)
	GetPVMInstanceByName(instanceName string) (instance *PVMInstance, err error)
	GetPVMInstanceByID(instanceID string) (instance *PVMInstance, err error)
	GetPVMInstanceDetails(instanceID string) (*models.PVMInstance, error)
	UpdateStoragePoolAffinity(instanceID string) error
	GetImageByID(imageID string) (image *PVMImage, err error)
	IsAttached(volumeID string, nodeID string) (attached bool, err error)
}
