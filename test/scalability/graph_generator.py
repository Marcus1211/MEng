import snap
import json

# Function to generate a graph and save it as JSON
def generate_rmat_graph(output_file):
    """
    Generate an R-MAT graph and save it as a JSON file.

    Args:
        output_file (str): Path to save the JSON file.
        num_nodes (int): Number of nodes in the graph.
        num_edges (int): Number of edges in the graph.
        a, b, c, d (float): R-MAT parameters for edge placement probabilities.
    """
    # Generate the R-MAT graph
    Rnd = snap.TRnd()
    graph = snap.GenRMat(5000000, 40000000, .45, .15, .15, Rnd)
    # Convert the graph to an adjacency list
    graph_dict = {}
    for node in graph.Nodes():
        graph_dict[node.GetId()] = [neighbor for neighbor in node.GetOutEdges()]

    # Save the graph to a JSON file
    with open(output_file, "w") as f:
        json.dump(graph_dict, f, indent=4)
    print(f"Graph saved to {output_file}")

# Main execution
if __name__ == "__main__":
    output_file = "rmat_5000000.json"
    generate_rmat_graph(output_file)
