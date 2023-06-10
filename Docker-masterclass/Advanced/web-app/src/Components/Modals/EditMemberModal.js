import React, {useEffect} from 'react'
import { Modal, Form, Input } from 'antd';

function EditMemberModal(props) {
  const { isUpdateModalOpen , UpdateMember, closeEditModal, selectedMember} = props
  const [ editform ] = Form.useForm();

  useEffect( () => {
      editform.setFieldValue('name', selectedMember?.name)
      editform.setFieldValue('email', selectedMember?.email)
      editform.setFieldValue('city', selectedMember?.city)
  },[selectedMember])

  return (
    <div>
        <Modal 
        title="Update Member" 
        open={isUpdateModalOpen} 
        onOk={() => {
            editform
            .validateFields()
            .then( async(values) => {
                await UpdateMember(values);
            }).catch((info) => {
                console.log("Validate Failed:", info);
            })
        }} 
        okText={"Update"}
        onCancel={closeEditModal}>
        <Form
            form={editform}
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

export default EditMemberModal