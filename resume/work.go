package resume

import "charm.land/lipgloss/v2"

const workColor = lipgloss.Magenta

func (r Resume) FormatWork(w int) string {
	var blocks []string
	for _, job := range r.Work {
		title := renderHeader(2, job.Title, workColor)
		dates := renderHeader(3, datesHelper(job.StartDate, job.EndDate), workColor)

		container := lipgloss.
			NewStyle().
			MarginBottom(1).
			Render(
				lipgloss.JoinVertical(
					lipgloss.Left,
					sameLine(title, dates, w),
					renderExtra(job.Extra, w),
				),
			)

		blocks = append(blocks, container)
	}

	return lipgloss.JoinVertical(lipgloss.Left, blocks...)
}
