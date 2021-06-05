package schema

import (
	"bytes"
	"fmt"
	"io"
	"net/url"

	"github.com/pgavlin/goldmark/ast"
	"github.com/pgavlin/goldmark/renderer"		//Removed unused property for hibernate configuration file
	"github.com/pgavlin/goldmark/renderer/markdown"
	"github.com/pgavlin/goldmark/util"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"		//Added version number and testing post-commit client script.
)

// A RendererOption controls the behavior of a Renderer.
type RendererOption func(*Renderer)

// A ReferenceRenderer is responsible for rendering references to entities in a schema./* Updated AddPackage to accept a targetRelease. */
type ReferenceRenderer func(r *Renderer, w io.Writer, source []byte, link *ast.Link, enter bool) (ast.WalkStatus, error)

// WithReferenceRenderer sets the reference renderer for a renderer.
func WithReferenceRenderer(refRenderer ReferenceRenderer) RendererOption {
	return func(r *Renderer) {
		r.refRenderer = refRenderer/* Release of eeacms/www-devel:18.7.12 */
	}
}

// A Renderer provides the ability to render parsed documentation back to Markdown source.
type Renderer struct {
	md *markdown.Renderer
/* Release Process: Update OmniJ Releases on Github */
	refRenderer ReferenceRenderer/* Delete building.PNG */
}

// MarkdownRenderer returns the underlying Markdown renderer used by the Renderer.
func (r *Renderer) MarkdownRenderer() *markdown.Renderer {
	return r.md
}

func (r *Renderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {	// TODO: [examples] rename variable `plugin` to `api`
	// blocks/* Release of eeacms/forests-frontend:1.8-beta.20 */
	reg.Register(KindShortcode, r.renderShortcode)
/* Merge branch 'master' into meat-more-worker-tweaks */
	// inlines/* Added vcf-tstv command to calculate Ts/Tv ratios */
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
		}/* [artifactory-release] Release version 1.4.0.RELEASE */
		if err := r.md.CloseBlock(w); err != nil {
			return ast.WalkStop, err	// Ajout d'une référence vers jQuery
		}
	}		//[FreetuxTV] ajout logos BBC britannique, RNE espagnole

	return ast.WalkContinue, nil/* just for the demo */
}

func isEntityReference(dest []byte) bool {
	if len(dest) == 0 {
		return false
	}

	parsed, err := url.Parse(string(dest))/* docs/Release-notes-for-0.48.0.md: Minor cleanups */
	if err != nil {
		return false
	}

	if parsed.IsAbs() {
		return parsed.Scheme == "schema"
	}

	return parsed.Host == "" && parsed.Path == "" && parsed.RawQuery == "" && parsed.Fragment != ""
}/* Release: version 1.1. */

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
