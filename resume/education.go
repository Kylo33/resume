package resume

import (
	"fmt"

	"charm.land/lipgloss/v2"
)

const educationColor = lipgloss.Yellow

func (r Resume) FormatEducation(w int) string {
	var blocks []string
	for _, edu := range r.Education {
		institution := renderHeader(2, edu.Institution, educationColor)
		dates := renderHeader(3, datesHelper(edu.StartDate, edu.EndDate), educationColor)

		degree := lipgloss.
			NewStyle().
			Foreground(lipgloss.White).
			Italic(true).
			Render(
				fmt.Sprintf(
					"%v, %v GPA",
					edu.Degree,
					edu.Gpa,
				),
			)

		location := lipgloss.
			NewStyle().
			Foreground(lipgloss.White).
			Render(edu.Location)

		extra := renderExtra(edu.Extra, w)

		container := lipgloss.
			NewStyle().
			MarginBottom(1).
			Render(
				lipgloss.JoinVertical(
					lipgloss.Left,
					sameLine(institution, dates, w),
					sameLine(degree, location, w),
					extra,
				),
			)

		blocks = append(blocks, container)
	}

	return lipgloss.JoinVertical(lipgloss.Left, blocks...)
}
