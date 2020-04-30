import React from 'react';
import { Form, Modal, Radio, Input, InputNumber} from 'antd';
import { addMenu } from '../../service'

const formItemLayout = {
  labelCol: {span: 6},
  wrapperCol: {span: 14},
};

const CreateForm = props => {
  const {modalVisible, onCancel} = props;
  const [form] = Form.useForm();
  const confirmAction = () => {
    form.validateFields().then(values => {
      addMenu(values).then(res => {
      })
    })
  }

  return (
    <Modal
      destroyOnClose
      title="新建菜单"
      visible={modalVisible}
      onCancel={() => onCancel()}
      onOk={() => confirmAction()}
    >
      <Form
        form={form}
        {...formItemLayout}
        initialValues={{type: 1}}
      >
        <Form.Item name='type' label="菜单类型">
          <Radio.Group>
            <Radio value={1}>一级菜单</Radio>
            <Radio value={2}>子菜单</Radio>
            <Radio value={3}>按钮/权限</Radio>
          </Radio.Group>


        </Form.Item>

        <Form.Item
          label="菜单名称"
          name='menuName'
          rules={[{required: true, message: '请输入菜单名!'}]}
        >
          <Input/>
        </Form.Item>

        <Form.Item
          label="菜单路径"
          name='path'
          rules={[{required: true, message: '请输入菜单路径!'}]}
        >
          <Input/>
        </Form.Item>

        <Form.Item
          label="前端组件"
          name="component"
          rules={[{required: true, message: '请输入前端组件!'}]}
        >
          <Input/>
        </Form.Item>

        <Form.Item
          label="菜单图标"
          name="icon"
        >
          <Input/>
        </Form.Item>

        <Form.Item
          name="order"
          label="排序"
        >
          <InputNumber min={1} max={10} defaultValue={1}/>
        </Form.Item>
      </Form>
    </Modal>
  );
};

export default CreateForm;
