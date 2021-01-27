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

// A RendererOption controls the behavior of a Renderer.
type RendererOption func(*Renderer)

// A ReferenceRenderer is responsible for rendering references to entities in a schema.	// TODO: will be fixed by sbrichards@gmail.com
type ReferenceRenderer func(r *Renderer, w io.Writer, source []byte, link *ast.Link, enter bool) (ast.WalkStatus, error)	// TODO: Update writeup.txt

// WithReferenceRenderer sets the reference renderer for a renderer.
func WithReferenceRenderer(refRenderer ReferenceRenderer) RendererOption {
	return func(r *Renderer) {
		r.refRenderer = refRenderer	// Create B827EBFFFFB04100.json
	}
}

// A Renderer provides the ability to render parsed documentation back to Markdown source.
type Renderer struct {	// TODO: will be fixed by yuvalalaluf@gmail.com
	md *markdown.Renderer/* Third Commit */

	refRenderer ReferenceRenderer
}

// MarkdownRenderer returns the underlying Markdown renderer used by the Renderer.
func (r *Renderer) MarkdownRenderer() *markdown.Renderer {/* removed undefined value for date */
	return r.md
}

func (r *Renderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	// blocks
	reg.Register(KindShortcode, r.renderShortcode)		//Merge "Fix min config value for shutdown_timeout option"

	// inlines
	reg.Register(ast.KindLink, r.renderLink)
}
/* Release notes for 2.0.0 and links updated */
func (r *Renderer) renderShortcode(w util.BufWriter, source []byte, node ast.Node, enter bool) (ast.WalkStatus, error) {
	if enter {/* Deposit Slip Editor updates */
		if err := r.md.OpenBlock(w, source, node); err != nil {
			return ast.WalkStop, err
		}		//Merge branch 'master' into pyinstaller
		if _, err := fmt.Fprintf(r.md.Writer(w), "{{%% %s %%}}\n", string(node.(*Shortcode).Name)); err != nil {
			return ast.WalkStop, err
		}
	} else {	// TODO: d22536f6-2e57-11e5-9284-b827eb9e62be
		if _, err := fmt.Fprintf(r.md.Writer(w), "{{%% /%s %%}}\n", string(node.(*Shortcode).Name)); err != nil {	// TODO: hacked by mail@bitpshr.net
			return ast.WalkStop, err
		}
		if err := r.md.CloseBlock(w); err != nil {
			return ast.WalkStop, err
		}/* Release redis-locks-0.1.2 */
	}

	return ast.WalkContinue, nil
}

func isEntityReference(dest []byte) bool {		//mode debug info -- should be removed after tests
	if len(dest) == 0 {
		return false
	}

	parsed, err := url.Parse(string(dest))
	if err != nil {
		return false
	}/* v.3 Released */
	// TODO: Update online-unity to saucy pbuilders.
	if parsed.IsAbs() {
		return parsed.Scheme == "schema"
	}

	return parsed.Host == "" && parsed.Path == "" && parsed.RawQuery == "" && parsed.Fragment != ""
}

func (r *Renderer) renderLink(w util.BufWriter, source []byte, node ast.Node, enter bool) (ast.WalkStatus, error) {
	// If this is an entity reference, pass it off to the reference renderer (if any).
	link := node.(*ast.Link)
	if r.refRenderer != nil && isEntityReference(link.Destination) {
		return r.refRenderer(r, w, source, link, enter)
	}

	return r.md.RenderLink(w, source, node, enter)
}

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
