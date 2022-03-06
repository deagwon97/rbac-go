import * as React from "react";
import { styled } from "@mui/system";
import Checkbox from "@mui/material/Checkbox";
import TablePaginationUnstyled from "@mui/base/TablePaginationUnstyled";

const label = { inputProps: { "aria-label": "Checkbox demo" } };

function createData(id, serviceName, name, action, object) {
  return { id, serviceName, name, action, object };
}

const rows = [
  createData("1", "bdg블로그", "게시판", "목록조회", "공지"),
  createData("2", "bdg블로그", "게시판", "목록조회", "비밀"),
  createData("3", "bdg블로그", "게시판", "목록조회", "자유"),
  createData("4", "bdg블로그", "게시판", "삭제", "공지"),
  createData("5", "bdg블로그", "게시판", "삭제", "비밀"),
  createData("6", "bdg블로그", "게시판", "삭제", "자유"),
  createData("7", "bdg블로그", "게시판", "상세조회", "공지"),
  createData("8", "bdg블로그", "게시판", "상세조회", "비밀"),
  createData("9", "bdg블로그", "게시판", "상세조회", "자유"),
  createData("10", "bdg블로그", "게시판", "수정", "공지"),
  createData("11", "bdg블로그", "게시판", "수정", "비밀"),
  createData("12", "bdg블로그", "게시판", "수정", "자유"),
  createData("12", "bdg블로그", "채팅", "목록조회", "VIP"),
  createData("13", "bdg블로그", "채팅", "목록조회", "도매"),
  createData("14", "bdg블로그", "채팅", "삭제", "VIP"),
  createData("15", "bdg블로그", "채팅", "삭제", "도매"),
  createData("16", "bdg블로그", "채팅", "상세조회", "VIP"),
  createData("17", "bdg블로그", "채팅", "상세조회", "도매"),
  createData("18", "bdg블로그", "채팅", "수정", "VIP"),
  createData("19", "bdg블로그", "채팅", "수정", "도매"),
].sort((a, b) => (a.calories < b.calories ? -1 : 1));

const blue = {
  200: "#A5D8FF",
  400: "#3399FF",
};

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

const CustomTablePagination = styled(TablePaginationUnstyled)(
  ({ theme }) => `
  & .MuiTablePaginationUnstyled-spacer {
    display: none;
  }
  & .MuiTablePaginationUnstyled-toolbar {
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    gap: 10px;

    @media (min-width: 768px) {
      flex-direction: row;
      align-items: center;
    }
  }
  & .MuiTablePaginationUnstyled-selectLabel {
    margin: 0;
  }
  & .MuiTablePaginationUnstyled-select {
    padding: 2px;
    border: 1px solid ${theme.palette.mode === "dark" ? grey[800] : grey[200]};
    border-radius: 50px;
    background-color: transparent;
    &:hover {
      background-color: ${theme.palette.mode === "dark" ? grey[800] : grey[50]};
    }
    &:focus {
      outline: 1px solid ${theme.palette.mode === "dark" ? blue[400] : blue[200]};
    }
  }
  & .MuiTablePaginationUnstyled-displayedRows {
    margin: 0;

    @media (min-width: 768px) {
      margin-left: auto;
    }
  }
  & .MuiTablePaginationUnstyled-actions {
    padding: 2px;
    border: 1px solid ${theme.palette.mode === "dark" ? grey[800] : grey[200]};
    border-radius: 50px;
    text-align: center;
  }
  & .MuiTablePaginationUnstyled-actions > button {
    margin: 0 8px;
    border: transparent;
    border-radius: 2px;
    background-color: transparent;
    &:hover {
      background-color: ${theme.palette.mode === "dark" ? grey[800] : grey[50]};
    }
    &:focus {
      outline: 1px solid ${theme.palette.mode === "dark" ? blue[400] : blue[200]};
    }
  }
  `,
);

export default function PermissionsOfRoleTable() {
  const [page, setPage] = React.useState(0);
  const [rowsPerPage, setRowsPerPage] = React.useState(-1);

  // Avoid a layout jump when reaching the last page with empty rows.
  const emptyRows = page > 0 ? Math.max(0, (1 + page) * rowsPerPage - rows.length) : 0;

  const handleChangePage = (event, newPage) => {
    setPage(newPage);
  };

  const handleChangeRowsPerPage = (event) => {
    setRowsPerPage(parseInt(event.target.value, 10));
    setPage(0);
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
            <th>확인</th>
          </tr>
        </thead>
        <tbody>
          {(rowsPerPage > 0 ? rows.slice(page * rowsPerPage, page * rowsPerPage + rowsPerPage) : rows).map(
            (row, idx) => (
              <tr key={row.idx}>
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
                <td style={{ width: 90 }} align="right">
                  <Checkbox {...label} defaultChecked />
                </td>
              </tr>
            ),
          )}

          {emptyRows > 0 && (
            <tr style={{ height: 41 * emptyRows }}>
              <td colSpan={5} />
            </tr>
          )}
        </tbody>
        <tfoot>
          <tr>
            <CustomTablePagination
              rowsPerPageOptions={[5, 10, 25, { label: "All", value: -1 }]}
              colSpan={5}
              count={rows.length}
              rowsPerPage={rowsPerPage}
              page={page}
              componentsProps={{
                select: {
                  "aria-label": "rows per page",
                },
                actions: {
                  showFirstButton: true,
                  showLastButton: true,
                },
              }}
              onPageChange={handleChangePage}
              onRowsPerPageChange={handleChangeRowsPerPage}
            />
          </tr>
        </tfoot>
      </table>
    </Root>
  );
}
