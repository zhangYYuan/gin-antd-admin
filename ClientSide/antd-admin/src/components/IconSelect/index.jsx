import React from 'react';
import {Modal, Tabs} from 'antd';
const { TabPane } = Tabs;
import icons from './icons';
import {
  HomeOutlined,
  SettingFilled,
  SmileOutlined,
  SyncOutlined,
  LoadingOutlined,
} from '@ant-design/icons';
const IconSelect = props => {
  const { modalVisible, onCancel } = props;

  const confirmAction = () => {

  }

  const onChange = () => {

  }

  return (
    <Modal
      destroyOnClose
      title="选择图标"
      visible={modalVisible}
      onCancel={() => onCancel()}
      onOk={() => confirmAction()}
    >
      <Tabs defaultActiveKey="1" onChange={onChange}>
        {
          icons.map(icon => (
            <TabPane tab={icon.title} key={icon.key}>
              <HomeOutlined />
              <SettingFilled />
              <SmileOutlined />
              <SyncOutlined spin />
              <SmileOutlined rotate={180} />
              <LoadingOutlined />
            </TabPane>
          ))
        }
      </Tabs>
    </Modal>
  )
}

export default IconSelect;
