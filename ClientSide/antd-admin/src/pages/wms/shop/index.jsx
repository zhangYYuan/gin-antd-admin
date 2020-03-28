import React, {useState} from 'react';
import {Button} from 'antd';
import ProTable from '@ant-design/pro-table'
import {PlusOutlined} from '@ant-design/icons';
import {queryStoreList} from '../service'
import CreateForm from './component/CreateForm'

const columns = [
  {
    title: '门店编号',
    dataIndex: 'warehouseId',
  },
  {
    title: '门店名称',
    dataIndex: 'storeName',
  },
  {
    title: '详细地址',
    dataIndex: 'address',
  },
  {
    title: '店长姓名',
    dataIndex: 'directorName',
  },
  {
    title: '店长电话',
    dataIndex: 'directorPhone',
  },
  {
    title: '门店状态',
    dataIndex: 'enabledStatus',
  }
]

const ShopList = () => {
  const [create, setCreate] = useState()
  return (
    <div>
      <ProTable
        headerTitle="门店管理"
        rowKey="id"
        request={async params => {
          const res = await queryStoreList(params)
          return {
            data: res.resultBody.data,
            page: params.current,
            success: true,
            total: res.resultBody.totalCount,
          }
        }}
        columns={columns}
        toolBarRender={() => [
          <Button type="primary" onClick={() => setCreate(true)}>
            <PlusOutlined/>
            新建
          </Button>
        ]}
      />
      <CreateForm modalVisible={create}/>
    </div>
  )
}
export default ShopList
