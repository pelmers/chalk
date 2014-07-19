package chalk

import "fmt"

// Color represents one of the ANSI color escape codes.
// http://en.wikipedia.org/wiki/ANSI_escape_code#Colors
type Color struct {
	value int
}

// Value returns the individual value for this color
// (Actually it's really just it's index in the list
// of color escape codes with the list being
// [black, red, green, yellow, blue, magenta, cyan, white].
func (c Color) Value() int {
	return c.value
}

// Color colors the foreground of the given string
// (whatever the previou background color was, it is
// left alone).
func (c Color) Color(val string) string {
	return fmt.Sprintf("%s%s%s", c, val, ResetColor)
}

func (c Color) String() string {
	return fmt.Sprintf("\u001b[%dm", 30+c.value)
}

// NewStyle creates a style with a foreground of the
// color we're creating the style from.
func (c Color) NewStyle() Style {
	return &style{foreground: c}
}

// A Style is how we want our text to look in the console.
// Consequently, we can set the foreground and background
// to specific colors, we can style specific strings and
// can also use this style in a builder pattern should we
// wish (these will be more useful once styles such as
// italics are supported).
type Style interface {
	// Foreground sets the foreground of the style to the specific color.
	Foreground(Color)
	// Background sets the background of the style to the specific color.
	Background(Color)
	// Style styles the given string with the curreny style.
	Style(string) string
	// WithBackground allows us to set the background in a builder
	// pattern style.
	WithBackground(Color) Style
	// WithForeground allows us to set the foreground in a builder
	// pattern style.
	WithForeground(Color) Style
}

type style struct {
	foreground Color
	background Color
	// TODO(ttacon): add styles at some point (when we care enough about them)
}

func (s *style) WithBackground(col Color) Style {
	s.Background(col)
	return s
}

func (s *style) WithForeground(col Color) Style {
	s.Foreground(col)
	return s
}

func (s *style) String() string {
	var toReturn string
	toReturn = fmt.Sprintf("\u001b[%dm", 40+s.background.Value())
	return toReturn + fmt.Sprintf("\u001b[%dm", 30+s.foreground.Value())
}

func (s *style) Style(val string) string {
	return fmt.Sprintf("%s%s%s", s, val, Reset)
}

func (s *style) Foreground(col Color) {
	s.foreground = col
}

func (s *style) Background(col Color) {
	s.background = col
}

var (
	nine = 9

	Black      = Color{0}
	Red        = Color{1}
	Green      = Color{2}
	Yellow     = Color{3}
	Blue       = Color{4}
	Magenta    = Color{5}
	Cyan       = Color{6}
	White      = Color{7}
	ResetColor = Color{9}

	Reset = &style{ResetColor, ResetColor}
)