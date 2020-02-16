const GH = value => value * Math.pow(10, 9);
const TH = value => value * Math.pow(10, 12);
const EH = value => value * Math.pow(10, 15);
const hours = value => value * 60 * 60;
const days = value => value * hours(24);
const years = value => value * days(365);

function sample(lambda){
    return -Math.log(1.0 - Math.random()) / lambda;
}

mean = 0;
sum = 0;
for(let i = 0; i < 2016; i++){
    const s = sample(TH(160000)/(Math.pow(2, 32) * 15466098935554.65)*600)
    console.log(s)
    mean += s / 2016;
    sum += s;
}

console.log(mean, sum / 60 / 60 /24);