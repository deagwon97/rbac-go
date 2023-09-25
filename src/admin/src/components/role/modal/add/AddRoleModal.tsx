import { useState } from 'react';
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

type AddRoleProps = {
  open: boolean;
  handleClose: () => void;
  reloadList: () => void;
  onChange: (data: Role) => void;
};

export default function AddRoleModal({
  open,
  handleClose,
  reloadList,
  onChange,
}: AddRoleProps) {
  const [roleName, setRoleName] = useState<string>('');
  const [roleDesc, setRoleDesc] = useState<string>('');

  const addRole = async (name: string, desc: string) => {
    const data = {
      Name: name,
      Description: desc,
    };
    if (name !== '') {
      alert('권한이 없습니다.');
      // const res = await axios.post(`${API_URL}/rbac/role`, data);
      // onChange(res.data);
      reloadList();
    }
    setRoleName('');
    setRoleDesc('');
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
            }}
            size="small"
            onClick={() => {
              addRole(roleName, roleDesc);
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
