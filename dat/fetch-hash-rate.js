fetch = require('node-fetch');

const URL_PARAMS = {
    TIMESPAN: {
        'ALL': 'all',
    },
    SCALE : {
        'LINEAR': '0'
    }
}

const hashRateURL = (timespan = URL_PARAMS.TIMESPAN.ALL, scale = URL_PARAMS.SCALE.LINEAR) => {
    return `https://www.blockchain.com/charts/hash-rate?timespan=${timespan}&scale=${scale}`
}

fetch(hashRateURL())
.then();
