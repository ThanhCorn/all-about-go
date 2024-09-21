package mediator

import (
	"fmt"
	"sync"
)

const DefaultMaxSize = 1

type StationManager struct {
	Size           uint64 // size of train can be arrival
	TrainQueue     []Train
	IsPlatformFree bool
	mu             *sync.Mutex
}

func NewStationManager() *StationManager {
	return &StationManager{
		IsPlatformFree: true,
		Size:           0,
		mu:             &sync.Mutex{},
	}
}

func (s *StationManager) canArrive(train Train) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.IsCanArrive() {
		fmt.Printf("TrainQueue %+v\n", s.TrainQueue)
		s.Size++
		if s.Size == DefaultMaxSize {
			s.IsPlatformFree = false
		}
		return true
	} else {
		s.TrainQueue = append(s.TrainQueue, train)
		fmt.Println("\nPlatform is busy, Train is waiting")
		return false
	}
}

func (s *StationManager) notifyAboutDeparture() {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.TrainQueue) > 0 {
		train := s.TrainQueue[0]
		s.TrainQueue = s.TrainQueue[1:]
		s.Size--
		train.PermitArrival()
	}
	if !s.IsPlatformFree && s.Size < DefaultMaxSize {
		s.IsPlatformFree = true
	}
}

func (s *StationManager) IsCanArrive() bool {
	return s.Size < DefaultMaxSize || s.IsPlatformFree == true
}
