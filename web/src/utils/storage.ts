const get = <T>(key: string): T => {
  const resp = sessionStorage.getItem(key)
  if (!resp) return {} as T
  else {
    const r = JSON.parse(resp)
    if (r.data) {
      return r.data as T
    } else {
      return {} as T
    }
  }
}

const set = <T>(key: string, data: T): void => {
  sessionStorage.setItem(
    key,
    JSON.stringify({
      type: typeof data,
      data: data
    })
  )
}
const clear = <T>(key: string): void => {
  sessionStorage.removeItem(key)
}

export default {
  get,
  set,
  clear
}
