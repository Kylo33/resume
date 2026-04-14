package resume

import (
	"fmt"
	"image/color"
	"log"
	"strings"

	// "github.com/charmbracelet/log"
	"charm.land/glamour/v2"
	"charm.land/lipgloss/v2"
	"charm.land/lipgloss/v2/list"
)

const (
	educationColor = lipgloss.Yellow
	workColor      = lipgloss.Magenta
	projectColor   = lipgloss.Green
)

func (r Resume) Format(w int) string {
	var (
		marginSize = 1
	)

	contactBlock := r.FormatContactBlock(w - marginSize*2)
	educationBlock := r.FormatEducation(w - marginSize*2)

	return lipgloss.NewStyle().
		Margin(1).
		Render(
			lipgloss.JoinVertical(
				lipgloss.Left,
				contactBlock,
				pageSection("Education", w-marginSize*2, educationColor),
				educationBlock,
				pageSection("Work Experience", w-marginSize*2, workColor),
				pageSection("Projects", w-marginSize*2, projectColor),
			),
		)

}

func (r Resume) FormatEducation(w int) string {
	var blocks []string
	for _, edu := range r.Education {
		institutiton := lipgloss.
			NewStyle().
			Foreground(educationColor).
			Bold(true).
			Render(edu.Institution)

		degreeLine := lipgloss.
			NewStyle().
			Foreground(lipgloss.White).
			Italic(true).
			Render(
				fmt.Sprintf("%v, %v GPA", edu.Degree, edu.Gpa),
			)

		extra := renderExtra(edu.Extra)

		container := lipgloss.
			NewStyle().
			MarginBottom(1).
			Render(
				lipgloss.JoinVertical(
					lipgloss.Left,
					institutiton,
					degreeLine,
					extra,
				),
			)

		blocks = append(blocks, container)
	}

	return lipgloss.JoinVertical(lipgloss.Left, blocks...)
}

func (r Resume) FormatContactBlock(w int) string {
	name := lipgloss.
		NewStyle().
		Foreground(lipgloss.Blue).
		Background(lipgloss.Black).
		PaddingLeft(1).
		PaddingRight(1).
		MarginBottom(1).
		Bold(true).
		Render(r.Name)

	name = lipgloss.PlaceHorizontal(w, lipgloss.Center, name)

	contactStyle := lipgloss.
		NewStyle().
		Italic(true).
		Foreground(lipgloss.White)

	contactLinkStyle := contactStyle.Underline(true)

	divider := lipgloss.
		NewStyle().
		Foreground(lipgloss.Black).
		MarginLeft(1).
		MarginRight(1).
		Render("|")

	contactBlock := lipgloss.JoinHorizontal(
		lipgloss.Bottom,
		contactStyle.Render(r.Phone),
		divider,
		contactStyle.Render(r.Location),
		divider,
		contactLinkStyle.Hyperlink(r.Email).Render(r.Email),
		divider,
		contactLinkStyle.Hyperlink(r.Website).Render(r.Website),
		divider,
		contactLinkStyle.Hyperlink(r.Github).Render(r.Github),
	)

	contactBlock = lipgloss.PlaceHorizontal(w, lipgloss.Center, contactBlock)

	return lipgloss.JoinVertical(lipgloss.Center, name, contactBlock)
}

func pageSection(name string, w int, color color.Color) string {
	sectionName := lipgloss.
		NewStyle().
		Foreground(color).
		Background(lipgloss.Black).
		PaddingLeft(1).
		PaddingRight(1).
		Bold(true).
		Render(name)

	container := lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderBottom(true).
		BorderForeground(lipgloss.Black).
		Render(
			lipgloss.PlaceHorizontal(w, lipgloss.Left, sectionName),
		)

	return container
}

func renderExtra(content []string) string {
	rn, err := glamour.NewTermRenderer(
		glamour.WithWordWrap(10000),
		glamour.WithStylePath("dark"),
	)
	if err != nil {
		log.Fatal("Failed to create glamour renderer", "err", err)
	}

	var renderedStrings []string
	for _, item := range content {
		renderedItem, err := rn.Render(item)
		// TODO: switch to charm logging library w/o crashing
		if err != nil {
			log.Fatal("Failed to parse markdown", "err", err)
		}
		renderedStrings = append(renderedStrings, strings.TrimSpace(renderedItem))
	}

	l := list.New()
	for _, renderedItem := range renderedStrings {
		l.Item(renderedItem)
	}
	return l.String()
}
