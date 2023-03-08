import { saveAs } from 'file-saver'

export const downloadFileByBlob = async (data: any, filename: string): Promise<any> => {
  let blob = new Blob([data], {
    type: 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet;charset=UTF-8'
  })
  saveAs(blob, filename)
}

// 不需要token ，可指定文件名
export function downloadFileByUrl(content, fileName) {
  const blob = new Blob([content]) //创建一个类文件对象：Blob对象表示一个不可变的、原始数据的类文件对象
  const url = window.URL.createObjectURL(blob) //URL.createObjectURL(object)表示生成一个File对象或Blob对象
  let dom = document.createElement('a') //设置一个隐藏的a标签，href为输出流，设置download
  dom.style.display = 'none'
  dom.href = url
  dom.setAttribute('download', fileName) //指示浏览器下载url,而不是导航到它；因此将提示用户将其保存为本地文件
  document.body.appendChild(dom)
  dom.click()
  dom.remove()
}

// 不需要token ，不需要指定文件名
export function downloadFileByUrlIframe(url) {
  try {
    let elemIF = document.createElement('iframe')
    elemIF.src = url
    elemIF.style.display = 'none'
    document.body.appendChild(elemIF)
  } catch (e) {
    console.log(e)
  }
}
