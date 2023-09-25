import { useState, useEffect } from 'react';
import axios from 'axios';
import { Button, TextField } from '@mui/material';
import Modal from '@mui/material/Modal';
import Box from '@mui/material/Box';
import { Role } from 'components/role/row/RoleRow';
import { API_URL } from 'App';

const style = {
  position: 'absolute',
  top: '50%',
  left: '50%',
  transform: 'translate(-50%, -50%)',
  width: 400,
  bgcolor: 'background.paper',
  border: '2px solid #000',
  boxShadow: 24,
  p: 4,
  display: 'flex',
  flexDirection: 'column',
};

type UpdateRoleProps = {
  open: boolean;
  role: Role;
  handleClose: () => void;
  reloadList: () => void;
  onChange: (data: Role) => void;
};

export default function UpdateRoleModal({
  open,
  role,
  handleClose,
  reloadList,
  onChange,
}: UpdateRoleProps) {
  const [roleName, setRoleName] = useState<string>('');
  const [roleDesc, setRoleDesc] = useState<string>('');

  useEffect(() => {
    setRoleName(role.Name);
    setRoleDesc(role.Description);
  }, [role]);

  const updateRole = async (roleID: number, name: string, desc: string) => {
    const data = {
      Name: name,
      Description: desc,
    };
    alert('권한이 없습니다.');
    // const res = await axios.patch(`${API_URL}/rbac/role/${roleID}`, data);
    // onChange(res.data);
    reloadList();
  };

  const deleteRole = async (roleID: number) => {
    alert('권한이 없습니다.');
    // const res = await axios.delete(`${API_URL}/rbac/role/${roleID}`);
    // if (res !== undefined) {
    //   reloadList();
    // }
  };

  return (
    <Modal
      open={open}
      onClose={handleClose}
      aria-labelledby="modal-modal-title"
      aria-describedby="modal-modal-description"
    >
      <Box sx={style}>
        <TextField
          id="outlined-search"
          size="small"
          label="역할"
          style={{ margin: '20px' }}
          value={roleName}
          onChange={(e) => {
            setRoleName(e.target.value);
          }}
          inputProps={{
            autoComplete: 'off',
          }}
        />
        <TextField
          id="outlined-search"
          size="small"
          label="설명"
          style={{ margin: '20px' }}
          value={roleDesc}
          onChange={(e) => {
            setRoleDesc(e.target.value);
          }}
          inputProps={{
            autoComplete: 'off',
          }}
        />
        <div
          style={{
            display: 'flex',
            flexDirection: 'row',
            justifyContent: 'center',
          }}
        >
          <Button
            variant="outlined"
            style={{
              width: '50px',
              margin: '20px',
              color: 'black',
              borderBlockColor: 'black',
              backgroundColor: 'rgb(255, 0, 0, 0.3)',
            }}
            size="small"
            onClick={() => {
              deleteRole(role.ID);
              handleClose();
            }}
          >
            삭제
          </Button>
          <Button
            variant="outlined"
            style={{
              width: '50px',
              margin: '20px',
              color: 'black',
              borderBlockColor: 'black',
              backgroundColor: '#E0E0E0',
            }}
            size="small"
            onClick={() => {
              updateRole(role.ID, roleName, roleDesc);
              handleClose();
            }}
          >
            저장
          </Button>
        </div>
      </Box>
    </Modal>
  );
}
