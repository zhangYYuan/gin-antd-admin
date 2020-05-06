import React, {useState} from 'react';
import { PageHeaderWrapper } from '@ant-design/pro-layout';
import ProTable from '@ant-design/pro-table';
import { queryRole } from '../service';
import {Button, Divider} from "antd";
import {PlusOutlined} from "@ant-design/icons";
import CreateForm from "@/pages/ListTableList/components/CreateForm";

const columns = [
  {
    title: '角色编码',
    dataIndex: 'roleCode',
    rules: [
      {
        required: true,
        message: '请输入角色编码',
      },
    ],
  },
  {
    title: '角色名称',
    dataIndex: 'roleName',
    rules: [
      {
        required: true,
        message: '请输入角色名称',
      },
    ],
  },
  {
    title: '角色描述',
    dataIndex: 'roleDesc',
  },
  {
    title: '创建时间',
    dataIndex: 'createTime',
    valueType: 'date',
    hideInForm: true,
    hideInSearch: true
  },
  {
    title: '操作',
    dataIndex: 'option',
    valueType: 'option',
    render: (_, record) => (
      <>
        <a>编辑</a>
        <Divider type="vertical" />
        <a href="">删除</a>
        <Divider type="vertical" />
        <a href="">权限</a>
      </>
    ),
  },
]
const RoleList = () => {
  const [createModalVisible, handleModalVisible] = useState(false);

  return (
    <PageHeaderWrapper title={false}>
      <ProTable
        bordered={true}
        rowKey="id"
        request={ async params => {
          const res = await queryRole(params)
          return {
            data: res.resultBody.data,
            page: params.current,
            success: true,
            total: res.resultBody.totalCount,
          }
        }}
        columns={columns}
        toolBarRender={() => [
          <Button type="primary" onClick={() => handleModalVisible(true)}>
            <PlusOutlined /> 新建
          </Button>
        ]}
      />

      <CreateForm onCancel={() => handleModalVisible(false)} modalVisible={createModalVisible}>
        <ProTable
          onSubmit={async value => {
            const success = await handleAdd(value);
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
export default RoleList

