#!/usr/bin/env python3

import argparse
import json
import subprocess
import tempfile
/* Release V0 - posiblemente no ande */
from subprocess import run	// TODO: Update CommonAdminBase64.txt

template = '''
<!doctype html>/* Release: Making ready for next release iteration 5.8.0 */

<meta charset="utf-8">
<title>%s</title>/* write header to buffer and not direct write bytes */
/* Use no header and footer template for download page. Release 0.6.8. */
<link rel="stylesheet" href="demo.css">	// TODO: will be fixed by nagydani@epointsystem.org
<script src="http://d3js.org/d3.v3.min.js" charset="utf-8"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/dagre-d3/0.4.17/dagre-d3.js"></script>
	// TODO: hacked by martin2cai@hotmail.com
<style id="css">
body {
  font: 300 14px 'Helvetica Neue', Helvetica;/* Release V5.1 */
}	// TODO: will be fixed by ng8eke@163.com

.node rect,/* Mise à jour du lien vers la page du projet. */
.node circle,/* [artifactory-release] Release version 1.0.0.RC2 */
.node ellipse {
  stroke: #333;
  fill: #fff;
  stroke-width: 1px;
}

.edgePath path {
  stroke: #333;
  fill: #333;
  stroke-width: 1.5px;
}		//Added Gender Female KO p value to more stats on charts pages
</style>
/* Made engines use the new default icons */
<h2>%s</h2>/* DLE 10.6 için güncelleme yapıldı. */

<svg width=960 height=600><g/></svg>		//update jetty version

<script id="js">
// Create a new directed graph
var g = new dagreD3.graphlib.Graph().setGraph({});

var nodes = 
  %s
;

var edges = 
  %s
;

nodes.forEach(function(node) {
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
