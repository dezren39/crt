package bubbletea

import (
	"github.com/BigJk/crt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"unicode"
)

type teaKey struct {
	key  tea.KeyType
	rune []rune
}

var ebitenToTeaKeys = map[ebiten.Key]teaKey{
	ebiten.KeyEnter:        {tea.KeyEnter, []rune{'\n'}},
	ebiten.KeyTab:          {tea.KeyTab, []rune{'\t'}},
	ebiten.KeySpace:        {tea.KeySpace, []rune{' '}},
	ebiten.KeyBackspace:    {tea.KeyBackspace, []rune{}},
	ebiten.KeyDelete:       {tea.KeyDelete, []rune{}},
	ebiten.KeyHome:         {tea.KeyHome, []rune{}},
	ebiten.KeyEnd:          {tea.KeyEnd, []rune{}},
	ebiten.KeyPageUp:       {tea.KeyPgUp, []rune{}},
	ebiten.KeyArrowUp:      {tea.KeyUp, []rune{}},
	ebiten.KeyArrowDown:    {tea.KeyDown, []rune{}},
	ebiten.KeyArrowLeft:    {tea.KeyLeft, []rune{}},
	ebiten.KeyArrowRight:   {tea.KeyRight, []rune{}},
	ebiten.KeyEscape:       {tea.KeyEscape, []rune{}},
	ebiten.Key1:            {tea.KeyRunes, []rune{'1'}},
	ebiten.Key2:            {tea.KeyRunes, []rune{'2'}},
	ebiten.Key3:            {tea.KeyRunes, []rune{'3'}},
	ebiten.Key4:            {tea.KeyRunes, []rune{'4'}},
	ebiten.Key5:            {tea.KeyRunes, []rune{'5'}},
	ebiten.Key6:            {tea.KeyRunes, []rune{'6'}},
	ebiten.Key7:            {tea.KeyRunes, []rune{'7'}},
	ebiten.Key8:            {tea.KeyRunes, []rune{'8'}},
	ebiten.Key9:            {tea.KeyRunes, []rune{'9'}},
	ebiten.Key0:            {tea.KeyRunes, []rune{'0'}},
	ebiten.KeyA:            {tea.KeyRunes, []rune{'a'}},
	ebiten.KeyB:            {tea.KeyRunes, []rune{'b'}},
	ebiten.KeyC:            {tea.KeyRunes, []rune{'c'}},
	ebiten.KeyD:            {tea.KeyRunes, []rune{'d'}},
	ebiten.KeyE:            {tea.KeyRunes, []rune{'e'}},
	ebiten.KeyF:            {tea.KeyRunes, []rune{'f'}},
	ebiten.KeyG:            {tea.KeyRunes, []rune{'g'}},
	ebiten.KeyH:            {tea.KeyRunes, []rune{'h'}},
	ebiten.KeyI:            {tea.KeyRunes, []rune{'i'}},
	ebiten.KeyJ:            {tea.KeyRunes, []rune{'j'}},
	ebiten.KeyK:            {tea.KeyRunes, []rune{'k'}},
	ebiten.KeyL:            {tea.KeyRunes, []rune{'l'}},
	ebiten.KeyM:            {tea.KeyRunes, []rune{'m'}},
	ebiten.KeyN:            {tea.KeyRunes, []rune{'n'}},
	ebiten.KeyO:            {tea.KeyRunes, []rune{'o'}},
	ebiten.KeyP:            {tea.KeyRunes, []rune{'p'}},
	ebiten.KeyQ:            {tea.KeyRunes, []rune{'q'}},
	ebiten.KeyR:            {tea.KeyRunes, []rune{'r'}},
	ebiten.KeyS:            {tea.KeyRunes, []rune{'s'}},
	ebiten.KeyT:            {tea.KeyRunes, []rune{'t'}},
	ebiten.KeyU:            {tea.KeyRunes, []rune{'u'}},
	ebiten.KeyV:            {tea.KeyRunes, []rune{'v'}},
	ebiten.KeyW:            {tea.KeyRunes, []rune{'w'}},
	ebiten.KeyX:            {tea.KeyRunes, []rune{'x'}},
	ebiten.KeyY:            {tea.KeyRunes, []rune{'y'}},
	ebiten.KeyZ:            {tea.KeyRunes, []rune{'z'}},
	ebiten.KeyComma:        {tea.KeyRunes, []rune{','}},
	ebiten.KeyPeriod:       {tea.KeyRunes, []rune{'.'}},
	ebiten.KeySlash:        {tea.KeyRunes, []rune{'/'}},
	ebiten.KeyBackslash:    {tea.KeyRunes, []rune{'\\'}},
	ebiten.KeySemicolon:    {tea.KeyRunes, []rune{';'}},
	ebiten.KeyApostrophe:   {tea.KeyRunes, []rune{'\''}},
	ebiten.KeyGraveAccent:  {tea.KeyRunes, []rune{'`'}},
	ebiten.KeyEqual:        {tea.KeyRunes, []rune{'='}},
	ebiten.KeyMinus:        {tea.KeyRunes, []rune{'-'}},
	ebiten.KeyLeftBracket:  {tea.KeyRunes, []rune{'['}},
	ebiten.KeyRightBracket: {tea.KeyRunes, []rune{']'}},
	ebiten.KeyF1:           {tea.KeyF1, []rune{}},
	ebiten.KeyF2:           {tea.KeyF2, []rune{}},
	ebiten.KeyF3:           {tea.KeyF3, []rune{}},
	ebiten.KeyF4:           {tea.KeyF4, []rune{}},
	ebiten.KeyF5:           {tea.KeyF5, []rune{}},
	ebiten.KeyF6:           {tea.KeyF6, []rune{}},
	ebiten.KeyF7:           {tea.KeyF7, []rune{}},
	ebiten.KeyF8:           {tea.KeyF8, []rune{}},
	ebiten.KeyF9:           {tea.KeyF9, []rune{}},
	ebiten.KeyF10:          {tea.KeyF10, []rune{}},
	ebiten.KeyF11:          {tea.KeyF11, []rune{}},
	ebiten.KeyF12:          {tea.KeyF12, []rune{}},
	ebiten.KeyShift:        {tea.KeyShiftLeft, []rune{}},
}

