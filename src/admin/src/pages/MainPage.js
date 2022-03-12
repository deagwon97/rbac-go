import { useState } from "react";
import PermissionsOfRoleTable from "components/PermissionsOfRoleTable";
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
      </div>
    </>
  );
};

export default MainPage;
