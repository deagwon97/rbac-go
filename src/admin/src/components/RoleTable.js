import { useState, useEffect } from "react";
import axios from "axios";
import Stack from "@mui/material/Stack";
import { styled } from "@mui/system";
import { Pagination, Button, TextField } from "@mui/material";
import Modal from "@mui/material/Modal";
import Box from "@mui/material/Box";
import { API_URL } from "App";
import "components/roleTransition.css";

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
const style = {
  position: "absolute",
  top: "50%",
  left: "50%",
  transform: "translate(-50%, -50%)",
  width: 400,
  bgcolor: "background.paper",
  border: "2px solid #000",
  boxShadow: 24,
  p: 4,
  display: "flex",
  flexDirection: "column",
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
  const [trans, setTrans] = useState("normal");
  const handleChangePermission = (e, role) => {
    props.onChange(role.ID, role);
    setTrans("transition");
  };
  const [open, setOpen] = useState(false);
  const handleOpen = () => setOpen(true);
  const handleClose = () => setOpen(false);

  const [roleName, setRoleName] = useState(props.row.name);
  const [roleDesc, setRoleDesc] = useState(props.row.description);

  const updateRole = async (roleID, name, desc) => {
    const data = {
      Name: name,
      Description: desc,
    };
    await axios.patch(`${API_URL}/rbac/role/${roleID}`, data);
    props.getRolePage(props.rolePage);
  };
  const deleteRole = async (roleID) => {
    await axios.delete(`${API_URL}/rbac/role/${roleID}`);
    props.getRolePage(props.rolePage);
  };

  const StyledTr = styled("tr")(
    () => `
    &:hover {
      background-color: rgb(24, 118, 209, 0.2);
    }
  `,
  );
  useEffect(() => {
    setRoleName(props.row.Name);
    setRoleDesc(props.row.Description);
  }, [props.row]);

  useEffect(() => {
    if (props.roleID == props.row.ID) {
      setTrans("transition");
    } else {
      setTrans("normal");
    }
  }, [props]);

  return (
    <>
      <StyledTr
        className={trans}
        style={{ height: "54px", boxSizing: "border-box" }}
        onClick={(e) => {
          handleChangePermission(e, props.row);
        }}
      >
        <td>{props.row.Name}</td>
        <td align="right">{props.row.Description}</td>
      </StyledTr>

      <div style={{ width: "0px", border: "none" }}>
        <div style={{ width: "350px", overflow: "visible", display: "flex" }}>
          <div className={trans} style={{ marginLeft: "auto" }}>
            <Button variant="outlined" size="small" onClick={handleOpen}>
              수정
            </Button>
          </div>
        </div>
      </div>
      <Modal
        open={open}
        onClose={handleClose}
        aria-labelledby="modal-modal-title"
        aria-describedby="modal-modal-description"
      >
        <Box sx={style}>
          <TextField
            id="outlined-search"
            size="small"
            label="역할"
            style={{ margin: "20px" }}
            value={roleName}
            onChange={(e) => {
              setRoleName(e.target.value);
            }}
            inputProps={{
              autoComplete: "off",
            }}
          />
          <TextField
            id="outlined-search"
            size="small"
            label="설명"
            style={{ margin: "20px" }}
            value={roleDesc}
            onChange={(e) => {
              setRoleDesc(e.target.value);
            }}
            inputProps={{
              autoComplete: "off",
            }}
          />
          <div style={{ display: "flex", flexDirection: "row", justifyContent: "center" }}>
            <Button
              variant="outlined"
              style={{
                width: "50px",
                margin: "20px",
                color: "black",
                borderBlockColor: "black",
                backgroundColor: "rgb(255, 0, 0, 0.3)",
              }}
              size="small"
              onClick={() => {
                deleteRole(props.roleID);
                handleClose();
              }}
            >
              삭제
            </Button>
            <Button
              variant="outlined"
              style={{
                width: "50px",
                margin: "20px",
                color: "black",
                borderBlockColor: "black",
                backgroundColor: "#E0E0E0",
              }}
              size="small"
              onClick={() => {
                updateRole(props.roleID, roleName, roleDesc);
                handleClose();
              }}
            >
              저장
            </Button>
          </div>
        </Box>
      </Modal>
    </>
  );
}

