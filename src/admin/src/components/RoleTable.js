import { useState, useEffect } from "react";
import axios from "axios";
import Stack from "@mui/material/Stack";
import { styled } from "@mui/system";
import { TextField, Pagination, Button } from "@mui/material";
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
      background-color: rgb(24, 118, 209, 0.2);
    }
  `,
  );

  return (
    <>
      {props.roleID == props.row.id ? (
        <tr style={{ backgroundColor: "rgb(24, 118, 209, 0.6)" }}>
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
    await axios.get(`${API_URL}/rbac/role/list?page=${page}&pageSize=10`).then((res) => setRolePage(res.data));
  };

  useEffect(() => {
    getRolePage(1);
  }, []);

  const handleChangePageNum = (event, value) => {
    getRolePage(value);
  };

  return (
    <Root sx={{ width: 350, maxWidth: "100%" }}>
      <h1>Role</h1>
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
        <TextField style={{ width: "360px" }} id="outlined-search" size="small" label="역할 검색" type="search" />
        <Button variant="outlined" sx={{ width: "20px" }}>
          검색
        </Button>
      </div>

      <div style={{ minHeight: "585px" }}>
        <div style={{ minHeight: "400px" }}>
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

        <div
          style={{
            display: "flex",
            alignItems: "baseline",
            flexDirection: "row",
            justifyContent: "left",
            marginBottom: "10px",
            alignSelf: "left",
          }}
        >
          <Button variant="outlined" size="small">
            삭제
          </Button>
          <Button variant="outlined" size="small">
            수정
          </Button>
        </div>
        <div
          style={{
            display: "flex",
            alignItems: "baseline",
            flexDirection: "row",
            justifyContent: "space-around",
            marginBottom: "10px",
            alignSelf: "center",
          }}
        >
          <TextField style={{ width: "40%" }} id="outlined-search" size="small" label="역할" type="search" />
          <TextField style={{}} id="outlined-search" size="small" label="설명" type="search" />
          <Button variant="outlined" style={{ width: "40px" }}>
            추가
          </Button>
        </div>
      </div>
      {rolePage && (
        <Stack spacing={2}>
          <Pagination
            sx={{ margin: "auto", marginTop: "10px" }}
            count={parseInt((rolePage.count + 1) / 10)}
            defaultPage={1}
            onChange={handleChangePageNum}
            shape="rounded"
          />
        </Stack>
      )}
    </Root>
  );
}
