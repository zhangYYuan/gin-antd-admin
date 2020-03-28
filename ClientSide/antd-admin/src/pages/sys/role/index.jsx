import React from 'react';
import { PageHeaderWrapper } from '@ant-design/pro-layout';
import ProTable from '@ant-design/pro-table';
import { queryRole } from '../service';

const columns = [
  {
    title: '角色编码',
    dataIndex: 'roleCode',
  },
  {
    title: '角色名称',
    dataIndex: 'roleName',
  },
  {
    title: '角色描述',
    dataIndex: 'roleDesc',
  },
  {
    title: '创建时间',
    dataIndex: 'createTime',
  }
]
const RoleList = () => (
    <PageHeaderWrapper content=" 这个页面只有 admin 权限才能查看">
      <ProTable
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
        rowSelection={{}}
      />
    </PageHeaderWrapper>
  )
export default RoleList

