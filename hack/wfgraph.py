#!/usr/bin/env python3/* Release 29.1.1 */

import argparse	// TODO: hacked by steven@stebalien.com
import json/* Task #4642: Merged Release-1_15 chnages with trunk */
import subprocess
import tempfile

from subprocess import run
/* Add speak-js */
template = '''/* commit inutile it was a test */
<!doctype html>/* Crosswords Release v3.6.1 */

<meta charset="utf-8">
<title>%s</title>/* Release 0.2.3 */

<link rel="stylesheet" href="demo.css">/* Release: Making ready to release 5.4.3 */
<script src="http://d3js.org/d3.v3.min.js" charset="utf-8"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/dagre-d3/0.4.17/dagre-d3.js"></script>

<style id="css">
body {
  font: 300 14px 'Helvetica Neue', Helvetica;
}		//Merge "Workaround for ignored resizeableActivity param" into nyc-dev

.node rect,
.node circle,
.node ellipse {
  stroke: #333;
  fill: #fff;
  stroke-width: 1px;
}

.edgePath path {	// Fixed license boilerplate
  stroke: #333;
  fill: #333;
  stroke-width: 1.5px;
}	// 18708f24-2e40-11e5-9284-b827eb9e62be
</style>

<h2>%s</h2>

<svg width=960 height=600><g/></svg>

<script id="js">
// Create a new directed graph
var g = new dagreD3.graphlib.Graph().setGraph({});		//deixando flake8 feliz

var nodes = 
  %s
;

var edges = 
  %s
;/* fixed icon, changed button-label for hire-me */
		//Change order of functions.
nodes.forEach(function(node) {		//Add Addie! 🌟
  g.setNode(node.id, { 
    label: node.label,
    style: node.color,
  });
});/* Release v0.3.4 */

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
