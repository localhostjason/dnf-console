import { ElMessage } from 'element-plus'

// 消息
const Message = (message: string, duration: number = 3000): any => {
  return ElMessage({
    showClose: true,
    message,
    duration
  })
}

// 成功
const successMessage = (message: string,duration: number = 3000): any => {
  return ElMessage({
    showClose: true,
    message,
    type: 'success',
    duration
  })
}

// 警告
const warnMessage = (message: string,duration: number = 3000): any => {
  return ElMessage({
    showClose: true,
    message,
    type: 'warning',
    duration
  })
}

// 失败
const errorMessage = (message: string,duration: number = 3000): any => {
  return ElMessage({
    showClose: true,
    message,
    type: 'error',
    duration
  })
}

export { Message, successMessage, warnMessage, errorMessage }
