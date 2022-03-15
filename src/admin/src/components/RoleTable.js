import { useState, useEffect } from "react";
import axios from "axios";
import Pagination from "@mui/material/Pagination";
import Stack from "@mui/material/Stack";
import { styled } from "@mui/system";
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

function RoleRow(props) {
  const handleChangePermission = (e, role) => {
    props.onChange(role.id, role);
  };
  const StyledTr = styled("tr")(
    () => `
    &:hover {
      background-color: rgba(201, 76, 76, 0.1);
    }
  `,
  );

  return (
    <>
      {props.roleID == props.row.id ? (
        <tr style={{ backgroundColor: "rgba(201, 76, 76, 0.3)" }}>
          <td>{props.row.name}</td>
          <td align="right">{props.row.description}</td>
        </tr>
      ) : (
        <StyledTr
          onClick={(e) => {
            handleChangePermission(e, props.row);
          }}
        >
          <td>{props.row.name}</td>
          <td align="right">{props.row.description}</td>
        </StyledTr>
      )}
    </>
  );
}

export default function RoleTable(props) {
  const [roleID, setRoleID] = useState();
  function handleRoleID(roleID, role) {
    setRoleID(roleID);
    props.onChange(role);
  }

  const [rolePage, setRolePage] = useState();

  const getRolePage = async (page) => {
    await axios.get(`${API_URL}/rbac/role/list?page=${page}&pageSize=5`).then((res) => setRolePage(res.data));
  };

  useEffect(() => {
    getRolePage(1);
  }, []);

  const handleChangePageNum = (event, value) => {
    getRolePage(value);
  };

  return (
    <Root sx={{ width: 500, maxWidth: "100%" }}>
      <div style={{ height: "200px" }}>
        <table aria-label="custom pagination table">
          <thead>
            <tr>
              <th>역할</th>
              <th>설명</th>
            </tr>
          </thead>
          <tbody>
            {rolePage &&
              rolePage.results.map((row, idx) => (
                <RoleRow key={idx} roleID={roleID} row={row} onChange={handleRoleID}></RoleRow>
              ))}
          </tbody>
        </table>
      </div>
      {rolePage && (
        <Stack spacing={2}>
          <Pagination
            sx={{ margin: "auto", marginTop: "10px" }}
            count={parseInt(rolePage.count / 5) + 1}
            defaultPage={1}
            onChange={handleChangePageNum}
            shape="rounded"
          />
        </Stack>
      )}
    </Root>
  );
}
