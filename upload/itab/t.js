// 使用模块的文件

// 引入你的模块
const itab = require('./itab'); // 根据文件路径可能需要调整

// 使用模块中的方法
// const result = hx.b64_hmac("5Y5b2hWisRnXXNUJOcPtkg1v2R9dZK", "eyJleHBpcmF0aW9uIjoiMjAyMy0xMi0wNVQwODowNzoyMS4xMjNaIiwiY29uZGl0aW9ucyI6W3siYnVja2V0IjoieGRsdW1pYTIifSx7ImtleSI6InVzZXItd2Vic2l0ZS1pY29uLzIwMjMxMjA0L1dOUTZPY19YcGNOdG11VGZxTk5INzA0ODUucG5nIn0sWyJjb250ZW50LWxlbmd0aC1yYW5nZSIsMCwxMDczNzQxODI0XV19");
// console.log(result);
console.log(itab.encryptOrHash('1','1'));
// console.log(hx.convertArrayToString(Array(1)));
// console.log(hx.safeAdd(1732584193,1655872086));
// console.log(hx.leftRotate(12345,5));
// console.log(hx.sha1F(30,123,456,789));
// console.log(hx.sha1K(30));
// console.log(hx.uriEncodeDecode('Hello World!@#$'));
// console.log(hx.utf8ToByteArray('abcd'));
// console.log(hx.base64Encode('Hello, World!','='));
