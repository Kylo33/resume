package resume

import "charm.land/lipgloss/v2"

const projectColor = lipgloss.Green

func (r Resume) FormatProjects(w int) string {
	var blocks []string

	for _, project := range r.Projects {
		name := renderHeader(2, project.Name, projectColor)
		dates := renderHeader(3, project.Date.String(), projectColor)
		 
	}

	return lipgloss.JoinVertical(lipgloss.Left, blocks...)
}
