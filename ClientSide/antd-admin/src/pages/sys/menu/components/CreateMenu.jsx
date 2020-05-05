import React, {useEffect, useState} from 'react';
import { Button, Space, Form, Modal, Radio, Input, InputNumber, TreeSelect } from 'antd';
import { EditOutlined } from "@ant-design/icons";
import IconSelect from '@/components/IconSelect';
import { formatterTreeData} from '@/utils';
import { connect } from 'dva';
import { addMenu } from '../../service';

const formItemLayout = {
  labelCol: {span: 6},
  wrapperCol: {span: 18},
};

const CreateMenu = props => {
  const [form] = Form.useForm();
  // 控制选择图标
  const [createModalVisible, handleModalVisible] = useState(false);
  const { modalVisible, onCancel } = props;
  const [treeData, setTreeData] = useState([])
  const {
    dispatch,
    sysMenu: { menuList },
  } = props;

  useEffect(() => {
    dispatch({
      type: 'sysMenu/fetch'
    });
  }, []);

  useEffect(() => {
    const tree =  formatterTreeData(menuList, 'menuName', 'menuCode')
    setTreeData(tree)
  }, [menuList]);

  const confirmAction = () => {
    form.validateFields().then(values => {
      addMenu(values).then(res => {
      })
    })
  }

  const onChange = () => {

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
        initialValues={{
          type: 1,
          order: 1
        }}
      >
        <Form.Item name='type' label="菜单类型">
          <Radio.Group>
            <Radio value={1}>一级菜单</Radio>
            <Radio value={2}>子菜单</Radio>
            <Radio value={3}>按钮/权限</Radio>
          </Radio.Group>
        </Form.Item>

        <Form.Item
          noStyle
          shouldUpdate={(prevValues, currentValues) => prevValues.type !== currentValues.type}
        >
          {({ getFieldValue}) => {
            return getFieldValue('type') === 2 ? (
              <Form.Item
                label="上级菜单"
                name='parentMenu'
                rules={[{required: true, message: '请选择上级菜单!'}]}
              >
                <TreeSelect
                  style={{ width: '100%' }}
                  dropdownStyle={{ maxHeight: 400, overflow: 'auto' }}
                  treeData={treeData}
                  placeholder="请选择上级菜单"
                  onChange={onChange}
                />
              </Form.Item>
            ) : null;
          }}
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
          <Space>
            <Input />
            <Button onClick={() => handleModalVisible(true)}>
              <EditOutlined />
              选择图标
            </Button>
          </Space>
        </Form.Item>

        <Form.Item
          name="order"
          label="排序"
        >
          <InputNumber min={1} max={10} />
        </Form.Item>
      </Form>
      <IconSelect onCancel={() => handleModalVisible(false)} modalVisible={createModalVisible} />
    </Modal>
  );
};

// export default CreateMenu;
export default connect(({ sysMenu }) => ({
  sysMenu,
}))(CreateMenu);
