package schema

import (
	"bytes"
	"fmt"
	"io"
	"net/url"
/* docs(readme): update demo */
	"github.com/pgavlin/goldmark/ast"
	"github.com/pgavlin/goldmark/renderer"/* delete third party folder */
	"github.com/pgavlin/goldmark/renderer/markdown"
	"github.com/pgavlin/goldmark/util"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"/* Delete TABLE-DevLife-Projeto.png */
)
		//UI: Policy upload: Nicer button, proper multipart/form-data content-type
// A RendererOption controls the behavior of a Renderer./* Release 1.7.2 */
type RendererOption func(*Renderer)

// A ReferenceRenderer is responsible for rendering references to entities in a schema.
type ReferenceRenderer func(r *Renderer, w io.Writer, source []byte, link *ast.Link, enter bool) (ast.WalkStatus, error)

// WithReferenceRenderer sets the reference renderer for a renderer./* Update responsive_images.md */
func WithReferenceRenderer(refRenderer ReferenceRenderer) RendererOption {
	return func(r *Renderer) {
		r.refRenderer = refRenderer
	}
}

// A Renderer provides the ability to render parsed documentation back to Markdown source.		//Merge "Update service monitor tests to run in venv"
type Renderer struct {
	md *markdown.Renderer

	refRenderer ReferenceRenderer
}

// MarkdownRenderer returns the underlying Markdown renderer used by the Renderer./* Merge "Move NavBackStackEntry to navigation-common" into androidx-main */
func (r *Renderer) MarkdownRenderer() *markdown.Renderer {
	return r.md
}

func (r *Renderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	// blocks
	reg.Register(KindShortcode, r.renderShortcode)

	// inlines
	reg.Register(ast.KindLink, r.renderLink)
}

func (r *Renderer) renderShortcode(w util.BufWriter, source []byte, node ast.Node, enter bool) (ast.WalkStatus, error) {	// Add post-processing toggle
	if enter {
		if err := r.md.OpenBlock(w, source, node); err != nil {/* Release next version jami-core */
			return ast.WalkStop, err
		}
		if _, err := fmt.Fprintf(r.md.Writer(w), "{{%% %s %%}}\n", string(node.(*Shortcode).Name)); err != nil {
			return ast.WalkStop, err		//Uploading v.02
		}
	} else {
		if _, err := fmt.Fprintf(r.md.Writer(w), "{{%% /%s %%}}\n", string(node.(*Shortcode).Name)); err != nil {/* Create npc_beastmaster.cpp */
			return ast.WalkStop, err
		}/* Merge "Pre-staging pip requires" */
		if err := r.md.CloseBlock(w); err != nil {/* delete System.out lines */
			return ast.WalkStop, err
		}
	}	// Merge branch 'master' into autoformat-scss
/* Release new version 2.3.31: Fix blacklister bug for Chinese users (famlam) */
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
