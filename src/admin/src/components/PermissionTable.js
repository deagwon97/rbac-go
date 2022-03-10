import * as React from "react";
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

export default function PermissionTable() {
  const [page, setPage] = React.useState(1);
  const [permissionPage, setPermissionPage] = useState();

  const getPermissionPage = async (page) => {
    await axios
      .get(`${API_URL}/rbac/permission/list?page=${page}&pageSize=5`)
      .then((res) => setPermissionPage(res.data));
  };

  useEffect(() => {
    getPermissionPage(1);
  }, []);

  const handleChangePageNum = (event, value) => {
    setPage(value);
    getPermissionPage(value);
  };

  return (
    <Root sx={{ width: 500, maxWidth: "100%" }}>
      <table aria-label="custom pagination table">
        <thead>
          <tr>
            <th>서비스</th>
            <th>권한</th>
            <th>행동</th>
            <th>대상</th>
          </tr>
        </thead>
        <tbody>
          {permissionPage &&
            permissionPage.results.map((row, idx) => (
              <tr key={idx}>
                <td>{row.name}</td>
                <td style={{ width: 120 }} align="right">
                  {row.name}
                </td>
                <td style={{ width: 120 }} align="right">
                  {row.action}
                </td>
                <td style={{ width: 120 }} align="right">
                  {row.object}
                </td>
              </tr>
            ))}
        </tbody>
      </table>
      {permissionPage && (
        <Stack spacing={3}>
          <Pagination
            sx={{ margin: "auto", marginTop: "10px" }}
            count={parseInt(permissionPage.count / 5)}
            defaultPage={page}
            onChange={handleChangePageNum}
            shape="rounded"
          />
        </Stack>
      )}
    </Root>
  );
}
