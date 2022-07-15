// pmm-agent
// Copyright 2019 Percona LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package connectionuptime

import (
	"math/big"
	"sync"
	"time"
)

// Service calculates connection uptime between agent and server
type Service struct {
	uptimeSeconds big.Int

	windowPeriodSeconds int64
	indexLastStatus     int64
	startTime           time.Time
	lastStatusTimestamp time.Time

	mx sync.Mutex
}

// NewService creates new instance of Service
func NewService(windowPeriod time.Duration) *Service {
	return &Service{
		windowPeriodSeconds: int64(windowPeriod.Seconds()),
	}
}

// RegisterConnectionStatus adds new connection status
func (s *Service) RegisterConnectionStatus(timestamp time.Time, connected bool) {
	s.mx.Lock()
	defer s.mx.Unlock()

	s.registerConnectionStatus(timestamp, connected)
}

func (s *Service) registerConnectionStatus(timestamp time.Time, connected bool) {
	if s.startTime.IsZero() {
		s.startTime = timestamp
		s.lastStatusTimestamp = timestamp
		s.uptimeSeconds.SetBit(&s.uptimeSeconds, 0, toUint(connected))
		s.indexLastStatus = 0

		return
	}

	secondsFromLastEvent := timestamp.Unix() - s.lastStatusTimestamp.Unix()
	endIndex := s.indexLastStatus + secondsFromLastEvent
	lastConnectedStatusBit := s.uptimeSeconds.Bit(int(s.indexLastStatus))

	for i := s.indexLastStatus + 1; i < endIndex; i++ {
		// set the same status to elements of previous connection status
		s.uptimeSeconds.SetBit(&s.uptimeSeconds, int(i%s.windowPeriodSeconds), lastConnectedStatusBit)
	}

	s.indexLastStatus = endIndex % s.windowPeriodSeconds
	s.uptimeSeconds.SetBit(&s.uptimeSeconds, int(s.indexLastStatus), toUint(connected))
	s.lastStatusTimestamp = timestamp
}

func toUint(b bool) uint {
	if b {
		return 1
	}
	return 0
}

// GetConnectedUpTimeSince calculates connected uptime between agent and server
// based on the connection statuses
func (s *Service) GetConnectedUpTimeSince(toTime time.Time) float32 {
	s.fillStatusesUntil(toTime)
	return s.calculateConnectionUpTime(toTime)
}

func (s *Service) calculateConnectionUpTime(toTime time.Time) float32 {
	numOfSeconds := s.getNumOfSecondsForCalculationUptime(toTime)
	startIndex := s.getStartIndex(numOfSeconds)
	connectedSeconds := s.getNumOfConnectedSeconds(startIndex, numOfSeconds)

	return float32(connectedSeconds) / float32(numOfSeconds) * 100
}

func (s *Service) getNumOfSecondsForCalculationUptime(toTime time.Time) int64 {
	numOfSeconds := s.windowPeriodSeconds
	diffInSecondsBetweenStartTimeAndToTime := toTime.Unix() - s.startTime.Unix()
	if diffInSecondsBetweenStartTimeAndToTime < s.windowPeriodSeconds {
		numOfSeconds = diffInSecondsBetweenStartTimeAndToTime
	}
	return numOfSeconds
}

func (s *Service) getStartIndex(size int64) int64 {
	startElement := s.indexLastStatus - size
	if startElement < 0 {
		startElement = s.windowPeriodSeconds + startElement
	}
	return startElement
}

func (s *Service) getNumOfConnectedSeconds(startIndex int64, totalNumOfSeconds int64) int {
	endIndex := startIndex + totalNumOfSeconds
	connectedSeconds := 0
	for i := startIndex; i < endIndex; i++ {
		if s.uptimeSeconds.Bit(int(i%s.windowPeriodSeconds)) == 1 {
			connectedSeconds++
		}
	}
	return connectedSeconds
}

// fill values in the slice until toTime
func (s *Service) fillStatusesUntil(toTime time.Time) {
	s.registerConnectionStatus(toTime, s.uptimeSeconds.Bit(int(s.indexLastStatus)) == 1)
}
