package models

import "fmt"

// StashEntry : A git stash entry
type StashEntry struct {
	Index   int
	Recency string
	Name    string
}

func (s *StashEntry) FullRefName() string {
	return "refs/" + s.RefName()
}

func (s *StashEntry) RefName() string {
	return fmt.Sprintf("stash@{%d}", s.Index)
}

func (s *StashEntry) ShortRefName() string {
	return s.RefName()
}

func (s *StashEntry) ParentRefName() string {
	return s.RefName() + "^"
}

func (s *StashEntry) ID() string {
	return s.RefName()
}

func (s *StashEntry) Description() string {
	return s.RefName() + ": " + s.Name
}
