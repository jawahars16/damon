package component

import (
	"fmt"

	"github.com/rivo/tview"

	"github.com/hcjulz/damon/primitives"
	"github.com/hcjulz/damon/state"
	"github.com/hcjulz/damon/styles"

	"github.com/hcjulz/damon/models"
)

var (
	labelNamespaceDropdown = fmt.Sprintf("%sNamespace <s>: ▾ %s",
		styles.HighlightSecondaryTag,
		styles.StandardColorTag,
	)
)

type Selections struct {
	Namespace DropDown
	Props     *SelectionsProps

	state *state.State
	slot  *tview.Flex
}

type SelectionsProps struct {
	Rerender func()
}

func NewSelections(state *state.State) *Selections {
	return &Selections{
		Namespace: primitives.NewDropDown(labelNamespaceDropdown),
		Props:     &SelectionsProps{},
		state:     state,
	}
}

func (s *Selections) Render() error {
	if s.Props.Rerender == nil {
		return ErrComponentPropsNotSet
	}

	if s.slot == nil {
		return ErrComponentNotBound
	}

	s.Namespace.SetOptions(convert(s.state.Namespaces), s.selected)
	s.Namespace.SetCurrentOption(len(s.state.Namespaces) - 1)
	s.Namespace.SetSelectedFunc(s.rerender)

	s.state.Elements.DropDownNamespace = s.Namespace.Primitive().(*tview.DropDown)
	s.slot.AddItem(s.Namespace.Primitive(), 0, 1, true)

	return nil
}

func (s *Selections) Bind(slot *tview.Flex) {
	s.slot = slot
}

func (s *Selections) selected(text string, index int) {
	s.state.SelectedNamespace = text
}

func (s *Selections) rerender(text string, index int) {
	s.state.SelectedNamespace = text
	s.Props.Rerender()
}

func convert(list []*models.Namespace) []string {
	var ns []string
	for _, n := range list {
		ns = append(ns, n.Name)
	}
	return ns
}
