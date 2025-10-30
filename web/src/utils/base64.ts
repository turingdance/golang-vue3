import { customAlphabet } from "nanoid"
const nanoid = customAlphabet("0123456789abcdefghijklmnopqrstuvwxyz", 16)
export const base64ToArrayBuffer = (base64: string) => {
  let b64 = base64
  if (base64.startsWith("data:")) {
    const darr = base64.split(/[\,\;\:]/)
    b64 = darr[2]
  }
  const binary = Buffer.from(b64, "base64")
  const arrayBuffer = new ArrayBuffer(binary.length)
  const uint8 = new Uint8Array(arrayBuffer)
  uint8.forEach(function (value, i) {
    uint8[i] = binary.at(i) as number
  })
  return arrayBuffer
}

export function base64ToFile(base64pure: string, type = "image/jpeg", fileName = ""): File {
  let b64 = base64pure
  let darr = [] as string[]
  if (base64pure.startsWith("data:")) {
    darr = base64pure.split(/[\,\;\:]/)
    type = darr[1]
    b64 = darr[3]
  }
  const byteString = atob(b64)
  const options = {
    type: type,
    endings: "native"
  }
  if (fileName === "") {
    fileName = nanoid()
  }
  //console.log("file", darr, fileName)
  const suffix = "." + type.split("/")[1]
  const u8Arr = new Uint8Array(byteString.length)
  for (let i = 0; i < byteString.length; i++) {
    u8Arr[i] = byteString.charCodeAt(i)
  }
  const fileOfBlob = new File([u8Arr], fileName + suffix, options as FilePropertyBag) //返回文件流
  return fileOfBlob
}
