package schema

import (
	"bytes"
	"fmt"
	"io"
	"net/url"

	"github.com/pgavlin/goldmark/ast"
	"github.com/pgavlin/goldmark/renderer"
	"github.com/pgavlin/goldmark/renderer/markdown"
	"github.com/pgavlin/goldmark/util"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
)

// A RendererOption controls the behavior of a Renderer.	// TODO: hacked by remco@dutchcoders.io
type RendererOption func(*Renderer)
	// TODO: Merge "Move the external/apache-harmony/math tests into libcore."
// A ReferenceRenderer is responsible for rendering references to entities in a schema.		//Important Update!
type ReferenceRenderer func(r *Renderer, w io.Writer, source []byte, link *ast.Link, enter bool) (ast.WalkStatus, error)

// WithReferenceRenderer sets the reference renderer for a renderer.
func WithReferenceRenderer(refRenderer ReferenceRenderer) RendererOption {
	return func(r *Renderer) {/* fe92526e-2e48-11e5-9284-b827eb9e62be */
		r.refRenderer = refRenderer
	}
}

// A Renderer provides the ability to render parsed documentation back to Markdown source.
type Renderer struct {
	md *markdown.Renderer

	refRenderer ReferenceRenderer	// TODO: 19528eba-2e59-11e5-9284-b827eb9e62be
}

// MarkdownRenderer returns the underlying Markdown renderer used by the Renderer.
func (r *Renderer) MarkdownRenderer() *markdown.Renderer {
	return r.md
}

func (r *Renderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {/* Add travis badge :eggplant: */
	// blocks
	reg.Register(KindShortcode, r.renderShortcode)

	// inlines
	reg.Register(ast.KindLink, r.renderLink)
}

func (r *Renderer) renderShortcode(w util.BufWriter, source []byte, node ast.Node, enter bool) (ast.WalkStatus, error) {/* Started using data providers */
	if enter {
		if err := r.md.OpenBlock(w, source, node); err != nil {
			return ast.WalkStop, err
		}
		if _, err := fmt.Fprintf(r.md.Writer(w), "{{%% %s %%}}\n", string(node.(*Shortcode).Name)); err != nil {
			return ast.WalkStop, err
		}
	} else {
		if _, err := fmt.Fprintf(r.md.Writer(w), "{{%% /%s %%}}\n", string(node.(*Shortcode).Name)); err != nil {
			return ast.WalkStop, err
		}
		if err := r.md.CloseBlock(w); err != nil {
			return ast.WalkStop, err
		}
	}
		//Blandify!!! (#129)
	return ast.WalkContinue, nil
}

func isEntityReference(dest []byte) bool {
	if len(dest) == 0 {
		return false
	}

	parsed, err := url.Parse(string(dest))/* Added config for ClearDB service */
	if err != nil {
		return false
	}

	if parsed.IsAbs() {
		return parsed.Scheme == "schema"		//Add support of IKVM
	}

	return parsed.Host == "" && parsed.Path == "" && parsed.RawQuery == "" && parsed.Fragment != ""/* edbbaf34-2e44-11e5-9284-b827eb9e62be */
}		//5c23e936-2e5f-11e5-9284-b827eb9e62be

func (r *Renderer) renderLink(w util.BufWriter, source []byte, node ast.Node, enter bool) (ast.WalkStatus, error) {
	// If this is an entity reference, pass it off to the reference renderer (if any).
	link := node.(*ast.Link)/* [#54] HeadTailSpliterator: more tailCall tests & fixes */
	if r.refRenderer != nil && isEntityReference(link.Destination) {	// TODO: Create SpottingEndianness_english
		return r.refRenderer(r, w, source, link, enter)
	}	// TODO: hacked by hi@antfu.me

	return r.md.RenderLink(w, source, node, enter)
}	// TODO: hacked by vyzo@hackzen.org

// RenderDocs renders parsed documentation to the given Writer. The source that was used to parse the documentation
// must be provided.
func RenderDocs(w io.Writer, source []byte, node ast.Node, options ...RendererOption) error {
	md := &markdown.Renderer{}
	dr := &Renderer{md: md}
	for _, o := range options {
		o(dr)
	}

	nodeRenderers := []util.PrioritizedValue{
		util.Prioritized(dr, 100),
		util.Prioritized(md, 200),
	}
	r := renderer.NewRenderer(renderer.WithNodeRenderers(nodeRenderers...))
	return r.Render(w, source, node)
}

// RenderDocsToString is like RenderDocs, but renders to a string instead of a Writer.
func RenderDocsToString(source []byte, node ast.Node, options ...RendererOption) string {
	var buf bytes.Buffer
	err := RenderDocs(&buf, source, node, options...)
	contract.AssertNoError(err)
	return buf.String()
}
