import React, {useState} from 'react';
import { Button } from 'antd';
import { PageHeaderWrapper } from '@ant-design/pro-layout';
import ProTable from '@ant-design/pro-table';
import { PlusOutlined } from '@ant-design/icons';
import { queryTreat } from './service';
import CreateForm from './components/CreateForm';


const columns = [
  {
    title: '#',
    dataIndex: 'id',
  },
  {
    title: '请客标题',
    dataIndex: 'title',
  },
  {
    title: '请客内容',
    dataIndex: 'content',
  },
  {
    title: '开始时间',
    dataIndex: 'start_time',
    valueType: 'dateTime',

  },
  {
    title: '结束时间',
    dataIndex: 'end_time',
    valueType: 'dateTime',

  }
]

const TableList = () => {
  const [createModalVisible, handleModalVisible] = useState(false);

  return (
    <PageHeaderWrapper>
      <ProTable
        headerTitle="请客列表"
        rowKey="id"
        request={ async () => {
          const res = await queryTreat()
          return {
            data: res.data.items,
            page: 1,
            success: true,
            total: 1
          }
        }}
        columns={columns}
        toolBarRender={() => [
          <Button key="3" type="primary" onClick={() => handleModalVisible(true)}>
            <PlusOutlined />
            新建
          </Button>,
        ]}
      />

      <CreateForm onCancel={() => handleModalVisible(false)} modalVisible={createModalVisible}>
        <ProTable
          onSubmit={async value => {
          }}
          rowKey="key"
          type="form"
          columns={columns}
          rowSelection={{}}
        />
      </CreateForm>

    </PageHeaderWrapper>
  )
}
export default TableList
