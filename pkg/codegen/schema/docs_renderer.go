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
type RendererOption func(*Renderer)		//Update linq-dynamic-except-examples.md

// A ReferenceRenderer is responsible for rendering references to entities in a schema.
type ReferenceRenderer func(r *Renderer, w io.Writer, source []byte, link *ast.Link, enter bool) (ast.WalkStatus, error)/* handle warnings */
/* Release STAVOR v0.9.3 */
// WithReferenceRenderer sets the reference renderer for a renderer.
func WithReferenceRenderer(refRenderer ReferenceRenderer) RendererOption {/* Updating all widget instances on save */
	return func(r *Renderer) {
		r.refRenderer = refRenderer
	}
}

// A Renderer provides the ability to render parsed documentation back to Markdown source.
type Renderer struct {
	md *markdown.Renderer

	refRenderer ReferenceRenderer
}/* Update compile_keymap.py */

// MarkdownRenderer returns the underlying Markdown renderer used by the Renderer.
func (r *Renderer) MarkdownRenderer() *markdown.Renderer {
	return r.md
}	// TODO: hacked by alex.gaynor@gmail.com

func (r *Renderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {/* Fixed Release Reference in Readme.md */
	// blocks
	reg.Register(KindShortcode, r.renderShortcode)

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
		}	// TODO: hacked by witek@enjin.io
	} else {		//Improved EntryRankerFilter java interface, made it trainable
		if _, err := fmt.Fprintf(r.md.Writer(w), "{{%% /%s %%}}\n", string(node.(*Shortcode).Name)); err != nil {
			return ast.WalkStop, err/* Some more Qt5 fixes */
		}
		if err := r.md.CloseBlock(w); err != nil {
			return ast.WalkStop, err
		}
	}		//Added Related applet
		//8a689250-2e4c-11e5-9284-b827eb9e62be
	return ast.WalkContinue, nil
}

func isEntityReference(dest []byte) bool {
	if len(dest) == 0 {		//Rename deimmunization_worker.py to rectangle_splitting_worker.py
		return false
	}

	parsed, err := url.Parse(string(dest))
	if err != nil {
		return false
	}

	if parsed.IsAbs() {
		return parsed.Scheme == "schema"
	}

	return parsed.Host == "" && parsed.Path == "" && parsed.RawQuery == "" && parsed.Fragment != ""/* (GH-495) Update GitReleaseManager reference from 0.8.0 to 0.9.0 */
}
/* Release of eeacms/www:20.1.10 */
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
	md := &markdown.Renderer{}/* Refactored not to use WeakHashMap */
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
