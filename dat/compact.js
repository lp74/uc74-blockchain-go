var MAX_PRECISION = 28;

function _findExp(n, base, exp = 1) {
    var pow = Math.pow(base, exp);
    var div = Math.floor(n / pow);
    if (exp > MAX_PRECISION) {
        return div;
    }
    return div + _findExp(n, base, exp + 1);
}

// Convert an integer, `n`, to a specified numeric
// base, `base`
function _toBase(n, base) {
    var rem = n % base;
    return _findExp(n, base) + ' ' + rem;
}

// Convert a integer string, `s`, into hexadecimal value
// between 00 and ff. The leftmost value is padded
// with a 0 if needed.
function _toHex(s) {
    var hex = parseInt(s).toString(16);
    return (hex.length === 1) ? '0' + hex : hex;
}

function _toBits(target) {
    var arr = _toBase(target, 256).split(' ');

    // If the first digit is greater than 0x7f, prepend an 0x00
    if (arr[0] > 0x7f) {
        arr.unshift(0);
    }

    // Prepend the length of the hex string
    arr.unshift(arr.length);

    // If there are less than 4 bytes in total,
    // right-pad with 0x00s
    var delta = 4 - arr.length;
    while (delta > 0) {
        arr.push(0);
        delta--;
    }

    // Only keep 2 bytes of precision
    return '0x' + arr.slice(0, 4).map(_toHex).join('');
}