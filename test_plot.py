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
			# heuristic.append(float(row[0]))
			# heuristic.append(float(row[1]))
			# heuristic.append(float(row[2]))
			heuristics.append(heuristic)
	return heuristics

# visualize plots min, mean & max values for each heuristic
def visualize(data):
	fig1, ax1 = plt.subplots()
	ax1.set_title('Solve time by heuristic')
	ax1.boxplot(data)
	# ax1.labels("manhattan" "nilsson" "outRowCol" "hamming" "euclidean")
# 	depth = list(range(1, len(test_accuracy) + 1))

# 	plt.plot(depth, train_accuracy, label='train accuracy')
# 	plt.plot(depth, test_accuracy, label='test accuracy')
# 	plt.title('Accuracy over Depth')
	plt.xlabel('Heuristic')
	plt.ylabel('Solve time (seconds)')
# 	plt.legend()
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






# Fixing random state for reproducibility
# np.random.seed(19680801)

# # fake up some data
# spread = np.random.rand(50) * 100
# center = np.ones(25) * 50
# flier_high = np.random.rand(10) * 100 + 100
# flier_low = np.random.rand(10) * -100
# data = np.concatenate((spread, center, flier_high, flier_low))
# Copy to clipboard
