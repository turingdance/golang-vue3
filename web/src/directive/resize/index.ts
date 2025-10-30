export const resize = (el: any, binding: any, vNode: any, pre: any) => {
  let resizeable = false
  const minW = 10
  const minH = 10
  let val = ""
  let clientY = 0
  let clientX = 0
  clientY = clientX
  el.addEventListener("mousedown", resizeMove)
  document.addEventListener("mousemove", resizeDomMove)
  document.addEventListener("mouseup", resizeDomUp)
  function getAddress(ev: any) {
    const xPos1 = ev.clientX
    const yPos1 = ev.clientY
    const offset = 10
    let cls = ""
    //console.log(el.offsetHeight, el.getBoundingClientRect().top, el.offsetHeight + el.getBoundingClientRect().top)
    //console.log(yPos1)

    // if (yPos2 < offset) cls += 'n';
    if (yPos1 > el.offsetHeight + el.getBoundingClientRect().top - offset) cls += "s"
    // if (xPos2 < offset) cls += 'w';
    if (xPos1 > el.offsetWidth + el.getBoundingClientRect().left - offset) cls += "e"

    return cls
  }
  function resizeMove(e: any) {
    e.stopPropagation()
    const d = getAddress(e)
    if (d !== "") {
      resizeable = true
      val = d
      clientX = e.clientX
      clientY = e.clientY
    }
  }
  function resizeDomMove(e: any) {
    const d = getAddress(e)
    let cursor
    if (d === "") cursor = "default"
    else if (d === "se") {
      cursor = d + "-resize"
    }
    el.style.cursor = cursor
    if (resizeable) {
      if (val === "se") {
        // 等比例缩放
        el.style.width = Math.max(minW, el.offsetWidth + (e.clientY - clientY)) + "px"
        el.style.height = Math.max(minH, el.offsetHeight + (e.clientY - clientY) * 0.5625) + "px"
        clientX = e.clientX
        clientY = e.clientY
      }
    }
  }
  function resizeDomUp(e: any) {
    resizeable = false
    const { width, height } = el.getBoundingClientRect()
    binding.w = width
    binding.h = height
  }
}
