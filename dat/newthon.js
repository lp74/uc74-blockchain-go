function nrRoot(f, derivative, guess, options) {
    options = options || {};
    var tolerance = options.tolerance || 0.00000001;
    var epsilon = options.epsilon || 0.0000000000001;
    var maxIterations = options.maxIterations || 20;
    var root;

    for (var i = 0; i < maxIterations; ++i) {
        var denominator = derivative(guess);
        if (Math.abs(denominator) < epsilon) {
            return false
        }

        root = guess - (f(guess) / denominator);

        var resultWithinTolerance = Math.abs(root - guess) < tolerance;
        if (resultWithinTolerance) {
            return root
        }

        guess = root;
    }

    return false;
}


const m = 2.29*1e-8;
const T2W = 14 * 24 * 60 * 60;
const f = x => x * Math.exp(m*x) - T2W;
const fprime = x => Math.exp(m*x) * (1 + m*x);

const t2016 = nrRoot(f, fprime, 1);
const tMean = t2016 / 2016;
const ratio = 600 / tMean;
console.log(t2016, f(t2016), tMean, ratio)

const movingAvg = w => (x, idx, xs) => {
    const iw = (w -1)/2
    const is = Math.max(0, idx-iw);
    const ie = Math.min(idx+iw, xs.length-1);
    const res = xs.slice(is, ie+1).reduce((a, c) => (a+c), 0)
    return res/(ie - is + 1);
}

const prime = w => (x, idx, xs) => {
    const iw = w / 2
    const is = Math.max(0, idx-iw);
    const ie = Math.min(idx+iw, xs.length-1);
    const res = (xs[ie] - xs[is]) / (w+1);
    return res;
}


const signm = w => (x, idx, xs) => {
    const iw = w / 2
    const is = idx;
    const ie = Math.min(idx+iw, xs.length-1);
    const mul = (xs[ie] * xs[is])
    const res = mul / Math.abs(mul);
    return res;
}

const w = 2;
const data = [-4, 0, 12, 8, 10, 12, 14, 18, 22, 26, 30];
const mavg = data.map(prime(w))

console.log(mavg)
console.log(data
    .map(prime(w))
    .map(prime(w))
    .map(prime(w))
)

console.log(data
    .map(prime(w))
    .map(prime(w))
    .map(prime(w))
)