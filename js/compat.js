function CompatToBig(nBitsStr){
  const bits = BigInt.asUintN(32,nBitsStr);
  let a = bits & BigInt.asUintN(32, "0x007fffff");
  const s = (bits & BigInt.asUintN(32, "0x00800000")) != 0
  const b = bits >> 24n

  if(b <= 3n ){
   a = a >> (BigInt(8) * (3n - b))
  }else{
     a = a << ( 8n * (b - 3n))
  }
  a = (-1n)**BigInt(s) * a

  return a;
}

//console.log(CompatToBig("0x1d00ffff").toString(16))
//console.log(CompatToBig("0x1b0404cb").toString(16))