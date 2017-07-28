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

package cpumanager

import (
	"k8s.io/kubernetes/pkg/kubelet/cpumanager/topology"
)

var (
	topoSingleSocketHT = &topology.CPUTopology{
		NumCPUs:        8,
		NumSockets:     1,
		NumCores:       4,
		HyperThreading: true,
		CPUtopoDetails: map[int]topology.CPUInfo{
			0: {CoreId: 0, SocketId: 0},
			1: {CoreId: 1, SocketId: 0},
			2: {CoreId: 2, SocketId: 0},
			3: {CoreId: 3, SocketId: 0},
			4: {CoreId: 0, SocketId: 0},
			5: {CoreId: 1, SocketId: 0},
			6: {CoreId: 2, SocketId: 0},
			7: {CoreId: 3, SocketId: 0},
		},
	}

	topoDualSocketHT = &topology.CPUTopology{
		NumCPUs:        12,
		NumSockets:     2,
		NumCores:       6,
		HyperThreading: true,
		CPUtopoDetails: map[int]topology.CPUInfo{
			0:  {CoreId: 0, SocketId: 0},
			1:  {CoreId: 1, SocketId: 1},
			2:  {CoreId: 2, SocketId: 0},
			3:  {CoreId: 3, SocketId: 1},
			4:  {CoreId: 4, SocketId: 0},
			5:  {CoreId: 5, SocketId: 1},
			6:  {CoreId: 0, SocketId: 0},
			7:  {CoreId: 1, SocketId: 1},
			8:  {CoreId: 2, SocketId: 0},
			9:  {CoreId: 3, SocketId: 1},
			10: {CoreId: 4, SocketId: 0},
			11: {CoreId: 5, SocketId: 1},
		},
	}

	topoDualSocketNoHT = &topology.CPUTopology{
		NumCPUs:        8,
		NumSockets:     2,
		NumCores:       8,
		HyperThreading: false,
		CPUtopoDetails: map[int]topology.CPUInfo{
			0: {CoreId: 0, SocketId: 0},
			1: {CoreId: 1, SocketId: 0},
			2: {CoreId: 2, SocketId: 0},
			3: {CoreId: 3, SocketId: 0},
			4: {CoreId: 4, SocketId: 1},
			5: {CoreId: 5, SocketId: 1},
			6: {CoreId: 6, SocketId: 1},
			7: {CoreId: 7, SocketId: 1},
		},
	}
)
