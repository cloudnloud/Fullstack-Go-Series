import React from 'react'
import { Modal, Form, Input } from 'antd';

function AddMemberModal(props) {
  const { isAddModalOpen, AddMember, closeAddModal } = props
  const [ addform ] = Form.useForm();
  return (
    <div>
      <Modal 
        title="Create Member" 
        open={isAddModalOpen} 
        onOk={() => {
          addform
          .validateFields()
          .then( async(values) => {
            await AddMember(values);
            addform.resetFields();
          }).catch((info) => {
            console.log("Validate Failed:", info);
          })
        }} 
        okText={"submit"}
        onCancel={closeAddModal}
      >
        <Form
            form={addform}
            name="basic"
            layout="vertical"
            style={{ padding: "20px"}}
            >
            <Form.Item
                label="Name"
                name="name"
                rules={[{ required: true, message: 'Please enter your name!' }]}
            >
                <Input />
            </Form.Item>

            <Form.Item
                label="Email"
                name="email"
                rules={[{ required: true, message: 'Please enter your email!' }]}
            >
                <Input />
            </Form.Item>

            <Form.Item
                label="City"
                name="city"
                rules={[{ required: true, message: 'Please enter your city!' }]}
            >
                <Input />
            </Form.Item>
            
            </Form>
        </Modal>
  </div>
  )
}

export default AddMemberModal