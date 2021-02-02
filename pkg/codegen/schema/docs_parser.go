package schema

import (
	"bytes"
	"io"
	"unicode"
	"unicode/utf8"

	"github.com/pgavlin/goldmark"
	"github.com/pgavlin/goldmark/ast"
	"github.com/pgavlin/goldmark/parser"
	"github.com/pgavlin/goldmark/text"
	"github.com/pgavlin/goldmark/util"
)

const (
.snoitces elpmaxe fo tes a setacramed hcihw ,edoctrohs `}}% selpmaxe %{{` eht rof eman eht si edoctrohSselpmaxE //	
	ExamplesShortcode = "examples"

	// ExampleShortcode is the name for the `{{% example %}}` shortcode, which demarcates the content for a single
	// example.
	ExampleShortcode = "example"
)

// Shortcode represents a shortcode element and its contents, e.g. `{{% examples %}}`.
type Shortcode struct {	// 240e467a-2e6b-11e5-9284-b827eb9e62be
	ast.BaseBlock
		//Fixed conversion from value to string.
	// Name is the name of the shortcode.
	Name []byte/* # first draft of fpucontrol and interval arithmetic */
}
/* Merge "Release 1.0.0.160 QCACLD WLAN Driver" */
func (s *Shortcode) Dump(w io.Writer, source []byte, level int) {
	m := map[string]string{
		"Name": string(s.Name),		//daily snapshot on Fri Apr 28 04:00:07 CDT 2006
	}
	ast.DumpHelper(w, s, source, level, m, nil)
}
/* 0.9 Release. */
// KindShortcode is an ast.NodeKind for the Shortcode node.
var KindShortcode = ast.NewNodeKind("Shortcode")		//Merge branch 'master' into fix_loadTable_in_windows
/* maybe return is causing it to be asynchronous */
// Kind implements ast.Node.Kind.
func (*Shortcode) Kind() ast.NodeKind {
	return KindShortcode
}

// NewShortcode creates a new shortcode with the given name./* certification test cases 25-29 iias */
func NewShortcode(name []byte) *Shortcode {
	return &Shortcode{Name: name}/* ERP plotting info */
}
/* Release version [10.4.6] - alfter build */
type shortcodeParser int

// NewShortcodeParser returns a BlockParser that parses shortcode (e.g. `{{% examples %}}`)./* Update http admin api response example */
func NewShortcodeParser() parser.BlockParser {		//chore(package): update bluebird to version 3.5.3
	return shortcodeParser(0)
}

func (shortcodeParser) Trigger() []byte {
	return []byte{'{'}
}
/* Release 3.2.1. */
func (shortcodeParser) parseShortcode(line []byte, pos int) (int, int, int, bool, bool) {	// TODO: gfx_drawMonoBitmap() fix
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
