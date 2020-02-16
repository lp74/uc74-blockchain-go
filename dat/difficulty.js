/* Example */


// const bits = value => parseInt('0x' + bits, 16);
// const exponent = bits >> 24;
// const mantissa = bits & 0xFFFFFF;
// const target = (mantissa * (2 ** (8 * (exponent - 3)))).toString('16');
        
const GH = value => value * Math.pow(10, 9);
const TH = value => value * Math.pow(10, 12);
const EH = value => value * Math.pow(10, 15);
const hours = value => value * 60 * 60;
const days = value => value * hours(24);
const years = value => value * days(365);

class Energy{
    constructor(){
        this.$price = 0.10 // USD / kWh
    }
}

class Hasher{
    constructor(price, hashrate, energy, power){
        this.$price = price;
        this.$h = hashrate;
        this.$e = energy;
        this.$p = power || (2.9 / TH(73) * this.$h);
    }
    rate(){
        return this.$h;
    }
    energy(t){
        return this.$p / hours(1) * t; // kWh
    }
    eCost(t){
        return this.energy(t) * this.$e.$price;
    }
    dCost(t){
        return this.$price /  years(2) * t;
    }
    cost(t){
        return this.eCost(t) + this.dCost(t);
    }
}

class PoissonMainer{
    constructor(difficulty, hasher, subside, btcusd){
        this.$diff = difficulty;
        this.$hasher = hasher;
        this.$subside = subside || 12.5 // BTC;
        this.$BTCUSD = btcusd, 9500;
    }
    lambda(){
        return this.$hasher.rate() / (Math.pow(2, 32) * this.$diff);
    }
    expectation(t = 1){
        return this.lambda() * t;
    }
    variance(t = 1){
        return this.lambda() * t;
    }
    std(t = 1){
        return Math.sqrt(this.variance(t));
    }
    runtime(probability){
        return (- Math.log(1 - probability) / this.lambda()) / days(1);
    }
    compute(t = 1){
        return `
        ————————————————————————————————————————————————————————————
        Expectation for t = ${t} seconds
        ————————————————————————————————————————————————————————————
        𝔼 = ${this.expectation(t)}
        𝜎 = ${this.std(t)}

        t | P = 0.80 = ${this.runtime(0.80)} days
        ————————————————————————————————————————————————————————————
        
        Subside: ${this.$subside} BTC - 1 BTC = ${this.$BTCUSD} USD
        ————————————————————————————————————————————————————————————
        𝔼 = ${this.expectation(t) * this.$subside} BTC
        𝜎 = ${this.std(t) * this.$subside} BTC

        𝔼 = ${this.expectation(t) * this.$subside * this.$BTCUSD} USD
        𝜎 = ${this.std(t) * this.$subside * this.$BTCUSD} USD
        ————————————————————————————————————————————————————————————

        Consumption
        ————————————————————————————————————————————————————————————
        E = ${this.$hasher.energy(t)} kWh
        
        device cost = ${this.$hasher.dCost(t)} USD
        energy cost = ${this.$hasher.eCost(t)} USD
        total cost = ${this.$hasher.cost(t)} USD
        ————————————————————————————————————————————————————————————
        `
    }
}

k = 1e5;

const energy = new Energy(0.10);
const hasher = new Hasher(k*2000, k*TH(73), energy)
const pm = new PoissonMainer(15466098935554.65, hasher, 12.5, 9850 );



const hasher0 = new Hasher(k * 2000, k * GH(1), energy)
const pm0 = new PoissonMainer(1690906, hasher0 , 50);

console.log(pm.compute(days(1)))
