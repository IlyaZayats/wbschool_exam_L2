package main

import "fmt"

type Daimyo interface {
	Visit(province Province)
}

type Province interface {
	Accept(daimyo Daimyo)
	Status()
}

type Kyoto struct {
	EmperorPalaceStatus bool
	NijoCastleStatus    bool
	EconomicsLevel      int
	CristianMission     bool
	IsCapital           bool
}

func (t *Kyoto) Accept(daimyo Daimyo) {
	daimyo.Visit(t)
}

func (t *Kyoto) SetEmperorPalace(status bool) {
	t.EmperorPalaceStatus = status
}

func (t *Kyoto) SetNijoCastle(status bool) {
	t.NijoCastleStatus = status
}

func (t *Kyoto) SetEconomicsLevel(level int) {
	t.EconomicsLevel = level
}

func (t *Kyoto) SetCristianMission(status bool) {
	t.CristianMission = status
}

func (t *Kyoto) RemoveCapitalStatus() {
	t.IsCapital = false
}

func (t *Kyoto) Status() {
	fmt.Printf("\nKyoto:\nEmperor palace: %v\nNijo Castle: %v\nEconomics level: %v\nCristian mission: %v\nIs capital: %v",
		t.EmperorPalaceStatus, t.NijoCastleStatus, t.EconomicsLevel, t.CristianMission, t.IsCapital)
}

func NewKyoto() *Kyoto {
	return &Kyoto{EmperorPalaceStatus: false, NijoCastleStatus: false, EconomicsLevel: 3, CristianMission: false, IsCapital: true}
}

type Edo struct {
	BakufuStatus    bool
	EdoCastleStatus bool
	EconomicsLevel  int
	IsCapital       bool
}

func (t *Edo) Accept(daimyo Daimyo) {
	daimyo.Visit(t)
}

func (t *Edo) SetBakufu(status bool) {
	t.BakufuStatus = status
}

func (t *Edo) SetEdoCastle(status bool) {
	t.EdoCastleStatus = status
}

func (t *Edo) SetEconomicsLevel(level int) {
	t.EconomicsLevel = level
}

func (t *Edo) SetCapitalStatus() {
	t.IsCapital = true
}

func (t *Edo) Status() {
	fmt.Printf("\nEdo:\nIs bakufu: %v\nEdo Castle: %v\nEconomics level: %v\nIs capital: %v",
		t.BakufuStatus, t.EdoCastleStatus, t.EconomicsLevel, t.IsCapital)
}

func NewEdo() *Edo {
	return &Edo{BakufuStatus: false, EdoCastleStatus: false, EconomicsLevel: 1, IsCapital: false}
}

type Oda struct {
}

func (d *Oda) Visit(province Province) {
	switch province.(type) {
	case *Kyoto:
		p := province.(*Kyoto)
		p.SetEmperorPalace(true)
		p.SetEconomicsLevel(10)
		p.SetCristianMission(true)
		p.SetNijoCastle(true)
	}
	province.Status()
}

type Tokugawa struct {
}

func (d *Tokugawa) Visit(province Province) {
	switch province.(type) {
	case *Kyoto:
		p := province.(*Kyoto)
		p.SetCristianMission(false)
		p.RemoveCapitalStatus()
	case *Edo:
		p := province.(*Edo)
		p.SetBakufu(true)
		p.SetEdoCastle(true)
		p.SetEconomicsLevel(8)
		p.SetCapitalStatus()
	}
	province.Status()
}

func main() {
	oda := &Oda{}
	tokugawa := &Tokugawa{}
	kyoto := NewKyoto()
	edo := NewEdo()

	kyoto.Status()
	edo.Status()

	kyoto.Accept(oda)

	kyoto.Accept(tokugawa)
	edo.Accept(tokugawa)
}
