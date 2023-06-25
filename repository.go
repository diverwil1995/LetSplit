package main

type GetGroupsFilter struct {
}
type GetMembersFilter struct {
}
type GetExpensesFilter struct {
}

type LetSplitRepository interface {
	AuthUser(*User) (bool, error)
	CreateUser(*User) error
	DeleteUser(*User) error
	CreateGroup(*Group) error
	GetGroups(*GetGroupsFilter) ([]Group, error)
	UpdateGroup(*Group) error
	DeleteGroup(*Group) error
	AddMember(*Member) error
	GetMembers(*GetMembersFilter) ([]Member, error)
	UpdateMember(*Member) error
	DeleteMember(*Member) error
	AddExpense(*Expense) error
	GetExpenses(*GetExpensesFilter) ([]Expense, error)
	UpdateExpense(*Expense) error
	DeleteExpense(*Expense) error
}
