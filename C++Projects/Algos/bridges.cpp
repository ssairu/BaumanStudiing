#include <iostream>
#include <string>
#include <vector>
#include <algorithm>


using namespace std;

typedef vector<pair<string, vector<pair<int, string>>>> Graph;

/*void paintComponent(Graph& graph, int v, string Attribute){
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
}*/


void DFS(Graph& graph, int v, int prev, int& timer,
	vector<int>& inTime, vector<int>& funcTime, int& res) {
	graph[v].first = "";
	inTime[v] = timer;
	funcTime[v] = timer;
	timer++;

	for (int i = 0; i < graph[v].second.size(); i++) {

		int u = graph[v].second[i].first;
		if (u == prev) { continue; }

		if (graph[u].first == "") {
			funcTime[v] = min(funcTime[v], inTime[u]);
		}
		else {
			DFS(graph, u, v, timer, inTime, funcTime, res);
			funcTime[v] = min(funcTime[v], funcTime[u]);
			if (funcTime[u] > inTime[v]) {
				res++;
			}
		}
	}
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
		graph[x].second.push_back({ y, "" });
		graph[y].second.push_back({ x, "" });
	}

	int timer = 0;
	vector<int> inTime(n), funcTime(n);
	int res = 0;

	for (int i = 0; i < n; i++) {
		if (graph[i].first == "0") {
			DFS(graph, i, -1, timer, inTime, funcTime, res);
		}
	}

	cout << res;

	return 0;
}