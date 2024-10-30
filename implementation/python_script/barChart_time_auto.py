import matplotlib.pyplot as plt
import re
import numpy as np
from scipy import stats

def calculate(file_name):
    pattern = re.compile(rf"{file_name}_output.*took (\d+) seconds")
    runtimes = {}
    with open("../all_time.log", "r") as file:
        for line in file:
            match = pattern.search(line)
            if match:
                runtime = int(match.group(1))
                if file_name not in runtimes:
                    runtimes[file_name] = []
                runtimes[file_name].append(runtime)
    for file_name, times in runtimes.items():
        avg_runtime = np.mean(times)
        confidence_interval = stats.t.interval(
                0.95, len(times) - 1, loc=avg_runtime, scale=stats.sem(times)
        )
    return [avg_runtime, confidence_interval]


average_times = []
confidence_intervals = []
categories = ['PTBR','FC','EEN','MGF','G31','S0811','EEU', 'WS','CA','A0505','WG','SPR', 'CLJ', 'LJ1']
file_names = ["musae_PTBR_features.json", "facebook_combined.txt.json", "email-Enron.txt.json", "musae_git_features.json", "p2p-Gnutella31.txt.json","soc-Slashdot0811.txt.json","email-EuAll.txt.json","web-Stanford.txt.json", "com-amazon.ungraph.txt.json", "amazon0505.txt.json","web-Google.txt.json", "soc-pokec-relationships.txt.json", "com-lj.ungraph.txt.json", "com-lj.ungraph.txt.json"]
vertices = [1912,4039,36692,37700,62586,77357,265214,281903,334863,410236,875713,1632803,3997962,4847571]

for file in file_names:
    result = calculate(file)
    average_times.append(result[0])
    confidence_intervals.append([result[0] - result[1][0], result[1][1] - result[0]])

fig = plt.figure(1, figsize=(9,9))
ax1  = fig.add_subplot(111)
ax1.set_ylabel('Time (Minutes)', fontsize=16)
ax2 = ax1.twinx()
ax2.plot(categories, vertices, color='red', marker='o')
ax2.set_ylabel('Number of Vertices', color='red', fontsize=16)

# Creating the bar chart
ax1.bar(categories, average_times, yerr=confidence_intervals, capsize=5, ecolor='black')

# Adding labels and title
plt.xlabel('Graphs', fontsize=16)
#plt.ylabel('Time (Minutes)', fontsize=16)
#plt.yscale('log')
#plt.title('Total Running Time', fontsize=16)
# Displaying the chart
plt.savefig("total_time_confidence.svg", format="svg")
