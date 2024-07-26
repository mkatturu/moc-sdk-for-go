// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the Apache v2.0 License.

package availabilityzone

import (
	"testing"

	"github.com/microsoft/moc-sdk-for-go/services/cloud"
	wssdcloudcompute "github.com/microsoft/moc/rpc/cloudagent/cloud"
	wssdcommon "github.com/microsoft/moc/rpc/common"
	"github.com/stretchr/testify/assert"
)

var id = "id"
var name = "avzone1"
var group = "testGroup1"
var location = "mocLocation"
var a1Name = "a1"
var a1Group = "ag1"
var a2Name = "a2"
var a2Group = "ag2"
var provisionoingstate = "CREATED"
var health = "OK"
var wssdnodes = []string {"node1", "node2"}
var wssdstatus = map[string]*string{
	"ProvisionState": &provisionoingstate,
	"HealthState":    &health,
}
var rpcstatus = wssdcommon.Status{
	Health:             &wssdcommon.Health{CurrentState: 1},
	ProvisioningStatus: &wssdcommon.ProvisionStatus{CurrentState: 2},
	Version:            &wssdcommon.Version{Number: "123"},
}

func Test_getRpcAvailabilityZone(t *testing.T) {
	result, err := getRpcAvailabilityZone(nil)
	assert.Error(t, err)
	assert.Nil(t, result)

	avzone := cloud.AvailabilityZone{
		Name:                     &name,
		Location:                 &location,
		Statuses:                 wssdstatus,
		Nodes:                    wssdnodes,
	}

	result, err = getRpcAvailabilityZone(&avzone)
	assert.Error(t, err)

	avzone.ID = &id
	result, err = getRpcAvailabilityZone(&avzone)
	assert.Nil(t, err)
	assert.Equal(t, name, result.Name)
	assert.Equal(t, location, result.LocationName)
}


func Test_getWssdAvailabilityZone(t *testing.T) {
	result, err := getWssdAvailabilityZone(nil)
	assert.Error(t, err)
	assert.Nil(t, result)

	avzone := wssdcloudcompute.AvailabilityZone{
		Name:                     name,
		LocationName:             location,
		Status:                   &rpcstatus,
		Nodes:                    wssdnodes,
	}

	result, err = getWssdAvailabilityZone(&avzone)
	assert.Nil(t, err)
	assert.EqualValues(t, name, *result.Name)
	assert.EqualValues(t, location, *result.Location)
}