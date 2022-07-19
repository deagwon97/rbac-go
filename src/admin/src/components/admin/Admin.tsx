import { useState } from 'react';
import { Role } from 'components/role/row/RoleRow';
import RoleTable from 'components/role/table/RoleTable';
import PermissionTable from 'components/permission/table/PermissionTable';
import Paper from '@mui/material/Paper';
import Grow from '@mui/material/Grow';
import SubjectTable from 'components/subject/table/SubjectTable';
import 'components/admin/Admin.scss';

function MainPage() {
  const initRole: Role = { ID: -1, Name: '', Description: '' };
  const [role, setRole] = useState<Role>(initRole);
  const [checked, setChecked] = useState<boolean>(false);

  const handleChange = (role: Role) => {
    setRole(role);
    setChecked(true);
  };

  return (
    <div className="background">
      {role && (
        <Grow in={checked}>
          <Paper elevation={0}>
            <SubjectTable role={role} />
          </Paper>
        </Grow>
      )}
      <div>
        <RoleTable onChange={handleChange} role={role} />
      </div>
      {role && (
        <Grow in={checked}>
          <Paper elevation={0}>
            <PermissionTable role={role} />
          </Paper>
        </Grow>
      )}
    </div>
  );
}

export default MainPage;
