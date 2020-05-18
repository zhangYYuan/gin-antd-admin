import React from 'react';
import {Modal, Tabs} from 'antd';
const { TabPane } = Tabs;
import data from './icons';

import * as AntdIcons from '@ant-design/icons';



const IconSelect = props => {
  const { modalVisible, onCancel, confirmAction } = props;

  const onChange = () => {

  }

  return (
    <Modal
      destroyOnClose
      title="选择图标"
      visible={modalVisible}
      onCancel={() => onCancel()}
    >
      <Tabs defaultActiveKey="1" onChange={onChange}>
        {
          data.map(d => (
            <TabPane tab={d.title} key={d.title}>
              {
                d.icons.map((i) => {
                  let str = `${i}Outlined`
                  const component = AntdIcons[str]
                  if (component) {
                      return (
                        <span onClick={() => confirmAction(i)}>
                          React.createElement(component)
                        </span>
                      )
                  }
                })
              }
            </TabPane>
          ))
        }
      </Tabs>
    </Modal>
  )
}

export default IconSelect;
