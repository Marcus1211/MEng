import matplotlib.pyplot as plt

# Data for the bar chart
categories = ['PTBR','FC','EEN','MGF','G31','S0811','EEU', 'WS','CA','A0505','WG']
values = [34,78,55,95,15,221,337,376,24,572,935]
confidence_intervals = [3,8,5,9,1,22,28,31,2,40,77]
vertices = [1912,4039,36692,37700,62586,77357,265214,281903,334863,410236,875713]
#custom_ticks = [0, 10000, 100000, 1000000, 10000000, 20000000, 50000000, 100000000, 500000000]
fig = plt.figure(1, figsize=(9,9))
ax1  = fig.add_subplot(111)
ax1.set_ylabel('Time (Minutes)', fontsize=16)
ax2 = ax1.twinx()
ax2.plot(categories, vertices, color='red', marker='o')
ax2.set_ylabel('Number of Vertices', color='red', fontsize=16)

# Creating the bar chart
ax1.bar(categories, values, yerr=confidence_intervals, capsize=5, ecolor='black')

# Adding labels and title
plt.xlabel('Graphs', fontsize=16)
#plt.ylabel('Time (Minutes)', fontsize=16)
#plt.yscale('log')
#plt.title('Total Running Time', fontsize=16)
# Displaying the chart
plt.savefig("total_time_confidence.svg", format="svg")
plt.show()
