import React from 'react';
import {Form, Modal} from 'antd';
// import AddressForm from '@/components/AForm/AddressForm'
import RightContent from '@/components/GlobalHeader/RightContent';


const CreateForm = props => {
  const { modalVisible, onCancel } = props;
  return (
    <Modal
      destroyOnClose
      title="新建门店"
      visible={modalVisible}
      onCancel={() => onCancel()}
      footer={null}
    >
      <Form>
        <Form.Item
          name="address"
          label="地址"
        >
          <RightContent />
        </Form.Item>
      </Form>
    </Modal>
  );
};

export default CreateForm;
