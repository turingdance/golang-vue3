import desensitizeApi from "@/utils/desensitize";
/**
 * 按钮权限
 */
export const string= {
  mounted(el: HTMLElement, binding: any) {
      let {value } = binding
      let {data,start,end,mask } = value
      el.innerHTML = desensitizeApi.string(data,start,end,mask)
  },
};

export const email= {
  mounted(el: HTMLElement, binding: any) {
      let {value } = binding
      el.innerHTML = desensitizeApi.email(value)
  },
};

export const phone= {
  mounted(el: HTMLElement, binding: any) {
      let {value } = binding
      el.innerHTML = desensitizeApi.phone(value)
  },
};
