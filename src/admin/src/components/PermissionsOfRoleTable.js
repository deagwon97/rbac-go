import * as React from "react";
import { useState, useEffect } from "react";
import axios from "axios";
import Pagination from "@mui/material/Pagination";
import Stack from "@mui/material/Stack";
import { styled } from "@mui/system";
import Checkbox from "@mui/material/Checkbox";
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
  const [checked, setChecked] = React.useState(props.row.is_allowed);

  useEffect(() => {
    setChecked(props.row.is_allowed);
  }, [props]);

  const handleChange = (event, roleID, permissionID) => {
    setChecked(event.target.checked);
    addPermissionAssignment(roleID, permissionID);
  };

  const addPermissionAssignment = async (roleID, permissionID) => {
    const data = {
      permission_id: permissionID,
      role_id: roleID,
    };
    await axios.post(`${API_URL}/rbac/permission-assignment`, data);
  };

  return (
    <>
      <tr key={props.idx}>
        <td>{props.row.service_name}</td>
        <td style={{ width: 90 }} align="right">
          {props.row.name}
        </td>
        <td style={{ width: 90 }} align="right">
          {props.row.action}
        </td>
        <td style={{ width: 90 }} align="right">
          {props.row.object}
        </td>
        <td style={{ width: 45, textAlign: "center" }}>
          <Checkbox checked={checked} onChange={(e) => handleChange(e, 31, props.row.id)} {...label} />
        </td>
      </tr>
    </>
  );
}

export default function PermissionsOfRoleTable() {
  const [page, setPage] = React.useState(1);
  const [permissionsOfRolePage, setPermissionsOfRolePage] = useState();

  const getPermissionsOfRolePage = async (page) => {
    await axios
      .get(`${API_URL}/rbac/role/31/permission?page=${page}&pageSize=5`)
      .then((res) => setPermissionsOfRolePage(res.data));
  };

  useEffect(() => {
    getPermissionsOfRolePage(1);
  }, []);

  const handleChangePageNum = (event, value) => {
    setPage(value);
    getPermissionsOfRolePage(value);
  };

  return (
    <Root sx={{ width: 500, maxWidth: "100%" }}>
      <table aria-label="custom pagination table">
        <thead>
          <tr>
            <th style={{ textAlign: "center" }}>서비스</th>
            <th style={{ textAlign: "center" }}>권한</th>
            <th style={{ textAlign: "center" }}>행동</th>
            <th style={{ textAlign: "center" }}>대상</th>
            <th style={{ textAlign: "center" }}>확인</th>
          </tr>
        </thead>
        <tbody>
          {permissionsOfRolePage &&
            permissionsOfRolePage.results.map((row, idx) => <PermissionRow idx={idx} row={row}></PermissionRow>)}
        </tbody>
      </table>
      {permissionsOfRolePage && (
        <Stack spacing={3}>
          <Pagination
            sx={{ margin: "auto", marginTop: "10px" }}
            count={parseInt(permissionsOfRolePage.count / 5)}
            defaultPage={page}
            onChange={handleChangePageNum}
            shape="rounded"
          />
        </Stack>
      )}
    </Root>
  );
}
