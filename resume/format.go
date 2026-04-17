package resume

import (
	"fmt"
	"image/color"
	"log"
	"strings"
	"time"

	"charm.land/glamour/v2"
	"charm.land/lipgloss/v2"
)

const (
	nameColor    = lipgloss.Green
	projectColor = lipgloss.Green
)

func (r Resume) Format(w int) string {
	var (
		marginSize = 2
	)

	contactBlock := r.FormatContactBlock(w - marginSize*2)
	educationBlock := r.FormatEducation(w - marginSize*2)
	workBlock := r.FormatWork(w - marginSize*2)

	spacer := "\n"

	return lipgloss.NewStyle().
		Margin(marginSize).
		Render(
			lipgloss.JoinVertical(
				lipgloss.Left,
				contactBlock,
				spacer,
				pageSection("Education", w-marginSize*2, educationColor),
				educationBlock,
				pageSection("Work Experience", w-marginSize*2, workColor),
				workBlock,
				pageSection("Projects", w-marginSize*2, projectColor),
			),
		)

}

func (r Resume) FormatContactBlock(w int) string {
	name := lipgloss.PlaceHorizontal(
		w,
		lipgloss.Center,
		renderHeader(1, r.Name, nameColor),
	)

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
	sectionName := renderHeader(1, name, color)
	container := lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderBottom(true).
		BorderForeground(lipgloss.Black).
		Render(
			lipgloss.PlaceHorizontal(w, lipgloss.Left, sectionName),
		)

	return container
}

func renderExtra(content []string, w int) string {
	bulletPoint := lipgloss.
		NewStyle().
		MarginRight(1).
		Render("•")

	rn, err := glamour.NewTermRenderer(
		glamour.WithWordWrap(w-lipgloss.Width(bulletPoint)),
		glamour.WithStylePath("resume/stylesheet.json"),
	)

	if err != nil {
		log.Fatal("Failed to create glamour renderer", "err", err)
	}

	var bullets []string
	for i, markdownItem := range content {
		renderedItem, err := rn.Render(markdownItem)
		renderedItem = strings.TrimSpace(renderedItem)

		if err != nil {
			log.Fatal("Failed to parse markdown", "err", err)
		}

		// Add spacing between each bullet
		containerStyle := lipgloss.NewStyle()
		// if i < len(content)-1 {
		// 	containerStyle = containerStyle.MarginBottom(1)
		// }
		_ = i

		bullet := containerStyle.Render(
			lipgloss.JoinHorizontal(
				lipgloss.Top,
				bulletPoint,
				renderedItem,
			),
		)

		bullets = append(bullets, bullet)
	}

	return lipgloss.JoinVertical(lipgloss.Left, bullets...)
}

func renderHeader(level int, text string, highlightColor color.Color) string {
	style := lipgloss.NewStyle()

	switch level {
	case 1:
		style = style.
			Background(lipgloss.Black).
			PaddingLeft(1).
			PaddingRight(1)
		fallthrough
	case 2:
		style = style.Bold(true)
		fallthrough
	case 3:
		style = style.Foreground(highlightColor)
	default:
		log.Fatal("Invalid header level", "level", level)
	}

	return style.Render(text)
}

func datesHelper(start, end YearMonthTime) string {
	startString := start.String()
	endString := end.String()

	// If end is before start, the string should be "Present"
	if time.Time(start).Compare(time.Time(end)) > 0 {
		endString = "Present"
	}

	return fmt.Sprintf("%v – %v", startString, endString)
}

func sameLine(start, end string, width int) string {
	spacer := strings.Repeat(" ", width - lipgloss.Width(start) - lipgloss.Width(end))
	return lipgloss.JoinHorizontal(lipgloss.Top, start, spacer, end)
}
