package main

import "fmt"

type Squad struct {
	SquadName    string
	SquadMembers []Soldier
}

func (s *Squad) SetSquadName(name string) {
	s.SquadName = name
}

func (s *Squad) AddSquadMember(soldier *Soldier) {
	s.SquadMembers = append(s.SquadMembers, *soldier)
}

func (s *Squad) SitRep() {
	fmt.Printf("\n%s:", s.SquadName)
	for _, member := range s.SquadMembers {
		member.SitRepSoldier()
	}
}

type Soldier struct {
	Role string
	Rank string
}

func (s *Soldier) SitRepSoldier() {
	fmt.Printf("\n%s - %s", s.Role, s.Rank)
}

type SquadBuilderInterface interface {
	NewSquad()
	GetSquad() *Squad
	SetName()
	BuildFormation()
}

type SquadBuilder struct {
	SquadOnBuild *Squad
}

func (b *SquadBuilder) NewSquad() {
	b.SquadOnBuild = &Squad{}
}

func (b *SquadBuilder) GetSquad() *Squad {
	return b.SquadOnBuild
}

type RangersBuilder struct {
	SquadBuilder
}

func (b *RangersBuilder) SetName() {
	b.SquadOnBuild.SetSquadName("Rangers Rifle Squad")
}

func (b *RangersBuilder) BuildFormation() {
	formation := [][]string{
		{"Squad Leader", "Staff Sergeant"},
		{"Team Leader Alpha", "Sergeant"},
		{"Automatic Rifleman Alpha", "Specialist"},
		{"Grenadier Alpha", "Specialist"},
		{"Rifleman Alpha", "Private"},
		{"Team Leader Bravo", "Sergeant"},
		{"Automatic Rifleman Bravo", "Specialist"},
		{"Grenadier Bravo", "Specialist"},
		{"Rifleman Bravo", "Private"},
	}

	for _, v := range formation {
		b.SquadOnBuild.AddSquadMember(&Soldier{Role: v[0], Rank: v[1]})
	}
}

type MarinesBuilder struct {
	SquadBuilder
}

func (b *MarinesBuilder) SetName() {
	b.SquadOnBuild.SetSquadName("Marine Rifle Squad")
}

func (b *MarinesBuilder) BuildFormation() {
	formation := [][]string{
		{"Squad Leader", "Sergeant"},
		{"Assistant Squad Leader", "Corporal"},
		{"Squad Systems Operator", "Private"},
		{"Team Leader Alpha", "Corporal"},
		{"Automatic Rifleman Alpha", "Private"},
		{"Grenadier Alpha", "Private"},
		{"Rifleman Alpha", "Private"},
		{"Team Leader Bravo", "Corporal"},
		{"Automatic Rifleman Bravo", "Private"},
		{"Grenadier Bravo", "Private"},
		{"Rifleman Bravo", "Private"},
		{"Team Leader Charlie", "Corporal"},
		{"Automatic Rifleman Charlie", "Private"},
		{"Grenadier Charlie", "Private"},
		{"Rifleman Charlie", "Private"},
	}

	for _, v := range formation {
		b.SquadOnBuild.AddSquadMember(&Soldier{Role: v[0], Rank: v[1]})
	}
}

type Pentagon struct {
	squadBuilder SquadBuilderInterface
}

func (f *Pentagon) SetSquadBuilder(builderInterface SquadBuilderInterface) {
	f.squadBuilder = builderInterface
}

func (f *Pentagon) GetSquad() *Squad {
	return f.squadBuilder.GetSquad()
}

func (f *Pentagon) BuildSquad() {
	f.squadBuilder.NewSquad()
	f.squadBuilder.SetName()
	f.squadBuilder.BuildFormation()
}

func main() {
	var pentagon Pentagon

	var rangers RangersBuilder
	pentagon.SetSquadBuilder(&rangers)
	pentagon.BuildSquad()
	squad := pentagon.GetSquad()
	squad.SitRep()

	fmt.Println()

	var marines MarinesBuilder
	pentagon.SetSquadBuilder(&marines)
	pentagon.BuildSquad()
	squad = pentagon.GetSquad()
	squad.SitRep()
}
