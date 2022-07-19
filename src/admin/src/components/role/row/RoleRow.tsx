import { useState, useEffect } from 'react';
// import { styled } from '@mui/system';
import 'components/role/row/RoleRow.scss';

export interface Role {
  ID: number;
  Name: string;
  Description: string;
}

export type RoleProps = {
  role: Role;
  isChecked: boolean;
  onChange: (role: Role) => void;
};

export default function RoleRow({ role, isChecked, onChange }: RoleProps) {
  const [roleName, setRoleName] = useState<string>(role.Name);
  const [roleDesc, setRoleDesc] = useState<string>(role.Description);
  const [rowName, setRowName] = useState<string>('no-check');

  const handleChangeSubject = (e: React.MouseEvent, role: Role) => {
    onChange(role);
  };

  useEffect(() => {
    setRoleName(role.Name);
    setRoleDesc(role.Description);
  }, [role]);

  useEffect(() => {
    if (isChecked) {
      setRowName('check');
    } else {
      setRowName('no-check');
    }
  }, [isChecked]);

  return (
    <tr
      className={`${rowName}`}
      style={{ height: '54px', boxSizing: 'border-box' }}
      onClick={(event: React.MouseEvent) => {
        return handleChangeSubject(event, role);
      }}
    >
      <td>{roleName}</td>
      <td align="right">{roleDesc}</td>
    </tr>
  );
}
