import { useState, useEffect } from "react";
import axios from "axios";
import Stack from "@mui/material/Stack";
import { styled } from "@mui/system";
import { Checkbox, Pagination } from "@mui/material";
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
  const [checked, setChecked] = useState(props.row.IsAllowed);

  useEffect(() => {
    setChecked(props.row.IsAllowed);
  }, [props]);

  const addSubjectAssignment = async (subjectID, roleID) => {
    const data = {
      SubjectID: subjectID,
      RoleID: roleID,
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
      <tr key={props.idx} style={{ height: "55px" }}>
        <td>{props.row.SubjectID}</td>
        <td>{props.row.Name}</td>
        <td style={{ width: 45, textAlign: "center" }}>
          <Checkbox checked={checked} onChange={(e) => handleChange(e, props.roleID, props.row.SubjectID)} {...label} />
        </td>
      </tr>
    </>
  );
}

export default function SubjectsOfRoleTable(props) {
  const [subjectsOfRolePage, setSubjectsOfRolePage] = useState();
  const [role, setRole] = useState(props.role);
  const [subjects, setSubjects] = useState();
  const rowSize = 6;

  const getSubjectsOfRolePage = async (page) => {
    if (role !== null) {
      await axios
        .get(`${API_URL}/rbac/role/${role.ID}/subject?page=${page}&pageSize=${rowSize}`)
        .then((res) => setSubjectsOfRolePage(res.data));
    }
  };

  const getSuibjectsName = async (IDList) => {
    let data = {
      IDList: [],
    };

    var i;
    if (IDList.length > 0) {
      for (i = 0; i < IDList.length; i++) {
        data.IDList[i] = IDList[i].SubjectID;
      }
    }
    await axios.post(`${API_URL}/account/name/list`, data).then((res) => setSubjects(res.data));
  };

  useEffect(() => {
    setRole(props.role);
  }, [props.role]);

  useEffect(() => {
    getSubjectsOfRolePage(1);
  }, [role]);

  useEffect(() => {
    if (subjectsOfRolePage !== undefined) {
      getSuibjectsName(subjectsOfRolePage.Results);
    }
  }, [subjectsOfRolePage]);

  const handleChangePageNum = (event, value) => {
    getSubjectsOfRolePage(value);
  };

  return (
    <Root sx={{ width: 400, maxWidth: "100%" }}>
      {role && (
        <>
          <h1>Subjects Of Role</h1>
          <div style={{ minHeight: "435px" }}>
            <table aria-label="custom pagination table">
              <thead>
                <tr>
                  <th style={{ textAlign: "center" }}>ID</th>
                  <th style={{ textAlign: "center" }}>이름</th>
                  <th style={{ textAlign: "center" }}>할당</th>
                </tr>
              </thead>
              <tbody>
                {subjectsOfRolePage &&
                  subjects &&
                  subjectsOfRolePage.Results.map((row, idx) => {
                    row.Name = subjects[idx].Name;
                    return <SubjectRow key={idx} idx={idx} row={row} roleID={role.ID}></SubjectRow>;
                  })}
              </tbody>
            </table>
          </div>
          {subjectsOfRolePage && (
            <Stack spacing={3}>
              <Pagination
                sx={{ margin: "auto", marginTop: "10px" }}
                count={parseInt((subjectsOfRolePage.Count + 1) / rowSize)}
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
