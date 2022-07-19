import { useState, useEffect } from 'react';
import axios from 'axios';
import Stack from '@mui/material/Stack';
import { styled } from '@mui/system';
import { Pagination } from '@mui/material';
import { Role } from 'components/role/row/RoleRow';
import PermissionRow from 'components/permission/row/PermissionRow';
import { API_URL } from 'App';

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

interface Permission {
  IsAllowed: boolean;
  ID: number;
  ServiceName: string;
  Name: string;
  Action: string;
  Object: string;
}

interface PermissionTableProps {
  role: Role;
}

export default function PermissionTable({ role }: PermissionTableProps) {
  const [count, setCount] = useState<number>(0);
  const [permissionArray, setPermissionArray] = useState<Permission[]>();
  const [page, setPage] = useState<number>(1);
  const rowSize = 6;

  const getPermissionsOfRolePage = async (page: number) => {
    if (role !== null) {
      await axios
        .get(
          `${API_URL}/rbac/role/${role.ID}/permission?page=${page}&pageSize=${rowSize}`,
        )
        .then((res) => {
          setPermissionArray(res.data.Results);
          setCount(res.data.Count);
        });
    }
  };

  useEffect(() => {
    getPermissionsOfRolePage(1);
  }, [role.ID]);

  useEffect(() => {
    if (permissionArray !== undefined) {
      setPage(parseInt(`${(count + 1) / rowSize}`, 10));
    }
  }, [permissionArray]);

  const handleChangePageNum = (
    event: React.ChangeEvent<unknown>,
    value: number,
  ) => {
    getPermissionsOfRolePage(value);
  };

  return (
    <Root sx={{ width: 400, maxWidth: '100%' }}>
      {role && (
        <>
          <h1>Permissions Of Role</h1>
          <div style={{ minHeight: '435px' }}>
            <table aria-label="custom pagination table">
              <thead>
                <tr>
                  <th style={{ textAlign: 'center' }}>할당</th>
                  <th style={{ textAlign: 'center' }}>서비스</th>
                  <th style={{ textAlign: 'center' }}>권한</th>
                  <th style={{ textAlign: 'center' }}>행동</th>
                  <th style={{ textAlign: 'center' }}>대상</th>
                </tr>
              </thead>
              <tbody>
                {permissionArray &&
                  permissionArray.map((permission: Permission) => (
                    <PermissionRow
                      key={permission.ID}
                      permission={permission}
                      roleID={role.ID}
                    />
                  ))}
              </tbody>
            </table>
          </div>
          {permissionArray && (
            <Stack spacing={3}>
              <Pagination
                sx={{ margin: 'auto', marginTop: '10px' }}
                count={page}
                defaultPage={1}
                onChange={handleChangePageNum}
                shape="rounded"
              />
            </Stack>
          )}
        </>
      )}
    </Root>
  );
}
