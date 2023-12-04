#include <iostream>
#include <string>
#include <vector>


using namespace std;

typedef vector<pair<string, vector<pair<int, string>>>> Graph;


void DFS(Graph& graph, int v, int& countV, int& countE) {
	graph[v].first = "1";
	countV++;
	countE += graph[v].second.size();
	for (int i = 0; i < graph[v].second.size(); i++) {
		int u = graph[v].second[i].first;
		if (graph[u].first == "0") {
			DFS(graph, u, countV, countE);
		}
	}
	graph[v].first = "";
}

void paintComponent(Graph& graph, int v, string Attribute){
	graph[v].first = "1";

	
	for (int i = 0; i < graph[v].second.size(); i++) {

		int u = graph[v].second[i].first; //определяем конец ребра
		graph[v].second[i].second = Attribute; //красим ребро


		if (graph[u].first == "") {
			paintComponent(graph, u, Attribute);
		}
	}
	graph[v].first = Attribute; //красим вершину
}

void printGraph(Graph& graph) {
	cout << "graph {\n";
	for (int i = 0; i < graph.size(); i++) {
		cout << "\t" << i << " " << graph[i].first << "\n";
	}
	for (int i = 0; i < graph.size(); i++) {
		for (int j = 0; j < graph[i].second.size(); j++) {
			if (i < graph[i].second[j].first) {
				cout << "\t" << i << " -- " << graph[i].second[j].first;
				cout << " " << graph[i].second[j].second << "\n";
			}
		}
	}
	cout << "}";
}

int main()
{
	int n, m;
	cin >> n >> m;

	Graph graph;

	for (int i = 0; i < n; i++) {
		graph.push_back({ "0", {} });
	}
	
	for (int i = 0; i < m; i++) {
		int x, y;
		cin >> x >> y;
		graph[x].second.push_back({y, ""});
		graph[y].second.push_back({x, ""});		
	}

	int res = 0;
	int maxV = 0;
	int maxE = 0;

	for (int i = 0; i < n; i++) {
		if (graph[i].first == "0") {
			int countV = 0, countE = 0;
			DFS(graph, i, countV, countE);
			if (countV > maxV){
				res = i;
				maxV = countV;
				maxE = countE;
			}
			else if (countV == maxV && countE > maxE) {
				res = i;
				maxE = countE;
			}
			//cout << countV << " - " << countE << " : ";
			//cout << maxV << " - " << maxE << "      " << res;
		}
	}

	paintComponent(graph, res, "[color = red]");
	printGraph(graph);


	return 0;
}