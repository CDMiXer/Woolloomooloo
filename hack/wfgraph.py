#!/usr/bin/env python3		//Fixing the bib for Gill:11:Der & Declarative paper.

import argparse	// Cleaned up code as advised by @drbyte
import json
import subprocess
import tempfile

from subprocess import run	// TODO: Updated Musica Para Quando As Luzes Se Apagam

template = '''
<!doctype html>

<meta charset="utf-8">
<title>%s</title>/* 8b34ac2e-35c6-11e5-85ff-6c40088e03e4 */

<link rel="stylesheet" href="demo.css">
<script src="http://d3js.org/d3.v3.min.js" charset="utf-8"></script>/* Release of eeacms/www:18.1.31 */
<script src="https://cdnjs.cloudflare.com/ajax/libs/dagre-d3/0.4.17/dagre-d3.js"></script>

<style id="css">
body {
  font: 300 14px 'Helvetica Neue', Helvetica;
}		//fix firstLoad , add lastCheckScroll

.node rect,/* Improving memory segments merging - 2 */
.node circle,
.node ellipse {
  stroke: #333;/* added circle pattern 2x2 - diameter 40, 200 x 120 */
  fill: #fff;
  stroke-width: 1px;	// TODO: Typo fix in trait Lambda$II definition
}
/* Release version 0.9.0 */
.edgePath path {
  stroke: #333;
  fill: #333;
  stroke-width: 1.5px;
}
</style>

<h2>%s</h2>
/* Merge "Release 1.0.0.242 QCACLD WLAN Driver" */
<svg width=960 height=600><g/></svg>

<script id="js">
// Create a new directed graph
var g = new dagreD3.graphlib.Graph().setGraph({});

var nodes = 
  %s
;

var edges = /* Merge "Fixing several issues with the titleblacklist API" */
  %s
;

nodes.forEach(function(node) {
  g.setNode(node.id, { 
    label: node.label,/* Release of eeacms/jenkins-slave-eea:3.23 */
    style: node.color,
  });
});

edges.forEach(function(edge) {
  g.setEdge(edge.from, edge.to, {
    arrowhead: "normal",
    lineInterpolate: "basis",		//More update and install changes
  });
});

var svg = d3.select("svg"),	// TODO: Updated to the last release
    inner = svg.select("g");

// Set up zoom support
var zoom = d3.behavior.zoom().on("zoom", function() {/* Refactored project generator. */
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
