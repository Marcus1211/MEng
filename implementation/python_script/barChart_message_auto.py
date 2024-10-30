import matplotlib.pyplot as plt
import re
import numpy as np
from scipy import stats

# Data for the bar chart
#categories = ['PTBR','FC','EEN','MGF','G31','S0811','EEU', 'WS','CA','A0505','WG']
#values = [34,78,55,95,15,221,337,376,24,572,935]
#confidence_intervals = [3,8,5,9,1,22,28,31,2,40,77]
#vertices = [1912,4039,36692,37700,62586,77357,265214,281903,334863,410236,875713]
#custom_ticks = [0, 10000, 100000, 1000000, 10000000, 20000000, 50000000, 100000000, 500000000]





categories = ['CA','A0505']
file_names = ["com-amazon.ungraph.txt.json", "amazon0505.txt.json"]
vertices = [334863,410236]
total_msg = []
for n in file_names:
    msg = []
    for i in range(1, 11):  # Loop from 1 to 10
        temp_count = 0
        filename = f'../{n}_output_run_{i}.txt'
        pattern = re.compile(r"\bNode\s+\d+\s+sent\s+(\d+)\s+messages\b")
        with open(filename, "r") as file:
            for line in file:
                match = pattern.search(line)
                if match:
                    msg_count = int(match.group(1))
                    temp_count += msg_count
        msg.append(temp_count)
    total_msg.append(np.mean(msg))

fig = plt.figure(1, figsize=(9,9))
ax1  = fig.add_subplot(111)
print(total_msg)
# Creating the bar chart
ax1.bar(categories, total_msg)
ax1.set_ylabel('Messages', fontsize=16)

ax2 = ax1.twinx()
ax2.plot(categories, vertices, color='red', marker='o')
ax2.set_ylabel('Number of Vertices', color='red', fontsize=16)

# Adding labels and title
plt.xlabel('Graphs', fontsize=16)
#plt.ylabel('Messages', fontsize=16)
#plt.yscale('log')
# Displaying the chart
plt.savefig("total_msg.svg", format="svg")
plt.show()
