// import { useNavigate } from "react-router-dom";
import PermissionsOfRoleTable from "components/PermissionsOfRoleTable";
import RoleTable from "components/RoleTable";
import PermissionTable from "components/PermissionTable";
import SubjectTable from "components/SubjectTable";

const MainPage = () => {
  return (
    <>
      <div style={{ width: "fit-content", marginLeft: "auto", marginRight: "auto", alignContent: "center" }}>
        <h1>Roles</h1>
        <RoleTable></RoleTable>
        <h1>Permissions</h1>
        <PermissionTable></PermissionTable>
        <h1>Permissions Of Role</h1>
        <PermissionsOfRoleTable></PermissionsOfRoleTable>
        <h1>Subjects Of Role</h1>
        <SubjectTable></SubjectTable>
      </div>
    </>
  );
};

export default MainPage;
