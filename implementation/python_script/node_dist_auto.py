from datetime import datetime
import re
import matplotlib.pyplot as plt
import numpy as np 


categories = ['PTBR','FC','EEN','MGF','G31','S0811','EEU', 'WS','CA','A0505','WG','SPR', 'CLJ', 'LJ1']
file_names = ["musae_PTBR_features.json", "facebook_combined.txt.json", "email-Enron.txt.json", "musae_git_features.json", "p2p-Gnutella31.txt.json","soc-Slashdot0811.txt.json","email-EuAll.txt.json","web-Stanford.txt.json", "com-amazon.ungraph.txt.json", "amazon0505.txt.json","web-Google.txt.json", "soc-pokec-relationships.txt.json", "com-lj.ungraph.txt.json", "soc-LiveJournal1.txt.json"]

for n in range(len(file_names)):
    final_time = [0,0,0,0,0,0,0,0,0]
    for i in range(1, 11):
        msg = []
        temp_count = 0
        filename = f'../{file_names[n]}_output_run_{i}.txt'
        pattern = re.compile(r"\bNode\s+\d+\s+sent\s+(\d+)\s+messages\b")
        #pattern = re.compile(r"\b(.*)\.(\d+) Node\s+\d+\s+sent\s+(\d+)\s+messages\b")
        with open(filename, "r") as file:
            for line in file:
                match = pattern.search(line)
                if match:
                    msg_count = int(match.group(1))
                    timestamp_str = line.split()[0] + " " + line.split()[1]
                    timestamp_dt = datetime.strptime(timestamp_str, "%Y/%m/%d %H:%M:%S.%f")
                    unix_time = timestamp_dt.timestamp()
                    msg.append((unix_time, msg_count))
        sorted_data = sorted(msg, key=lambda x: x[0])

        interval = (sorted_data[-1][0] - sorted_data[1][0])/8
        y_0 = y_1 = y_2 = y_3 = y_4 = y_5 = y_6 = y_7 = y_8 = 0
        for i in sorted_data:
            if i[0] < sorted_data[1][0] + interval:
                y_0 += 1
            elif i[0] >= sorted_data[1][0] + interval and i[0] < sorted_data[1][0] + interval * 2:
                y_1 += 1
            elif i[0] >= sorted_data[1][0] + interval * 2 and i[0] < sorted_data[1][0] + interval * 3:
                y_2 += 1
            elif i[0] >= sorted_data[1][0] + interval * 3 and i[0] < sorted_data[1][0] + interval * 4:
                y_3 += 1
            elif i[0] >= sorted_data[1][0] + interval * 4 and i[0] < sorted_data[1][0] + interval * 5:
                y_4 += 1
            elif i[0] >= sorted_data[1][0] + interval * 5 and i[0] < sorted_data[1][0] + interval * 6:
                y_5 += 1
            elif i[0] >= sorted_data[1][0] + interval * 6 and i[0] < sorted_data[1][0] + interval * 7:
                y_6 += 1
            elif i[0] >= sorted_data[1][0] + interval * 7 and i[0] < sorted_data[1][0] + interval * 8:
                y_7 += 1
            else:
                y_8 += 1
        final_time[0] += y_0
        final_time[1] += y_1
        final_time[2] += y_2
        final_time[3] += y_3
        final_time[4] += y_4
        final_time[5] += y_5
        final_time[6] += y_6
        final_time[7] += y_7
        final_time[8] += y_8
    
    y = [final_time[0]/10, final_time[1]/10, final_time[2]/10, final_time[3]/10, final_time[4]/10, final_time[5]/10, final_time[6]/10, final_time[7]/10, final_time[8]/10]
    plt.plot(y) 
    plt.xticks(np.arange(0,9,1)) 


    # Adding labels and title
    plt.xlabel('Time Interval', fontsize=16)
    plt.ylabel('Number of Active Nodes', fontsize=16)
    plt.title(f"{categories[n]}", fontsize=16)
    plt.savefig(f"{file_names[n]}_node_over_time.svg", format="svg")
