package progress_bar

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	// padding  = 2
	maxWidth = 80
)

var quitMessage = tea.Sequence(tea.ShowCursor, tea.Quit)

var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF")).Render

type model struct {
	duration   time.Duration
	amount     float64
	percentage float64
	p          progress.Model
	// channel    chan bool
}

func New(duration time.Duration) *model {
	return &model{
		duration:   duration,
		amount:     1 / duration.Seconds(),
		percentage: 0,
		p:          progress.New(progress.WithSolidFill("#FFFFFF")),
		// channel:    make(chan bool, 1),
	}
}

func (m model) Init() tea.Cmd {
	return tickCommand()
}

func (m model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch message := message.(type) {
	case tea.KeyMsg:
		return m, quitMessage
	case tea.WindowSizeMsg:
		m.p.Width = message.Width
		if m.p.Width > maxWidth {
			m.p.Width = maxWidth
		}
		return m, nil
	case time.Time:
		m.percentage += m.amount
		if m.percentage >= 1.0 {
			return m, quitMessage
		}
		return m, tickCommand()
	default:
		return m, nil
	}

}

func (m model) View() string {
	return fmt.Sprintf("%s\n%s/%s\n%s", m.p.ViewAs(m.percentage), time.Duration(float64(m.duration)*m.percentage).Round(time.Second), m.duration, helpStyle("Press any key to quit"))
}

func tickCommand() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return t
	})
}
