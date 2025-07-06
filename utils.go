package koiApi

import "fmt"

func (a *Album) Summary() string {
	return fmt.Sprintf("%-40s %s", a.Title, a.ID)
}

func (cl *ChoiceList) Summary() string {
	return fmt.Sprintf("%-40s %s", cl.Name, cl.ID)
}

func (c *Collection) Summary() string {
	return fmt.Sprintf("%-40s %s", c.Title, c.ID)
}

func (d *Datum) Summary() string {
	return fmt.Sprintf("%-40s %s", d.Label, d.ID)
}

func (f *Field) Summary() string {
	return fmt.Sprintf("%-40s %s", f.Name, f.ID)
}

func (i *Inventory) Summary() string {
	return fmt.Sprintf("%-40s %s", i.Name, i.ID)
}

func (i *Item) Summary() string {
	return fmt.Sprintf("%-40s %s", i.Name, i.ID)
}

func (l *Loan) Summary() string {
	return fmt.Sprintf("%-40s %s", l.LentTo, l.ID)
}

func (l *Log) Summary() string {
	return fmt.Sprintf("%-40s %s", l.ObjectLabel, l.ID)
}

func (m *Metrics) Summary() string {
	return fmt.Sprintf("%-40s %s", string(*m), "")
}

func (p *Photo) Summary() string {
	return fmt.Sprintf("%-40s %s", p.Title, p.ID)
}

func (t *Tag) Summary() string {
	return fmt.Sprintf("%-40s %s", t.Label, t.ID)
}

func (tc *TagCategory) Summary() string {
	return fmt.Sprintf("%-40s %s", tc.Label, tc.ID)
}

func (t *Template) Summary() string {
	return fmt.Sprintf("%-40s %s", t.Name, t.ID)
}

func (u *User) Summary() string {
	return fmt.Sprintf("%-40s %s", u.Username, u.ID)
}

func (w *Wish) Summary() string {
	return fmt.Sprintf("%-40s %s", w.Name, w.ID)
}

func (w *Wishlist) Summary() string {
	return fmt.Sprintf("%-40s %s", w.Name, w.ID)
}
