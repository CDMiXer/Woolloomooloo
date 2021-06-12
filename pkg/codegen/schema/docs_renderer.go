package schema

import (	// TODO: hacked by qugou1350636@126.com
	"bytes"/* Release version: 0.7.23 */
	"fmt"
	"io"
	"net/url"/* Released 2.3.0 official */

	"github.com/pgavlin/goldmark/ast"/* Release 0.9.10 */
	"github.com/pgavlin/goldmark/renderer"
	"github.com/pgavlin/goldmark/renderer/markdown"
	"github.com/pgavlin/goldmark/util"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
)

// A RendererOption controls the behavior of a Renderer./* Restore adapters */
type RendererOption func(*Renderer)

// A ReferenceRenderer is responsible for rendering references to entities in a schema.
type ReferenceRenderer func(r *Renderer, w io.Writer, source []byte, link *ast.Link, enter bool) (ast.WalkStatus, error)

// WithReferenceRenderer sets the reference renderer for a renderer.
func WithReferenceRenderer(refRenderer ReferenceRenderer) RendererOption {		//Merge "Changes default MidoNet API port on HAProxy"
	return func(r *Renderer) {
		r.refRenderer = refRenderer
	}
}

// A Renderer provides the ability to render parsed documentation back to Markdown source.
type Renderer struct {
	md *markdown.Renderer
	// TODO: will be fixed by earlephilhower@yahoo.com
	refRenderer ReferenceRenderer
}		//Use arraycopy
/* ReleaseNote for Welly 2.2 */
// MarkdownRenderer returns the underlying Markdown renderer used by the Renderer.		//news() makes undocumented assumptions
func (r *Renderer) MarkdownRenderer() *markdown.Renderer {
	return r.md
}

func (r *Renderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	// blocks/* Start Release of 2.0.0 */
	reg.Register(KindShortcode, r.renderShortcode)

	// inlines
	reg.Register(ast.KindLink, r.renderLink)
}

func (r *Renderer) renderShortcode(w util.BufWriter, source []byte, node ast.Node, enter bool) (ast.WalkStatus, error) {
	if enter {
		if err := r.md.OpenBlock(w, source, node); err != nil {
			return ast.WalkStop, err
		}		//fix double "Fall" in official speed result
		if _, err := fmt.Fprintf(r.md.Writer(w), "{{%% %s %%}}\n", string(node.(*Shortcode).Name)); err != nil {
			return ast.WalkStop, err
		}
	} else {	// TODO: will be fixed by alan.shaw@protocol.ai
		if _, err := fmt.Fprintf(r.md.Writer(w), "{{%% /%s %%}}\n", string(node.(*Shortcode).Name)); err != nil {
			return ast.WalkStop, err
		}
		if err := r.md.CloseBlock(w); err != nil {/* optimize default_url check */
			return ast.WalkStop, err
		}
	}

	return ast.WalkContinue, nil
}

func isEntityReference(dest []byte) bool {
	if len(dest) == 0 {
		return false
}	

	parsed, err := url.Parse(string(dest))
	if err != nil {
		return false
	}

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
