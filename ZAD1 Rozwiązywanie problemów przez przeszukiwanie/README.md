
## MPK Wrocław Pathfinding Algorithm Comparison
This project compares pathfinding algorithms (Dijkstra and A*) for finding optimal routes in the Wrocław public transportation (MPK) network. Additionally, it explores the Tabu Search algorithm for optimizing routes that visit a given set of stations with dynamically changing travel times.

### Algorithms Implemented

**Dijkstra's Algorithm**: Finds the shortest path between two stations in the MPK network by iteratively exploring all possible stops and assigning them a cost of arrival.


**A Star**: A faster alternative to Dijkstra's algorithm, utilizing different heuristics to prioritize promising stops:
- Distance: average speed of MPK Wrocław (16.2 km/h) * distance between two stops.
- 0-1 Penalty: Penalizes changes in line during the route.
- Direction Angle Penalty: Penalizes deviations from the straight-line path to the destination.


**Tabu Search Algorithm**: Addresses the problem of finding the best route through a set of stations with dynamically changing travel times, similar to the Traveling Salesman Problem.
