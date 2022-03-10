import * as React from "react";
import { useState, useEffect } from "react";
import axios from "axios";
import Pagination from "@mui/material/Pagination";
import Stack from "@mui/material/Stack";
import Button from "@mui/material/Button";
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

const StyledButton = styled(Button)({
  background: "#1876D1",
  "&:hover": {
    background: "#F6FAFD",
    color: "#1876D1",
  },
  border: 0,
  borderRadius: 3,
  marginLeft: "5px",
  float: "right",
  color: "white",
  height: "20px",
  fontSize: 12,
  textTransform: "none",
});

export default function RoleTable() {
  const [page, setPage] = React.useState(1);
  const [rolePage, setRolePage] = useState();

  const getRolePage = async (page) => {
    console.log(`${API_URL}/rbac/role/list`);
    await axios.get(`${API_URL}/rbac/role/list?page=${page}&pageSize=5`).then((res) => setRolePage(res.data));
  };

  useEffect(() => {
    getRolePage(1);
  }, []);

  const handleChangePageNum = (event, value) => {
    setPage(value);
    getRolePage(value);
  };

  return (
    <Root sx={{ width: 500, maxWidth: "100%" }}>
      <div>
        <table aria-label="custom pagination table">
          <thead>
            <tr>
              <th>역할</th>
              <th>설명</th>
              <th> </th>
            </tr>
          </thead>
          <tbody>
            {rolePage &&
              rolePage.results.map((row, idx) => (
                <tr key={idx}>
                  <td>{row.name}</td>
                  <td align="right">{row.description}</td>
                  <td style={{ width: "150px", textAlign: "center" }}>
                    <StyledButton size="small" aria-label="fingerprint">
                      Subject
                    </StyledButton>
                    <StyledButton size="small" color="secondary" aria-label="fingerprint">
                      Permission
                    </StyledButton>
                  </td>
                </tr>
              ))}
          </tbody>
        </table>
        {rolePage && (
          <Stack spacing={3}>
            <Pagination
              sx={{ margin: "auto", marginTop: "10px" }}
              count={parseInt(rolePage.count / 5) + 1}
              defaultPage={page}
              onChange={handleChangePageNum}
              shape="rounded"
            />
          </Stack>
        )}
      </div>
    </Root>
  );
}
