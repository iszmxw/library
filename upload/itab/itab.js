// hmac.js

function b64_hmac(str1, str2) {
    return base64Encode(encryptOrHash(str1, str2), "=");
}

/**
 * 将输入字符串进行 Base64 编码
 * @param {string} input - 输入字符串
 * @param {string} paddingChar - Base64 填充字符，默认为 '='
 * @returns {string} - Base64 编码后的字符串
 */
function base64Encode(input, paddingChar = "=") {
    // 如果未提供填充字符，默认使用 '='
    void 0 === paddingChar && (paddingChar = "=");

    var result = "";
    var inputLength = input.length;

    for (var i = 0; i < inputLength; i += 3) {
        // 将三个字符合并成一个数字
        var combined =
            (input.charCodeAt(i) << 16) |
            (i + 1 < inputLength ? input.charCodeAt(i + 1) << 8 : 0) |
            (i + 2 < inputLength ? input.charCodeAt(i + 2) : 0);

        for (var j = 0; j < 4; j += 1) {
            // 判断是否需要填充
            if (8 * i + 6 * j > 8 * input.length) {
                result += paddingChar;
            } else {
                // 获取 Base64 字符并拼接到结果中
                result +=
                    "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/".charAt(
                        (combined >>> (6 * (3 - j))) & 63
                    );
            }
        }
    }
    return result;
}

/**
 * 对输入字符串进行加密或哈希处理
 * @param {string} str1 - 第一个输入字符串
 * @param {string} str2 - 第二个输入字符串
 * @returns {string} - 处理后的字符串结果
 */
function encryptOrHash(str1, str2) {
    // 如果条件为真，则对输入字符串进行 URI 编码和解码
    str1 = true ? uriEncodeDecode(str1) : str1;
    str2 = true ? uriEncodeDecode(str2) : str2;

    // 将 UTF-8 编码的字符串转换为字节数组
    let bytes1 = utf8ToByteArray(str1);

    // 如果字节数组长度大于16，则进行填充操作
    if (bytes1.length > 16) {
        bytes1 = padInput(bytes1, 8 * str1.length);
    }

    // 初始化两个数组
    let xorArray1 = Array(16);
    let xorArray2 = Array(16);

    // 对数组进行异或运算
    for (let i = 0; i < 16; i += 1) {
        xorArray1[i] = 909522486 ^ bytes1[i];
        xorArray2[i] = 1549556828 ^ bytes1[i];
    }

    // 将第二个字符串转换为字节数组，并与填充后的第一个数组拼接
    var concats = xorArray1.concat(utf8ToByteArray(str2));
    var concatsLen = 512 + 8 * str2.length;
    // console.log(concatsLen, concats);
    let bytes2 = padInput(concats, concatsLen);
    console.log(bytes2);
    console.log(bytes2);

    // 将两个数组拼接并填充到672位，然后转换为字符串返回结果
    return convertArrayToString(padInput(xorArray2.concat(bytes2), 672));
}

/**
 * 将字符串转换为 UTF-8 编码的字节数组
 * @param {string} input - 输入字符串
 * @returns {Array} - UTF-8 编码的字节数组
 */
function utf8ToByteArray(input) {
    var bitLength = 8 * input.length;
    var byteArray = Array(input.length >> 2);
    var byteArrayLength = byteArray.length;

    // 初始化字节数组
    for (var i = 0; i < byteArrayLength; i += 1) {
        byteArray[i] = 0;
    }

    // 将字符转换为 UTF-8 编码的字节数组
    for (var j = 0; j < bitLength; j += 8) {
        var charCodeAts = input.charCodeAt(j / 8);
        var index = j >> 5;
        if (index > byteArray.length - 1) {
            // continue;
        }
        byteArray[index] |= (255 & charCodeAts) << (24 - (j % 32));
    }

    return byteArray;
}

// 对字符串进行两次 URI 编码和解码，用于确保特殊字符正确处理
function uriEncodeDecode(str) {
    var doubleEncodedStr = encodeURIComponent(decodeURIComponent(str));
    return doubleEncodedStr;
}

/**
 * 执行 SHA-1 哈希算法的填充操作
 * @param {Array} message - 输入消息的字节数组
 * @param {number} length - 输入消息的位长度
 * @returns {Array} - 填充后的消息数组
 */
