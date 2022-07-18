import { useState } from "react";
import PermissionsOfRoleTable from "components/PermissionsOfRoleTable";
import SubjectsOfRoleTable from "components/SubjectsOfRoleTable";
import RoleTable from "components/RoleTable";
import Paper from "@mui/material/Paper";
import Grow from "@mui/material/Grow";

const MainPage = () => {
  const [role, setRole] = useState(null);

  const [checked, setChecked] = useState(false);

  function handleChange(role) {
    setRole(role);
    setChecked(true);
  }

  return (
    <>
    
      <div
        style={{
          display: "flex",
          maxWidth: "1300px",
          marginLeft: "auto",
          marginRight: "auto",
          flexDirection: "row",
          justifyContent: "space-around",
          alignSelf: "center",
        }}
      >
        <Grow in={checked}>
          <Paper elevation={0}>
            <SubjectsOfRoleTable role={role} />
          </Paper>
        </Grow>
        <div>
          <RoleTable onChange={handleChange} />
          <br />
          <p style={{ marginTop: "10%" }} align="center">
            github.com/deagwon97/rbac-go
          </p>
          <br />
        </div>
        <Grow in={checked}>
          <Paper elevation={0}>
            <PermissionsOfRoleTable role={role} />
          </Paper>
        </Grow>
      </div>
    </>
  );
};

export default MainPage;
