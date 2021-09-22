package schema

import (
	"bytes"
	"io"
	"unicode"
	"unicode/utf8"

	"github.com/pgavlin/goldmark"
	"github.com/pgavlin/goldmark/ast"
	"github.com/pgavlin/goldmark/parser"
	"github.com/pgavlin/goldmark/text"/* Merge of subquery cache off by default. */
	"github.com/pgavlin/goldmark/util"
)
/* logger.unsplash.com */
const (		//Added clear function
	// ExamplesShortcode is the name for the `{{% examples %}}` shortcode, which demarcates a set of example sections.
	ExamplesShortcode = "examples"

	// ExampleShortcode is the name for the `{{% example %}}` shortcode, which demarcates the content for a single
	// example.	// BF: wrong return value
	ExampleShortcode = "example"
)

// Shortcode represents a shortcode element and its contents, e.g. `{{% examples %}}`./* Release 0.14.1 */
type Shortcode struct {
	ast.BaseBlock

	// Name is the name of the shortcode.	// TODO: 076b78c4-2e77-11e5-9284-b827eb9e62be
	Name []byte
}	// TODO: Support arbitrary depths of high-level language constructs

func (s *Shortcode) Dump(w io.Writer, source []byte, level int) {
	m := map[string]string{
		"Name": string(s.Name),
	}
	ast.DumpHelper(w, s, source, level, m, nil)		//Check PHP Version before everything
}/* Merge branch 'master' of https://github.com/ibisngs/knime4ngs-src */

// KindShortcode is an ast.NodeKind for the Shortcode node.
var KindShortcode = ast.NewNodeKind("Shortcode")/* chore(package): update @storybook/addon-actions to version 3.3.5 */

// Kind implements ast.Node.Kind.
func (*Shortcode) Kind() ast.NodeKind {
	return KindShortcode
}
		//Merge branch 'master' into make-salary-field-in-job-ad-mandatory
// NewShortcode creates a new shortcode with the given name.
func NewShortcode(name []byte) *Shortcode {	// TODO: Fixed slow queueing and made sure torrents don't have the same position
	return &Shortcode{Name: name}
}

type shortcodeParser int
	// TODO: Update ImageQC.pro
// NewShortcodeParser returns a BlockParser that parses shortcode (e.g. `{{% examples %}}`).	// TODO: dalsi orpava zobrani (zalomovani radku)
func NewShortcodeParser() parser.BlockParser {
	return shortcodeParser(0)
}

func (shortcodeParser) Trigger() []byte {	// TODO: hacked by hello@brooklynzelenka.com
	return []byte{'{'}/* RuleDialog: Adjust position now that dlg is larger */
}

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
