import React from 'react'
import { Button, Space, Table, Modal, Form, Input, notification} from 'antd';
import axios from "axios"
import { useEffect, useState } from 'react';

import AddMemberModal from '../../Components/Modals/AddMemberModal';
import EditMemberModal from '../../Components/Modals/EditMemberModal';

import { getAllMembersApi, createMemberApi, updateMemberApi, deleteMemberApi } from '../../Services/Members';


function Dashboard() {
  const [membersList, setMembersList] = useState([])
  const [isAddModalOpen, setIsAddModalOpen] = useState(false)
  const [isUpdateModalOpen, setIsUpdateModalOpen] = useState(false)
  const [selectedMember, setSelectedMember] = useState(null);

  useEffect(() => {
    getMembers()
  },[])

  const columns = [
    {
      title: 'Name',
      dataIndex: 'name',
      key: 'name',
      render: (text) => <a>{text}</a>,
    },
    {
      title: 'Email',
      dataIndex: 'email',
      key: 'email',
    },
    {
      title: 'City',
      dataIndex: 'city',
      key: 'city',
    },
    {
      title: 'Action',
      key: 'action',
      render: (_, record) => (
        <Space size="middle">
          <a onClick={() => selectMember(record)}>Edit </a>
          <a onClick={() => DeleteMember(record)}>Delete</a>
        </Space>
      ),
    },
  ];

  const selectMember = (member) => {
    setSelectedMember(member);
    setIsUpdateModalOpen(true)
  }

  const getMembers = async() =>{
    try {
      const response = await getAllMembersApi();
      const { data }  = response?.data?.data
      if(data?.length >0){
        setMembersList(data)
      }else{
        console.log("No members found");
      }
    } catch (error) {
      console.log(error)
    } 
  }

  const AddMember = async (values) => {
    try {
        const response = await createMemberApi(values);
        const { data , message, status } = response?.data
        if(status ==  201){
          await getMembers();
          setIsAddModalOpen(false);
          notification.success({message : message})
        }else{
          notification.warning({message : message})
        }
    } catch (error) {
      console.log(error)
      notification.warning({message : error})
    }
    
  }


  const UpdateMember = async(values) => {
    try {
      const response = await updateMemberApi(selectedMember?.id, values);
      const { data , message, status } = response?.data
      if(status ===  200){
          await getMembers();
          setIsUpdateModalOpen(false);
          notification.success({message : message})
      }else{
          notification.warning({message : message})
      }
    } catch (error) {
      console.log("Err => ", error)
    }
  }

  const DeleteMember = async (member) => {
    try {
      const response = await deleteMemberApi(member?.id);
      console.log("response => ", response);
      const { data , message, status } = response?.data
      if(status ==  200){
        await getMembers();
        notification.success({message : message})
      }else{
        notification.warning({message : message})
      }
    } catch (error) {
      console.log(error)
    } 
  }

  const closeAddModal = () => {
    setIsAddModalOpen(false)
  }

  const closeEditModal = () => {
    setIsUpdateModalOpen(false)
  }

  return (
    <div>
        <div style={{ textAlign: "center"}}>
            <h2>Cloudnloud Team Members</h2>
        </div>
        <Button type='primary' style={{ float: "right", marginBottom: "10px"}} onClick={() => setIsAddModalOpen(true)}>Add Member</Button>
        <Table columns={columns} dataSource={membersList} />

        <AddMemberModal
            isAddModalOpen={isAddModalOpen}
            AddMember={AddMember}
            closeAddModal={closeAddModal}
        />

        <EditMemberModal
            isUpdateModalOpen={isUpdateModalOpen}
            UpdateMember={UpdateMember}
            closeEditModal={closeEditModal}
            selectedMember={selectedMember}
        />
    </div>
  )
}

export default Dashboard