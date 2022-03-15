import { useState, useEffect } from "react";
import axios from "axios";
import Stack from "@mui/material/Stack";
import { styled } from "@mui/system";
import { Checkbox, TextField, Pagination } from "@mui/material";
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

function SubjectRow(props) {
  const label = { inputProps: { "aria-label": "Checkbox demo" } };
  const [checked, setChecked] = useState(props.row.is_allowed);

  useEffect(() => {
    setChecked(props.row.is_allowed);
  }, [props]);

  const addSubjectAssignment = async (subjectID, roleID) => {
    const data = {
      subject_id: subjectID,
      role_id: roleID,
    };
    await axios.post(`${API_URL}/rbac/subject-assignment`, data);
  };
  const deleteSubjectAssignment = async (subjectID, roleID) => {
    await axios.delete(`${API_URL}/rbac/subject-assignment?subjectID=${subjectID}&roleID=${roleID}`);
  };

  const handleChange = (event, roleID, subjectID) => {
    setChecked(event.target.checked);

    if (event.target.checked === true) {
      addSubjectAssignment(subjectID, roleID);
    } else {
      deleteSubjectAssignment(subjectID, roleID);
    }
  };

  return (
    <>
      <tr key={props.idx}>
        <td>{props.row.subject_id}</td>
        <td style={{ width: 45, textAlign: "center" }}>
          <Checkbox
            checked={checked}
            onChange={(e) => handleChange(e, props.roleID, props.row.subject_id)}
            {...label}
          />
        </td>
      </tr>
    </>
  );
}

export default function SubjectsOfRoleTable(props) {
  const [subjectsOfRolePage, setSubjectsOfRolePage] = useState();
  const [role, setRole] = useState(props.role);

  const getSubjectsOfRolePage = async (page) => {
    if (role !== null) {
      await axios
        .get(`${API_URL}/rbac/role/${role.id}/subject?page=${page}&pageSize=5`)
        .then((res) => setSubjectsOfRolePage(res.data));
    }
  };

  useEffect(() => {
    setRole(props.role);
  }, [props.role]);

  useEffect(() => {
    getSubjectsOfRolePage(1);
  }, [role]);

  const handleChangePageNum = (event, value) => {
    getSubjectsOfRolePage(value);
  };

  return (
    <Root sx={{ width: 300, maxWidth: "100%" }}>
      {role && (
        <>
          <h1>Subjects Of Role</h1>

          <TextField
            style={{ width: "100%", marginBottom: "10px" }}
            id="outlined-search"
            size="small"
            label="사용자 검색"
            type="search"
          />

          <div style={{ minHeight: "310px" }}>
            <table aria-label="custom pagination table">
              <thead>
                <tr>
                  <th style={{ textAlign: "center" }}>유저 인덱스</th>
                  <th style={{ textAlign: "center" }}>할당</th>
                </tr>
              </thead>
              <tbody>
                {subjectsOfRolePage &&
                  subjectsOfRolePage.results.map((row, idx) => (
                    <SubjectRow key={idx} idx={idx} row={row} roleID={role.id}></SubjectRow>
                  ))}
              </tbody>
            </table>
          </div>
          {subjectsOfRolePage && (
            <Stack spacing={3}>
              <Pagination
                sx={{ margin: "auto", marginTop: "10px" }}
                count={parseInt((subjectsOfRolePage.count + 1) / 5)}
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
