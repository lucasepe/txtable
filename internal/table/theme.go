package table

var (
	// ThemeClassic - MySql-like table style:
	//
	// +---+------------------+------------+
	// | # |       NAME       |    TAX     |
	// +---+------------------+------------+
	// | 1 | Newton G. Goetz  |   $ 532.70 |
	// | 2 | Rebecca R. Edney |  $ 1423.25 |
	// | 3 | John R. Jackson  |  $ 7526.12 |
	// | 4 | Ron J. Gomes     |   $ 123.84 |
	// | 5 | Penny R. Lewis   |  $ 3221.11 |
	// +---+------------------+------------+
	ThemeClassic = &Theme{
		Border: &BorderStyle{
			TopLeft:            "+",
			Top:                "-",
			TopRight:           "+",
			Right:              "|",
			BottomRight:        "+",
			Bottom:             "-",
			BottomLeft:         "+",
			Left:               "|",
			TopIntersection:    "+",
			BottomIntersection: "+",
		},
		Divider: &DividerStyle{
			Left:         "+",
			Center:       "-",
			Right:        "+",
			Intersection: "+",
		},
		Cell: "|",
	}

	// ThemeCompact - compact table style:
	//
	//  #         NAME            TAX
	// --- ------------------ ------------
	//  1   Newton G. Goetz      $ 532.70
	//  2   Rebecca R. Edney    $ 1423.25
	//  3   John R. Jackson     $ 7526.12
	//  4   Ron J. Gomes         $ 123.84
	//  5   Penny R. Lewis      $ 3221.11
	// --- ------------------ ------------
	ThemeCompact = &Theme{
		Border: &BorderStyle{
			TopLeft:            "",
			Top:                "",
			TopRight:           "",
			Right:              "",
			BottomRight:        "",
			Bottom:             "",
			BottomLeft:         "",
			Left:               "",
			TopIntersection:    "",
			BottomIntersection: "",
		},
		Divider: &DividerStyle{
			Left:         "",
			Center:       "-",
			Right:        "",
			Intersection: " ",
		},
		Cell: " ",
	}

	// ThemeLight - unicode light border table theme:
	ThemeLight = &Theme{
		Border: &BorderStyle{
			TopLeft:            "┌",
			Top:                "─",
			TopRight:           "┐",
			Right:              "│",
			BottomRight:        "┘",
			Bottom:             "─",
			BottomLeft:         "└",
			Left:               "│",
			TopIntersection:    "┬",
			BottomIntersection: "┴",
		},
		Divider: &DividerStyle{
			Left:         "├",
			Center:       "─",
			Right:        "┤",
			Intersection: "┼",
		},
		Cell: "│",
	}

	// ThemeHeavy - unicode heavy border table theme:
	ThemeHeavy = &Theme{
		Border: &BorderStyle{
			TopLeft:            "┏",
			Top:                "━",
			TopRight:           "┓",
			Right:              "┃",
			BottomRight:        "┛",
			Bottom:             "━",
			BottomLeft:         "┗",
			Left:               "┃",
			TopIntersection:    "┯", //"┳",
			BottomIntersection: "┷", // "┻",
		},
		Divider: &DividerStyle{
			Left:         "┠", //"┣",
			Center:       "─",
			Right:        "┨", //"┫",
			Intersection: "┼", //"╋",
		},
		Cell: "│",
	}

	// ThemeRounded - unicode rounded border table theme:
	ThemeRounded = &Theme{
		Border: &BorderStyle{
			TopLeft:            "╭",
			Top:                "─",
			TopRight:           "╮",
			Right:              "│",
			BottomRight:        "╯",
			Bottom:             "─",
			BottomLeft:         "╰",
			Left:               "│",
			TopIntersection:    "┬",
			BottomIntersection: "┴",
		},
		Divider: &DividerStyle{
			Left:         "├",
			Center:       "─",
			Right:        "┤",
			Intersection: "┼",
		},
		Cell: "│",
	}

	Themes = map[string]*Theme{
		"classic": ThemeClassic,
		"compact": ThemeCompact,
		"light":   ThemeLight,
		"heavy":   ThemeHeavy,
		"rounded": ThemeRounded,
	}
)

// BorderStyle defines table border style
type BorderStyle struct {
	TopLeft            string
	Top                string
	TopRight           string
	Right              string
	BottomRight        string
	Bottom             string
	BottomLeft         string
	Left               string
	TopIntersection    string
	BottomIntersection string
}

// DividerStyle defines table divider style
type DividerStyle struct {
	Left         string
	Center       string
	Right        string
	Intersection string
}

// Theme is a table style (borders, dividers etc)
type Theme struct {
	Border  *BorderStyle
	Divider *DividerStyle
	Cell    string // Symbol between cells
}