var ebitenToTeaMouse = map[ebiten.MouseButton]tea.MouseEventType{
	ebiten.MouseButtonLeft:   tea.MouseLeft,
	ebiten.MouseButtonMiddle: tea.MouseMiddle,
	ebiten.MouseButtonRight:  tea.MouseRight,
}

// Options are used to configure the adapter.
type Options func(*Adapter)

// WithFilterMousePressed filters the MousePressed event and only emits MouseReleased events.
func WithFilterMousePressed(filter bool) Options {
	return func(b *Adapter) {
		b.filterMousePressed = filter
	}
}

// Adapter represents a bubbletea adapter for the crt package.
type Adapter struct {
	prog               *tea.Program
	filterMousePressed bool
}

// NewAdapter creates a new bubbletea adapter.
func NewAdapter(prog *tea.Program, options ...Options) *Adapter {
	b := &Adapter{prog: prog, filterMousePressed: true}

	for i := range options {
		options[i](b)
	}

	return b
}

func (b *Adapter) HandleMouseMotion(motion crt.MouseMotion) {
	b.prog.Send(tea.MouseMsg{
		X:    motion.X,
		Y:    motion.Y,
		Alt:  false,
		Ctrl: false,
		Type: tea.MouseMotion,
	})
}

func (b *Adapter) HandleMouseButton(button crt.MouseButton) {
	// Filter this event or two events will be sent for one click in the current bubbletea version.
	if b.filterMousePressed && button.JustPressed {
		return
	}

	b.prog.Send(tea.MouseMsg{
		X:    button.X,
		Y:    button.Y,
		Alt:  ebiten.IsKeyPressed(ebiten.KeyAlt),
		Ctrl: ebiten.IsKeyPressed(ebiten.KeyControl),
		Type: ebitenToTeaMouse[button.Button],
	})
}

func (b *Adapter) HandleMouseWheel(wheel crt.MouseWheel) {
	if wheel.DY > 0 {
		b.prog.Send(tea.MouseMsg{
			X:    wheel.X,
			Y:    wheel.Y,
			Alt:  ebiten.IsKeyPressed(ebiten.KeyAlt),
			Ctrl: ebiten.IsKeyPressed(ebiten.KeyControl),
			Type: tea.MouseWheelUp,
		})
	} else if wheel.DY < 0 {
		b.prog.Send(tea.MouseMsg{
			X:    wheel.X,
			Y:    wheel.Y,
			Alt:  ebiten.IsKeyPressed(ebiten.KeyAlt),
			Ctrl: ebiten.IsKeyPressed(ebiten.KeyControl),
			Type: tea.MouseWheelDown,
		})
	}
}

func (b *Adapter) HandleKeyPress() {
	var keys []ebiten.Key
	keys = inpututil.AppendJustReleasedKeys(keys)

	for _, k := range keys {

		switch k {
		case ebiten.KeyC:
			if ebiten.IsKeyPressed(ebiten.KeyControl) {
				b.prog.Send(tea.KeyMsg{
					Type:  tea.KeyCtrlC,
					Runes: []rune{},
					Alt:   false,
				})

				continue
			}
		}

		if val, ok := ebitenToTeaKeys[k]; ok {
			runes := make([]rune, len(val.rune))
			copy(runes, val.rune)

			if ebiten.IsKeyPressed(ebiten.KeyShift) {
				for i := range runes {
					runes[i] = unicode.ToUpper(runes[i])
				}
			}

			b.prog.Send(tea.KeyMsg{
				Type:  val.key,
				Runes: runes,
				Alt:   ebiten.IsKeyPressed(ebiten.KeyAlt),
			})
		}
	}
}

func (b *Adapter) HandleWindowSize(size crt.WindowSize) {
	b.prog.Send(tea.WindowSizeMsg{
		Width:  size.Width,
		Height: size.Height,
	})
}
