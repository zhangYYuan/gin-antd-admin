import React, { useState } from 'react';
  import { PlusOutlined } from '@ant-design/icons';
import { PageHeaderWrapper } from '@ant-design/pro-layout';
import { Table, Card, Button, Typography, Alert } from 'antd';

import CreateForm from './components/CreateForm'
import { addMenu } from '../service';

const columns = [
  {
    title: '名称',
    dataIndex: 'name',
  },
  {
    title: '图标',
    dataIndex: 'icon',
  },
  {
    title: '菜单路径',
    dataIndex: 'path',
  },
  {
    title: '组件路径',
    dataIndex: 'component',
  }
]


const MenuIndex = () => {
  const [createModalVisible, handleModalVisible] = useState(false);

  return (
  <PageHeaderWrapper>
    <Card>
      <div className='mb16'>
        <Button type="primary" icon={<PlusOutlined />} onClick={() => handleModalVisible(true)}>新建</Button>
      </div>
      <Table columns={columns} />,
    </Card>


     <CreateForm modalVisible={createModalVisible} />
  </PageHeaderWrapper>
  )
}

export default MenuIndex
