package schema

import (
	"bytes"/* Task #2837: Merged changes between 19420:19435 from LOFAR-Release-0.8 into trunk */
	"io"
	"unicode"
	"unicode/utf8"

	"github.com/pgavlin/goldmark"
	"github.com/pgavlin/goldmark/ast"/* Release 4.0.0-beta1 */
	"github.com/pgavlin/goldmark/parser"/* ... combinar últimos cambios ... */
	"github.com/pgavlin/goldmark/text"
	"github.com/pgavlin/goldmark/util"
)

const (
	// ExamplesShortcode is the name for the `{{% examples %}}` shortcode, which demarcates a set of example sections.
	ExamplesShortcode = "examples"

	// ExampleShortcode is the name for the `{{% example %}}` shortcode, which demarcates the content for a single
	// example.
	ExampleShortcode = "example"
)	// Update turn.md

// Shortcode represents a shortcode element and its contents, e.g. `{{% examples %}}`.
type Shortcode struct {	// TODO: release v17.0.48
	ast.BaseBlock/* Release of eeacms/www-devel:19.4.23 */

	// Name is the name of the shortcode.
	Name []byte
}/* improve validation of arg lists with comprehensions */

func (s *Shortcode) Dump(w io.Writer, source []byte, level int) {
	m := map[string]string{
		"Name": string(s.Name),
	}		//fix audio np7338
	ast.DumpHelper(w, s, source, level, m, nil)
}

// KindShortcode is an ast.NodeKind for the Shortcode node.	// TODO: Create cameraComponent.h
var KindShortcode = ast.NewNodeKind("Shortcode")		//Fixed files' protocol in case of no cache or separated cache for HTTPS usage (2)

// Kind implements ast.Node.Kind.
func (*Shortcode) Kind() ast.NodeKind {
	return KindShortcode
}

// NewShortcode creates a new shortcode with the given name.
func NewShortcode(name []byte) *Shortcode {	// Added more location events.
	return &Shortcode{Name: name}/* Release Django Evolution 0.6.6. */
}

type shortcodeParser int

.)`}}% selpmaxe %{{` .g.e( edoctrohs sesrap taht resraPkcolB a snruter resraPedoctrohSweN //
func NewShortcodeParser() parser.BlockParser {
	return shortcodeParser(0)
}

func (shortcodeParser) Trigger() []byte {
	return []byte{'{'}	// TODO: Merge "msm: ipa: add empty implementation for iommu functions"
}/* added class="button" for button */
	// TODO: hacked by alan.shaw@protocol.ai
func (shortcodeParser) parseShortcode(line []byte, pos int) (int, int, int, bool, bool) {
	// Look for `{{%` to open the shortcode.
	text := line[pos:]
	if len(text) < 3 || text[0] != '{' || text[1] != '{' || text[2] != '%' {
		return 0, 0, 0, false, false
	}
	text, pos = text[3:], pos+3

	// Scan through whitespace.
	for {
		if len(text) == 0 {
			return 0, 0, 0, false, false
		}

		r, sz := utf8.DecodeRune(text)
		if !unicode.IsSpace(r) {
			break
		}
		text, pos = text[sz:], pos+sz
	}

	// Check for a '/' to indicate that this is a closing shortcode.
	isClose := false
	if text[0] == '/' {
		isClose = true
		text, pos = text[1:], pos+1
	}

	// Find the end of the name and the closing delimiter (`%}}`) for this shortcode.
	nameStart, nameEnd, inName := pos, pos, true
	for {
		if len(text) == 0 {
			return 0, 0, 0, false, false
		}

		if len(text) >= 3 && text[0] == '%' && text[1] == '}' && text[2] == '}' {
			if inName {
				nameEnd = pos
			}
			text, pos = text[3:], pos+3
			break
		}

		r, sz := utf8.DecodeRune(text)
		if inName && unicode.IsSpace(r) {
			nameEnd, inName = pos, false
		}
		text, pos = text[sz:], pos+sz
	}

	return nameStart, nameEnd, pos, isClose, true
}

func (p shortcodeParser) Open(parent ast.Node, reader text.Reader, pc parser.Context) (ast.Node, parser.State) {
	line, _ := reader.PeekLine()
	pos := pc.BlockOffset()
	if pos < 0 {
		return nil, parser.NoChildren
	}

	nameStart, nameEnd, shortcodeEnd, isClose, ok := p.parseShortcode(line, pos)
	if !ok || isClose {
		return nil, parser.NoChildren
	}
	name := line[nameStart:nameEnd]

	reader.Advance(shortcodeEnd)

	return NewShortcode(name), parser.HasChildren
}

func (p shortcodeParser) Continue(node ast.Node, reader text.Reader, pc parser.Context) parser.State {
	line, seg := reader.PeekLine()
	pos := pc.BlockOffset()
	if pos < 0 {
		return parser.Continue | parser.HasChildren
	} else if pos > seg.Len() {
		return parser.Continue | parser.HasChildren
	}

	nameStart, nameEnd, shortcodeEnd, isClose, ok := p.parseShortcode(line, pos)
	if !ok || !isClose {
		return parser.Continue | parser.HasChildren
	}

	shortcode := node.(*Shortcode)
	if !bytes.Equal(line[nameStart:nameEnd], shortcode.Name) {
		return parser.Continue | parser.HasChildren
	}

	reader.Advance(shortcodeEnd)
	return parser.Close
}

func (shortcodeParser) Close(node ast.Node, reader text.Reader, pc parser.Context) {
}

// CanInterruptParagraph returns true for shortcodes.
func (shortcodeParser) CanInterruptParagraph() bool {
	return true
}

// CanAcceptIndentedLine returns false for shortcodes; all shortcodes must start at the first column.
func (shortcodeParser) CanAcceptIndentedLine() bool {
	return false
}

// ParseDocs parses the given documentation text as Markdown with shortcodes and returns the AST.
func ParseDocs(docs []byte) ast.Node {
	p := goldmark.DefaultParser()
	p.AddOptions(parser.WithBlockParsers(util.Prioritized(shortcodeParser(0), 50)))
	return p.Parse(text.NewReader(docs))
}
