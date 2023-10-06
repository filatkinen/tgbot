package lorry

import (
	"errors"
	"fmt"
	"github.com/filatkinen/tgbot/internal/model/car/lorry"
	"sync"
)

var (
	ErrWrongIndexSlice = errors.New("index is out of range")
	ErrWrongIndexValue = errors.New("element with this index not found")
)

type LorryService interface {
	Describe(lorryID uint64) (lorry.Lorry, error)
	List(cursor uint64, limit uint64) ([]lorry.Lorry, error)
	Create(lorry lorry.Lorry) (uint64, error)
	Update(lorryID uint64, lorry lorry.Lorry) error
	Remove(lorryID uint64) (bool, error)
}

type DummyLorryService struct {
	lorries []lorry.Lorry
	// map with pointer to the idx element of slices by index element.
	inc  uint64
	lock sync.RWMutex
}

func NewDummyLorryService(a ...int) *DummyLorryService {
	return &DummyLorryService{
		lorries: []lorry.Lorry{},
		lock:    sync.RWMutex{}}
}

func (l *DummyLorryService) Describe(lorryID uint64) (lorry.Lorry, error) {
	l.lock.RLock()
	defer l.lock.RUnlock()
	ok := false
	idx := 0
	for k := range l.lorries {
		if l.lorries[k].ID == lorryID {
			ok = true
			idx = k
			break
		}
	}
	if !ok {
		return lorry.Lorry{}, fmt.Errorf("%w, id:%d", ErrWrongIndexValue, lorryID)
	}
	return l.lorries[idx], nil
}

func (l *DummyLorryService) List(cursor uint64, limit uint64) ([]lorry.Lorry, error) {
	l.lock.RLock()
	defer l.lock.RUnlock()

	if len(l.lorries) == 0 {
		return []lorry.Lorry{}, nil
	}

	if cursor < 1 || cursor > uint64(len(l.lorries)) {
		return nil, fmt.Errorf("%w, id:%d", ErrWrongIndexSlice, cursor)
	}
	if cursor+limit-1 > uint64(len(l.lorries)) {
		limit = uint64(len(l.lorries)) - cursor + 1
	}
	lorriesOut := make([]lorry.Lorry, limit)
	copy(lorriesOut, l.lorries[cursor-1:cursor+limit-1])
	return lorriesOut, nil
}

func (l *DummyLorryService) Create(lorryIn lorry.Lorry) (uint64, error) {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.inc++
	l.lorries = append(l.lorries, lorry.Lorry{
		Model: lorryIn.Model,
		ID:    l.inc,
	})
	return l.inc, nil
}

func (l *DummyLorryService) Update(lorryID uint64, lorryIn lorry.Lorry) error {
	l.lock.Lock()
	defer l.lock.Unlock()

	ok := false
	idx := 0
	for k := range l.lorries {
		if l.lorries[k].ID == lorryID {
			ok = true
			idx = k
			break
		}
	}
	if !ok {
		return fmt.Errorf("%w, id:%d", ErrWrongIndexValue, lorryID)
	}
	l.lorries[idx].Model = lorryIn.Model

	return nil
}

func (l *DummyLorryService) Remove(lorryID uint64) (bool, error) {
	l.lock.Lock()
	l.lock.Unlock()

	ok := false
	idx := 0
	for k := range l.lorries {
		if l.lorries[k].ID == lorryID {
			ok = true
			idx = k
			break
		}
	}
	if !ok {
		return false, fmt.Errorf("%w, id:%d", ErrWrongIndexValue, lorryID)
	}
	l.lorries = append(l.lorries[0:idx], l.lorries[idx+1:]...)

	return true, nil
}
