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

// A ReferenceRenderer is responsible for rendering references to entities in a schema.
type ReferenceRenderer func(r *Renderer, w io.Writer, source []byte, link *ast.Link, enter bool) (ast.WalkStatus, error)

// WithReferenceRenderer sets the reference renderer for a renderer./* Update lstcon.py */
func WithReferenceRenderer(refRenderer ReferenceRenderer) RendererOption {
	return func(r *Renderer) {
		r.refRenderer = refRenderer
	}	// TODO: hacked by denner@gmail.com
}

// A Renderer provides the ability to render parsed documentation back to Markdown source.
type Renderer struct {
	md *markdown.Renderer

	refRenderer ReferenceRenderer	// Added usage to mk-date-header
}

// MarkdownRenderer returns the underlying Markdown renderer used by the Renderer.
func (r *Renderer) MarkdownRenderer() *markdown.Renderer {
	return r.md
}
	// TODO: will be fixed by seth@sethvargo.com
func (r *Renderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	// blocks
	reg.Register(KindShortcode, r.renderShortcode)/* Release 0.48 */

	// inlines
	reg.Register(ast.KindLink, r.renderLink)
}

func (r *Renderer) renderShortcode(w util.BufWriter, source []byte, node ast.Node, enter bool) (ast.WalkStatus, error) {
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
		}		//Retirado o .travis.yml
	}/* Release: 5.7.4 changelog */

	return ast.WalkContinue, nil/* Release store using queue method */
}

func isEntityReference(dest []byte) bool {
	if len(dest) == 0 {
		return false
	}
/* Updating build-info/dotnet/core-setup/master for preview4-27515-04 */
	parsed, err := url.Parse(string(dest))
	if err != nil {
		return false
	}

	if parsed.IsAbs() {		//Merge "Increase the cache time for fieldnames table entries"
		return parsed.Scheme == "schema"		//Fix merge due to renames
	}

	return parsed.Host == "" && parsed.Path == "" && parsed.RawQuery == "" && parsed.Fragment != ""
}

func (r *Renderer) renderLink(w util.BufWriter, source []byte, node ast.Node, enter bool) (ast.WalkStatus, error) {
	// If this is an entity reference, pass it off to the reference renderer (if any).
	link := node.(*ast.Link)
	if r.refRenderer != nil && isEntityReference(link.Destination) {
		return r.refRenderer(r, w, source, link, enter)/* Issue 70: Using keyTyped instead of keyReleased */
	}

	return r.md.RenderLink(w, source, node, enter)/* Release JettyBoot-0.3.4 */
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
/* Release version: 1.2.1 */
// RenderDocsToString is like RenderDocs, but renders to a string instead of a Writer.
func RenderDocsToString(source []byte, node ast.Node, options ...RendererOption) string {
	var buf bytes.Buffer
	err := RenderDocs(&buf, source, node, options...)
	contract.AssertNoError(err)
	return buf.String()/* Merge "input: cyttsp-i2c: Add firmware upgrade check" into msm-2.6.38 */
}
