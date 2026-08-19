package ui

import "github.com/charmbracelet/lipgloss"

var (
	// Palette Colors
	ColorPrimary   = lipgloss.Color("208") // Mozilla Orange
	ColorSuccess   = lipgloss.Color("42")  // Green
	ColorFailure   = lipgloss.Color("196") // Red
	ColorPending   = lipgloss.Color("240") // Dark Gray
	ColorHighlight = lipgloss.Color("51")  // Cyan / Bright Blue
	ColorBg        = lipgloss.Color("234") // Dark Background
	ColorCardBg    = lipgloss.Color("235") // Card Background

	// Main Header & Outer Frame
	HeaderStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(ColorPrimary).
			Background(ColorBg).
			Padding(0, 1).
			MarginBottom(0)

	TitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("255")).
			Background(ColorPrimary).
			Padding(0, 2)

	ProgressBarFullStyle = lipgloss.NewStyle().
				Foreground(ColorSuccess)

	ProgressBarEmptyStyle = lipgloss.NewStyle().
				Foreground(ColorPending)

	// Panel Styles Base
	LeftPanelStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("240")).
			Padding(0, 1).
			MarginRight(1)

	RightPanelStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("240")).
			Padding(0, 1)

	// Card Matrix Styles
	CardWidth  = 16
	CardHeight = 4

	CardPassedStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(ColorSuccess).
			Foreground(ColorSuccess).
			Background(ColorCardBg).
			Padding(0, 1).
			Width(CardWidth).
			Height(CardHeight).
			Align(lipgloss.Center, lipgloss.Center)

	CardFailedStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(ColorFailure).
			Foreground(ColorFailure).
			Background(ColorCardBg).
			Padding(0, 1).
			Width(CardWidth).
			Height(CardHeight).
			Align(lipgloss.Center, lipgloss.Center)

	CardPendingStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(ColorPending).
			Foreground(lipgloss.Color("246")).
			Background(ColorCardBg).
			Padding(0, 1).
			Width(CardWidth).
			Height(CardHeight).
			Align(lipgloss.Center, lipgloss.Center)

	CardFocusedStyle = lipgloss.NewStyle().
			Border(lipgloss.DoubleBorder()).
			BorderForeground(ColorHighlight).
			Bold(true).
			Foreground(lipgloss.Color("255")).
			Background(lipgloss.Color("237")).
			Padding(0, 1).
			Width(CardWidth).
			Height(CardHeight).
			Align(lipgloss.Center, lipgloss.Center)

	// Detail View Styles
	SectionTitleStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(ColorHighlight).
				MarginBottom(1)

	BadgeBeginnerStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("42")).
				Bold(true)

	BadgeIntermediateStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("220")).
				Bold(true)

	BadgeAdvancedStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("202")).
				Bold(true)

	FoxBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(ColorPrimary).
			Foreground(ColorPrimary).
			Padding(0, 1).
			MarginTop(1).
			MarginBottom(1).
			Align(lipgloss.Center)

	// Sidebar Styles
	ChapterColWidth = 18
	LessonColWidth  = 26

	SidebarStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(ColorPending).
			Padding(0, 1).
			MarginRight(1)

	SidebarFocusedStyle = lipgloss.NewStyle().
				Border(lipgloss.DoubleBorder()).
				BorderForeground(ColorHighlight).
				Padding(0, 1).
				MarginRight(1)

	SidebarItemStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("246"))

	SidebarItemFocusStyle = lipgloss.NewStyle().
				Foreground(ColorHighlight).
				Background(lipgloss.Color("237"))

	SidebarSelectedStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("255")).
				Background(ColorPrimary)

	ColumnTitleStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(ColorHighlight).
				Underline(true).
				MarginBottom(1)

	// Footer / Controls
	FooterStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("240")).
			Foreground(lipgloss.Color("248")).
			Padding(0, 1).
			MarginTop(0)
)
