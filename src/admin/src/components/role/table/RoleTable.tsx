import { useState, useEffect } from 'react';
import axios from 'axios';
import Stack from '@mui/material/Stack';
import { styled } from '@mui/system';
import { Pagination, Button } from '@mui/material';
import { API_URL } from 'App';
import RoleRow, { Role } from 'components/role/row/RoleRow';
import AddRoleModal from 'components/role/modal/add/AddRoleModal';
import UpdateRoleModal from 'components/role/modal/update/UpdateRoleModal';

const grey = {
  50: '#F3F6F9',
  100: '#E7EBF0',
  200: '#E0E3E7',
  300: '#CDD2D7',
  400: '#B2BAC2',
  500: '#A0AAB4',
  600: '#6F7E8C',
  700: '#3E5060',
  800: '#2D3843',
  900: '#1A2027',
};

const Root = styled('div')(
  ({ theme }) => `
  table {
    font-family: IBM Plex Sans, sans-serif;
    font-size: 0.875rem;
    border-collapse: collapse;
    width: 100%;
  }

  td,
  th {
    border: 1px solid ${theme.palette.mode === 'dark' ? grey[800] : grey[200]};
    text-align: left;
    padding: 6px;
  }

  th {
    background-color: ${theme.palette.mode === 'dark' ? grey[900] : grey[100]};
  }
  `,
);

type RoleTableProps = {
  onChange: (data: Role) => void;
  role: Role;
};

interface RolePage {
  Results: Array<Role>;
  Count: number;
}

export default function RoleTable({ onChange, role }: RoleTableProps) {
  const [addOpen, setAddOpen] = useState<boolean>(false);
  const [updateOpen, setUpdateOpen] = useState<boolean>(false);

  const handleAddOpen = () => setAddOpen(true);
  const handleAddClose = () => setAddOpen(false);

  const handleUpdateOpen = () => setUpdateOpen(true);
  const handleUpdateClose = () => setUpdateOpen(false);

  const [selectedRole, setSelectedRole] = useState<Role>(role);

  const rowSize = 5;
  const [rolePage, setRolePage] = useState<RolePage>();
  const [page, setPage] = useState<number>(1);

  useEffect(() => {
    if (rolePage !== undefined) {
      setPage(parseInt(`${(rolePage.Count + 1) / rowSize}`, 10));
    }
  }, [rolePage]);

  const handleRole = (role: Role) => {
    setSelectedRole(role);
    onChange(role);
  };

  const getPage = async (page: number) => {
    await axios
      .get(`${API_URL}/rbac/role/list?page=${page}&pageSize=${rowSize}`)
      .then((res) => setRolePage(res.data));
  };

  useEffect(() => {
    getPage(1);
    setSelectedRole(role);
  }, []);

  const handleChangePageNum = (
    event: React.ChangeEvent<unknown>,
    value: number,
  ) => {
    getPage(value);
  };

  const reloadList = () => {
    getPage(1);
  };

  return (
    <Root sx={{ width: 350, maxWidth: '100%' }}>
      <h1>Role</h1>
      <div style={{ minHeight: '435px' }}>
        <div style={{ minHeight: '400px' }}>
          <table aria-label="custom pagination table">
            <thead>
              <tr>
                <th style={{ width: '110px' }}>역할</th>
                <th>설명</th>
              </tr>
            </thead>
            <tbody>
              {rolePage &&
                rolePage.Results.map((rowRole: Role) => (
                  <RoleRow
                    key={rowRole.ID}
                    role={rowRole}
                    isChecked={rowRole === selectedRole}
                    onChange={handleRole}
                  />
                ))}
            </tbody>
          </table>
        </div>
        <div style={{ alignSelf: 'center' }}>
          <Button variant="outlined" size="small" onClick={handleAddOpen}>
            추가하기
          </Button>
          <Button variant="outlined" size="small" onClick={handleUpdateOpen}>
            수정
          </Button>
        </div>
      </div>
      {rolePage && (
        <Stack spacing={2}>
          <Pagination
            sx={{ margin: 'auto', marginTop: '10px' }}
            count={page}
            defaultPage={1}
            onChange={handleChangePageNum}
            shape="rounded"
          />
        </Stack>
      )}
      <AddRoleModal
        open={addOpen}
        reloadList={reloadList}
        handleClose={handleAddClose}
        onChange={onChange}
      />
      {selectedRole && (
        <UpdateRoleModal
          open={updateOpen}
          role={selectedRole}
          reloadList={reloadList}
          handleClose={handleUpdateClose}
          onChange={onChange}
        />
      )}
    </Root>
  );
}
