import React, {useEffect, useState} from 'react';
import ProTable from '@ant-design/pro-table';
import { connect } from 'dva';
import { Card, Button } from 'antd';
import { PlusOutlined } from '@ant-design/icons';
import { PageHeaderWrapper } from '@ant-design/pro-layout';
import { queryMenu } from '../service';
import CreateMenu from './components/CreateMenu';
import CreateForm from "@/pages/ListTableList/components/CreateForm";

const columns = [
  {
    title: '菜单编号',
    dataIndex: 'menuCode',
  },
  {
    title: '名称',
    dataIndex: 'menuName',
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
    title: '组件名',
    dataIndex: 'component',
  },
  {
    title: '权限',
    dataIndex: 'permission',
  }
]

const MenuIndex = props => {
  const [createModalVisible, handleModalVisible] = useState(false);

  return (
    <PageHeaderWrapper>
      <Card>
        <ProTable
          rowKey="menuCode"
          search={false}
          columns={columns}
          request={async params => {
            const res = await queryMenu(params)
            return {
              data: res.resultBody,
              page: params.current,
              success: true,
              total: res.resultBody.totalCount,
            }
          }}
          toolBarRender={() => [
            <Button key="3" type="primary" onClick={() => handleModalVisible(true)}>
              <PlusOutlined />
              新建
            </Button>,
          ]}
        >
        </ProTable>

      </Card>
      <CreateMenu onCancel={() => handleModalVisible(false)} modalVisible={createModalVisible} />
    </PageHeaderWrapper>
  )
}



export default connect(({ sysMenu }) => ({
  sysMenu,
}))(MenuIndex);

