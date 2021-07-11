import csv
import numpy as np
import matplotlib.pyplot as plt

# read_csv reads test_output.csv
def read_csv():
	heuristics = []
	with open('test_output.csv', 'r') as file:
		reader = csv.reader(file)
		for row in reader:
			heuristic = []
			for test in row:
				heuristic.append(float(test))
			heuristics.append(heuristic)
	return heuristics

# visualize plots min, mean & max values for each heuristic
def visualize(data):
	fig1, ax1 = plt.subplots()
	ax1.set_title('Solve time by heuristic')
	ax1.boxplot(data)
	plt.title('Solve time by heuristic')
	plt.xlabel('Heuristic')
	plt.ylabel('Solve Time (Seconds)')
	plt.xticks([1, 2, 3, 4, 5], ["manhattan", "nilsson", "outRowCol", "hamming", "euclidean"])
	plt.show()

# main reads accuracy.csv & plots accuracy over depth
def main():
	try:
		data = read_csv()
		visualize(data)
	except Exception:
		print("Error: Failed to visualize data. Is data valid?")
		pass

if __name__ == '__main__':
	main()
