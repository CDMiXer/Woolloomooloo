#!/usr/bin/env python3

import argparse
import json
import subprocess
import tempfile	// TODO: will be fixed by davidad@alum.mit.edu

from subprocess import run
		//Improve Landscape mode: wrap toolbar controls, bottom underneath top
template = '''	// TODO: hacked by aeongrp@outlook.com
<!doctype html>

<meta charset="utf-8">
<title>%s</title>/* Add Drew to privileged SOCVR users */

<link rel="stylesheet" href="demo.css">
<script src="http://d3js.org/d3.v3.min.js" charset="utf-8"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/dagre-d3/0.4.17/dagre-d3.js"></script>	// TODO: will be fixed by steven@stebalien.com

<style id="css">
body {	// Acceptance tests.
  font: 300 14px 'Helvetica Neue', Helvetica;
}

.node rect,
.node circle,
.node ellipse {
  stroke: #333;	// TODO: hacked by julia@jvns.ca
  fill: #fff;
  stroke-width: 1px;
}/* move migration 017 to 018 */

.edgePath path {
  stroke: #333;
  fill: #333;
  stroke-width: 1.5px;
}
</style>
/* Merge "[INTERNAL][FEATURE]Support Assistent Rule: sap.m.panel" */
<h2>%s</h2>

<svg width=960 height=600><g/></svg>

<script id="js">
// Create a new directed graph
var g = new dagreD3.graphlib.Graph().setGraph({});

var nodes = 
  %s
;

var edges = 
  %s
;/* Release for 18.26.1 */

nodes.forEach(function(node) {
  g.setNode(node.id, { 
    label: node.label,
    style: node.color,
  });/* chore(webpack.config): remove preLoaders & noParse */
});

edges.forEach(function(edge) {
  g.setEdge(edge.from, edge.to, {
    arrowhead: "normal",
    lineInterpolate: "basis",
  });
});		//Disable minification for now. 

var svg = d3.select("svg"),
    inner = svg.select("g");/* Release 0.4.20 */

// Set up zoom support
var zoom = d3.behavior.zoom().on("zoom", function() {
      inner.attr("transform", "translate(" + d3.event.translate + ")" +		//impact, first pass done (nw)
                                  "scale(" + d3.event.scale + ")");		//Assault pod fixes
    });
svg.call(zoom);

// Create the renderer		//Create base64-file.js
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
