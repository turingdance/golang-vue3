export const pure = (el: any, binding: any, vNode: any, pre: any) => {
  if (!el.hasAttribute("readonly")) {
    el.hasAttribute("readonly", "readonly")
  }
  setTimeout(() => {
    el.removeAttribute("readonly")
  }, 300)
}
