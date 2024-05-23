package main

import "fmt"

type Soldier interface {
	SitRep()
	SetStatus(status string)
}

type SquadMember struct {
	Name   string
	Rank   string
	Status string
}

func NewSquadMember(name, rank string) SquadMember {
	return SquadMember{Name: name, Rank: rank, Status: "awaiting orders"}
}

func (s *SquadMember) SitRep() {
	fmt.Printf("\n%s - %s - sitrep - %s", s.Name, s.Rank, s.Status)
}

func (s *SquadMember) SetStatus(status string) {
	s.Status = status
	s.SitRep()
}

type SquadLeader struct {
	SquadMember
}

func (sl *SquadLeader) Command(soldier Soldier, command string) {
	soldier.SetStatus(command)
}

type Gunner struct {
	SquadMember
}

type Marksman struct {
	SquadMember
}

type MachineGunner struct {
	SquadMember
}

type Squad struct {
	SQSquadLeader   *SquadLeader
	SQMachineGunner *MachineGunner
	SQMarksman      *Marksman
	SQGunner1       *Gunner
	SQGunner2       *Gunner
}

func (sq *Squad) PrepareToMission() {
	sq.SQSquadLeader = &SquadLeader{NewSquadMember("Sergeant", "Adams")}
	sq.SQMachineGunner = &MachineGunner{NewSquadMember("Private First Class", "Johns")}
	sq.SQMarksman = &Marksman{NewSquadMember("Corporal", "Park")}
	sq.SQGunner1 = &Gunner{NewSquadMember("Private", "Davidson")}
	sq.SQGunner2 = &Gunner{NewSquadMember("Private", "Griffin")}
	sq.SQSquadLeader.SitRep()
	sq.SQMachineGunner.SitRep()
	sq.SQMarksman.SitRep()
	sq.SQGunner1.SitRep()
	sq.SQGunner2.SitRep()
}

func (sq *Squad) StartMission() {
	sq.SQSquadLeader.Command(sq.SQMarksman, "recon")
	sq.SQSquadLeader.Command(sq.SQGunner1, "hold position")
	sq.SQSquadLeader.Command(sq.SQGunner2, "hold position")
	sq.SQSquadLeader.Command(sq.SQMachineGunner, "cover")
}

func main() {
	squad := Squad{}
	squad.PrepareToMission()
	squad.StartMission()
}
