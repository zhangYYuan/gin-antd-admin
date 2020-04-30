import React, {useEffect, useState} from 'react';
import { connect } from 'dva';
import { Table, Card, Button } from 'antd';
import { PlusOutlined } from '@ant-design/icons';
import { PageHeaderWrapper } from '@ant-design/pro-layout';

import CreateForm from './components/CreateForm'

const columns = [
  {
    title: '菜单编号',
    dataIndex: 'menuId',
  },
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

const MenuIndex = props => {
  const {
    dispatch,
    sysMenu: { menuList },
  } = props;

  useEffect(() => {
    console.log(props)
    dispatch({
      type: 'sysMenu/fetch',
      payload: {
        count: 5,
      },
    });
  }, []);

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



export default connect(({ sysMenu }) => ({
  sysMenu,
}))(MenuIndex);

