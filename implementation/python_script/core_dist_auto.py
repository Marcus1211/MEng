import re
from collections import defaultdict
import matplotlib.pyplot as plt

categories = ['PTBR','FC','EEN','MGF','G31','S0811','EEU', 'WS','CA','A0505','WG','SPR', 'CLJ', 'LJ1']
file_names = ["musae_PTBR_features.json", "facebook_combined.txt.json", "email-Enron.txt.json", "musae_git_features.json", "p2p-Gnutella31.txt.json","soc-Slashdot0811.txt.json","email-EuAll.txt.json","web-Stanford.txt.json", "com-amazon.ungraph.txt.json", "amazon0505.txt.json","web-Google.txt.json", "soc-pokec-relationships.txt.json", "com-lj.ungraph.txt.json", "com-lj.ungraph.txt.json"]

data = []
for n in file_names:
    core_count = defaultdict(int)
    filename = f'../{n}_output_run_1.txt'
    pattern = re.compile(r"final core number of\s+(\d+)")
    with open(filename, "r") as file:
        for line in file:
            match = pattern.search(line)
            if match:
                core_number = int(match.group(1))
                core_count[core_number] += 1
    temp_core = []
    temp_count = []
    for core_number, count in sorted(core_count.items()):
        temp_core.append(core_number)
        temp_count.append(count)

    data.append((temp_core,temp_count))

custom_ticks = [0, 100, 1000, 10000, 100000, 200000, 500000, 1000000]
fig = plt.figure(1, figsize=(9,9))
ax1  = fig.add_subplot(111)
# Plotting the data
for x, y in data:
    ax1.plot(x, y)

#ax1.yaxis.set_ticks(custom_ticks)
#ax1.set_yscale('log', basex=10)
#ax1.yaxis.set_ticklabels(custom_ticks)

# Adding labels and title
plt.xlabel('Core Numbers', fontsize=16)
plt.ylabel('Number of Vertices', fontsize=16)
plt.yscale('log')
# Displaying the plot
plt.legend(['CA', 'A0505'])  # Add legend for each line
plt.savefig("core_number_distribution.svg", format="svg")
plt.show()