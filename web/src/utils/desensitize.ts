// 基础脱敏函数，保留首尾固定长度的字符，中间脱敏
export function string(str, headLen = 3, tailLen = 4, maskChar = '*') {
if (!str || str.length <= headLen + tailLen) {
      // 如果字符串为空或者长度小于等于保留的首尾长度之和，则直接返回原字符串或空字符串
    return str || '';
    }
const maskedPartLength = str.length - headLen - tailLen;
 // 创建脱敏部分的字符串
let maskedPart = new Array(maskedPartLength + 1).join(maskChar);
 // 返回拼接后的脱敏字符串
 return str.slice(0, headLen) + maskedPart + str.slice(-tailLen);
}


export const email = (input)=>{
  if(!input){return "*"}
  var avg;
  var splitted;
  var email1;
  var email2;
  splitted = input.split('@');
  email1 = splitted[0];
  avg = email1.length / 2;
  email1 = email1.substring(0, email1.length - avg);
  email2 = splitted[1];
  return email1 + '***@' + email2; // 输出为81226***@qq.com
}

function phone(input) {
  if(!input){return "*"}
  let reg = /^(1[3-9][0-9])\d{4}(\d{4}$)/; // 定义手机号正则表达式
  input = input.replace(reg, '$1****$2');
  return input; // 185****6696
}

function idcard(card) {
  if(!card){
    return "*"
  }
  const reg = /^(.{6})(?:\d+)(.{4})$/; // 匹配身份证号前6位和后4位的正则表达式
  const maskedIdCard = card.replace(reg, '$1******$2'); // 身份证号脱敏，将中间8位替换为“*”
  return maskedIdCard; // 输出：371782******5896
}

export default {
  phone,idcard,email,string
}
