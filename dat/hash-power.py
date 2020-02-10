import json

import matplotlib.pyplot as plt


with open('hash-power.json') as json_file:
    data = json.load(json_file)
    print data['values']
    plt.scatter('x', 'y', c='c', s='d', data=data['values'][0:10])
    plt.xlabel('entry a')
    plt.ylabel('entry b')
    plt.show()