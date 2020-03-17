import React, {useState} from 'react';
import { Button } from 'antd';
import { PageHeaderWrapper } from '@ant-design/pro-layout';
import ProTable from '@ant-design/pro-table';
import { PlusOutlined } from '@ant-design/icons';
import { queryTreatType } from './service';
import CreateForm from './components/CreateForm';


/**
 * CreatedAt: "2020-03-14T17:37:27+08:00"
 DeletedAt: null
 ID: 1
 UpdatedAt: "2020-03-14T17:37:27+08:00"
 describe: "美食王国"
 level: 3
 name: "美食系"
 region: 2
 status: 0
 * @type {({dataIndex: string, title: string}|{dataIndex: string, title: string}|{dataIndex: string, title: string}|{dataIndex: string, valueType: string, title: string}|{dataIndex: string, valueType: string, title: string})[]}
 */
const columns = [
  {
    title: '#',
    dataIndex: 'ID',
  },
  {
    title: '名称',
    dataIndex: 'name',
  },
  {
    title: '简介',
    dataIndex: 'describe',
  },
  {
    title: '开始时间',
    dataIndex: 'UpdatedAt',
    valueType: 'dateTime',

  },
  {
    title: '结束时间',
    dataIndex: 'CreatedAt',
    valueType: 'dateTime',

  }
]


const TableList = () => {
  const [createModalVisible, handleModalVisible] = useState(false);


  return (
      <PageHeaderWrapper>
        <ProTable
          headerTitle="请客列表"
          rowKey="ID"
          columns={columns}
          request={async () => {
            const res = await queryTreatType({ page: 1, pageSize: 10 })
            console.log(res.data.list)
            return {
              data: res.data.list,
              page: 1,
              success: true,
              total: 1
            }
          }}
          toolBarRender={() => [
            <Button key="3" type="primary" onClick={() => handleModalVisible(true)}>
              <PlusOutlined />
              新建
            </Button>,
          ]}
         />
      </PageHeaderWrapper>
  )
}
export default TableList