function padInput(message, length) {
    const w = new Array(80).fill(0);
    let [h0, h1, h2, h3, h4] = [
        1732584193, -271733879, -1732584194, 271733878, -1009589776,
    ];

    // 进行填充
    const index = 15 + (((length + 64) >> 9) << 4);
    message[length >> 5] |= 128 << (24 - (length % 32));
    message[index] = length;

    for (let block = 0; block < message.length; block += 16) {
        // 初始化哈希值
        let [a, b, c, d, e] = [h0, h1, h2, h3, h4];

        // 处理每个 512 位块
        for (let i = 0; i < 80; i++) {
            const isLessThan16 = i < 16;
            w[i] = isLessThan16
                ? message[block + i]
                : leftRotate(w[i - 3] ^ w[i - 8] ^ w[i - 14] ^ w[i - 16], 1);

            const tempW = w[i];
            const tempA = leftRotate(a, 5);
            const tempB = sha1F(i, b, c, d);
            const tempC = safeAdd(e, tempW);
            const tempD = sha1K(i);

            const temp = safeAdd(safeAdd(tempA, tempB), safeAdd(tempC, tempD));

            e = d;
            d = c;
            c = leftRotate(b, 30);
            b = a;
            a = temp;
        }

        // 更新哈希值
        h0 = safeAdd(h0, a);
        h1 = safeAdd(h1, b);
        h2 = safeAdd(h2, c);
        h3 = safeAdd(h3, d);
        h4 = safeAdd(h4, e);
    }

    // 返回最终的哈希值数组
    return [h0, h1, h2, h3, h4];
}

// 辅助函数：左移
function leftRotate(value, shift) {
    return (value << shift) | (value >>> (32 - shift));
}

/**
 * 辅助函数：安全的加法，处理位溢出
 * @param {number} x - 第一个加数
 * @param {number} y - 第二个加数
 * @returns {number} - 加法结果
 */
function safeAdd(x, y) {
    // 获取低16位的和
    var lsw = (x & 0xffff) + (y & 0xffff);

    // 获取高16位的和，并考虑低位的进位
    var msw = (x >>> 16) + (y >>> 16) + (lsw >>> 16);
    // 将结果组合成32位整数
    return (msw << 16) | (lsw & 0xffff);
}

/**
 * SHA-1 算法中使用的函数 F
 * @param {number} t - 轮次
 * @param {number} b - B
 * @param {number} c - C
 * @param {number} d - D
 * @returns {number} - 函数 F 的计算结果
 */
function sha1F(t, b, c, d) {
    if (t < 20) return (b & c) | (~b & d); // 0 <= t < 20
    if (t < 40) return b ^ c ^ d; // 20 <= t < 40
    if (t < 60) return (b & c) | (b & d) | (c & d); // 40 <= t < 60
    return b ^ c ^ d; // 60 <= t < 80
}

/**
 * SHA-1 算法中使用的常量 K
 * @param {number} t - 轮次
 * @returns {number} - 常量 K 的值
 */
function sha1K(t) {
    const constants = [1518500249, 1859775393, -1894007588, -899497514];
    if (t < 20) return constants[0];
    if (t < 40) return constants[1];
    if (t < 60) return constants[2];
    return constants[3];
}

/**
 * 将输入的字节数组转换为字符串
 * @param {Array} array - 输入的字节数组
 * @returns {string} - 转换后的字符串
 */
function convertArrayToString(array) {
    // 计算字符串的长度，每个元素占32位
    const length = 32 * array.length;
    let result = "";
    // 迭代每8位，从数组中提取相应的字节，并构建字符串
    for (let i = 0; i < length; i += 8) {
        // 获取当前字节的字符编码
        const charCode = (array[i >> 5] >>> (24 - (i % 32))) & 255;
        // 将字符编码转换为字符并添加到结果字符串
        var ress = String.fromCharCode(charCode);
        result += ress;
    }
    // 返回最终的字符串
    return result;
}

// 导出模块中的方法
module.exports = {
    b64_hmac,
    safeAdd,
    leftRotate,
    sha1F,
    sha1K,
    uriEncodeDecode,
    utf8ToByteArray,
    encryptOrHash,
    convertArrayToString,
};
