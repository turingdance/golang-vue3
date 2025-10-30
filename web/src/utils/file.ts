import { lib, SHA256 } from "crypto-js"
import fileToArrayBuffer from "file-to-array-buffer"
export const base64ToArrayBuffer = (base64: string) => {
  const binary = Buffer.from(base64, "base64")
  const arrayBuffer = new ArrayBuffer(binary.length)
  const uint8 = new Uint8Array(arrayBuffer)
  uint8.forEach(function (value, i) {
    uint8[i] = binary.at(i) as number
  })
  return arrayBuffer
}
// 获取文件后缀
export function fileSuffix(filename) {
  const pos = filename.lastIndexOf(".")
  let suffix = ""
  if (pos !== -1) {
    suffix = filename.substring(pos)
  }
  return suffix
}

export function arrayBufferToWordArray(ab) {
  const i8a = new Uint8Array(ab)
  const a = [] as number[]
  for (let i = 0; i < i8a.length; i += 4) {
    a.push((i8a[i] << 24) | (i8a[i + 1] << 16) | (i8a[i + 2] << 8) | i8a[i + 3])
  }
  return lib.WordArray.create(a, i8a.length)
}

// 获取文件Hash
export async function fileHash(file) {
  const buffer = await fileToArrayBuffer(file)
  return SHA256(arrayBufferToWordArray(buffer)).toString()
}
// 判断是否是
export function shareUrl(key){
  
}