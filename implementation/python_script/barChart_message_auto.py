import matplotlib.pyplot as plt
import re
import numpy as np
from scipy import stats
import sys

categories = ['PTBR','FC','EEN','MGF','G31','S0811','EEU', 'WS','CA','A0505','WG','SPR', 'CLJ', 'LJ1']
file_names = ["musae_PTBR_features.json", "facebook_combined.txt.json", "email-Enron.txt.json", "musae_git_features.json", "p2p-Gnutella31.txt.json","soc-Slashdot0811.txt.json","email-EuAll.txt.json","web-Stanford.txt.json", "com-amazon.ungraph.txt.json", "amazon0505.txt.json","web-Google.txt.json", "soc-pokec-relationships.txt.json", "com-lj.ungraph.txt.json", "soc-LiveJournal1.txt.json"]
vertices = [1912,4039,36692,37700,62586,77357,265214,281903,334863,410236,875713,1632803,3997962,4847571]

total_msg = []

runs = sys.argv[1]
for n in file_names:
    msg = []
    for i in range(1, int(runs) + 1):  # Loop all files with logs
        temp_count = 0
        filename = f'../{n}_output_run_{i}_logEnabled.txt'
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
plt.yscale('log')
# Displaying the chart
plt.savefig("total_msg.svg", format="svg")

