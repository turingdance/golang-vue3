import eventbus from "./eventbus"
const wsUrl = import.meta.env.VITE_WS_API // websocket 默认连接地址
let websocket: any // 用于存储实例化后websocket
let isConnect = false // 连接标识 避免重复连接
let rec: any // 断线重连后，延迟5秒重新创建WebSocket连接  rec用来存储延迟请求的代码
let globaldata: Record<string, any> = {} as Record<string, any>
// 创建websocket
function buildurl() {
  const keys = Object.keys(globaldata)
  const tmp = [] as string[]
  for (let i = 0; i < keys.length; i++) {
    const k = keys[i]
    tmp.push(k + "=" + globaldata[k])
  }
  return [wsUrl, tmp.join("&")].join("?")
}
function start(query: any | undefined) {
  globaldata = query || {}
  try {
    initWebSocket() // 初始化websocket连接
  } catch (e) {
    reConnect() // 如果无法连接上 webSocket 那么重新连接！可能会因为服务器重新部署，或者短暂断网等导致无法创建连接
  }
}

// 初始化websocket
function initWebSocket() {
  websocket = new WebSocket(buildurl())
 
  websocket.onopen = function (e: any) {
    websocketOpen(e)
  }

  // 接收
  websocket.onmessage = function (e: any) {
    websocketonmessage(e)
  }

  // 连接发生错误
  websocket.onerror = function (res: any) {
    eventbus.emit("ws.erroe", res)
    console.error("WebSocket连接发生错误", res)
    isConnect = false // 连接断开修改标识
    reConnect() // 连接错误 需要重连
  }

  websocket.onclose = function (e: any) {
    websocketclose(e)
  }
}

// 定义重连函数
const reConnect = () => {
  if (isConnect) return // 如果已经连上就不在重连了
  rec && clearTimeout(rec)
  rec = setTimeout(function () {
    // 延迟5秒重连  避免过多次过频繁请求重连
    start(buildurl())
  }, 5000)
}

// 创建连接
function websocketOpen(e: any) {

  eventbus.emit("ws.open")
}
// 数据接收
function websocketonmessage(e: any) {
  eventbus.emit("ws.message", e.data)
  // let data = JSON.parse(decodeUnicode(e.data))
}
// 关闭
function websocketclose(e: any) {
  isConnect = false // 断开后修改标识
  //console.log("connection closed (" + e.code + ")")
}

// 数据发送
function websocketsend(data: any) {
  //console.log("发送的数据", data, JSON.stringify(data))
  websocket.send(JSON.stringify(data))
}

// 实际调用的方法==============

// 发送
function publish(data: any) {
  if (websocket.readyState === websocket.OPEN) {
    // 开启状态
    websocketsend(data)
  } else {
    // 若 未开启 / 正在开启 状态 ，则等待1s后重新调用
    setTimeout(function () {
      publish(data)
    }, 1000)
  }
}

// 关闭
const stop = () => {
  websocket.close()
}

export default {
  publish,
  start,
  stop
}
