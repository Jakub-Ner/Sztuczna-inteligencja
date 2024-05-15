import re
import matplotlib.pyplot as plt


def extract_data(file_path, pattern):
    finished_in = []

    # Define the regular expression pattern for extracting "Finished in" times
    finished_pattern = re.compile(pattern +r'(\d+)')

    with open(file_path, 'r') as file:
        for line in file:
            match = finished_pattern.search(line)
            if match:
                finished_in.append(int(match.group(1)))

    return finished_in

def average_players(data):
    avg = []
    n = (len(data) // 2) * 2
    for i in range(0, n, 2):
      avg.append((data[i] + data[i+1]) / 2)
    return avg

# Plot all depths on the same graph
def plot_all_depths(depth_3, depth_4, depth_5, ylabel, title):
    plt.plot(depth_3, label='Depth 3')
    plt.plot(depth_4, label='Depth 4')
    plt.plot(depth_5, label='Depth 5')
    plt.xlabel('Round')
    plt.ylabel(ylabel)
    plt.title(title)
    plt.legend()

    plt.show()

DEPTH_3_FILENAME = './depth eq 3.html'
DEPTH_4_FILENAME = './depth eq 4.html'
DEPTH_5_FILENAME = './depth eq 5.html'

FINISHED_IN_PATTERN = 'Finished in: '
FINISHED_IN_YLABEL = 'Time [ms]'
FINISHED_IN_PLOT_TITLE = 'Time taken to finish rounds'

VISITED_NODES_PATTERN = 'Visited nodes: '
VISITED_NODES_YLABEL = 'Visited nodes'
VISITED_NODES_PLOT_TITLE = 'Number of visited nodes in rounds'

PRUNED_NODES_PATTERN = 'Pruned nodes: '
PRUNED_NODES_YLABEL = 'Pruned nodes'
PRUNED_NODES_PLOT_TITLE = 'Number of pruned nodes in rounds'

def main(pattern, ylabel, title):
  depth_3_data = extract_data(DEPTH_3_FILENAME, pattern)
  depth_4_data = extract_data(DEPTH_4_FILENAME, pattern)
  depth_5_data = extract_data(DEPTH_5_FILENAME, pattern)

  depth_3_avg = average_players(depth_3_data)
  depth_4_avg = average_players(depth_4_data)
  depth_5_avg = average_players(depth_5_data)

  plot_all_depths(depth_3_avg, depth_4_avg, depth_5_avg, ylabel, title)

main(FINISHED_IN_PATTERN, FINISHED_IN_YLABEL, FINISHED_IN_PLOT_TITLE)
main(VISITED_NODES_PATTERN, VISITED_NODES_YLABEL, VISITED_NODES_PLOT_TITLE)
main(PRUNED_NODES_PATTERN, PRUNED_NODES_YLABEL, PRUNED_NODES_PLOT_TITLE)