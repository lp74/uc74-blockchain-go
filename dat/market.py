import requests
import matplotlib.pyplot as plt
import numpy as np
import pandas as pd
from scipy.stats import norm
import seaborn as sns
from termcolor import colored
import time
import os

sns.set()
# ethusdt
# btcusdt
# bchusdt
# bsvusdt
# ltcusdt

# Chiedere i dati una volta sola

clear = lambda: os.system('clear')

def getData(symbol, period):
    url = 'https://api.huobi.pro/market/history/kline?symbol='+symbol+'&period='+period+'&size=1440'
    response = requests.get(url)
    data = response.json()['data']
    return data

def compute(data, first, last, symbol):
    delta = [(x['close'] - x['open'])/x['open'] * 100 for x in data[0:first+1]]

    nDelta = np.array(delta)

    mu, std = norm.fit(nDelta)

    #fig, ax1 = plt.subplots()
    #ax1.grid(False)
    #ax2 = ax1.twinx()
    #ax2.grid(False)

    #ax1.hist(nDelta, color='#d7d7d7', bins=30)


    # Plot the PDF.
    xmin, xmax = plt.xlim()
    x = np.linspace(-.5, .5)
    p = norm.pdf(x, mu, std)
    #ax2.plot(x, p, 'k', linewidth=2, c='#1E35E6')
    #title = symbol.upper() + " : MEAN = %.3f,  DEV = %.3f" % (mu, std)
    #plt.title(title)

    #plt.vlines(0, 0, 10, colors='#3c3c3c')
    #plt.vlines(mu, 0, norm.pdf(mu, mu, std), colors='#d90000')
    #plt.vlines(3*std, 0, norm.pdf(mu, mu, std), colors='#d90000')
    
    #plt.fill_between(np.linspace(mu, 3*std),norm.pdf(np.linspace(mu, 3*std), mu, std), alpha=0.5)
    
    #plt.show()

    tOpen = data[first]['open']
    tClose = data[last]['close']
    tDelta = formatColor(tClose - tOpen)
    tRate = formatRateColor((tClose - tOpen)/tOpen * 100) 
    cdf = (norm.cdf(3*std, mu, std) - norm.cdf(0, mu, std)) * 100
    rate = formatColor(mu * (60))
    hours = (first + 1)/60.
    cdfInv = (norm.cdf(0, mu, std) - norm.cdf(-3*std, mu, std)) * 100
    
    print ('| {:7s} | {:4.1f} h | {:12.3f} | {:12.3f} | {} | {} | {:12.3f} | {:12.3f} | {} | {:12.1f} | {:12.1f} |'.format(symbol.upper(), hours, tOpen, tClose, tDelta, tRate, mu, std, rate, cdf, cdfInv))


def formatColor(value):
    if value < 0: 
        value = colored('{:12.3f}'.format(value), 'red') 
    else:
        value = colored('{:12.3f}'.format(value), 'green')
    return value

def formatRateColor(value):
    res = None
    if value < 0:
        if value < -2:
            res = colored('{:12.3f}'.format(value), 'red', attrs=['bold']) 
        else:
            res = colored('{:12.3f}'.format(value), 'red') 
    
    if value >= 0:
        if value > 2:
            res = colored('{:12.3f}'.format(value), 'green', attrs=['reverse']) 
        else:
            res = colored('{:12.3f}'.format(value), 'green') 
    return res


def checkNormal(symbol, period, size):
    data = getData(symbol, period)
    compute(data, 14, 0, symbol)
    compute(data, 29, 0, symbol)
    compute(data, 59, 0, symbol)
    compute(data, 119, 0, symbol)
    compute(data, 479, 0, symbol)
    compute(data, 1439, 0, symbol)

symbols = ['btcusdt', 'bchusdt', 'bsvusdt', 'ltcusdt', 'ethusdt', 'xmrusdt', 'xrpusdt']

while(True):
    clear()
    print ('| SYM     |  SIZE  | OPEN         | CLOSE        | DELTA        | VARIA        | MEAN         | SIGMA        | RATE         | PROBA +      | PROBA -      |')
    print('-' * 155)
    for sym in symbols:
        checkNormal(sym, '1min', '1440')
        print('-' * 155)
    time.sleep(60)

