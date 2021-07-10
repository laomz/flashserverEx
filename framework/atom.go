package framework

import (
	"common"
	"errors"
	"sync"
)

type AtomCloseCallback func(uint64) error

type AtomInterface interface {
	Read() error
	Write() error
	Close() error
	release() error
	BeActive() bool
	GetAtomID() uint64
	//out interface
	Send([]byte) error
}

type atom struct {
	id       uint64
	beActive common.SafeBool
	once     sync.Once
	atomImp  AtomInterface
	wg       sync.WaitGroup
}

func (a *atom) StartAtom(id uint64, atomImp AtomInterface) error {
	if id == 0 {
		return errors.New("atom:StartAtom: id = 0")
	}
	if atomImp == nil {
		return errors.New("atom:StartAtom: atomImp = nil")
	}
	a.id = id
	a.atomImp = atomImp
	a.beActive.Set(true)
	a.wg.Add(1)
	go a.read()
	a.wg.Add(1)
	go a.write()
	return nil
}

func (a *atom) read() {
	defer func() {

		a.wg.Done()
		a.wg.Wait()
		a.closeAtom()
	}()

	for {
		if !a.beActive.Get() {
			break
		}

		if err := a.atomImp.Read(); err != nil {
			break
		}
	}
}

func (a *atom) write() {
	defer func() {
		a.wg.Done()
		a.wg.Wait()
		a.closeAtom()
	}()

	for {
		if !a.beActive.Get() {
			break
		}

		if err := a.atomImp.Write(); err != nil {
			break
		}
	}
}

func (a *atom) closeAtom() {
	a.once.Do(
		func() {
			a.atomImp.release()
		},
	)
}
