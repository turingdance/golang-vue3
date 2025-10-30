import xlsxUtil from "@/utils/xlsx"
interface Input {
  table?: string | Element | HTMLElement | Object
  data?: Array<any>
  provider?: Function
  title?: string
  header: Array<any>
}
const handleClick = (input: Input) => {
  if (typeof input.table == "string") {
    xlsxUtil.exportFromDom(document.querySelector(input.table), input.title)
  } else if (typeof input.table == "object") {
    xlsxUtil.exportFromDom(input.table, input.title)
  } else if (typeof input.data == "object") {
    xlsxUtil.exportFromData(input.data, input.header, input.title)
  } else if (typeof input.provider == "function") {
    xlsxUtil.exportFromServer(input.provider, input.header, input.title)
  }
}
export const xlsx = (el: Element | HTMLElement, binding: any, vNode: any, pre: any) => {
  //{title:"","data":{table:"",function(){},data:(){ rows:[],"header":{} }}}
  // {title:"",table:"",data:[],"header":{},"provider":""}}}

  const input = binding.value
  const _handleClick = (e) => {
    handleClick(input)
  }
  if (!el.hasAttribute("xlsx")) {
    el.setAttribute("xlsx", "on")
    el.addEventListener("click", _handleClick)
  }
}
