import { useState, useEffect } from "react";
import axios from "axios";
import Stack from "@mui/material/Stack";
import { styled } from "@mui/system";
import { Checkbox, TextField, Pagination, Button } from "@mui/material";
import { API_URL } from "App";

const grey = {
  50: "#F3F6F9",
  100: "#E7EBF0",
  200: "#E0E3E7",
  300: "#CDD2D7",
  400: "#B2BAC2",
  500: "#A0AAB4",
  600: "#6F7E8C",
  700: "#3E5060",
  800: "#2D3843",
  900: "#1A2027",
};

const Root = styled("div")(
  ({ theme }) => `
  table {
    font-family: IBM Plex Sans, sans-serif;
    font-size: 0.875rem;
    border-collapse: collapse;
    width: 100%;
  }

  td,
  th {
    border: 1px solid ${theme.palette.mode === "dark" ? grey[800] : grey[200]};
    text-align: left;
    padding: 6px;
  }

  th {
    background-color: ${theme.palette.mode === "dark" ? grey[900] : grey[100]};
  }
  `,
);

function PermissionRow(props) {
  const label = { inputProps: { "aria-label": "Checkbox demo" } };
  const [checked, setChecked] = useState(props.row.is_allowed);

  useEffect(() => {
    setChecked(props.row.is_allowed);
  }, [props]);

  const addPermissionAssignment = async (permissionID, roleID) => {
    const data = {
      permission_id: permissionID,
      role_id: roleID,
    };
    await axios.post(`${API_URL}/rbac/permission-assignment`, data);
  };
  const deletePermissionAssignment = async (permissionID, roleID) => {
    await axios.delete(`${API_URL}/rbac/permission-assignment?permissionID=${permissionID}&roleID=${roleID}`);
  };

  const handleChange = (event, roleID, permissionID) => {
    setChecked(event.target.checked);

    if (event.target.checked === true) {
      addPermissionAssignment(permissionID, roleID);
    } else {
      deletePermissionAssignment(permissionID, roleID);
    }
  };

  return (
    <>
      <tr key={props.idx}>
        <td style={{ width: 35, textAlign: "center" }}>
          <Checkbox checked={checked} onChange={(e) => handleChange(e, props.roleID, props.row.id)} {...label} />
        </td>
        <td>{props.row.service_name}</td>
        <td style={{ width: 70 }} align="right">
          {props.row.name}
        </td>
        <td style={{ width: 70 }} align="right">
          {props.row.action}
        </td>
        <td style={{ width: 70 }} align="right">
          {props.row.object}
        </td>
      </tr>
    </>
  );
}

export default function PermissionsOfRoleTable(props) {
  const [permissionsOfRolePage, setPermissionsOfRolePage] = useState();
  const [role, setRole] = useState(props.role);

  const getPermissionsOfRolePage = async (page) => {
    if (role !== null) {
      await axios
        .get(`${API_URL}/rbac/role/${role.id}/permission?page=${page}&pageSize=10`)
        .then((res) => setPermissionsOfRolePage(res.data));
    }
  };

  useEffect(() => {
    setRole(props.role);
  }, [props.role]);

  useEffect(() => {
    getPermissionsOfRolePage(1);
  }, [role]);

  const handleChangePageNum = (event, value) => {
    getPermissionsOfRolePage(value);
  };

  return (
    <Root sx={{ width: 400, maxWidth: "100%" }}>
      {role && (
        <>
          <h1>Permissions Of Role</h1>
          <div
            style={{
              display: "flex",
              alignItems: "baseline",
              flexDirection: "row",
              justifyContent: "space-between",
              marginBottom: "10px",
              alignSelf: "center",
            }}
          >
            <TextField style={{ width: "360px" }} id="outlined-search" size="small" label="권한 검색" type="search" />
            <Button variant="outlined" sx={{ width: "20px" }}>
              검색
            </Button>
          </div>

          <div style={{ minHeight: "585px" }}>
            <table aria-label="custom pagination table">
              <thead>
                <tr>
                  <th style={{ textAlign: "center" }}>할당</th>
                  <th style={{ textAlign: "center" }}>서비스</th>
                  <th style={{ textAlign: "center" }}>권한</th>
                  <th style={{ textAlign: "center" }}>행동</th>
                  <th style={{ textAlign: "center" }}>대상</th>
                </tr>
              </thead>
              <tbody>
                {permissionsOfRolePage &&
                  permissionsOfRolePage.results.map((row, idx) => (
                    <PermissionRow key={idx} idx={idx} row={row} roleID={role.id}></PermissionRow>
                  ))}
              </tbody>
            </table>
          </div>
          {permissionsOfRolePage && (
            <Stack spacing={3}>
              <Pagination
                sx={{ margin: "auto", marginTop: "10px" }}
                count={parseInt((permissionsOfRolePage.count + 1) / 10)}
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
