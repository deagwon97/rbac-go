import { useState, useEffect } from 'react';
import axios from 'axios';
import { Checkbox } from '@mui/material';
import { API_URL } from 'App';

interface Permission {
  IsAllowed: boolean;
  ID: number;
  ServiceName: string;
  Name: string;
  Action: string;
  Object: string;
}

type PermissionRowProps = {
  permission: Permission;
  roleID: number;
};

export default function PermissionRow({
  permission,
  roleID,
}: PermissionRowProps) {
  //   const label = { inputProps: { 'aria-label': 'Checkbox demo' } };
  const [checked, setChecked] = useState(permission.IsAllowed);

  useEffect(() => {
    setChecked(permission.IsAllowed);
  }, [roleID, permission]);

  const addPermissionAssignment = async (
    permissionID: number,
    roleID: number,
  ) => {
    const data = {
      PermissionID: permissionID,
      RoleID: roleID,
    };
    await axios.post(`${API_URL}/rbac/permission-assignment`, data);
  };
  const deletePermissionAssignment = async (
    permissionID: number,
    roleID: number,
  ) => {
    await axios.delete(
      `${API_URL}/rbac/permission-assignment?permissionID=${permissionID}&roleID=${roleID}`,
    );
  };

  const handleChange = (
    event: React.ChangeEvent<HTMLInputElement>,
    roleID: number,
    permissionID: number,
  ) => {
    setChecked(event.target.checked);
    if (event.target.checked === true) {
      addPermissionAssignment(permissionID, roleID);
    } else {
      deletePermissionAssignment(permissionID, roleID);
    }
  };

  return (
    <tr key={permission.ID} style={{ height: '55px' }}>
      <td style={{ width: 35, textAlign: 'center' }}>
        <Checkbox
          checked={checked}
          onChange={(e) => handleChange(e, roleID, permission.ID)}
        />
      </td>
      <td>{permission.ServiceName}</td>
      <td style={{ width: 70 }} align="right">
        {permission.Name}
      </td>
      <td style={{ width: 70 }} align="right">
        {permission.Action}
      </td>
      <td style={{ width: 70 }} align="right">
        {permission.Object}
      </td>
    </tr>
  );
}
