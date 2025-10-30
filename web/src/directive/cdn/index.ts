export const cdn = (el: Element | HTMLElement, binding: any, vNode: any, pre: any) => {
  const tag = el.tagName.toLowerCase()
  const cdnhost = import.meta.env.VITE_CND_HOST
  const dsturl = cdnhost + "/" + binding.value
  if (tag == "img") {
    binding.value && el.setAttribute("src", dsturl)
  }
}
