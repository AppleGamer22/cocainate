package progressbar_test

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	padding  = 2
	maxWidth = 80
)

var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")).Render

type model struct {
	duration   time.Duration
	amount     float64
	percentage float64
	p          progress.Model
	channel    chan bool
}

func New(duration time.Duration) *model {
	return &model{
		duration:   duration,
		amount:     float64(duration/time.Second) / float64(duration),
		percentage: 0.0,
		p:          progress.New(progress.WithScaledGradient("#FF7CCB", "#FDFF8C")),
		channel:    make(chan bool, 1),
	}
}

func (m model) Init() tea.Cmd {
	return tickCommand()
}

func (m model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch message := message.(type) {
	case tea.KeyMsg:
		return m, tea.Quit
	case tea.WindowSizeMsg:
		m.p.Width = message.Width
		if m.p.Width > maxWidth {
			m.p.Width = maxWidth
		}
		return m, nil
	case time.Time:
		m.percentage += m.amount
		if m.percentage > 1.0 {
			m.percentage = 1.0
			return m, tea.Quit
		}
		return m, tickCommand()
	default:
		return m, nil
	}

}

func (m model) View() string {
	return fmt.Sprintf("\n%s\n%s/%s\n%s", m.p.ViewAs(m.percentage), time.Duration(float64(m.duration)*m.percentage), m.duration, helpStyle("Press any key to quit"))
}

func tickCommand() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return t
	})
}