export default function RoleTable(props) {
  const [roleID, setRoleID] = useState();

  const [open, setOpen] = useState(false);
  const handleOpen = () => setOpen(true);
  const handleClose = () => setOpen(false);

  const [roleName, setRoleName] = useState();
  const [roleDesc, setRoleDesc] = useState();

  const addRole = async (name, desc) => {
    const data = {
      Name: name,
      Description: desc,
    };
    if (name !== null) {
      console.log(name);
      const res = await axios.post(`${API_URL}/rbac/role`, data);
      setRoleID(res.data.ID);
      getRolePage(rolePage);
      props.onChange(res.data);
    }
    setRoleName(null);
    setRoleDesc(null);
  };

  function handleRoleID(roleID, role) {
    setRoleID(roleID);
    props.onChange(role);
  }
  const rowSize = 5;

  const [rolePage, setRolePage] = useState();

  const getRolePage = async (page) => {
    await axios.get(`${API_URL}/rbac/role/list?page=${page}&pageSize=${rowSize}`).then((res) => setRolePage(res.data));
  };

  useEffect(() => {
    getRolePage(rolePage);
  }, []);

  const handleChangePageNum = (event, value) => {
    getRolePage(value);
  };

  return (
    <Root sx={{ width: 350, maxWidth: "100%" }}>
      <h1>Role</h1>
      <div style={{ minHeight: "435px" }}>
        <div style={{ minHeight: "400px" }}>
          <table aria-label="custom pagination table">
            <thead>
              <tr>
                <th style={{ width: "110px" }}>역할</th>
                <th>설명</th>
              </tr>
            </thead>
            <tbody>
              {rolePage &&
                rolePage.Results.map((row, idx) => (
                  <RoleRow
                    key={idx}
                    roleID={roleID}
                    row={row}
                    rolePage={rolePage}
                    getRolePage={getRolePage}
                    onChange={handleRoleID}
                  ></RoleRow>
                ))}
            </tbody>
          </table>
        </div>
        <div style={{ alignSelf: "center" }}>
          <Button variant="outlined" size="small" onClick={handleOpen}>
            추가하기
          </Button>
        </div>
      </div>
      {rolePage && (
        <Stack spacing={2}>
          <Pagination
            sx={{ margin: "auto", marginTop: "10px" }}
            count={parseInt((rolePage.Count + 1) / rowSize)}
            defaultPage={1}
            onChange={handleChangePageNum}
            shape="rounded"
          />
        </Stack>
      )}
      <Modal
        open={open}
        onClose={handleClose}
        aria-labelledby="modal-modal-title"
        aria-describedby="modal-modal-description"
      >
        <Box sx={style}>
          <TextField
            id="outlined-search"
            size="small"
            label="역할"
            style={{ margin: "20px" }}
            value={roleName}
            onChange={(e) => {
              setRoleName(e.target.value);
            }}
            inputProps={{
              autoComplete: "off",
            }}
          />
          <TextField
            id="outlined-search"
            size="small"
            label="설명"
            style={{ margin: "20px" }}
            value={roleDesc}
            onChange={(e) => {
              setRoleDesc(e.target.value);
            }}
            inputProps={{
              autoComplete: "off",
            }}
          />
          <div style={{ display: "flex", flexDirection: "row", justifyContent: "center" }}>
            <Button
              variant="outlined"
              style={{
                width: "50px",
                margin: "20px",
              }}
              size="small"
              onClick={() => {
                addRole(roleName, roleDesc);
                handleClose();
              }}
            >
              저장
            </Button>
          </div>
        </Box>
      </Modal>
    </Root>
  );
}
