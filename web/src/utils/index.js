import dayjs from "dayjs"
import eventbus from "./eventbus"

/** 格式化时间 */
export const datetimeformat = (time, fmt = "YYYY-MM-DD HH:mm:ss") => {
  if (!time) {
    return "N/A"
  }
  const date = new Date(time)
  return dayjs(date).format(fmt)
}

/** 将全局 CSS 变量导入 JS 中使用 */
export const getCssVariableValue = (cssVariableName) => {
  let cssVariableValue = ""
  try {
    // 没有拿到值时，会返回空串
    cssVariableValue = getComputedStyle(document.documentElement).getPropertyValue(cssVariableName)
  } catch (error) {
    console.error(error)
  }
  return cssVariableValue
}

/**
 * Check if an element has a class
 * @param {HTMLElement} ele
 * @param {string} cls
 * @returns {boolean}
 */
export function hasClass(ele, cls) {
  return !!ele.className.match(new RegExp("(\\s|^)" + cls + "(\\s|$)"));
}

/**
 * Add class to element
 * @param {HTMLElement} ele
 * @param {string} cls
 */
export function addClass(ele, cls) {
  if (!hasClass(ele, cls)) ele.className += " " + cls;
}

/**
 * Remove class from element
 * @param {HTMLElement} ele
 * @param {string} cls
 */
export function removeClass(ele, cls) {
  if (hasClass(ele, cls)) {
    const reg = new RegExp("(\\s|^)" + cls + "(\\s|$)");
    ele.className = ele.className.replace(reg, " ");
  }
}

/**
 * 判断是否是外部链接
 *
 * @param {string} path
 * @returns {Boolean}
 */
export function isExternal(path) {
  const isExternal = /^(https?:|http?:|mailto:|tel:)/.test(path);
  return isExternal;
}


export function camelToUnderscore(str) {
  return str.replace(/([A-Z])/g, "_$1").toLowerCase();
}

export function showlogin(){
  eventbus.emit("showlogin")
}

export function  parseUrl(url) {
  let urlObj = {
    pathname:url,
    protocal:location.protocol,
    host:location.host,
  } 
  try {
     urlObj = new URL(url);
  } catch (error) {
    
  }  
  urlObj.pathname = urlObj.pathname.replace(/\/+/g, '/')
  return urlObj
  
}
export function buildUrl(url){
  if(/^http(s):\/\//.test(url)){
    return url;
  }else{
    return location.protocol +"//" + location.host +"/"+ url
  }
}

export function resolve(...paths){
          let resolvePath = '';
          let isAbsolutePath = false;
          for(let i = paths.length-1; i > -1; i--){
              let path = paths[i];
              if(isAbsolutePath){
                  break;
              }
              if(!path){
                  continue
              }
              resolvePath = path + '/' + resolvePath;
              isAbsolutePath = path.charCodeAt(0) === 47;
          }
          if(/^\/+$/.test(resolvePath)){
              resolvePath = resolvePath.replace(/(\/+)/,'/')
          }else{
              resolvePath = resolvePath.replace(/(?!^)\w+\/+\.{2}\//g, '')
              .replace(/(?!^)\.\//g,'')
              .replace(/\/+$/, '')
          }
          return resolvePath;
      }

  export function getFileType(fileName) {
    let suffix = ''; // 后缀获取
    let result = ''; // 获取类型结果
    if (fileName) {
      const flieArr = fileName.split('.'); // 根据.分割数组
      suffix = flieArr[flieArr.length - 1]; // 取最后一个
    }
    if (!suffix) return false; // fileName无后缀返回false
    suffix = suffix.toLocaleLowerCase(); // 将后缀所有字母改为小写方便操作
    // 匹配图片
    const imgList = ['png', 'jpg', 'jpeg', 'bmp', 'gif']; // 图片格式
    result = imgList.find(item => item === suffix);
    if (result) return 'image';
    // 匹配txt
    const txtList = ['txt'];
    result = txtList.find(item => item === suffix);
    if (result) return 'txt';
    // 匹配excel
    const excelList = ['xls', 'xlsx'];
    result = excelList.find(item => item === suffix);
    if (result) return 'xlsx';
    // 匹配word
    const wordList = ['doc', 'docx'];
    result = wordList.find(item => item === suffix);
    if (result) return 'docx';
    // 匹配pdf
    const pdfList = ['pdf'];
    result = pdfList.find(item => item === suffix);
    if (result) return 'pdf';
    // 匹配ppt
    const pptList = ['ppt', 'pptx'];
    result = pptList.find(item => item === suffix);
    if (result) return 'pptx';
    // 匹配zip
    const zipList = ['rar', 'zip', '7z'];
    result = zipList.find(item => item === suffix);
    if (result) return 'zip';
    // 匹配视频
    const videoList = ['mp4', 'm2v', 'mkv', 'rmvb', 'wmv', 'avi', 'flv', 'mov', 'm4v'];
    result = videoList.find(item => item === suffix);
    if (result) return 'video';
    // 匹配音频
    const radioList = ['mp3', 'wav', 'wmv'];
    result = radioList.find(item => item === suffix);
    if (result) return 'audio';
    // 其他文件类型
    return 'unkown';
}