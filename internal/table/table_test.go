package table

import (
	"fmt"
	"testing"
)

var (
	data = [][]interface{}{
		{1, "Newton G. Goetz", "252-585-5166", "NewtonGGoetz@dayrep.com", 10},
		{2, "Rebecca R. Edney", "865-475-4171", "RebeccaREdney@armyspy.com", 12},
		{3, "John R. Jackson", "810-325-1417", "JohnRJackson@armyspy.com", 15},
		{4, "Ron J. Gomes", "217-450-8568", "RonJGomes@rhyta.com", 25},
		{5, "Penny R. Lewis", "870-794-1666", "PennyRLewis@rhyta.com", 5},
		{6, "Sofia J. Smith", "770-333-7379", "SofiaJSmith@armyspy.com", 3},
		{7, "Karlene D. Owen", "231-242-4157", "KarleneDOwen@jourrapide.com", 12},
		{8, "Daniel L. Love", "978-210-4178", "DanielLLove@rhyta.com", 44},
		{9, "Julie T. Dial", "719-966-5354", "JulieTDial@jourrapide.com", 8},
		{10, "Juan J. Kennedy", "908-910-8893", "JuanJKennedy@dayrep.com", 16},
	}
)

func TestTable(t *testing.T) {
	tt := New()

	tt.Header = &Header{
		Cells: []*Cell{
			{Align: AlignCenter, Text: "#"},
			{Align: AlignCenter, Text: "NAME"},
			{Align: AlignCenter, Text: "PHONE"},
			{Align: AlignCenter, Text: "EMAIL"},
			{Align: AlignCenter, Text: "QTTY"},
		},
	}

	subtotal := 0
	for _, row := range data {
		r := []*Cell{
			{Align: AlignRight, Text: fmt.Sprintf("%d", row[0].(int))},
			{Text: row[1].(string)},
			{Text: row[2].(string)},
			{Text: row[3].(string)},
			{Align: AlignRight, Text: fmt.Sprintf("%d", row[4])},
		}

		tt.Body.Cells = append(tt.Body.Cells, r)
		subtotal += row[4].(int)
	}

	tt.Footer = &Footer{
		Cells: []*Cell{
			{Align: AlignRight, Span: 4, Text: "Subtotal"},
			{Align: AlignRight, Text: fmt.Sprintf("%d", subtotal)},
		},
	}

	for n, s := range Themes {
		fmt.Println(n)

		tt.SetTheme(s)
		tt.Println()

		fmt.Println()
	}
}
