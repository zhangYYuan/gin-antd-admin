import React, { useEffect, useState } from 'react';
import { Button, Form, Input, InputNumber, Select, Tag } from 'antd';
import { connect } from "umi";
import { addMenu } from '../../service';

const formItemLayout = {
  labelCol: {span: 4 },
  wrapperCol: {span: 18},
};

const tailLayout = {
  wrapperCol: { offset: 8, span: 16 },
};

function tagRender(props) {
  const { label, value, closable, onClose } = props;

  return (
    <Tag color='green' closable={closable} onClose={onClose} style={{ marginRight: 3 }}>
      {label}
    </Tag>
  );
}

const CreateDepart = props => {
  const [form] = Form.useForm();
  console.log(props)
  const {
    dispatch,
    sysRole: { roleList },
  } = props;

  useEffect(() => {
    dispatch({
      type: 'sysRole/fetch'
    });
  }, []);

  const confirmAction = () => {
    form.validateFields().then(values => {
      addMenu(values).then(res => {
      })
    })
  }

  const onReset = () => {

  }

  return (
    <Form
      form={form}
      {...formItemLayout}
      initialValues={{
        parent: '无'
      }}
    >
      <Form.Item
        label="上级部门"
        name='parent'
      >
        <Input disabled />
      </Form.Item>

      <Form.Item
        label="部门名称"
        name='departName'
        rules={[{required: true, message: '请输入部门名称!'}]}
      >
        <Input/>
      </Form.Item>

      <Form.Item
        label="部门负责人"
        name='deptLeader'
        rules={[{required: true, message: '请输入部门负责人!'}]}
      >
        <Select
          mode="multiple"
          tagRender={tagRender}
          style={{ width: '100%' }}
          options={roleList}
        />
      </Form.Item>

      <Form.Item
        name="order"
        label="排序"
      >
        <InputNumber min={1} max={10} />
      </Form.Item>

      <Form.Item
        {...tailLayout}
        style={{ marginTop: 32 }}
      >
        <Button style={{ marginRight: 8 }} htmlType="button" onClick={() => onReset()}>
          重置
        </Button>
        <Button type="primary" htmlType="submit">
          保存并提交
        </Button>
      </Form.Item>
    </Form>
  );
};

export default connect(({ sysRole }) => ({
  sysRole
}))(CreateDepart);
