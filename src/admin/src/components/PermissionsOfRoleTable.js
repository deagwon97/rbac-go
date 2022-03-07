import * as React from "react";
import { styled } from "@mui/system";
import Checkbox from "@mui/material/Checkbox";

function createData(id, serviceName, name, action, object, checked) {
  return { id, serviceName, name, action, object, checked };
}

const rows = [
  createData("1", "bdg블로그", "게시판", "목록조회", "공지", false),
  createData("2", "bdg블로그", "게시판", "목록조회", "비밀", false),
  createData("3", "bdg블로그", "게시판", "목록조회", "자유", false),
  createData("4", "bdg블로그", "게시판", "삭제", "공지", false),
  createData("5", "bdg블로그", "게시판", "삭제", "비밀", false),
  createData("6", "bdg블로그", "게시판", "삭제", "자유", false),
  createData("7", "bdg블로그", "게시판", "상세조회", "공지", false),
  createData("8", "bdg블로그", "게시판", "상세조회", "비밀", false),
  createData("9", "bdg블로그", "게시판", "상세조회", "자유", false),
  createData("10", "bdg블로그", "게시판", "수정", "공지", false),
  createData("11", "bdg블로그", "게시판", "수정", "비밀", false),
  createData("12", "bdg블로그", "게시판", "수정", "자유", false),
  createData("12", "bdg블로그", "채팅", "목록조회", "VIP", false),
  createData("13", "bdg블로그", "채팅", "목록조회", "도매", false),
  createData("14", "bdg블로그", "채팅", "삭제", "VIP", false),
  createData("15", "bdg블로그", "채팅", "삭제", "도매", false),
  createData("16", "bdg블로그", "채팅", "상세조회", "VIP", false),
  createData("17", "bdg블로그", "채팅", "상세조회", "도매", false),
  createData("18", "bdg블로그", "채팅", "수정", "VIP", false),
  createData("19", "bdg블로그", "채팅", "수정", "도매", false),
].sort((a, b) => (a.calories < b.calories ? -1 : 1));

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

export default function PermissionsOfRoleTable() {
  const page = React.useState(0)[0];
  const setChecked = React.useState(0)[1];
  const rowsPerPage = React.useState(-1)[0];

  const label = { inputProps: { "aria-label": "Checkbox demo" } };

  const handleChange = (event, value) => {
    console.log(value);
    console.log(event.target.checked);
    setChecked(event.target.checked);
  };

  // const handleChange = function (value) {
  //   console.log(value);
  //   setChecked(event.target.checked);
  // };

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
          {(rowsPerPage > 0 ? rows.slice(page * rowsPerPage, page * rowsPerPage + rowsPerPage) : rows).map(
            (row, idx) => (
              <tr key={idx}>
                <td>{row.serviceName}</td>
                <td style={{ width: 90 }} align="right">
                  {row.name}
                </td>
                <td style={{ width: 90 }} align="right">
                  {row.action}
                </td>
                <td style={{ width: 90 }} align="right">
                  {row.object}
                </td>
                <td style={{ width: 45, textAlign: "center" }}>
                  <Checkbox onChange={(e) => handleChange(e, row.id)} {...label} />
                </td>
              </tr>
            ),
          )}
        </tbody>
      </table>
    </Root>
  );
}
