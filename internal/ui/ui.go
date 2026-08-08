package ui

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/fatih/color"
)

var (
	Red    = color.New(color.FgRed, color.Bold).SprintFunc()
	Green  = color.New(color.FgGreen, color.Bold).SprintFunc()
	Yellow = color.New(color.FgYellow).SprintFunc()
	Bold   = color.New(color.Bold).SprintFunc()
	Cyan   = color.New(color.FgCyan, color.Bold).SprintFunc()
	Gray   = color.New(color.FgHiBlack).SprintFunc()
)

func PrintError(msg string) {
	fmt.Fprintf(os.Stderr, "❌ %s\n", Red(msg))
}

func PrintSuccess(msg string) {
	fmt.Printf("%s %s\n", Green("✔"), msg)
}

func PrintInfo(msg string) {
	fmt.Println(msg)
}

type TableRow struct {
	Active   string
	Username string
	Name     string
	Email    string
}

var ansiRegexp = regexp.MustCompile(`\x1b\[[0-9;]*[a-zA-Z]`)

func visibleLen(s string) int {
	clean := ansiRegexp.ReplaceAllString(s, "")
	return utf8.RuneCountInString(clean)
}

func padRight(s string, width int) string {
	vLen := visibleLen(s)
	if vLen >= width {
		return s
	}
	return s + strings.Repeat(" ", width-vLen)
}

func padCenter(s string, width int) string {
	vLen := visibleLen(s)
	if vLen >= width {
		return s
	}
	totalPad := width - vLen
	leftPad := totalPad / 2
	rightPad := totalPad - leftPad
	return strings.Repeat(" ", leftPad) + s + strings.Repeat(" ", rightPad)
}

func PrintAccountTable(w io.Writer, headers []string, rows []TableRow) {
	if len(headers) < 4 {
		return
	}

	widths := []int{
		visibleLen(headers[0]),
		visibleLen(headers[1]),
		visibleLen(headers[2]),
		visibleLen(headers[3]),
	}

	for _, r := range rows {
		activeStr := r.Active
		userStr := r.Username

		if visibleLen(activeStr) > widths[0] {
			widths[0] = visibleLen(activeStr)
		}
		if visibleLen(userStr) > widths[1] {
			widths[1] = visibleLen(userStr)
		}
		if visibleLen(r.Name) > widths[2] {
			widths[2] = visibleLen(r.Name)
		}
		if visibleLen(r.Email) > widths[3] {
			widths[3] = visibleLen(r.Email)
		}
	}

	if widths[0] < 6 {
		widths[0] = 6
	}
	if widths[1] < 10 {
		widths[1] = 10
	}
	if widths[2] < 12 {
		widths[2] = 12
	}
	if widths[3] < 15 {
		widths[3] = 15
	}

	border := Gray("─")
	topLine := Gray("┌") + strings.Repeat(border, widths[0]+2) + Gray("┬") + strings.Repeat(border, widths[1]+2) + Gray("┬") + strings.Repeat(border, widths[2]+2) + Gray("┬") + strings.Repeat(border, widths[3]+2) + Gray("┐")
	sepLine := Gray("├") + strings.Repeat(border, widths[0]+2) + Gray("┼") + strings.Repeat(border, widths[1]+2) + Gray("┼") + strings.Repeat(border, widths[2]+2) + Gray("┼") + strings.Repeat(border, widths[3]+2) + Gray("┤")
	botLine := Gray("└") + strings.Repeat(border, widths[0]+2) + Gray("┴") + strings.Repeat(border, widths[1]+2) + Gray("┴") + strings.Repeat(border, widths[2]+2) + Gray("┴") + strings.Repeat(border, widths[3]+2) + Gray("┘")

	vline := Gray("│")

	fmt.Fprintln(w, topLine)
	fmt.Fprintf(w, "%s %s %s %s %s %s %s %s %s\n",
		vline, padCenter(Cyan(headers[0]), widths[0]),
		vline, padRight(Cyan(headers[1]), widths[1]),
		vline, padRight(Cyan(headers[2]), widths[2]),
		vline, padRight(Cyan(headers[3]), widths[3]),
		vline,
	)
	fmt.Fprintln(w, sepLine)

	for _, r := range rows {
		activeDisp := " "
		if r.Active == "*" {
			activeDisp = Green("*")
		}
		userDisp := r.Username
		if r.Active == "*" {
			userDisp = Green(r.Username)
		}

		fmt.Fprintf(w, "%s %s %s %s %s %s %s %s %s\n",
			vline, padCenter(activeDisp, widths[0]),
			vline, padRight(userDisp, widths[1]),
			vline, padRight(r.Name, widths[2]),
			vline, padRight(r.Email, widths[3]),
			vline,
		)
	}

	fmt.Fprintln(w, botLine)
}
