/**
 * 格式化树形菜单数据
 * @param data
 * @returns {*}
 */
export const formatterTreeSelect = (data, title, value) => {
  return data.map(m => {
    if (m.children) {
      m.children = formatterTreeSelect(m.children, title, value)
    }
    return {
      title: m[title],
      value: m[value],
      children: m.children || []
    }
  })
}

/**
 * 格式化树形控件数据
 * @param data
 * @returns {*}
 */
export const formatterTreeNode = (data, title, key) => {
  if (data && data.length) {
    return data.map(m => {
      if (m.children) {
        m.children = formatterTreeNode(m.children, title, key)
      }
      return {
        title: m[title],
        key: m[key],
        children: m.children || []
      }
    })
  }
}

export const formatterSelect = (data, value) => {
  if (data && data.length) {
    return data.map(m => ({ value: m[value] }))
  }
}
