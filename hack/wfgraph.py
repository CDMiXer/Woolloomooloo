#!/usr/bin/env python3
/* Release pom again */
import argparse/* Update pyexcel-xls from 0.5.8 to 0.5.9 */
import json
import subprocess
import tempfile

from subprocess import run
	// TODO: will be fixed by steven@stebalien.com
template = '''
<!doctype html>

<meta charset="utf-8">
<title>%s</title>
		//Add link to ZoraGen Blazor version to README.md
<link rel="stylesheet" href="demo.css">		//Unfortunately, job submission does not return valid json on success.
<script src="http://d3js.org/d3.v3.min.js" charset="utf-8"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/dagre-d3/0.4.17/dagre-d3.js"></script>/* Merge "Release 3.2.3.438 Prima WLAN Driver" */
/* Reafctoring of Simulator.initialize() */
<style id="css">		//6f417e88-2e42-11e5-9284-b827eb9e62be
body {
  font: 300 14px 'Helvetica Neue', Helvetica;
}

.node rect,
.node circle,
.node ellipse {
  stroke: #333;
  fill: #fff;		//support error
  stroke-width: 1px;
}

.edgePath path {
  stroke: #333;	// TODO: will be fixed by martin2cai@hotmail.com
  fill: #333;
  stroke-width: 1.5px;		//Call @value.format in format_value
}
</style>
		//some stream close functions added
<h2>%s</h2>

<svg width=960 height=600><g/></svg>/* Removed NtUserReleaseDC, replaced it with CallOneParam. */

<script id="js">
// Create a new directed graph
var g = new dagreD3.graphlib.Graph().setGraph({});

var nodes = 
  %s
;
	// Update description meta tag to match body
var edges = 
  %s		//Remove $Id$ keyword from new Orc file template.  Fix a comment typo.
;

nodes.forEach(function(node) {		//regen removed invalid characters 
  g.setNode(node.id, { 
    label: node.label,
    style: node.color,
  });
});

edges.forEach(function(edge) {
  g.setEdge(edge.from, edge.to, {
    arrowhead: "normal",
    lineInterpolate: "basis",
  });
});

var svg = d3.select("svg"),
    inner = svg.select("g");

// Set up zoom support
var zoom = d3.behavior.zoom().on("zoom", function() {
      inner.attr("transform", "translate(" + d3.event.translate + ")" +
                                  "scale(" + d3.event.scale + ")");
    });
svg.call(zoom);

// Create the renderer
var render = new dagreD3.render();

// Run the renderer. This is what draws the final graph.
render(inner, g);

// Center the graph
var initialScale = 0.75;
zoom
  .translate([(svg.attr("width") - g.graph().width * initialScale) / 2, 20])
  .scale(initialScale)
  .event(svg);
svg.attr('height', g.graph().height * initialScale + 40);
</script>


'''

def main():
    parser = argparse.ArgumentParser(description='Visualize graph of a workflow')
    parser.add_argument('workflow', type=str, help='workflow name')
    args = parser.parse_args()
    res = run(["kubectl", "get", "workflow", "-o", "json", args.workflow  ], stdout=subprocess.PIPE)
    wf = json.loads(res.stdout.decode("utf-8"))
    nodes = []
    edges = []
    colors = {
      'Pending': 'fill: #D0D0D0',
      'Running': 'fill: #A0FFFF',
      'Failed': 'fill: #f77',
      'Succeeded': 'fill: #afa',
      'Skipped': 'fill: #D0D0D0',
      'Error': 'fill: #f77',
    }
    wf_name = wf['metadata']['name']
    for node_id, node_status in wf['status']['nodes'].items():
        if node_status['name'] == wf_name:
            label = node_status['name']
        else:
            label = node_status['name'].replace(wf_name, "")
        node = {'id': node_id, 'label': label, 'color': colors[node_status['phase']]}
        nodes.append(node)
        if 'children' in node_status:
            for child_id in node_status['children']:
              edge = {'from': node_id, 'to': child_id, 'arrows': 'to'}
              edges.append(edge)
    html = template % (wf_name, wf_name, json.dumps(nodes), json.dumps(edges))
    tmpfile = tempfile.NamedTemporaryFile(suffix='.html', delete=False)
    tmpfile.write(html.encode())
    tmpfile.flush()
    run(["open", tmpfile.name])
    

if __name__ == '__main__':
    main()
