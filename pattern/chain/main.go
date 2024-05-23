package main

import "fmt"

type GermanTank struct {
	Name         string
	FrontalArmor int
}

func NewGermanTank(name string, frontalArmor int) *GermanTank {
	return &GermanTank{
		Name:         name,
		FrontalArmor: frontalArmor,
	}
}

type SovietATInterface interface {
	Aim(tank *GermanTank)
	Destroy(tank string)
	SetNext(atInterface SovietATInterface) SovietATInterface
}

type SovietAT struct {
	Name        string
	Penetration int
	GunCaliber  float32
	Next        SovietATInterface
}

func (at *SovietAT) SetNext(atInterface SovietATInterface) SovietATInterface {
	at.Next = atInterface
	return at
}

func (at *SovietAT) Aim(tank *GermanTank) {
	if at.Penetration < tank.FrontalArmor {
		if at.Next != nil {
			at.Next.Aim(tank)
		} else {
			fmt.Printf("\nWe need ISU-152 here! It s %s!", tank.Name)
		}
	} else {
		at.Destroy(tank.Name)
	}
}

func (at *SovietAT) Destroy(tank string) {
	fmt.Printf("\nGlorious %s have just destoyed %s tank with %vmm gun!", at.Name, tank, at.GunCaliber)
}

func NewSovietAT1(name string, penetration int, gunCaliber float32) *SovietAT1 {
	return &SovietAT1{SovietAT{Name: name, Penetration: penetration, GunCaliber: gunCaliber}}
}

func NewSovietAT2(name string, penetration int, gunCaliber float32) *SovietAT2 {
	return &SovietAT2{SovietAT{Name: name, Penetration: penetration, GunCaliber: gunCaliber}}
}

func NewSovietAT3(name string, penetration int, gunCaliber float32) *SovietAT3 {
	return &SovietAT3{SovietAT{Name: name, Penetration: penetration, GunCaliber: gunCaliber}}
}

type SovietAT1 struct {
	SovietAT
}

type SovietAT2 struct {
	SovietAT
}

type SovietAT3 struct {
	SovietAT
}

func main() {
	var sovietATSquad SovietATInterface
	sovietATSquad = NewSovietAT1("53-K", 40, 45)
	sovietATSquad.SetNext(NewSovietAT2("ZIS-3", 60, 76.2).SetNext(NewSovietAT3("BS-3", 120, 100)))
	sovietATSquad.Aim(NewGermanTank(`Pz.Kpfw. VI Ausf.H`, 100))
	sovietATSquad.Aim(NewGermanTank(`Pz.Kpfw. IV Ausf.J`, 80))
	sovietATSquad.Aim(NewGermanTank(`Pz.Kpfw. III Ausf.M`, 50))
	sovietATSquad.Aim(NewGermanTank(`Pz.Kpfw. II Ausf.L`, 30))
	sovietATSquad.Aim(NewGermanTank(`Pz.Kpfw. VI Ausf.B`, 150))
}
