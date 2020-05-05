/**
 * 格式化树形菜单数据
 * @param data
 * @returns {*}
 */
export const formatterTreeData = (data, title, value) => {
  return data.map(m => {
    if (m.children) {
      m.children = formatterTreeData(m.children, title, value)
    }
    return {
      title: m[title],
      value: m[value],
      children: m.children || []
    }
  })
}
