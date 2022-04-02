package view

import (
	"context"

	"github.com/derailed/k9s/internal/client"
	"github.com/derailed/k9s/internal/dao"
	"github.com/derailed/k9s/internal/model"
	"github.com/derailed/k9s/internal/render"
	"github.com/derailed/k9s/internal/ui"
	"github.com/gdamore/tcell/v2"
)

// Helm represents a helm chart view.
type Helm struct {
	ResourceViewer
}

// NewHelm returns a new alias view.
func NewHelm(gvr client.GVR) ResourceViewer {
	c := Helm{
		ResourceViewer: NewBrowser(gvr),
	}
	c.GetTable().SetColorerFn(render.Helm{}.ColorerFunc())
	c.GetTable().SetBorderFocusColor(tcell.ColorMediumSpringGreen)
	c.GetTable().SetSelectedStyle(tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorMediumSpringGreen).Attributes(tcell.AttrNone))
	c.AddBindKeysFn(c.bindKeys)
	c.SetContextFn(c.chartContext)

	return &c
}

func (c *Helm) chartContext(ctx context.Context) context.Context {
	return ctx
}

func (c *Helm) bindKeys(aa ui.KeyActions) {
	aa.Delete(ui.KeyShiftA, ui.KeyShiftN, tcell.KeyCtrlS, tcell.KeyCtrlSpace, ui.KeySpace)
	aa.Add(ui.KeyActions{
		ui.KeyShiftN: ui.NewKeyAction("Sort Name", c.GetTable().SortColCmd(nameCol, true), false),
		ui.KeyShiftS: ui.NewKeyAction("Sort Status", c.GetTable().SortColCmd(statusCol, true), false),
		ui.KeyShiftA: ui.NewKeyAction("Sort Age", c.GetTable().SortColCmd(ageCol, true), false),
		ui.KeyV:      ui.NewKeyAction("Values Get", c.getValsCmd(false), true),
		ui.KeyA:      ui.NewKeyAction("Values Get All", c.getValsCmd(true), true),
	})
}

func (c *Helm) getHelmDao() *dao.Helm {
	return model.Registry["helm"].DAO.(*dao.Helm)
}

func (c *Helm) getValsCmd(allValues bool) func(evt *tcell.EventKey) *tcell.EventKey {
	return func(evt *tcell.EventKey) *tcell.EventKey {
		path := c.GetTable().GetSelectedItem()
		vals, _ := c.getHelmDao().GetValues(path, allValues)
		v := NewLiveView(c.App(), "Values", model.NewValues(c.GVR(), path, vals))
		if err := v.app.inject(v); err != nil {
			v.app.Flash().Err(err)
		}
		return nil
	}
}
