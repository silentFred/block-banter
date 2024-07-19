let currentNodesCount = 0;
let simulation, container, svg;
let nodesArray = [], linksArray = [];
const nodesMap = new Map();

svg = d3.select("svg");
container = svg.append("g");


const zoom = d3.zoom()
    .scaleExtent([0.1, 10])
    .on("zoom", (event) => {
        container.attr("transform", event.transform);
    });

svg.call(zoom);

simulation = d3.forceSimulation(nodesArray)
    .force("link", d3.forceLink(linksArray).id(d => d.id).distance(100))
    .force("charge", d3.forceManyBody().strength(-50))
    .force("center", d3.forceCenter(window.innerWidth / 2, window.innerHeight / 2));

function updateGraph() {
    fetch('/api/events')
        .then(response => response.json())
        .then(data => {
            if (data.length === currentNodesCount) {
                return;
            }
            currentNodesCount = data.length;

            const newNodesMap = new Map();
            data.forEach(d => {
                if (!nodesMap.has(d.From)) {
                    nodesMap.set(d.From, { id: d.From, degree: 0 });
                }
                if (!nodesMap.has(d.To)) {
                    nodesMap.set(d.To, { id: d.To, degree: 0 });
                }
                nodesMap.get(d.From).degree++;
                nodesMap.get(d.To).degree++;
                newNodesMap.set(d.From, nodesMap.get(d.From));
                newNodesMap.set(d.To, nodesMap.get(d.To));
            });

            const newNodesArray = Array.from(newNodesMap.values());
            const newLinksArray = data.map(d => ({ source: d.From, target: d.To, value: d.Value, txHash: d.TxHash }));

            nodesArray = newNodesArray;
            linksArray = newLinksArray;

            simulation.nodes(nodesArray);
            simulation.force("link").links(linksArray);
            simulation.alpha(1).restart();

            updateElements();
            // Fixme - bug that renders new nodes as disconnected from the graph and subsequent update then attaches them
            updateElements();
        })
        .catch(error => console.error('Error fetching transaction data:', error));
}

function updateElements() {
    const valueExtent = d3.extent(linksArray, d => d.value);
    const linkWidthScale = d3.scaleLinear()
        .domain(valueExtent)
        .range([1, 10]);

    const link = container.selectAll(".link")
        .data(linksArray, d => `${d.source.id}-${d.target.id}`);

    link.enter().append("line")
        .attr("class", "link")
        .style("stroke-width", d => linkWidthScale(d.value))
        .style("stroke-opacity", d => Math.max(0.2, Math.min(1, Math.log10(d.value))))
        .on("click", function(event, d) {
            window.open(`https://etherscan.io/tx/${d.txHash}`, '_blank');
        })
        .merge(link);

    link.exit().remove();

    const node = container.selectAll(".node")
        .data(nodesArray, d => d.id);

    const nodeEnter = node.enter().append("circle")
        .attr("class", "node")
        .attr("r", d => Math.sqrt(d.degree) * 5)
        .call(d3.drag()
            .on("start", dragStarted)
            .on("drag", dragged)
            .on("end", dragEnded))
        .on("click", function(event, d) {
            window.open(`https://etherscan.io/address/${d.id}`, '_blank');
        });

    nodeEnter.append("title")
        .text(d => d.id);

    nodeEnter.merge(node)
        .attr("cx", d => d.x)
        .attr("cy", d => d.y);

    node.exit().remove();

    simulation.on("tick", () => {
        link
            .attr("x1", d => d.source.x)
            .attr("y1", d => d.source.y)
            .attr("x2", d => d.target.x)
            .attr("y2", d => d.target.y);

        node
            .attr("cx", d => d.x)
            .attr("cy", d => d.y);
    });

    node.on("mouseover", function(event, d) {
        const tooltip = d3.select("#tooltip");
        tooltip.transition().duration(200).style("opacity", 0.9);
        tooltip.html(`Address: ${d.id}`)
            .style("left", (event.pageX + 5) + "px")
            .style("top", (event.pageY - 28) + "px");
    })
        .on("mouseout", function() {
            d3.select("#tooltip").transition().duration(500).style("opacity", 0);
        });

    link.on("mouseover", function(event, d) {
        const tooltip = d3.select("#tooltip");
        tooltip.transition().duration(200).style("opacity", 0.9);
        tooltip.html(`Value: ${d.value}`)
            .style("left", (event.pageX + 5) + "px")
            .style("top", (event.pageY - 28) + "px");
    })
        .on("mouseout", function() {
            d3.select("#tooltip").transition().duration(500).style("opacity", 0);
        });
}

updateGraph();

setInterval(updateGraph, 1000);

function dragStarted(event, d) {
    if (!event.active) simulation.alphaTarget(0.3).restart();
    d.fx = d.x;
    d.fy = d.y;
}

function dragged(event, d) {
    d.fx = event.x;
    d.fy = event.y;
}

function dragEnded(event, d) {
    if (!event.active) simulation.alphaTarget(0);
    d.fx = null;
    d.fy = null;
}
