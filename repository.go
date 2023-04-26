package main

type GetGroupsFilter struct {
}

type LetSplitRepository interface {
	CreateGroup(*Group) error
	GetGroups(*GetGroupsFilter) ([]Group, error)
	DeleteGroup(*Group) error
}
