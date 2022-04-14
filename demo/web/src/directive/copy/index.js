import { Message } from 'element-ui'
const Clipboard = require('clipboard')
const key = Symbol()

function createCopyButton() {
  // const btn = document.createElement('button')
  // btn.setAttribute('type', 'button')
  // btn.style.marginRight = '5px'
  // btn.className = 'el-button el-tooltip el-button--text el-button--small'
  // btn.innerHTML = `<i class="el-icon-document-copy"></i>`
  const btn = document.createElement('i')
  btn.className = 'el-icon-document-copy link-type'
  btn.style.marginRight = '3px'
  return btn
}

function createClipboard(btn, text) {
  const clipboard = new Clipboard(btn, {
    text: function(trigger) {
      return text
    }
  })
  clipboard.on('success', function(e) {
    Message.success(text + ' 复制成功')
  })
  clipboard.on('error', function(e) {
    Message.error('复制失败')
  })
  return clipboard
}

export default {
  install(Vue) {
    Vue.directive('copy', {
      inserted(el, binding) {
        const { modifiers: { debug }} = binding
        debug && console.log('v-copy inserted')
        const text = el.innerText.trim()
        if (text) {
          const btn = createCopyButton()
          el.parentNode.insertBefore(btn, el)
          el[key] = createClipboard(btn, text)
        }
      },
      update(el, binding) {
        const text = el.innerText.trim()
        if (text) {
          if (!el[key]) {
            const btn = createCopyButton()
            el.parentNode.insertBefore(btn, el)
            el[key] = createClipboard(btn, text)
          }
        }
      },
      unbind(el, binding) {
        const { modifiers: { debug }} = binding
        debug && console.log('v-copy unbind')
        const clipboard = el[key]
        clipboard && clipboard.destroy()
      }
    })
  }
}
