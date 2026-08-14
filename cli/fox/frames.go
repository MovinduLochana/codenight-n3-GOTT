package fox

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type TickMsg time.Time

func TickCmd() tea.Cmd {
	return tea.Tick(120*time.Millisecond, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

var FoxFrames = []string{
	`
      /\_/\   
    =( °.o )=  ~~> [ Running tests... ]
      (   )─/ 
`,
	`
      /\_/\   
    =( ⊙.⊙ )=   ~~> [ Running tests... ]
      /   \  
`,
	`
      /\_/\   
    =( ^.^ )=    ~~> [ Running tests... ]
     /  |  \ 
`,
	`
      /\_/\   
    =( >.< )=     ~~> [ Running tests... ]
      \   /  
`,
}

func GetFrame(index int) string {
	if len(FoxFrames) == 0 {
		return ""
	}
	return FoxFrames[index%len(FoxFrames)]
}
