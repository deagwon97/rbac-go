import { useState } from "react";
import PermissionsOfRoleTable from "components/PermissionsOfRoleTable";
import SubjectsOfRoleTable from "components/SubjectsOfRoleTable";
import RoleTable from "components/RoleTable";
import PermissionTable from "components/PermissionTable";
// import SubjectTable from "components/SubjectTable";

const MainPage = () => {
  const [role, setRole] = useState(null);

  function handleChange(role) {
    setRole(role);
  }
  return (
    <>
      <div style={{ width: "fit-content", marginLeft: "auto", marginRight: "auto", alignContent: "center" }}>
        <h1>Permissions</h1>
        <PermissionTable />
        <h1>Roles</h1>
        <RoleTable onChange={handleChange} />
        <PermissionsOfRoleTable role={role} />
        <SubjectsOfRoleTable role={role} />
      </div>
      <div>
        <br />
        <br />
        <p align="center">github.com/deagwon97/rbac-go</p>
        <br />
        <br />
      </div>
    </>
  );
};

export default MainPage;
