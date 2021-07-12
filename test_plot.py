import csv
import numpy as np
import matplotlib.pyplot as plt

# read_csv reads solve_time.csv
def read_csv(filepath):
	heuristics = []
	with open(filepath, 'r') as file:
		reader = csv.reader(file)
		for row in reader:
			heuristic = []
			for test in row:
				heuristic.append(float(test))
			heuristics.append(heuristic)
	return heuristics

# visualize plots min, mean & max values for each heuristic
def visualize(data, title, y_label):
	_, ax1 = plt.subplots()
	ax1.boxplot(data)
	plt.title(title)
	plt.ylabel(y_label)
	plt.xlabel('Heuristic')
	if len(data) == 5:
		plt.xticks([1, 2, 3, 4, 5], ["manhattan", "nilsson", "outRowCol", "hamming", "euclidean"])
	if len(data) == 2:
		plt.xticks([1, 2], ["manhattan", "nilsson"])
	plt.show()

# main reads .csv & visualizes test output
def main():
	try:
		solve_time = read_csv('solve_time.csv')
		visualize(solve_time, 'Solve time by heuristic', 'Solve Time (Seconds)')

		size_complexity = read_csv('size_complexity.csv')
		visualize(size_complexity, 'Size complexity by heuristic', 'Size complexity')

		time_complexity = read_csv('time_complexity.csv')
		visualize(time_complexity, 'Time complexity by heuristic', 'Time complexity')

		moves = read_csv('moves.csv')
		visualize(moves, 'Moves by heuristic', 'Moves')

	except Exception:
		print("Error: Failed to visualize data. Is data valid?")
		pass

if __name__ == '__main__':
	main()
