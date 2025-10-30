//导入依赖项
import FileSaver from "file-saver"
import * as XLSX from "xlsx"
//导入uuid.js工具
import { nanoid } from "nanoid"
const tableId = () => {
  return nanoid(15)
}
export interface ITableExportHeader {
  label?: string
  field: string
  format?: Function
  width?: number
}
export const exportFromDom = (domtable, title = tableId()) => {
  //根据table生成Book工作簿
  const wb = XLSX.utils.table_to_book(domtable)
  //将Book工作簿作为输出
  const wbout = XLSX.write(wb, {
    bookType: "xlsx",
    bookSST: true,
    type: "array"
  })
  //尝试将当前table内容保存为excel文件
  try {
    FileSaver.saveAs(
      //被导出的blob二进制对象
      new Blob([wbout], { type: "application/octet-stream" }),
      //导出文件的名称+后缀名
      title + ".xlsx"
    )
  } catch (e) {
    if (typeof console != "undefined") {
      console.log(e, wbout)
    }
  }
}
export const format = {
  txt: (input: any): any => {
    return "" + input
  }
}
export const exportFromData = (rows, header: ITableExportHeader[] | undefined, title = tableId()) => {
  /* generate worksheet and workbook */
  //
  const resultArr = [] as any[]
  const titleArr = header?.map((o) => o.field)
  const labelArr = header?.map((o) => o.label)
  const headerMap = {} as Record<string, ITableExportHeader>
  header?.forEach((hh) => {
    headerMap[`${hh.field}`] = hh
  })
  rows.forEach((element) => {
    const tmp = new Object()
    titleArr?.forEach((tt) => {
      const fmt = headerMap[`${tt}`]
      tmp[`${tt}`] = element[`${tt}`]
      if (fmt.format) {
        tmp[`${tt}`] = fmt.format(element[`${tt}`])
      }
    })
    resultArr.push(tmp)
  })
  const worksheet = XLSX.utils.json_to_sheet(resultArr)
  const workbook = XLSX.utils.book_new()
  XLSX.utils.book_append_sheet(workbook, worksheet, title)
  //   workbook.SheetNames.push(title);
  //   workbook.Sheets[`${title}`] = worksheet;
  //console.log(labelArr, resultArr)

  /* fix headers */
  if (typeof labelArr == "object") {
    XLSX.utils.sheet_add_aoa(worksheet, [labelArr], { origin: "A1" })
  }

  /* calculate column width */
  // const max_width = rows.reduce((w, r) => Math.max(w, r.name.length), 10);
  //[ { wch: 12 },{ wch: 150 }, ];
  worksheet["!cols"] = header?.map((hh) => {
    return {
      wch: hh.width || 10
    }
  })
  /* create an XLSX file and try to save to Presidents.xlsx */
  const defaultCellStyle = {
    font: { name: "Verdana", sz: 13, color: "FF00FF88" },
    fill: { fgColor: { rgb: "FFFFAA00" } }
  }

  const wt = {
    bookType: "xlsx",
    bookSST: false,
    type: "array",
    cellStyles: true,
    defaultCellStyle: defaultCellStyle,
    showGridLines: false
  }
  const wbout = XLSX.write(workbook, wt as any)
  try {
    FileSaver.saveAs(
      //被导出的blob二进制对象
      new Blob([wbout], { type: "application/octet-stream" }),
      //导出文件的名称+后缀名
      title + ".xlsx"
    )
  } catch (e) {
    if (typeof console != "undefined") {
      console.log(e, wbout)
    }
  }
}

export const exportFromServer = async (resolve, headerarr: ITableExportHeader[] | undefined, title = tableId()) => {
  /* generate worksheet and workbook */

  resolve().then((res) => {
    return exportFromData(res.rows, headerarr, title)
  })
}

export default {
  exportFromDom,
  exportFromData,
  exportFromServer,
  format
}
