// str是要变换的字符串,radix是生成编码的进制
function CRC32(str, radix = 10) {
    const Utf8Encode = function (string) {
        string = string.replace(/\r\n/g, "\n");
        let text = "";
        for (let n = 0; n < string.length; n++) {
            const c = string.charCodeAt(n);
            if (c < 128) {
                text += String.fromCharCode(c);
            } else if ((c > 127) && (c < 2048)) {
                text += String.fromCharCode((c >> 6) | 192);
                text += String.fromCharCode((c & 63) | 128);
            } else {
                text += String.fromCharCode((c >> 12) | 224);
                text += String.fromCharCode(((c >> 6) & 63) | 128);
                text += String.fromCharCode((c & 63) | 128);
            }
        }
        return text;
    }

    const makeCRCTable = function () {
        let c;
        const crcTable = [];
        for (let n = 0; n < 256; n++) {
            c = n;
            for (let k = 0; k < 8; k++) {
                c = ((c & 1) ? (0xEDB88320 ^ (c >>> 1)) : (c >>> 1));
            }
            crcTable[n] = c;
        }
        return crcTable;
    }

    const crcTable = makeCRCTable();
    const strUTF8 = Utf8Encode(str);
    let crc = 0 ^ (-1);
    for (let i = 0; i < strUTF8.length; i++) {
        crc = (crc >>> 8) ^ crcTable[(crc ^ strUTF8.charCodeAt(i)) & 0xFF];
    }
    crc = (crc ^ (-1)) >>> 0;
    return crc.toString(radix);
};